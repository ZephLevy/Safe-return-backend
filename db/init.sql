-- IMPORTANT: After changing this script, run docker compose down -v to
-- delete volumes to make sure that the next time docker compose up is run,
-- the new version is run again
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    passwordHash VARCHAR(50) NOT NULL
);
