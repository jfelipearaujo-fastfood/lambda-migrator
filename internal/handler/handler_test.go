package handler

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	testQuery           = "SELECT 1"
	testInfoTablesQuery = "SELECT (.+) FROM (.+)?information_schema.tables(.+)?"
)

func TestHandle(t *testing.T) {
	t.Run("Should run the migrations", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		ordersDbService := mocks.NewMockDatabaseSQLService(t)
		ordersDbService.On("GetDbName").Return("orders").Once()
		ordersDbService.On("GetInstance").Return(db).Once()

		paymentsDbService := mocks.NewMockDatabaseSQLService(t)
		paymentsDbService.On("GetDbName").Return("payments").Once()
		paymentsDbService.On("GetInstance").Return(db).Once()

		productionsDbService := mocks.NewMockDatabaseSQLService(t)
		productionsDbService.On("GetDbName").Return("productions").Once()
		productionsDbService.On("GetInstance").Return(db).Once()

		mock.ExpectExec(testQuery).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(testInfoTablesQuery).
			WillReturnRows(sqlmock.NewRows([]string{"table_name"}).AddRow("orders"))

		mock.ExpectExec(testQuery).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(testInfoTablesQuery).
			WillReturnRows(sqlmock.NewRows([]string{"table_name"}).AddRow("payments"))

		mock.ExpectExec(testQuery).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(testInfoTablesQuery).
			WillReturnRows(sqlmock.NewRows([]string{"table_name"}).AddRow("productions"))

		handler := NewHandler(
			ordersDbService,
			paymentsDbService,
			productionsDbService,
			testQuery,
			testQuery,
			testQuery,
		)

		// Act
		err = handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		ordersDbService.AssertExpectations(t)
		paymentsDbService.AssertExpectations(t)
		productionsDbService.AssertExpectations(t)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Should log error if a migration error occurs", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		ordersDbService := mocks.NewMockDatabaseSQLService(t)
		ordersDbService.On("GetDbName").Return("orders").Once()
		ordersDbService.On("GetInstance").Return(db).Once()

		paymentsDbService := mocks.NewMockDatabaseSQLService(t)
		paymentsDbService.On("GetDbName").Return("payments").Once()
		paymentsDbService.On("GetInstance").Return(db).Once()

		productionsDbService := mocks.NewMockDatabaseSQLService(t)
		productionsDbService.On("GetDbName").Return("productions").Once()
		productionsDbService.On("GetInstance").Return(db).Once()

		mock.ExpectExec(testQuery).WillReturnError(assert.AnError)
		mock.ExpectExec(testQuery).WillReturnError(assert.AnError)
		mock.ExpectExec(testQuery).WillReturnError(assert.AnError)

		handler := NewHandler(
			ordersDbService,
			paymentsDbService,
			productionsDbService,
			testQuery,
			testQuery,
			testQuery,
		)

		// Act
		err = handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		ordersDbService.AssertExpectations(t)
		paymentsDbService.AssertExpectations(t)
		productionsDbService.AssertExpectations(t)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
