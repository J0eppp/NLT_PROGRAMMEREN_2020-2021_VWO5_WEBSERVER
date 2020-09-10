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