CREATE TABLE products (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `barcode` TEXT NOT NULL,
    `title` TEXT NOT NULL,
    `mainCategory` TEXT NOT NULL,
    `subCategory` TEXT NOT NULL,
    `brand` TEXT NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE images (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `barcode` TEXT NOT NULL,
    `width` INT NOT NULL,
    `height` INT NOT NULL,
    `URL` TEXT NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE recipes (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `URL` TEXT NOT NULL,
    `imageURL` TEXT NOT NULL,
    `name` TEXT NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE ingredients (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `name` TEXT NOT NULL,
    `recipeID` INT NOT NULL,
    PRIMARY KEY (ID)
);