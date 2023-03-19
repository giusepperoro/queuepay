create extension if not exists pg_stat_statements;

-- Creation of accounts table
CREATE TABLE IF NOT EXISTS accounts (
    client_id BIGINT,
    balance int NOT NULL,
    PRIMARY KEY(client_id)
    );

-- Creation of orders table
CREATE TABLE IF NOT EXISTS orders (
    client_id BIGINT NOT NULL,
    amount int,
    create_at timestamp NOT NULL,
    looked_at timestamp,
    status int
);

INSERT INTO accounts VALUES (1, 500);
INSERT INTO accounts VALUES (2, 1500);
INSERT INTO accounts VALUES (3, 3500);
INSERT INTO accounts VALUES (4, 4500);
INSERT INTO accounts VALUES (5, 5500);
