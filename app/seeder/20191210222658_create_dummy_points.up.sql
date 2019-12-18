START TRANSACTION;
set
  foreign_key_checks = 0;
DELETE FROM point_categories;
DELETE FROM point_logs;
INSERT INTO point_categories (
    id,
    name,
    point,
    life_time_month
  )
VALUES
  (1, 'Submitted Feedback', 25, 3),
  (
    2,
    'Recieved positive feedback',
    50,
    3
  );
INSERT INTO point_logs(
    user_id,
    point_category_id,
    feedback_id,
    expire_date
  )
VALUES
  (4, 1, 6, "2019-12-31"),
  (2, 2, 6, "2019-12-31");
UPDATE users
SET
  total_points = 25
WHERE
  id = 4;
UPDATE users
SET
  total_points = 50
WHERE
  id = 2;
set
  foreign_key_checks = 1;
COMMIT;