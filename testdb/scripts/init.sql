DROP TABLE IF EXISTS tb_users;
CREATE TABLE tb_users (
	`account_id` INT NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(50) NOT NULL,
    `password` VARCHAR(50) NOT NULL,
    `real_name` VARCHAR(50) NOT NULL DEFAULT '',
    `phone_number` VARCHAR(50) NULL DEFAULT '',
    `address` VARCHAR(200) NULL DEFAULT '',
    `major` VARCHAR(200) NULL DEFAULT '',
    `gender` int not null default 1,
    `age` int not null default 1,
    `user_type` int not null default 1
    PRIMARY KEY(`account_id`)
)DEFAULT CHARSET=utf8;

INSERT INTO tb_users (account_id, user_name, password, user_type) VALUES (1,'admin','111111', 0);

DROP TABLE IF EXISTS tb_users_roles;
CREATE TABLE tb_users_roles (
	`row_id` INT NOT NULL AUTO_INCREMENT,
    `account_id` INT NOT NULL,
    `role_id` INT NOT NULL
    PRIMARY KEY(`rowId`)
)DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_functions;
CREATE TABLE tb_functions (
	`function_id` INT NOT NULL AUTO_INCREMENT,
	`number` INT NOT NULL,
	`order` int NOT NULL,
	`name` VARCHAR(200) NULL,
	`path` VARCHAR(200) NULL,
	`type` TINYINT NOT NULL DEFAULT 0,
	`parent_function_id` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`function_id`)
)DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_functions_items;
CREATE TABLE tb_functions_items (
	`function_item_id` INT NOT NULL AUTO_INCREMENT,
	`function_id` INT NOT NULL DEFAULT 0,
	`item_number` INT NOT NULL DEFAULT 0,
	`item_name` VARCHAR(200) NULL,
  PRIMARY KEY (`function_item_id`)
)DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_roles;
CREATE TABLE tb_roles (
	`role_id` INT NOT NULL AUTO_INCREMENT,
	`code` VARCHAR(10) NOT NULL,
	`name` VARCHAR(200) NULL,
	`status` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`role_id`)
)DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_roles_functions;
CREATE TABLE tb_roles_functions (
    `row_id` INT NOT NULL AUTO_INCREMENT,
    `role_id` INT NOT NULL,
    `function_id` INT NOT NULL,
    `status` TINYINT NOT NULL DEFAULT 0,
    PRIMARY KEY (`row_id`)
)DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_roles_items;
CREATE TABLE tb_roles_items (
    `row_id` INT NOT NULL AUTO_INCREMENT,
    `role_id` INT NOT NULL,
    `item_id` INT NOT NULL,
    `status` TINYINT NOT NULL DEFAULT 0,
    PRIMARY KEY (`row_id`)
)DEFAULT CHARSET=utf8;

#-----------

DROP TABLE IF EXISTS tb_roles_items;
CREATE TABLE tb_roles_items (
    `row_id` INT NOT NULL AUTO_INCREMENT,
    `role_id` INT NOT NULL,
    `item_id` INT NOT NULL,
    `status` TINYINT NOT NULL DEFAULT 0,
    PRIMARY KEY (`row_id`)
)DEFAULT CHARSET=utf8;

