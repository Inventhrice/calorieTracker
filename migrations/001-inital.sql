CREATE TABLE IF NOT EXISTS `users` (
  `id` uuid NOT NULL,
  `firstname` varchar(300) DEFAULT "''" NOT NULL,
  `lastname` varchar(300) DEFAULT  "''" NOT NULL,
  `pronouns` varchar(50) DEFAULT   "''" NOT NULL,
  `username` varchar(200) DEFAULT  "''" NOT NULL,
  `password` varchar(200) DEFAULT  "''" NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `food_info` (
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

CREATE TABLE IF NOT EXISTS `entries` (
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

CREATE TABLE IF NOT EXISTS `goals` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `dateRecord` date NOT NULL,
  `goalLbs` decimal(10,2) DEFAULT NULL,
  `multiplier` int(11) DEFAULT NULL,
  `acceptablePercent` decimal(3,2) DEFAULT NULL,
  `goalsPerMeal` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`goalsPerMeal`)),
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

CREATE TABLE IF NOT EXISTS `settings` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `keyName` varchar(30) DEFAULT NULL,
  `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`value`)),
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

CREATE TABLE IF NOT EXISTS `weightTrack` (
  `dateRecord` date NOT NULL,
  `kg` decimal(10,2) NOT NULL DEFAULT 0.00,
  `lbs` decimal(10,2) GENERATED ALWAYS AS (`kg` * 2.2) VIRTUAL,
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`dateRecord`),
  UNIQUE KEY `weightTrack_unique` (`dateRecord`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

CREATE TABLE IF NOT EXISTS `metadata` (
  `key` VARCHAR(200) NOT NULL,
  `value` VARCHAR(2000),
  PRIMARY KEY (`key`)
);

INSERT INTO metadata (`key`, value) VALUES ('migrations', '1');