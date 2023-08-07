ALTER TABLE bonus_transaction
    ADD COLUMN transaction_code VARCHAR DEFAULT 'UNKNOWN';

UPDATE bonus_transaction
    SET "transaction_code" = CASE
        WHEN "transaction_bounds" ->> 'reason' LIKE '%<d94a4842>%'
            THEN 'PURCHASE_CHARGE'
        WHEN "transaction_bounds" ->> 'reason' LIKE '%<35bd860f>%'
            THEN 'PURCHASE_INCOME'
        WHEN "transaction_bounds" ->> 'reason' LIKE '%<e3b2bbfd>%'
            THEN 'PARTIAL_REFUND_RECALCULATION_CHARGE'
        WHEN "transaction_bounds" ->> 'reason' LIKE '%<c5835800>%'
            THEN 'PARTIAL_REFUND_RECALCULATION_INCOME'
        WHEN "transaction_bounds" ->> 'reason' LIKE '%<ddfa4e60>%'
            THEN 'SCRIPT_OPERATION'
        WHEN "transaction_bounds" ->> 'reason' LIKE '%Центр-инвест%'
            THEN 'PARTNER_INCOME'
        WHEN "transaction_bounds" ->> 'transactionType' = 'EXPIRATION_CHARGE'
            THEN 'EXPIRATION_CHARGE'
        ELSE 'UNKNOWN'
        END