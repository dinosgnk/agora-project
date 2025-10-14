-- Create orders schema and tables

-- Create the orders schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS orders;

-- Grant permissions to admin user
GRANT ALL PRIVILEGES ON SCHEMA orders TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA orders TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA orders TO admin;

-- Create orders tables
DROP TABLE IF EXISTS orders.t_ordered_product;
DROP TABLE IF EXISTS orders.t_order;

CREATE TABLE orders.t_order (
    id UUID PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL,
    shipping_address TEXT,
    payment_method VARCHAR(100),
    status VARCHAR(50) NOT NULL,
    total_amount DECIMAL(12,2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders.t_ordered_product (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    code VARCHAR(36) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    quantity INTEGER NOT NULL,
    subtotal DECIMAL(12,2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ordered_product_order 
        FOREIGN KEY (order_id) 
        REFERENCES orders.t_order(id) 
        ON DELETE CASCADE
);
