START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS slacks;
CREATE TABLE slacks (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    token VARCHAR(500) NOT NULL,
    url VARCHAR(500) NOT NULL default "https://slack.com/api/chat.postMessage",
    text text NOT NULL,
    deleted boolean,
    deleted_at DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
  );
ALTER TABLE users
ADD
  slack_member_id VARCHAR(50);
set
  foreign_key_checks = 1;
COMMIT;