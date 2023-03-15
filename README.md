# HESTIA
Was ist das Projekt Hestia:
	Ich möchte gerne eine eigene Applikation schreiben, um meine Gesundheitsdaten im Überblick zu behalten. 

### Wie wird mein Programm geschrieben:
	Es wird eine Webapplikation, welche in zwei verschiedene Container gebaut wird. 

### Backend:
	Mein Backend enthält eine Postgres Datenbank. Mehr zum Backend in instructions.md

### Frontend:
	Frontend wird bei mir ein node.js projekt

# Frontend
	Mein Frontend wird mit PUG und Node.js geschrieben. Sobald dies fertig ist, werde ich es versuchen in einen Container zu verpacken. 

# Backend
	Das Backen ist eine Postgres Datenbank. Das Dockerfile dazu befindet sich im Backendordner. Zusammen mit dem Dockerimage wird gerade noch die Datenbank im Image erstellt. Das SQL Script welches dazu verwendet wird ist das [[init.sql]]. Genauere anweisungen zum Backend sind im [[instructions.md]] zu finden. 

# Aktueller stand
- [+] Daten über Internetseite zur Tag-Tabelle hinzufügen. 
- [] Daten über Internetseite von Tag-Tabelle entfernen.
- [] Daten im [[init.sql]] für Essen Tabelle hinzufügen
- [] Daten im [[init.sql]] für Kraftsport Tabelle hinzufügen
- [] Daten im [[init.sql]] für Ausdauer Tabelle hinzufügen