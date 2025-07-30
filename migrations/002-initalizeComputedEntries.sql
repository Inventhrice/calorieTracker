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