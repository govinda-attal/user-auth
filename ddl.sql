CREATE SCHEMA IF NOT EXISTS users;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE USERS.ACCOUNT(
    user_id UUID DEFAULT uuid_generate_v4 (),
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP,
    PRIMARY KEY (user_id)
);

CREATE TABLE USERS.MFA(
    rec_id UUID DEFAULT uuid_generate_v4 (),
    user_id UUID,
    mfa_enabled BOOLEAN NOT NULL,
    user_secret VARCHAR (255) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_updated TIMESTAMP,
    PRIMARY KEY (rec_id),
    CONSTRAINT usr_mfa_user_id_fkey FOREIGN KEY (user_id) REFERENCES USERS.ACCOUNT (user_id) MATCH SIMPLE
)