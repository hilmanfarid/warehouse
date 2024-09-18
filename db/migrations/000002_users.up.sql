CREATE TABLE users (
    id INT unsigned NOT NULL AUTO_INCREMENT,
    email varchar(64) NOT NULL,
    crypted_password varchar(255) NOT NULL,
    secret varchar(255) DEFAULT NULL,
    token varchar(255) DEFAULT NULL,
    status tinyint unsigned DEFAULT NULL,
    role varchar(255) DEFAULT NULL,
    created_at datetime(6) NOT NULL,
    updated_at datetime(6) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY index_users_on_username (email)
);

INSERT INTO users
(
    `email`,
    `crypted_password`,
    `secret`,
    `status`,
    `role`,
    `created_at`,
    `updated_at`
)
VALUES(
  "admin@warehouse.com",
  "276321d9f7553dd4d55d8b224b5731f4da5f57d35a489330587cb1cbe1b59245",
  "d1439fee57e8ef56564048517de217c0",
  2,
  "admin",
   now(),
   now()
)