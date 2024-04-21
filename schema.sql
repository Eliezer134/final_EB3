DROP DATABASE IF EXISTS `odontologia`;

CREATE DATABASE IF NOT EXISTS `odontologia`;

USE `odontologia`;

DROP TABLE IF EXISTS `dentists`;
CREATE TABLE IF NOT EXISTS `odontologia`.`dentists` (
  `id_dentist` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `lastname` VARCHAR(45) NULL,
  `registration_number` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_dentist`, `registration_number`)  
  );

DROP TABLE IF EXISTS `patients`;
CREATE TABLE IF NOT EXISTS `odontologia`.`patients` (
  `id_patient` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `lastname` VARCHAR(45) NULL,
  `address` VARCHAR(100) NULL,
  `dni` VARCHAR(45) NOT NULL,
  `registration_date` DATE NULL,
  PRIMARY KEY (`id_patient`, `dni`)
  );

ALTER TABLE `dentists` ADD UNIQUE (`registration_number`);
ALTER TABLE `patients` ADD UNIQUE (`dni`);

DROP TABLE IF EXISTS `appointments`;
CREATE TABLE IF NOT EXISTS `odontologia`.`appointments` (
  `id_appointment` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(100) NULL,
  `registration_number` varchar(45) NOT NULL,
  `dni` VARCHAR(45) NOT NULL,
  `datetime` DATE NULL,
  PRIMARY KEY (`id_appointment`),
  FOREIGN KEY (`registration_number`) REFERENCES `dentists`(`registration_number`),
  FOREIGN KEY (`dni`) REFERENCES `patients`(`dni`)
  )