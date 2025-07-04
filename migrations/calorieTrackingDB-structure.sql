-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: 192.168.0.11    Database: calorieTracking
-- ------------------------------------------------------
-- Server version	11.4.5-MariaDB-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `entries`
--

DROP TABLE IF EXISTS `entries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=2779 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `food_info`
--

DROP TABLE IF EXISTS `food_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `goals`
--

DROP TABLE IF EXISTS `goals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goals` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `dateRecord` date NOT NULL,
  `goalLbs` decimal(10,2) DEFAULT NULL,
  `multiplier` int(11) DEFAULT NULL,
  `acceptablePercent` decimal(3,2) DEFAULT NULL,
  `goalsPerMeal` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`goalsPerMeal`)),
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Temporary view structure for view `processed_entries`
--

DROP TABLE IF EXISTS `processed_entries`;
/*!50001 DROP VIEW IF EXISTS `processed_entries`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `processed_entries` AS SELECT 
 1 AS `ID`,
 1 AS `dateRecord`,
 1 AS `meal`,
 1 AS `foodname`,
 1 AS `foodID`,
 1 AS `grams`,
 1 AS `cal`,
 1 AS `protein`,
 1 AS `fat`,
 1 AS `carbs`,
 1 AS `notes`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `settings`
--

DROP TABLE IF EXISTS `settings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `settings` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `keyName` varchar(30) DEFAULT NULL,
  `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`value`)),
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` uuid NOT NULL,
  `firstname` varchar(300) DEFAULT NULL,
  `lastname` varchar(300) DEFAULT NULL,
  `pronouns` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `weightTrack`
--

DROP TABLE IF EXISTS `weightTrack`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `weightTrack` (
  `dateRecord` date NOT NULL,
  `kg` decimal(10,2) NOT NULL DEFAULT 0.00,
  `lbs` decimal(10,2) GENERATED ALWAYS AS (`kg` * 2.2) VIRTUAL,
  `userid` uuid DEFAULT NULL,
  PRIMARY KEY (`dateRecord`),
  UNIQUE KEY `weightTrack_unique` (`dateRecord`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping routines for database 'calorieTracking'
--

--
-- Final view structure for view `processed_entries`
--

/*!50001 DROP VIEW IF EXISTS `processed_entries`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `processed_entries` AS select `e`.`ID` AS `ID`,`e`.`dateRecord` AS `dateRecord`,`e`.`meal` AS `meal`,if(`e`.`foodID` is null,`e`.`foodname`,`fi`.`name`) AS `foodname`,`e`.`foodID` AS `foodID`,`e`.`grams` AS `grams`,if(`e`.`foodID` is null,`e`.`cal`,`e`.`grams` * `fi`.`calPerG`) AS `cal`,if(`e`.`foodID` is null,`e`.`protein`,`e`.`grams` * `fi`.`proteinPerG`) AS `protein`,if(`e`.`foodID` is null,`e`.`fat`,`e`.`grams` * `fi`.`fatPerG`) AS `fat`,if(`e`.`foodID` is null,`e`.`carbs`,`e`.`grams` * `fi`.`carbPerG`) AS `carbs`,`e`.`notes` AS `notes` from (`entries` `e` left join `food_info` `fi` on(`e`.`foodID` = `fi`.`ID`)) order by `e`.`dateRecord` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-04 14:48:13
