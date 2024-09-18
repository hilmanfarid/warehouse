CREATE TABLE inventory (
   id INT unsigned NOT NULL AUTO_INCREMENT,
   product_id INT unsigned,
   warehouse_id INT unsigned,
   quantity INT NOT NULL,
   status INT,
   created_at datetime NOT NULL,
   updated_at datetime NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (product_id) REFERENCES products(id),
   FOREIGN KEY (warehouse_id) REFERENCES warehouses(id),
   UNIQUE KEY `idx_inventory_product_id_warehouse_id` (`product_id`,`warehouse_id`)
);