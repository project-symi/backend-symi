START TRANSACTION;
DELETE FROM users;
DELETE FROM permissions;
DELETE FROM genders;
DELETE FROM departments;
INSERT INTO permissions (
    id,
    name,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (1, 'CEO', false, NOW(), NOW()),
  (2, 'employee', false, NOW(), NOW()),
  (3, 'admin', false, NOW(), NOW());
INSERT INTO genders (
    id,
    gender,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (1, 'male', false, NOW(), NOW()),
  (2, 'female', false, NOW(), NOW());
INSERT INTO departments (
    id,
    name,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (1, 'CEO', false, NOW(), NOW()),
  (2, 'Accounting', false, NOW(), NOW()),
  (3, 'Administration', false, NOW(), NOW()),
  (4, 'Advertising', false, NOW(), NOW()),
  (5, 'Audit', false, NOW(), NOW()),
  (6, 'Communications', false, NOW(), NOW()),
  (7, 'Development', false, NOW(), NOW()),
  (8, 'Distribution', false, NOW(), NOW()),
  (9, 'Engineering', false, NOW(), NOW()),
  (10, 'Marketing', false, NOW(), NOW()),
  (11, 'Sales', false, NOW(), NOW()),
  (12, 'Planning', false, NOW(), NOW()),
  (13, 'Materials', false, NOW(), NOW()),
  (14, 'Trading', false, NOW(), NOW());
INSERT INTO users (
    employee_id,
    name,
    mail,
    birthday,
    gender_id,
    department_id,
    permission_id,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (
    'A000001',
    'Mini',
    'samadova.minira@gmail.com',
    '20191205',
    2,
    1,
    1,
    false,
    NOW(),
    NOW()
  ),
  (
    'B000300',
    'Igor',
    'igor.m.byak@gmail.com',
    '20191205',
    1,
    7,
    3,
    false,
    NOW(),
    NOW()
  ),
  (
    'B000500',
    'Steffie',
    'steffie.harner@gmail.com',
    '20191205',
    2,
    10,
    3,
    false,
    NOW(),
    NOW()
  ),
  (
    'X009999',
    'Yukio',
    'triangle.pillow@gmail.com',
    '20191205',
    1,
    9,
    2,
    false,
    NOW(),
    NOW()
  );
COMMIT;