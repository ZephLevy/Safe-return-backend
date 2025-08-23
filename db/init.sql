-- IMPORTANT: After changing this script, run docker compose down -v to
-- delete volumes to make sure that the next time docker compose up is run,
-- the new version is run again
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    refresh_token TEXT
);

CREATE TABLE emergency_contacts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    phone_number TEXT NOT NULL
);
