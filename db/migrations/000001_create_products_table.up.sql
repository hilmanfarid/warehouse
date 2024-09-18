CREATE TABLE products (
    id int unsigned NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    code varchar(100) NOT NULL,
    status INT,
    price DECIMAL(10,2),
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `idx_products_name` (`name`),
    UNIQUE KEY `idx_products_code` (`code`)
)