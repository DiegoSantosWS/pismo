-- USER postgres;
-- DROP DATABASE IF EXISTS pismodb;

-- CREATE DATABASE pismodb;

-- CREATE TABLE
CREATE TABLE account(id serial NOT NULL, doc_number varchar(20), created_at TIMESTAMP);
INSERT INTO account (doc_number, created_at) VALUES ('01558809930', '2022-10-05T13:33:29.386170889-03:00');