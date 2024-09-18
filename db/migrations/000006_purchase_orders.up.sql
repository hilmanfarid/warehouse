CREATE TABLE purchase_orders (
    id INT unsigned NOT NULL AUTO_INCREMENT,
    user_id INT unsigned,
    shop_id INT unsigned,
    status INT NOT NULL,
    total_amount DECIMAL(10, 2),
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    process_at datetime DEFAULT NULL,
    success_at datetime DEFAULT NULL,
    failed_at datetime DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (shop_id) REFERENCES shops(id)
);