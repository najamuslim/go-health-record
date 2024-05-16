CREATE TABLE users (
    user_id VARCHAR(255) PRIMARY KEY,
    nip NUMERIC(13) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    identity_card_scan_img VARCHAR(2083),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);