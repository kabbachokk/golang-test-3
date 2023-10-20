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

ALTER TABLE `product_order` ADD CONSTRAINT `product_order_fk0` FOREIGN KEY (`order_id`) REFERENCES `order`(`id`);

ALTER TABLE `product_order` ADD CONSTRAINT `product_order_fk1` FOREIGN KEY (`product_id`) REFERENCES `product`(`id`);

ALTER TABLE `product_rack` ADD CONSTRAINT `product_rack_fk0` FOREIGN KEY (`rack_id`) REFERENCES `rack`(`id`);

ALTER TABLE `product_rack` ADD CONSTRAINT `product_rack_fk1` FOREIGN KEY (`product_id`) REFERENCES `product`(`id`);

ALTER TABLE `product_rack` ADD UNIQUE `unique_index` (`product_id`, `is_primary`);

