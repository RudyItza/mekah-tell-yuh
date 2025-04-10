-- Up migration to create the users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,               -- Auto-incrementing primary key for the user
    username VARCHAR(255) NOT NULL UNIQUE,   -- Unique username for the user
    email VARCHAR(255) NOT NULL UNIQUE,      -- Unique email for the user
    password_hash TEXT NOT NULL,            -- Hashed password
    role VARCHAR(50) DEFAULT 'user',        -- User role (e.g., 'user', 'admin')
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),  -- Timestamp when the user is created
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()   -- Timestamp when the user is last updated
);