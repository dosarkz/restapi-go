CREATE TABLE users
(
    id                SERIAL PRIMARY KEY,
    name              VARCHAR(255)                                      NOT NULL,
    password          VARCHAR(255)                                       NOT NULL,
    phone             VARCHAR(255) NULL,
    email             VARCHAR(255) UNIQUE                                NOT NULL,
    email_verified_at TIMESTAMP NULL,
    role_id           INTEGER                                            NOT NULL REFERENCES roles (id),
    remember_token    VARCHAR(255) NULL,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at        TIMESTAMP NULL,
    deleted_at        TIMESTAMP NULL,
    telegram_key      VARCHAR(255) UNIQUE NULL,
    status_id         INTEGER                  DEFAULT 0                 NOT NULL
);

CREATE INDEX idx_users_role_id ON users (role_id);
