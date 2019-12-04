START TRANSACTION;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS types;
DROP TABLE IF EXISTS departments;
CREATE TABLE types (
  id INT AUTO_INCREMENT NOT NULL,
  name VARCHAR(20),
  created_at TIMESTAMP,
  modified_at TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TABLE departments (
  id INT AUTO_INCREMENT NOT NULL,
  name VARCHAR(20),
  created_at TIMESTAMP,
  modified_at TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TABLE users (
  id INT AUTO_INCREMENT NOT NULL,
  name VARCHAR(20),
  mail VARCHAR(256),
  birthday DATE,
  department_id INT NOT NULL,
  type_id INT NOT NULL,
  created_at TIMESTAMP,
  modified_at TIMESTAMP,
  FOREIGN KEY (department_id) REFERENCES departments(id),
  FOREIGN KEY (type_id) REFERENCES types(id),
  PRIMARY KEY (id)
);
COMMIT;