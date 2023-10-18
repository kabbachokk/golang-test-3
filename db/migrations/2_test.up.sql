#Товары
INSERT INTO `product` VALUES (1, "Ноутбук", NOW());
INSERT INTO `product` VALUES (2, "Телевизор", NOW());
INSERT INTO `product` VALUES (3, "Телефон", NOW());
INSERT INTO `product` VALUES (4, "Системный блок", NOW());
INSERT INTO `product` VALUES (5, "Часы", NOW());
INSERT INTO `product` VALUES (6, "Микрофон", NOW());

#Стеллажи
INSERT INTO `rack` VALUES (1, "А");
INSERT INTO `rack` VALUES (2, "Б");
INSERT INTO `rack` VALUES (3, "В");
INSERT INTO `rack` VALUES (4, "Г");
INSERT INTO `rack` VALUES (5, "Д");
INSERT INTO `rack` VALUES (6, "Е");
INSERT INTO `rack` VALUES (7, "Ж");
INSERT INTO `rack` VALUES (8, "З");

#Товар->Стеллаж
INSERT INTO `product_rack` VALUES (1, 1, true); #Ноутбук 
INSERT INTO `product_rack` VALUES (2, 1, true); #Телевизор 

INSERT INTO `product_rack` VALUES (3, 2, true); #Телефон 
INSERT INTO `product_rack` VALUES (3, 8, NULL); #Телефон (доп)
INSERT INTO `product_rack` VALUES (3, 3, NULL); #Телефон (доп)

INSERT INTO `product_rack` VALUES (4, 7, true); #Системный блок

INSERT INTO `product_rack` VALUES (5, 7, true); #Часы
INSERT INTO `product_rack` VALUES (5, 1, NULL); #Часы (доп)

INSERT INTO `product_rack` VALUES (6, 7, true); #Микрофон

#Заказы
INSERT INTO `order` VALUES (10, NOW());
INSERT INTO `order` VALUES (11, NOW());
INSERT INTO `order` VALUES (12, NOW());
INSERT INTO `order` VALUES (13, NOW());
INSERT INTO `order` VALUES (14, NOW());
INSERT INTO `order` VALUES (15, NOW());

#Заказ->Товар
INSERT INTO `product_order` VALUES (6, 10, 1); #Микрофон
INSERT INTO `product_order` VALUES (3, 10, 1); #Телефон
INSERT INTO `product_order` VALUES (1, 10, 1); #Ноутбук

INSERT INTO `product_order` VALUES (2, 11, 1); #Телевизор 

INSERT INTO `product_order` VALUES (1, 14, 1); #Ноутбук 

INSERT INTO `product_order` VALUES (5, 15, 1); #Часы


