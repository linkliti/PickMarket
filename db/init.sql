-- PostgreSQL script converted from MySQL

-- -----------------------------------------------------
-- Initialize database
-- -----------------------------------------------------
\set ON_ERROR_STOP on
CREATE USER pmworker WITH PASSWORD 'securepass';
CREATE DATABASE pmdb WITH OWNER pmworker;
\c pmdb

-- -----------------------------------------------------
-- Table pmdb.Marketplaces
-- -----------------------------------------------------
CREATE TABLE Marketplaces (
  marketName VARCHAR(45) PRIMARY KEY,
  marketURL VARCHAR(128) UNIQUE NOT NULL,
  marketParseDate TIMESTAMP NULL
);
ALTER TABLE Marketplaces OWNER TO pmworker;
-- -----------------------------------------------------
-- Table pmdb.Categories
-- -----------------------------------------------------
CREATE TABLE Categories (
  categoryURL VARCHAR(128) PRIMARY KEY,
  Marketplaces_marketName VARCHAR(45) NOT NULL,
  Categories_parentURL VARCHAR(128) NULL,
  categoryName VARCHAR(128) NOT NULL,
  categoryFilters JSONB NULL,
  categoryParseDate TIMESTAMP NULL,
  FOREIGN KEY (Marketplaces_marketName) REFERENCES Marketplaces (marketName) ON DELETE NO ACTION ON UPDATE NO ACTION,
  FOREIGN KEY (Categories_parentURL) REFERENCES Categories (categoryURL) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE Categories OWNER TO pmworker;
-- -----------------------------------------------------
-- Table pmdb.Items
-- -----------------------------------------------------
CREATE TABLE Items (
  itemURL VARCHAR(255) PRIMARY KEY,
  Marketplaces_marketName VARCHAR(45) NOT NULL,
  itemChars JSONB NOT NULL,
  itemParseDate TIMESTAMP NOT NULL,
  FOREIGN KEY (Marketplaces_marketName) REFERENCES Marketplaces (marketName) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE Items OWNER TO pmworker;

-- -----------------------------------------------------
-- PreInsert data
-- -----------------------------------------------------
INSERT INTO Marketplaces (marketName, marketURL, marketParseDate) VALUES ('OZON', 'https://www.ozon.ru', NULL);
