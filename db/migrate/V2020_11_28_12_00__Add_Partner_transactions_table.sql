CREATE TABLE IF NOT EXISTS partner_transactions
(
    id VARCHAR PRIMARY KEY,
    user_hash VARCHAR NOT NULL,
    transaction_value DECIMAL NOT NULL,
    transaction_bounds JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP,
    attempt_count INT NOT NULL,
    is_succeed BOOLEAN NOT NULL
);