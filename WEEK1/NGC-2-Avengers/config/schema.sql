CREATE DATABASE avenger;

CREATE TABLE Heroes (
  HeroID SERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Universe VARCHAR(255) NOT NULL,
  Skill VARCHAR(255) NOT NULL,
  Imageurl VARCHAR(255) NOT NULL
);

CREATE TABLE Villain (
  VillainID SERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Universe VARCHAR(255) NOT NULL,
  Imageurl VARCHAR(255) NOT NULL
);

INSERT INTO Heroes (Name, Universe, Skill, Imageurl) VALUES
('Captain America', 'Earth 201', 'Powerfull Human', 'https://www.avengers.com/captain-america'),
('Iron Man', 'Earth 201', 'Billionare', 'https://www.avengers.com/iron-man'),
('Thor', 'Asgard', 'God', 'https://www.avengers.com/thor');

INSERT INTO Villain (Name, Universe, Imageurl) VALUES
('Loki', 'Asgard','https://www.avengers.com/lk'),
('Ultron', 'Earth 201', 'https://www.avengers.com/ultron'),
('Thanos', 'Earth 616', 'https://www.avengers.com/thor');