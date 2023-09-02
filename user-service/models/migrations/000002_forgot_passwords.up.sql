CREATE TABLE IF NOT EXISTS forget_passwords (
    id VARCHAR(100) PRIMARY KEY,
    otp VARCHAR(10) NOT NULL,
    expired_otp BIGINT,
    user_id VARCHAR(100) REFERENCES users(id),
    created_at BIGINT,
    updated_at BIGINT
);
