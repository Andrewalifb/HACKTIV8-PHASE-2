CREATE DATABASE avenger_criminal_report;

CREATE TABLE heroes (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  superpower VARCHAR(255) NOT NULL
);

CREATE TABLE villains (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  nemesis_id INTEGER REFERENCES heroes(id)
);

CREATE TABLE criminal_reports (
  id SERIAL PRIMARY KEY,
  hero_id INTEGER REFERENCES heroes(id),
  villain_id INTEGER REFERENCES villains(id),
  description TEXT,
  time_of_incident TIMESTAMP
);


INSERT INTO heroes (name, superpower)
VALUES ('Iron Man', 'Genius, Power Suit'),
       ('Captain America', 'Super Soldier Serum, Peak Human'),
       ('Thor', 'God of Thunder, Superhuman Strength'),
       ('Black Widow', 'Expert Martial Artist, Spy'),
       ('Hulk', 'Superhuman Strength, Regeneration');

INSERT INTO villains (name, nemesis_id)
VALUES ('Loki', 1),  
       ('Red Skull', 2), 
       ('Thanos', 5),  
       ('Black Widow (Imposter)', 4), 
       ('Abomination', 5);  

INSERT INTO criminal_reports (hero_id, villain_id, description, time_of_incident)
VALUES (1, 1, 'Loki attempts to steal the Tesseract from S.H.I.E.L.D.', '2012-04-15 08:00:00'),
       (2, 2, 'Red Skull leads HYDRA forces against the Allied Forces', '1945-05-08 12:00:00'),
       (3, 3, 'Thanos attacks Earth to collect the Infinity Stones', '2018-10-14 16:00:00'),
       (4, 5, 'Black Widow imposter infiltrates the Avengers', '2024-02-20 10:00:00'),
       (5, 5, 'Hulk battles Abomination in Harlem', '2008-06-13 18:00:00');

