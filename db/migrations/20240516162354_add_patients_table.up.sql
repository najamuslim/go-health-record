CREATE TABLE patients (
    identity_number NUMERIC(16) UNIQUE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    birth_date VARCHAR(25) NOT NULL,
    gender VARCHAR(6) NOT NULL,
    identity_card_scan_img VARCHAR(2083),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);