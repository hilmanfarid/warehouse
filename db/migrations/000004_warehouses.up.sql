CREATE TABLE warehouses (
    id INT unsigned NOT NULL AUTO_INCREMENT,
    shop_id INT unsigned,
    name VARCHAR(100) NOT NULL,
    status INT,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (shop_id) REFERENCES shops(id)
);