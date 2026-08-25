CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_code VARCHAR(7) NOT NULL UNIQUE,
    destination_url TEXT NOT NULL,
    click_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ 
);
