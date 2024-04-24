CREATE DATABASE hestia;

\c hestia;

CREATE TABLE Food (
	Food_id SERIAL PRIMARY KEY,
	Name VARCHAR(255) NOT NULL,
	Calories INT NOT NULL,
 	Proteine INT NOT NULL
);

CREATE TABLE Weight (
	Weight_id SERIAL PRIMARY KEY,
	Name VARCHAR(255) NOT NULL,
	Reps INT NOT NULL,
	Sets INT NOT NULL
);

CREATE TABLE Cardio (
	Cardio_id SERIAL PRIMARY KEY,
	Name VARCHAR(255) NOT NULL,
	Distance INT NOT NULL,
	Duration INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Day (
	Day_id SERIAL PRIMARY KEY,
	Name VARCHAR(255) NOT NULL,
	Date DATE NOT NULL,
	Distance INT,
	Food_id INT,
	Weight_id INT,
	Cardio_id INT,
	FOREIGN KEY (Food_id) REFERENCES Food (Food_id),
	FOREIGN KEY (Weight_id) REFERENCES Weight (Weight_id),
	FOREIGN KEY (Cardio_id) REFERENCES Cardio (Cardio_id)
);

CREATE TABLE IF NOT EXISTS "DayWeight" (
	Day_id INT NOT NULL,
	Weight_id INT NOT NULL,
	PRIMARY KEY (Day_id, Weight_id),
	FOREIGN KEY (Day_id) REFERENCES Day(day_id),
	FOREIGN KEY (Weight_id) REFERENCES Weight(Weight_id)
);

CREATE TABLE IF NOT EXISTS "DayCardio" (
	Day_id INT NOT NULL,
	Cardio_id INT NOT NULL,
	PRIMARY KEY (Day_id, Cardio_id),
	FOREIGN KEY (Day_id) REFERENCES Day(Day_id),
	FOREIGN KEY (Cardio_id) REFERENCES Cardio(Cardio_id)
);

CREATE TABLE IF NOT EXISTS "DayFood" (
	Day_id INT NOT NULL,
	Food_id INT NOT NULL,
	PRIMARY KEY (Day_id, Food_id),
	FOREIGN KEY (Day_id) REFERENCES Day(day_id),
	FOREIGN KEY (Food_id) REFERENCES Food(Food_id)
);

CREATE TABLE IF NOT EXISTS ToDo (
	ToDo_id SERIAL PRIMARY KEY,
	TaskName VARCHAR(255) NOT NULL,
	Deadline DATE NOT NULL,
	Urgency VARCHAR(255) NOT NULL,
	Description VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS ToDo_Work (
	ToDo_Work_id SERIAL PRIMARY KEY,
	TaskName VARCHAR(255) NOT NULL,
	Deadline DATE NOT NULL,
	Urgency VARCHAR(255) NOT NULL,
	Description VARCHAR(255) NOT NULL
);

INSERT INTO Food (Name, Calories, Proteine)
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

INSERT INTO Weight (Name, Reps, Sets)
VALUES	('Benchpress', 12, 4),
	('Leggpress', 12, 4),
	('Leggcurls', 12, 4);

INSERT INTO Cardio (Name, Distance, Duration)
VALUES	('running 5KM', 5, 30),
	('running 10KM', 10, 60),
	('row', 1, 5);