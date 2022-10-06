-- USER postgres;
-- DROP DATABASE IF EXISTS pismodb;

-- CREATE DATABASE pismodb;

-- CREATE TABLE ACCOUNT
CREATE TABLE account(id serial primary key NOT NULL, doc_number varchar(20), created_at TIMESTAMP);
INSERT INTO account (doc_number, created_at) VALUES ('01558809930', '2022-10-05T13:33:29.386170889-03:00');

-- CREATE TABLE operation_types
CREATE TABLE operation_types (id serial primary key NOT NULL, description text);
INSERT INTO operation_types(description) VALUES ('COMPRA A VISTA');
INSERT INTO operation_types(description) VALUES ('COMPRA PARCELADA');
INSERT INTO operation_types(description) VALUES ('SAQUE');
INSERT INTO operation_types(description) VALUES ('PAGAMENTO');

-- CREATE TABLE trasactions
CREATE TABLE transactions (id serial primary key NOT NULL, account_id INTEGER, operation_type_id INTEGER, amount FLOAT, event_date TIMESTAMP);
INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (1, 1, -12.5, '2022-10-05T13:33:29.386170889-03:00');
INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (2, 1, -50.5, '2022-10-05T13:33:29.386170889-03:00');
INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (3, 1, -25.5, '2022-10-05T13:33:29.386170889-03:00');
INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (4, 1, 1000.00, '2022-10-05T13:33:29.386170889-03:00');