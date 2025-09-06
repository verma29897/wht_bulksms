CREATE TABLE whatsapp_accounts (
    id SERIAL PRIMARY KEY,
    business_id VARCHAR(100),
    waba_id VARCHAR(100),
    phone_number_id VARCHAR(100),
    access_token TEXT,
    user_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(200) NOT NULL UNIQUE,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
