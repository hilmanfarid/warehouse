CREATE TABLE purchase_order_details (
    id INT unsigned NOT NULL AUTO_INCREMENT,
    purchase_order_id INT unsigned,
    product_id INT unsigned,
    warehouse_id INT unsigned,
    quantity INT NOT NULL,
    status INT NOT NULL,
    price_per_unit DECIMAL(10, 2),
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    success_at datetime DEFAULT NULL,
    failed_at datetime DEFAULT NULL,
    refunded_at datetime DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);