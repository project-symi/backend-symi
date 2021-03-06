START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS feelings;
DROP TABLE IF EXISTS feedbacks;
DROP TABLE IF EXISTS categories;
CREATE TABLE feelings (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(20),
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY (id)
  );
CREATE TABLE categories (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(20),
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
  );
CREATE TABLE feedbacks (
    id INT AUTO_INCREMENT NOT NULL,
    user_id INT NOT NULL,
    feeling_id INT NOT NULL,
    category_id INT NOT NULL,
    recipient_id INT default 0,
    news_id INT default 0,
    feedback_note TEXT,
    seen boolean default false,
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (feeling_id) REFERENCES feelings(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    PRIMARY KEY (id)
  );
set
  foreign_key_checks = 1;
COMMIT;