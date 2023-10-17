CREATE TABLE `product` (
	`id` int NOT NULL,
	`name` varchar(200) NOT NULL,
	`created_at` TIMESTAMP NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `order` (
	`id` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `product_order` (
	`order_id` int NOT NULL,
	`product_id` int NOT NULL,
	`qty` int NOT NULL
);

CREATE TABLE `product_rack` (
	`rack_id` int NOT NULL,
	`product_id` int NOT NULL,
	`name` varchar(200) NOT NULL
);

CREATE TABLE `rack` (
	`id` int NOT NULL,
	`name` varchar(200) NOT NULL,
	PRIMARY KEY (`id`)
);

ALTER TABLE `product_order` ADD CONSTRAINT `product_order_fk0` FOREIGN KEY (`order_id`) REFERENCES `order`(`id`);

ALTER TABLE `product_order` ADD CONSTRAINT `product_order_fk1` FOREIGN KEY (`product_id`) REFERENCES `product`(`id`);

ALTER TABLE `product_rack` ADD CONSTRAINT `product_rack_fk0` FOREIGN KEY (`rack_id`) REFERENCES `rack`(`id`);

ALTER TABLE `product_rack` ADD CONSTRAINT `product_rack_fk1` FOREIGN KEY (`product_id`) REFERENCES `product`(`id`);






