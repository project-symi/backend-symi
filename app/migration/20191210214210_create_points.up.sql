START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS point_logs;
DROP TABLE IF EXISTS point_categories;
CREATE TABLE point_categories (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    name VARCHAR(50) UNIQUE,
    point INT NOT NULL,
    life_time_month INT NOT NULL,
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY (id)
  );
CREATE TABLE point_logs (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    user_id INT NOT NULL,
    point_category_id INT NOT NULL,
    feedback_id INT NOT NULL,
    expired boolean default false,
    expire_date DATE,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (point_category_id) REFERENCES point_categories(id)
  );
ALTER TABLE users
ADD
  total_points INT default 0 NOT NULL;
set
  foreign_key_checks = 1;
COMMIT;