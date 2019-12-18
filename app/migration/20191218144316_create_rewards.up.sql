START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS rewards;
CREATE TABLE rewards (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    points INT NOT NULL,
    url VARCHAR(150) NOT NULL,
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
  );
set
  foreign_key_checks = 1;
COMMIT;