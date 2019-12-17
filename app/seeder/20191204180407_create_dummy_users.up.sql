START TRANSACTION;
set
  foreign_key_checks = 0;
DELETE FROM users;
DELETE FROM genders;
DELETE FROM departments;
DELETE FROM permissions;
INSERT INTO permissions (id, name)
VALUES
  (1, 'CEO'),
  (2, 'employee'),
  (3, 'admin');
INSERT INTO genders (id, gender)
VALUES
  (1, 'male'),
  (2, 'female');
INSERT INTO departments (id, name)
VALUES
  (1, 'CEO'),
  (2, 'Accounting'),
  (3, 'Administration'),
  (4, 'Advertising'),
  (5, 'Audit'),
  (6, 'Communications'),
  (7, 'Development'),
  (8, 'Distribution'),
  (9, 'Engineering'),
  (10, 'Marketing'),
  (11, 'Sales'),
  (12, 'Planning'),
  (13, 'Materials'),
  (14, 'Trading');
INSERT INTO users (
    employee_id,
    name,
    password,
    mail,
    birthday,
    gender_id,
    department_id,
    permission_id,
    total_points
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
    0
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
    50
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
    0
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
    35
  );
set
  foreign_key_checks = 1;
COMMIT;