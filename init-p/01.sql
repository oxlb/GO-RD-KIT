CREATE TABLE todos (
	id serial PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	completed BOOLEAN DEFAULT false
);

INSERT INTO todos (title, description, completed)
VALUES ('Pay bills', 'Pay the electric and water bills', false);