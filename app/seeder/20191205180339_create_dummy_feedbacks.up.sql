START TRANSACTION;
DELETE FROM feelings;
DELETE FROM categories;
DELETE FROM feedbacks;
INSERT INTO feelings (
    id,
    name,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (1, 'Good', false, NOW(), NOW()),
  (2, 'OK', false, NOW(), NOW()),
  (3, 'Bad', false, NOW(), NOW());
INSERT INTO categories (
    id,
    name,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (1, 'Work/Life Balance', false, NOW(), NOW()),
  (2, 'Benefits', false, NOW(), NOW()),
  (3, 'Holidays', false, NOW(), NOW()),
  (4, 'Job Satisfaction', false, NOW(), NOW()),
  (5, 'Company Policy', false, NOW(), NOW()),
  (6, 'Other', false, NOW(), NOW());
INSERT INTO feedbacks (
    user_id,
    feeling_id,
    category_id,
    feedback_note,
    deleted,
    created_at,
    modified_at
  )
VALUES
  (
    2,
    1,
    1,
    'I want to work more.',
    false,
    NOW(),
    NOW()
  ),
  (
    2,
    2,
    1,
    'I need money.',
    false,
    NOW(),
    NOW()
  ),
  (
    4,
    1,
    3,
    'I don\'t want to work',
    false,
    NOW(),
    NOW()
  ),
  (
    2,
    1,
    1,
    'I want to work more',
    false,
    NOW(),
    NOW()
  ),
  (
    4,
    3,
    1,
    'I want more holiday',
    false,
    NOW(),
    NOW()
  ),
  (
    3,
    1,
    6,
    'I\'m beautiful model',
    false,
    NOW(),
    NOW()
  );
COMMIT;