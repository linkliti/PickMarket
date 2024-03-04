-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema pmdb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema pmdb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `pmdb` DEFAULT CHARACTER SET utf8 ;
USE `pmdb` ;

-- -----------------------------------------------------
-- Table `pmdb`.`Marketplaces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `pmdb`.`Marketplaces` (
  `marketName` VARCHAR(45) NOT NULL,
  `marketURL` VARCHAR(128) NOT NULL,
  `marketParseDate` DATETIME NOT NULL,
  PRIMARY KEY (`marketName`),
  UNIQUE INDEX `marketURL_UNIQUE` (`marketURL` ASC) VISIBLE)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `pmdb`.`Categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `pmdb`.`Categories` (
  `categoryURL` VARCHAR(128) NOT NULL,
  `Marketplaces_marketName` VARCHAR(45) NOT NULL,
  `Categories_parentURL` VARCHAR(128) NOT NULL,
  `categoryName` VARCHAR(128) NOT NULL,
  `categoryFilters` JSON NULL,
  `categoryParseDate` DATETIME NOT NULL,
  PRIMARY KEY (`categoryURL`),
  INDEX `fk_categories_marketplaces_idx` (`Marketplaces_marketName` ASC) VISIBLE,
  INDEX `fk_Categories_Categories1_idx` (`Categories_parentURL` ASC) VISIBLE,
  UNIQUE INDEX `categoryURL_UNIQUE` (`categoryURL` ASC) VISIBLE,
  CONSTRAINT `fk_categories_marketplaces`
    FOREIGN KEY (`Marketplaces_marketName`)
    REFERENCES `pmdb`.`Marketplaces` (`marketName`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Categories_Categories1`
    FOREIGN KEY (`Categories_parentURL`)
    REFERENCES `pmdb`.`Categories` (`categoryURL`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `pmdb`.`Items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `pmdb`.`Items` (
  `itemURL` VARCHAR(255) NOT NULL,
  `Marketplaces_marketName` VARCHAR(45) NOT NULL,
  `itemName` VARCHAR(255) NOT NULL,
  `itemChars` JSON NOT NULL,
  `itemParseDate` DATETIME NOT NULL,
  INDEX `fk_ItemData_Marketplaces1_idx` (`Marketplaces_marketName` ASC) VISIBLE,
  UNIQUE INDEX `itemURL_UNIQUE` (`itemURL` ASC) VISIBLE,
  PRIMARY KEY (`itemURL`),
  CONSTRAINT `fk_ItemData_Marketplaces1`
    FOREIGN KEY (`Marketplaces_marketName`)
    REFERENCES `pmdb`.`Marketplaces` (`marketName`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
