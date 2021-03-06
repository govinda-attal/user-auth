CREATE SCHEMA IF NOT EXISTS users;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";




CREATE OR REPLACE TYPE USERS.ACCT_STATUS AS ENUM ('PENDING_CONFIRMATION', 'ACTIVE', 'INACTIVE');

CREATE OR REPLACE FUNCTION USERS.Trigger_SetTimestamp()  
RETURNS TRIGGER AS $$  
BEGIN  
  NEW.last_updated = NOW();
  RETURN NEW;
END;  
$$ LANGUAGE plpgsql;

DROP TABLE USERS.MFA;
DROP TABLE USERS.ACCT_CONFIRMATION;
DROP TABLE USERS.ACCOUNT;



CREATE TABLE USERS.ACCOUNT(
    user_id UUID DEFAULT uuid_generate_v4 (),
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    status USERS.ACCT_STATUS DEFAULT 'PENDING_CONFIRMATION',
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    last_login TIMESTAMP,
    last_updated TIMESTAMP,
    PRIMARY KEY (user_id)
);

CREATE TABLE USERS.ACCT_CONFIRMATION(
    rec_id UUID DEFAULT uuid_generate_v4 (),
    user_id UUID,
    token VARCHAR (255) UNIQUE NOT NULL,
    confirmed BOOLEAN NOT NULL DEFAULT FALSE,
    confirmed_on TIMESTAMP,
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_on TIMESTAMP NOT NULL,
    last_updated TIMESTAMP,
    PRIMARY KEY (rec_id),
    CONSTRAINT usr_acct_conf_user_id_fkey FOREIGN KEY (user_id) REFERENCES USERS.ACCOUNT (user_id) MATCH SIMPLE
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
);

CREATE TRIGGER AccountSetTimestamp  
BEFORE UPDATE ON USERS.ACCOUNT  
FOR EACH ROW  
EXECUTE PROCEDURE USERS.Trigger_SetTimestamp();

CREATE TRIGGER MfaSetTimestamp  
BEFORE UPDATE ON USERS.MFA  
FOR EACH ROW  
EXECUTE PROCEDURE USERS.Trigger_SetTimestamp(); 

CREATE TRIGGER AccountConfSetTimestamp  
BEFORE UPDATE ON USERS.ACCT_CONFIRMATION  
FOR EACH ROW  
EXECUTE PROCEDURE USERS.Trigger_SetTimestamp();     