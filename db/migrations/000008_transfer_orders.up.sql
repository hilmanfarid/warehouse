CREATE TABLE transfer_orders (
     id INT unsigned NOT NULL AUTO_INCREMENT,
     user_id INT unsigned,
     product_id INT unsigned,
     status INT NOT NULL,
     source_warehouse INT unsigned,
     destination_warehouse INT unsigned,
     quantity INT NOT NULL,
     created_at datetime NOT NULL,
     updated_at datetime NOT NULL,
     success_at datetime DEFAULT NULL,
     failed_at datetime DEFAULT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (user_id) REFERENCES users(id),
     FOREIGN KEY (product_id) REFERENCES products(id),
     FOREIGN KEY (source_warehouse) REFERENCES warehouses(id),
     FOREIGN KEY (destination_warehouse) REFERENCES warehouses(id)

);