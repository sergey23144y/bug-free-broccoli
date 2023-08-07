CREATE TABLE IF NOT EXISTS loyalty_program
(
    legality_documentation_id VARCHAR DEFAULT NULL,
    legality_documentation_link VARCHAR DEFAULT NULL,
    loyalty_currency_name JSONB DEFAULT NULL,
    rule_ids VARCHAR[] NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE loyalty_rule
(
    id VARCHAR PRIMARY KEY,
    title JSONB DEFAULT NULL,
    priority NUMERIC DEFAULT 0,
    description JSONB DEFAULT NULL,
    descriptor_id VARCHAR NOT NULL,
    segment_ids VARCHAR[] DEFAULT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

ALTER TABLE loyalty_descriptor DROP CONSTRAINT loyalty_descriptor_pkey;
ALTER TABLE loyalty_descriptor
    ADD COLUMN id VARCHAR PRIMARY KEY DEFAULT 'will-be-updated',
    ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now(),
    ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;