CREATE TABLE password_resets
(
    id         SERIAL PRIMARY KEY,
    token      VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP + INTERVAL '24 hours'
);
CREATE INDEX idx_password_resets_email ON password_resets (email);