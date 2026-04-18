CREATE TABLE IF NOT EXISTS harga (
    id SERIAL PRIMARY KEY,
    tanggal DATE NOT NULL,
    harga NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS forecasts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    alpha NUMERIC NOT NULL,
    beta NUMERIC NOT NULL,
    gamma NUMERIC NOT NULL,
    season_length INTEGER NOT NULL,
    periods INTEGER NOT NULL,
    result JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tes_perhitungan (
    id SERIAL PRIMARY KEY,
    forecast_id INTEGER REFERENCES forecasts(id) ON DELETE CASCADE,
    periode INTEGER NOT NULL,
    tanggal DATE,
    harga_asli NUMERIC,
    level NUMERIC,
    trend NUMERIC,
    seasonal NUMERIC,
    forecast NUMERIC
);