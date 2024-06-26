CREATE TABLE medical_records (
    id SERIAL PRIMARY KEY,
    identity_number NUMERIC(16) NOT NULL,
    symptoms VARCHAR(2000) NOT NULL,
    medications VARCHAR(2000) NOT NULL,
    created_by VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);