START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS genders;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS departments;
CREATE TABLE permissions (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    name VARCHAR(20) UNIQUE,
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY (id)
  );
CREATE TABLE genders (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    gender VARCHAR(20) UNIQUE,
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
    employee_id VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(20),
    password VARCHAR(40),
    mail VARCHAR(256) UNIQUE,
    birthday DATE,
    gender_id INT NOT NULL,
    department_id INT NOT NULL,
    permission_id INT NOT NULL,
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    currentToken VARCHAR(36),
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (gender_id) REFERENCES genders(id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id),
    PRIMARY KEY (id)
  );
set
  foreign_key_checks = 1;
COMMIT;