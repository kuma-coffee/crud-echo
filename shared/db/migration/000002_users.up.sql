CREATE TABLE users (
    user_id bigserial PRIMARY KEY,
	username VARCHAR ( 255 ) NOT NULL,
	email VARCHAR ( 255 ) NOT NULL,
	address TEXT NOT NULL,
	password VARCHAR ( 255 ) NOT NULL
);