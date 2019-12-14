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
    total_points,
    created_at,
    modified_at
  )
VALUES
  (
    'A000001',
    'Mini',
    '$2a$10$1bPVBUofakXGPsU9phR2q.CNbcAoXjv0pOzhYwgM3qX3moc/2kK4i',
    'samadova.minira@gmail.com',
    '20191205',
    2,
    1,
    1,
    0,
    NOW(),
    NOW()
  ),
  (
    'B000300',
    'Igor',
    '$2a$10$B0nZWE.h6V.PyGCr/8RUkuwRMRGMM3gPEVlt6NNVBjRdZWAfZ2a.G',
    'igor.m.byak@gmail.com',
    '20191205',
    1,
    7,
    3,
    50,
    NOW(),
    NOW()
  ),
  (
    'B000500',
    'Steffie',
    '$2a$10$xlVenMv3ssNZBWX1M3VVzev8T6oyvrGomZVrwZ2i1kh4e/5cQUvWa',
    'steffie.harner@gmail.com',
    '20191205',
    2,
    10,
    3,
    0,
    NOW(),
    NOW()
  ),
  (
    'X009999',
    'Yukio',
    '$2a$10$PRZEjcqzQbxmfnFw/QcRXuNc/R/ez.icPgm5nxjXf07iuCzWvRSxy',
    'triangle.pillow@gmail.com',
    '20191205',
    1,
    9,
    2,
    35,
    NOW(),
    NOW()
  );
COMMIT;