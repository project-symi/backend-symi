DELETE FROM point_categories;
DELETE FROM point_logs;
INSERT INTO point_categories (
    id,
    name,
    point,
    life_time_month,
    created_at,
    modified_at
  )
VALUES
  (1, 'Poll', 10, 3, NOW(), NOW()),
  (2, 'Submitted Feedback', 25, 3, NOW(), NOW()),
  (
    3,
    'Recieved positive feedback',
    50,
    3,
    NOW(),
    NOW()
  );
INSERT INTO point_logs(
    user_id,
    point_category_id,
    feedback_id,
    expire_date,
    created_at,
    modified_at
  )
VALUES
  (4, 2, 6, "2019-12-31", NOW(), NOW()),
  (4, 1, 7, "2019-12-31", NOW(), NOW()),
  (2, 3, 6, "2019-12-31", NOW(), NOW());
UPDATE users
SET
  total_points = 35
WHERE
  id = 4;
UPDATE users
SET
  total_points = 50
WHERE
  id = 2;