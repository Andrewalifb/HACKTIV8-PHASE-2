-- Create Avenger Inventory Database
CREATE DATABASE avenger_inventory;

-- Create Items Table
CREATE TABLE items (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  item_code VARCHAR(255) NOT NULL UNIQUE,
  stock INTEGER NOT NULL,
  description TEXT,
  category_id INTEGER REFERENCES categories(id),
  brand_id INTEGER REFERENCES brands(id),
  location_id INTEGER REFERENCES locations(id),
  status VARCHAR(10) NOT NULL CHECK (status IN ('active', 'broken'))
);

-- Create Categories Table
CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- Create Brands Table
CREATE TABLE brands (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- Create Locations Table
CREATE TABLE locations (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

INSERT INTO items (name, item_code, stock, description, category_id, status, brand_id, location_id)
VALUES ('Mjolnir', 'MJOLNIR-001', 1, 'Thor''s hammer made from Uru', 1, 'active', 1, 1),
       ('Infinity Gauntlet', 'IG-001', 1, 'Thanos'' glove that can hold Infinity Stones', 2, 'broken', 5, 5),
       ('Iron Man Suit', 'IM-001', 1, 'Tony Stark''s suit made from vibranium', 3, 'active', 2, 2),
       ('Captain America''s Shield', 'CA-001', 1, 'Captain America''s shield made from vibranium', 3, 'active', 2, 2),
       ('Black Panther Suit', 'BP-001', 1, 'Black Panther''s suit made from vibranium', 3, 'active', 3, 3),
       ('Hawkeye''s Bow', 'HB-001', 1, 'Hawkeye''s bow made from wood and metal', 4, 'active', 2, 2),
       ('Black Widow''s Batons', 'BW-001', 2, 'Black Widow''s batons made from vibranium', 4, 'active', 2, 2),
       ('Loki''s Scepter', 'LS-001', 1, 'Loki''s scepter made from Chitauri metal', 2, 'broken', 4, 5);

INSERT INTO categories (name)
VALUES ('Weapons'),
       ('Jewelry'),
       ('Armor'),
       ('Equipment');

INSERT INTO locations (name)
VALUES ('Asgard'),
       ('Midgard'),
       ('Wakanda'),
       ('Earth'),
       ('Titan');

INSERT INTO brands (name)
VALUES ('Asgardian'),
       ('Stark Industries'),
       ('Wakandan Design Group'),
       ('Chitauri'),
       ('Asgard Art');
