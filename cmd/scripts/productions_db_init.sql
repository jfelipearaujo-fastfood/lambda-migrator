DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;

CREATE TABLE IF NOT EXISTS orders (
    order_id varchar(255),
    state INT,
    state_updated_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (order_id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id varchar(255),
    order_id varchar(255),
    name varchar(255),
    quantity int,
    PRIMARY KEY (id, order_id)
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
);