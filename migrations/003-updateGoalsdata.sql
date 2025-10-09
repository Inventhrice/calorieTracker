ALTER TABLE goals ADD proteinGPerLBS DECIMAL(3,2);
ALTER TABLE goals ADD fatGPerLBS DECIMAL(3,2);

UPDATE metadata SET value="3" WHERE `key`="migrations";