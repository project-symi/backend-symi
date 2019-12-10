START TRANSACTION;
DELETE FROM users;
DELETE FROM permissions;
DELETE FROM genders;
DELETE FROM departments;
INSERT INTO permissions (
    id,
    name,
    created_at,
    modified_at
  )
VALUES
  (1, 'CEO', NOW(), NOW()),
  (2, 'employee', NOW(), NOW()),
  (3, 'admin', NOW(), NOW());
INSERT INTO genders (
    id,
    gender,
    created_at,
    modified_at
  )
VALUES
  (1, 'male', NOW(), NOW()),
  (2, 'female', NOW(), NOW());
INSERT INTO departments (
    id,
    name,
    created_at,
    modified_at
  )
VALUES
  (1, 'CEO', NOW(), NOW()),
  (2, 'Accounting', NOW(), NOW()),
  (3, 'Administration', NOW(), NOW()),
  (4, 'Advertising', NOW(), NOW()),
  (5, 'Audit', NOW(), NOW()),
  (6, 'Communications', NOW(), NOW()),
  (7, 'Development', NOW(), NOW()),
  (8, 'Distribution', NOW(), NOW()),
  (9, 'Engineering', NOW(), NOW()),
  (10, 'Marketing', NOW(), NOW()),
  (11, 'Sales', NOW(), NOW()),
  (12, 'Planning', NOW(), NOW()),
  (13, 'Materials', NOW(), NOW()),
  (14, 'Trading', NOW(), NOW());
INSERT INTO users (
    employee_id,
    name,
    password,
    mail,
    birthday,
    gender_id,
    department_id,
    permission_id,
    created_at,
    modified_at
  )
VALUES
  (
    'A000001',
    'Mini',
    'abc123',
    'samadova.minira@gmail.com',
    '20191205',
    2,
    1,
    1,
    NOW(),
    NOW()
  ),
  (
    'B000300',
    'Igor',
    'password1',
    'igor.m.byak@gmail.com',
    '20191205',
    1,
    7,
    3,
    NOW(),
    NOW()
  ),
  (
    'B000500',
    'Steffie',
    'nopassword',
    'steffie.harner@gmail.com',
    '20191205',
    2,
    10,
    3,
    NOW(),
    NOW()
  ),
  (
    'X009999',
    'Yukio',
    'nabe-monster',
    'triangle.pillow@gmail.com',
    '20191205',
    1,
    9,
    2,
    NOW(),
    NOW()
  );
COMMIT;