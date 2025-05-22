CREATE TYPE order_status AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL CHECK (user_id > 0),
    total DECIMAL(10,2) NOT NULL,
    status order_status NOT NULL DEFAULT 'pending',
    address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
