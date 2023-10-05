CREATE DATABASE hestia;

\c hestia;

CREATE TABLE Essen (
	Essen_id SERIAL PRIMARY KEY,
	Name VARCHAR(255) NOT NULL,
	Kalorien INT NOT NULL,
 	Proteine INT NOT NULL
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
	Zeit INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Tag (
	Tag_id SERIAL PRIMARY KEY,
	Name VARCHAR(255) NOT NULL,
	Datum DATE NOT NULL,
	Kilometer INT,
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

INSERT INTO Essen (Name, Kalorien, Proteine)
VALUES	('Steak', 600, 25),
	('Big Mac', 498, 12),
	('Reiswaffel', 90, 1),
	('Toastbrot', 60, 2.9),
	('Haferflocken mit Obst und Joghurt', 296, 30),
	('Fleischvögel mit Kartoffelstock', 598, 7),
	('Pommes Frites gross', 470, 3.4),
	('McFlurry Classic', 500, 4),
	('Crispy Chicken', 530, 17),
	('Coca-Cola 100ml', 175, 0.1),
	('BK KING Fries Cheese & Bacon', 594, 5.2),
    ('BK Hot Brownie', 344, 10.2);

INSERT INTO Kraftsport (Name, Reps, Sets)
VALUES	('Benchpress', 12, 4),
	('Leggpress', 12, 4),
	('Leggcurls', 12, 4);

INSERT INTO Ausdauer (Name, Distanz, Zeit)
VALUES	('running 5KM', 5, 30),
	('running 10KM', 10, 60),
	('row', 1, 5);