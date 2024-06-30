
-- Order Service
CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    method VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    user_id INT NOT NULL,
    cart_id INT NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    payment_method_id INT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    address VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_items (
    id SERIAL PRIMARY KEY,
    transaction_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);



