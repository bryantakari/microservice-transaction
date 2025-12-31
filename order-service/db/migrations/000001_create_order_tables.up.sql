


CREATE TABLE orders (
    id      VARCHAR(26) PRIMARY KEY,
    user_id VARCHAR(64) NOT NULL,

    status VARCHAR(16) NOT NULL
      CHECK (status IN ('CREATED', 'CONFIRMED', 'CANCELLED')),

    total_amount BIGINT NOT NULL CHECK (total_amount >= 0),
    currency CHAR(3) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE order_items (
    id      VARCHAR(26) PRIMARY KEY,
    order_id VARCHAR(26) NOT NULL REFERENCES orders(id) ON DELETE CASCADE,

    product_id VARCHAR(64) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    unit_price BIGINT NOT NULL CHECK (unit_price >= 0),

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);