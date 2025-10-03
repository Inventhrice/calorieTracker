CREATE TABLE IF NOT EXISTS `entryTemplates` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` uuid NOT NULL,
  `meal` varchar(30) DEFAULT NULL,
  `food_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `food_id` (`food_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `entryTemplates_ibfk_1` FOREIGN KEY (`food_id`) REFERENCES `food_info` (`ID`),
  CONSTRAINT `entryTemplates_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

UPDATE metadata SET value="4" WHERE `key`="migrations";