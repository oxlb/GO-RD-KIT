CREATE TABLE todos (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  completed BOOLEAN NOT NULL
);


INSERT INTO todos (title, description, completed)
VALUES ('Pay bills', 'Pay the electric and water bills', false);
