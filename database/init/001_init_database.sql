-- Initialize database and schema
-- Create schema for products
CREATE SCHEMA IF NOT EXISTS products;

-- Grant permissions to admin user
GRANT ALL PRIVILEGES ON SCHEMA products TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA products TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA products TO admin;
