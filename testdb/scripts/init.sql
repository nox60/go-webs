
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
	`function_id` INT NOT NULL,
	`name` VARCHAR(200) NULL,
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