CREATE DATABASE Hestia;

USE Hestia;

CREATE TABLE Essen (
  Essen_id INT(11) NOT NULL AUTO_INCREMENT,
  Name VARCHAR(255) NOT NULL,
  Kalorien INT(11) NOT NULL,
  PRIMARY KEY (Essen_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE Kraftsport (
  Kraftsport_id INT(11) NOT NULL AUTO_INCREMENT,
  Name VARCHAR(255) NOT NULL,
  Reps INT(11) NOT NULL,
  Sets INT(11) NOT NULL,
  PRIMARY KEY (Kraftsport_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE Ausdauer (
  Ausdauer_id INT(11) NOT NULL AUTO_INCREMENT,
  Name VARCHAR(255) NOT NULL,
  Distanz INT(11) NOT NULL,
  Zeit VARCHAR(255) NOT NULL,
  PRIMARY KEY (Ausdauer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS Tag (
    Tag_id INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Datum DATE NOT NULL,
    Essen_id INT,
    Kraftsport_id INT,
    Ausdauer_id INT,
    FOREIGN KEY (Essen_id) REFERENCES Essen (Essen_id),
    FOREIGN KEY (Kraftsport_id) REFERENCES Kraftsport (Kraftsport_id),
    FOREIGN KEY (Ausdauer_id) REFERENCES Ausdauer (Ausdauer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `Tag-Kraft` (
    `Tag_id` INT NOT NULL,
    `Kraftsport_id` INT NOT NULL,
    PRIMARY KEY (`Tag_id`, `Kraftsport_id`),
    FOREIGN KEY (`Tag_id`) REFERENCES `Tag`(`Tag_id`),
    FOREIGN KEY (`Kraftsport_id`) REFERENCES `Kraftsport`(`Kraftsport_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `Tag-Aus` (
    `Tag_id` INT NOT NULL,
    `Ausdauer_id` INT NOT NULL,
    PRIMARY KEY (`Tag_id`, `Ausdauer_id`),
    FOREIGN KEY (`Tag_id`) REFERENCES `Tag`(`Tag_id`),
    FOREIGN KEY (`Ausdauer_id`) REFERENCES `Ausdauer`(`Ausdauer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
