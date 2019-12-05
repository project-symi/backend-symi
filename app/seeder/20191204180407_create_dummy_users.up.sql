START TRANSACTION;
DELETE FROM users;
DELETE FROM types;
DELETE FROM departments;
INSERT INTO types (
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
    name,
    mail,
    birthday,
    department_id,
    type_id,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (
    'Mini',
    'samadova.minira@gmail.com',
    '20191205',
    1,
    1,
    false,
    NOW(),
    NOW()
  ),
  (
    'Igor',
    'igor.m.byak@gmail.com',
    '20191205',
    7,
    3,
    false,
    NOW(),
    NOW()
  ),
  (
    'Steffie',
    'steffie.harner@gmail.com',
    '20191205',
    10,
    3,
    false,
    NOW(),
    NOW()
  ),
  (
    'Yukio',
    'triangle.pillow@gmail.com',
    '20191205',
    9,
    2,
    false,
    NOW(),
    NOW()
  );
COMMIT;