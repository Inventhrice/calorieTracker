DROP TABLE IF EXISTS `entries`;
CREATE TABLE `entries` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `dateRecord` date NOT NULL,
  `meal` varchar(30) NOT NULL,
  `foodID` int(11) DEFAULT NULL,
  `foodname` varchar(200) NOT NULL DEFAULT '',
  `grams` decimal(10,2) NOT NULL,
  `cal` decimal(10,2) NOT NULL DEFAULT 0.00,
  `protein` decimal(10,2) NOT NULL DEFAULT 0.00,
  `fat` decimal(10,2) NOT NULL DEFAULT 0.00,
  `carbs` decimal(10,2) NOT NULL DEFAULT 0.00,
  `notes` varchar(3000) NOT NULL DEFAULT '',
  `quantity` int(11) DEFAULT NULL,
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `foodID` (`foodID`),
  CONSTRAINT `entries_ibfk_1` FOREIGN KEY (`foodID`) REFERENCES `food_info` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


DROP TABLE IF EXISTS `food_info`;
CREATE TABLE `food_info` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `calPerG` decimal(10,2) NOT NULL DEFAULT 0.00,
  `proteinPerG` decimal(10,2) NOT NULL DEFAULT 0.00,
  `fatPerG` decimal(10,2) NOT NULL DEFAULT 0.00,
  `carbPerG` decimal(10,2) NOT NULL DEFAULT 0.00,
  `notes` varchar(3000) NOT NULL,
  `source` varchar(3000) NOT NULL,
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `userid` (`userid`),
  CONSTRAINT `food_info_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


DROP TABLE IF EXISTS `goals`;
CREATE TABLE `goals` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `dateRecord` date NOT NULL,
  `goalLbs` decimal(10,2) DEFAULT NULL,
  `multiplier` int(11) DEFAULT NULL,
  `acceptablePercent` decimal(3,2) DEFAULT NULL,
  `goalsPerMeal` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`goalsPerMeal`)),
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


DROP TABLE IF EXISTS `settings`;
CREATE TABLE `settings` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `keyName` varchar(30) DEFAULT NULL,
  `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`value`)),
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` uuid NOT NULL,
  `firstname` varchar(300) DEFAULT NULL,
  `lastname` varchar(300) DEFAULT NULL,
  `pronouns` varchar(50) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DROP TABLE IF EXISTS `weightTrack`;
CREATE TABLE `weightTrack` (
  `dateRecord` date NOT NULL,
  `kg` decimal(10,2) NOT NULL DEFAULT 0.00,
  `lbs` decimal(10,2) GENERATED ALWAYS AS (`kg` * 2.2) VIRTUAL,
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`dateRecord`),
  UNIQUE KEY `weightTrack_unique` (`dateRecord`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

DROP VIEW IF EXISTS `processed_entries`;

CREATE VIEW `processed_entries` AS 
	select `e`.`ID` AS `ID`,`e`.`dateRecord` AS `dateRecord`,`e`.`meal` AS `meal`,
	if(`e`.`foodID` is null,`e`.`foodname`,`fi`.`name`) AS `foodname`,`e`.`foodID` AS `foodID`,`e`.`grams` AS `grams`,
	if(`e`.`foodID` is null,`e`.`cal`,`e`.`grams` * `fi`.`calPerG`) AS `cal`,
	if(`e`.`foodID` is null,`e`.`protein`,`e`.`grams` * `fi`.`proteinPerG`) AS `protein`,
	if(`e`.`foodID` is null,`e`.`fat`,`e`.`grams` * `fi`.`fatPerG`) AS `fat`,
	if(`e`.`foodID` is null,`e`.`carbs`,`e`.`grams` * `fi`.`carbPerG`) AS `carbs`,
	`e`.`notes` AS `notes`,`e`.`userid` AS `userid` from 
	(`entries` `e` left join `food_info` `fi` on(`e`.`foodID` = `fi`.`ID`)) order by `e`.`dateRecord`;
