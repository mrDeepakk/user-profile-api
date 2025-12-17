
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

CREATE INDEX idx_users_name ON users(name);

COMMENT ON TABLE users IS 'Stores user information with name and date of birth';
COMMENT ON COLUMN users.id IS 'Auto-incrementing primary key';
COMMENT ON COLUMN users.name IS 'User full name';
COMMENT ON COLUMN users.dob IS 'User date of birth (used to calculate age in application)';
