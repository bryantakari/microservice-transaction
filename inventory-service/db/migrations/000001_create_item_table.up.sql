
CREATE TABLE items (
    id      VARCHAR(26) PRIMARY KEY,
    item_name VARCHAR(255) NOT NULL,
    item_code VARCHAR(64) NOT NULL,
    item_description TEXT NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    reserved_quantity INT NOT NULL CHECK (reserved_quantity >= 0),
    price BIGINT NOT NULL CHECK (price >= 0),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);