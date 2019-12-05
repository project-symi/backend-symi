START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS types;
DROP TABLE IF EXISTS departments;
CREATE TABLE types (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    name VARCHAR(20) UNIQUE,
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY (id)
  );
CREATE TABLE departments (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    name VARCHAR(20) UNIQUE,
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY (id)
  );
CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    name VARCHAR(20),
    mail VARCHAR(256) UNIQUE,
    birthday DATE,
    department_id INT NOT NULL,
    type_id INT NOT NULL,
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (type_id) REFERENCES types(id),
    PRIMARY KEY (id)
  );
set
  foreign_key_checks = 1;
COMMIT;