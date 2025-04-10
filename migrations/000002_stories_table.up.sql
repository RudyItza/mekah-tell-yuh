-- Up migration to create the stories table with a foreign key reference to users
CREATE TABLE IF NOT EXISTS stories (
    id BIGSERIAL PRIMARY KEY,                  -- Auto-incrementing primary key for the story
    title VARCHAR(255) NOT NULL,               -- Title of the story
    content TEXT NOT NULL,                     -- Content of the story
    language CITEXT NOT NULL,                  -- Language of the story (case-insensitive text)
    location VARCHAR(255),                     -- Location associated with the story
    category VARCHAR(255),                     -- Category of the story
    user_id BIGINT NOT NULL,                   -- Foreign key to the users table
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),  -- Timestamp when the story is created
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),  -- Timestamp when the story is last updated
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE  -- Foreign key constraint
);
