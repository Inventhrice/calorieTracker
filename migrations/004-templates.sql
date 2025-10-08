CREATE TABLE IF NOT EXISTS `entryTemplates` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `userID` uuid NOT NULL,
  `meal` varchar(30) DEFAULT NULL,
  `foodID` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `foodID` (`foodID`),
  KEY `userID` (`userID`),
  CONSTRAINT `entryTemplates_ibfk_1` FOREIGN KEY (`foodID`) REFERENCES `food_info` (`ID`),
  CONSTRAINT `entryTemplates_ibfk_2` FOREIGN KEY (`userID`) REFERENCES `users` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

UPDATE metadata SET value="4" WHERE `key`="migrations";