package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func handler(ctx context.Context, s3Event events.S3Event) error {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to load default config", "error", err)
		return err
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	slog.InfoContext(ctx, "processing request", "num_of_records", len(s3Event.Records))

	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.URLDecodedKey

		// check if file extension is .sql
		if len(key) < 4 || key[len(key)-4:] != ".sql" {
			slog.WarnContext(ctx, "skipping file, maybe not a .sql", "bucket", bucket, "key", key)
			continue
		}

		// get the object
		raw, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		})
		if err != nil {
			slog.ErrorContext(ctx, "error while trying to get the object", "bucket", bucket, "key", key, "error", err)
			return err
		}

		// read the object data
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(raw.Body)
		if err != nil {
			slog.ErrorContext(ctx, "error while trying to read the object", "bucket", bucket, "key", key, "error", err)
			return err
		}

		data := buf.String()

		// load env vars
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")

		// connect to the database
		connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		conn, err := sql.Open("postgres", connectionStr)
		if err != nil {
			slog.ErrorContext(ctx, "error while trying to connect to the database", "error", err)
			return err
		}
		defer conn.Close()

		// execute the sql
		_, err = conn.Exec(data)
		if err != nil {
			slog.ErrorContext(ctx, "error while trying to execute the query", "error", err)
		}

		// delete the object
		// _, err = s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		// 	Bucket: &bucket,
		// 	Key:    &key,
		// })
		// if err != nil {
		// 	slog.ErrorContext(ctx, "error while trying to delete the object", "bucket", bucket, "key", key, "error", err)
		// 	return err
		// }
	}

	return nil
}

func main() {
	lambda.Start(handler)
}