/* enterprise edition */
DROP TABLE IF EXISTS `feedback`;

CREATE TABLE IF NOT EXISTS `feedback` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`refid` CHAR(16) NOT NULL COLLATE utf8_bin,
	`orgid` CHAR(16) NOT NULL COLLATE utf8_bin,
	`documentid` CHAR(16) NOT NULL COLLATE utf8_bin,
	`userid` CHAR(16) DEFAULT '' COLLATE utf8_bin,
	`email` VARCHAR(250) NOT NULL DEFAULT '',
	`feedback` LONGTEXT,
	`created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_id PRIMARY KEY (id))
DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin
ENGINE =  InnoDB;
