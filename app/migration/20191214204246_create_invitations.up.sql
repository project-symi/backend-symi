START TRANSACTION;
set
  foreign_key_checks = 0;
DROP TABLE IF EXISTS invitation_status_categories;
DROP TABLE IF EXISTS invitations;
CREATE TABLE invitation_status_categories (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    status VARCHAR(20),
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
  );
CREATE TABLE invitations (
    id INT AUTO_INCREMENT NOT NULL UNIQUE,
    sender_id INT NOT NULL,
    employee_id INT NOT NULL,
    comments TEXT,
    invitation_status_category_id INT NOT NULL default 1,
    reply VARCHAR(5000) default "",
    seen boolean default false,
    invitation_date DATETIME,
    deleted boolean default false,
    deleted_at DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (employee_id) REFERENCES users(id),
    FOREIGN KEY (invitation_status_category_id) REFERENCES invitation_status_categories(id),
    PRIMARY KEY (id)
  );
set
  foreign_key_checks = 1;
COMMIT;