-- Example schema file
-- Replace or modify these tables according to your needs

-- Example: Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Example: Add your custom tables below
-- CREATE TABLE IF NOT EXISTS your_table (
--     id SERIAL PRIMARY KEY,
--     field1 VARCHAR(255),
--     field2 INTEGER,
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
-- );

-- Example: Add your indexes below
-- CREATE INDEX IF NOT EXISTS idx_users_username ON users(username); 