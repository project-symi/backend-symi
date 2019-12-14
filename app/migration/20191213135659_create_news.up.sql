START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS news;
CREATE TABLE news (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(150) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    photo_link VARCHAR(150) NOT NULL,
    hidden boolean default false,
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    PRIMARY KEY (id)
  );
set
  foreign_key_checks = 1;
COMMIT;