CREATE TABLE IF NOT EXISTS bonus_transaction
(
    id VARCHAR PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    is_canceled BOOLEAN NOT NULL,
    transaction_value DECIMAL NOT NULL,
    transaction_bounds JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS loyalty_descriptor
(
    active_period_days NUMERIC,
    price_coverage_percentage DECIMAL(5,2),
    accrual_percentage DECIMAL(5,2),
    max_expenditure_per_transaction_percentage DECIMAL(5,2),
    applicable_with_discount BOOLEAN,
    legality_documentation_id VARCHAR,
    updated_at TIMESTAMP NOT NULL,
    is_active BOOLEAN PRIMARY KEY
);

CREATE INDEX by_transaction_value ON bonus_transaction(transaction_value);
CREATE INDEX by_user_id ON bonus_transaction(user_id);

INSERT INTO loyalty_descriptor(
    active_period_days,
    price_coverage_percentage,
    accrual_percentage,
    max_expenditure_per_transaction_percentage,
    applicable_with_discount,
    legality_documentation_id,
    updated_at,
    is_active
) VALUES (
    null,
    100,
    0,
    100,
    true,
    null,
    now(),
    false
);