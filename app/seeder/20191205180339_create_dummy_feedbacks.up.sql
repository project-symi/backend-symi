START TRANSACTION;
set
  foreign_key_checks = 0;
DELETE FROM feelings;
DELETE FROM categories;
DELETE FROM feedbacks;
INSERT INTO feelings (id, name)
VALUES
  (1, 'good'),
  (2, 'meh'),
  (3, 'sad');
INSERT INTO categories (id, name)
VALUES
  (1, 'Work/Life Balance'),
  (2, 'Benefits'),
  (3, 'Holidays'),
  (4, 'Job Satisfaction'),
  (5, 'Company Policy'),
  (6, 'News'),
  (7, 'Employee'),
  (8, 'Other');
INSERT INTO feedbacks (
    user_id,
    feeling_id,
    category_id,
    recipient_id,
    news_id,
    feedback_note
  )
VALUES
  (
    2,
    1,
    1,
    0,
    0,
    'I want to work more.'
  ),
  (
    2,
    2,
    1,
    0,
    0,
    'I need money.'
  ),
  (
    4,
    1,
    3,
    0,
    0,
    'I don\'t want to work'
  ),
  (
    2,
    1,
    1,
    0,
    0,
    'I want to work more'
  ),
  (
    4,
    3,
    1,
    0,
    0,
    'I want more holiday'
  ),
  (
    4,
    3,
    7,
    2,
    0,
    'Awesome Igor'
  ),
  (
    4,
    3,
    6,
    0,
    1,
    'Awesome news'
  ),
  (
    3,
    1,
    6,
    0,
    0,
    'I\'m beautiful model'
  );
set
  foreign_key_checks = 1;
COMMIT;