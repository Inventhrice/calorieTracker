CREATE TABLE IF NOT EXISTS `entryTemplates` (
	ID int NOT NULL AUTO_INCREMENT,
	user_id UUID NOT NULL,
	meal VARCHAR(30),
	food_id int NOT NULL,
	quantity int NOT NULL,
	FOREIGN KEY(food_id) REFERENCES food_info (ID),
	FOREIGN KEY(user_id) REFERENCES users (id),
	PRIMARY KEY(ID)
);

UPDATE metadata SET value="4" WHERE `key`="migrations";
