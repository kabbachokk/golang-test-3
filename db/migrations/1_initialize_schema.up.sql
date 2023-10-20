CREATE TABLE `product` (
	`id` int NOT NULL,
	`name` varchar(200) NOT NULL,
	`created_at` TIMESTAMP NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `order` (
	`id` int NOT NULL,
	`created_at` TIMESTAMP NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `product_order` (
	`product_id` int NOT NULL,
	`order_id` int NOT NULL,
	`qty` int NOT NULL,
	PRIMARY KEY (`order_id`, `product_id`)
);

CREATE TABLE `rack` (
	`id` int NOT NULL,
	`name` varchar(200) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `product_rack` (
	`product_id` int NOT NULL,
	`rack_id` int NOT NULL,
	`is_primary` boolean DEFAULT true
);

CREATE TABLE `product_rack_names` (
	`product_id` int NOT NULL,
	`p_name` varchar(200) NOT NULL,
	`s_name` varchar(200) NULL
);

ALTER TABLE `product_order` ADD CONSTRAINT `product_order_fk0` FOREIGN KEY (`order_id`) REFERENCES `order`(`id`);

ALTER TABLE `product_order` ADD CONSTRAINT `product_order_fk1` FOREIGN KEY (`product_id`) REFERENCES `product`(`id`);

ALTER TABLE `product_rack` ADD CONSTRAINT `product_rack_fk0` FOREIGN KEY (`rack_id`) REFERENCES `rack`(`id`);

ALTER TABLE `product_rack` ADD CONSTRAINT `product_rack_fk1` FOREIGN KEY (`product_id`) REFERENCES `product`(`id`);

ALTER TABLE `product_rack` ADD UNIQUE `unique_index` (`product_id`, `is_primary`);

ALTER TABLE `product_rack_names` ADD CONSTRAINT `product_rack_names_fk0` FOREIGN KEY (`product_id`) REFERENCES `product`(`id`);

# ?multiStatements=true
CREATE TRIGGER after_product_rack_insert
AFTER INSERT
ON `product_rack` FOR EACH ROW
BEGIN
  	IF NEW.is_primary IS true THEN
    	INSERT INTO `product_rack_names` SET
    	product_id = NEW.product_id,
    	p_name = (SELECT name FROM rack WHERE id = NEW.rack_id LIMIT 1),
		s_name = NULL;
	ELSE
		UPDATE `product_rack_names` SET
		s_name = CONCAT_WS(',', s_name, (SELECT name FROM rack WHERE id = NEW.rack_id LIMIT 1))
		WHERE product_id = NEW.product_id;
	END IF;
END;

