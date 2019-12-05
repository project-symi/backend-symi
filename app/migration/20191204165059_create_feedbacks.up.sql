START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS feedbacks;
DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(20),
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY (id)
  );
CREATE TABLE feedbacks (
    id INT AUTO_INCREMENT NOT NULL,
    feedback_note VARCHAR(5000),
    user_id INT NOT NULL,
    category_id INT NOT NULL,
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    PRIMARY KEY (id)
  );
set
  foreign_key_checks = 1;
COMMIT;