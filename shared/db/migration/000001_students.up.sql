CREATE TABLE students (
    id bigserial PRIMARY KEY,
	fullname VARCHAR ( 255 ) NOT NULL,
	address VARCHAR ( 255 ) NOT NULL,
	birthdate DATE NOT NULL,
	class VARCHAR ( 50 ) NOT NULL,
	batch INT NOT NULL,
	school_name VARCHAR ( 255 ) NOT NULL
);