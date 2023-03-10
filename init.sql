CREATE DATABASE hestia;

\c hestia;

CREATE TABLE Essen (
  Essen_id SERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Kalorien INT NOT NULL
);

CREATE TABLE Kraftsport (
  Kraftsport_id SERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Reps INT NOT NULL,
  Sets INT NOT NULL
);

CREATE TABLE Ausdauer (
  Ausdauer_id SERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Distanz INT NOT NULL,
  Zeit INTERVAL NOT NULL
);

CREATE TABLE IF NOT EXISTS Tag (
  Tag_id SERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Datum DATE NOT NULL,
  Essen_id INT,
  Kraftsport_id INT,
  Ausdauer_id INT,
  FOREIGN KEY (Essen_id) REFERENCES Essen (Essen_id),
  FOREIGN KEY (Kraftsport_id) REFERENCES Kraftsport (Kraftsport_id),
  FOREIGN KEY (Ausdauer_id) REFERENCES Ausdauer (Ausdauer_id)
);

CREATE TABLE IF NOT EXISTS "Tag-Kraft" (
  Tag_id INT NOT NULL,
  Kraftsport_id INT NOT NULL,
  PRIMARY KEY (Tag_id, Kraftsport_id),
  FOREIGN KEY (Tag_id) REFERENCES Tag(Tag_id),
  FOREIGN KEY (Kraftsport_id) REFERENCES Kraftsport(Kraftsport_id)
);

CREATE TABLE IF NOT EXISTS "Tag-Aus" (
  Tag_id INT NOT NULL,
  Ausdauer_id INT NOT NULL,
  PRIMARY KEY (Tag_id, Ausdauer_id),
  FOREIGN KEY (Tag_id) REFERENCES Tag(Tag_id),
  FOREIGN KEY (Ausdauer_id) REFERENCES Ausdauer(Ausdauer_id)
);

INSERT INTO Essen (Name, Kalorien)
VALUES	('Steak', 600),
	('Salmon', 400),
	('Rice', 200),
	('Broccoli', 50),
	('Big Mac', 498),
	('Reiswaffel', 90),
	('Toastbrot', 60),
	('Pommes Frites gross', 470),
	('McFlurry Classic', 500),
	('Crispy Chicken', 530),
	('Coca-Cola 100ml', 175),
	('BK KING Fries Cheese & Bacon', 593),
       	('BK Hot Brownie', 344);

