CREATE TABLE shops (
   id INT unsigned NOT NULL AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   status INT,
   created_at datetime NOT NULL,
   updated_at datetime NOT NULL,
   PRIMARY KEY (id),
   UNIQUE KEY index_shops_on_name (name)
);