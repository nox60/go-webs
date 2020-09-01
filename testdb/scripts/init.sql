
DROP TABLE IF EXISTS tb_functions;
CREATE TABLE tb_functions (
	`function_id` INT NOT NULL AUTO_INCREMENT,
	`number` INT NOT NULL,
	`order` int NOT NULL,
	`name` VARCHAR(200) NULL,
	`path` VARCHAR(200) NULL,
	`parent_function_id` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`function_id`)
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