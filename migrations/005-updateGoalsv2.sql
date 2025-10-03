ALTER TABLE goals MODIFY COLUMN `dateRecord` date NOT NULL DEFAULT NOW();
ALTER TABLE goals MODIFY COLUMN `goalLbs` decimal(10,2) NOT NULL DEFAULT 100;
ALTER TABLE goals MODIFY COLUMN `multiplier` int(11) NOT NULL DEFAULT 10;
ALTER TABLE goals MODIFY COLUMN `acceptablePercent` decimal(3,2) NOT NULL DEFAULT 0.1;
ALTER TABLE calorieTracking.users CHANGE id ID uuid NOT NULL;
UPDATE metadata SET value="5" WHERE `key`="migrations";