START TRANSACTION;
DELETE FROM feelings;
DELETE FROM categories;
DELETE FROM feedbacks;
INSERT INTO feelings (
    id,
    name,
    created_at,
    modified_at
  )
VALUES
  (1, 'good', NOW(), NOW()),
  (2, 'meh', NOW(), NOW()),
  (3, 'sad', NOW(), NOW());
INSERT INTO categories (
    id,
    name,
    created_at,
    modified_at
  )
VALUES
  (1, 'Work/Life Balance', NOW(), NOW()),
  (2, 'Benefits', NOW(), NOW()),
  (3, 'Holidays', NOW(), NOW()),
  (4, 'Job Satisfaction', NOW(), NOW()),
  (5, 'Company Policy', NOW(), NOW()),
  (6, 'News', NOW(), NOW()),
  (7, 'Employee', NOW(), NOW()),
  (8, 'Other', NOW(), NOW());
INSERT INTO feedbacks (
    user_id,
    feeling_id,
    category_id,
    recipient_id,
    news_id,
    feedback_note,
    created_at,
    modified_at
  )
VALUES
  (
    2,
    1,
    1,
    0,
    0,
    'I want to work more.',
    NOW(),
    NOW()
  ),
  (
    2,
    2,
    1,
    0,
    0,
    'I need money.',
    NOW(),
    NOW()
  ),
  (
    4,
    1,
    3,
    0,
    0,
    'I don\'t want to work',
    NOW(),
    NOW()
  ),
  (
    2,
    1,
    1,
    0,
    0,
    'I want to work more',
    NOW(),
    NOW()
  ),
  (
    4,
    3,
    1,
    0,
    0,
    'I want more holiday',
    NOW(),
    NOW()
  ),
  (
    4,
    3,
    7,
    2,
    0,
    'Awesome Igor',
    NOW(),
    NOW()
  ),
  (
    4,
    3,
    6,
    0,
    1,
    'Awesome news',
    NOW(),
    NOW()
  ),
  (
    3,
    1,
    6,
    0,
    0,
    'I\'m beautiful model',
    NOW(),
    NOW()
  );
COMMIT;