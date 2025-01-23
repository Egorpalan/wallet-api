CREATE TABLE transactions (
                              id UUID PRIMARY KEY,
                              wallet_id UUID REFERENCES wallets(id),
                              operation_type VARCHAR(10) NOT NULL,
                              amount BIGINT NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);