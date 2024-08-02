CREATE TABLE IF NOT EXISTS customers (
    id varchar(255),
    document_id varchar(255),
    document_type int,
    is_anonymous boolean,
    password varchar(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS customer_deletion_requests (
    id varchar(255),
    customer_id varchar(255),
    name varchar(255),
    address varchar(255),
    phone varchar(255),
    executed boolean,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);