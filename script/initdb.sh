#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  GRANT ALL PRIVILEGES ON DATABASE bigproject_dev TO docker;
  CREATE TABLE ws_user (
  	user_id BIGSERIAL NOT NULL,
  	full_name VARCHAR(40) NOT NULL,
  	msisdn VARCHAR(40) NOT NULL,
  	user_email VARCHAR(40) NOT NULL,
  	birth_date DATE NOT NULL,
  	PRIMARY KEY (user_id)
  );


  INSERT INTO ws_user(user_id, full_name, msisdn, user_email, birth_date) VALUES (1, 'Dhanar J Kusuma', '000000001', 'dhanar.j.kusuma@gmail.com', '1980-01-01');
  INSERT INTO ws_user(user_id, full_name, msisdn, user_email, birth_date) VALUES (2, 'Arif Sholehddin', '000000002', 'arif.sholehuddin@gmail.com', '1980-01-01');
  INSERT INTO ws_user(user_id, full_name, msisdn, user_email, birth_date) VALUES (3, 'Fachrul W', '000000003', 'fachrul.w@gmail.com', '1980-01-01');
  INSERT INTO ws_user(user_id, full_name, msisdn, user_email, birth_date) VALUES (4, 'Adi Nurlemma', '000000004', 'adi.n@gmail.com', '1980-01-01');
  INSERT INTO ws_user(user_id, full_name, msisdn, user_email, birth_date) VALUES (5, 'Deko Picasso', '000000005', 'deko.p@gmail.com', '1980-01-01');
  INSERT INTO ws_user(user_id, full_name, msisdn, user_email, birth_date) VALUES (6, 'Verdana', '000000006', 'verdana@gmail.com', '1980-01-01');
  INSERT INTO ws_user(user_id, full_name, msisdn, user_email, birth_date) VALUES (7, 'Eko Priyanto', '000000007', 'eko.p@gmail.com', '1980-01-01');
EOSQL