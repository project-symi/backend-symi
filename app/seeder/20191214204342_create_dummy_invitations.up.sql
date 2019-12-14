START TRANSACTION;
set
  foreign_key_checks = 0;
DELETE FROM invitation_status_categories;
DELETE FROM invitations;
INSERT INTO invitation_status_categories (id, status)
VALUES
  (1, "pending"),
  (2, "accepted"),
  (3, "refused");
INSERT INTO invitations (
    sender_id,
    employee_id,
    comments,
    invitation_status_category_id,
    reply,
    invitation_date
  )
VALUES
  (
    1,
    2,
    'Christmas party!!!',
    1,
    '',
    '2019-12-25'
  ),
  (
    1,
    3,
    'Let''s discuss about beauty',
    2,
    'Why not!',
    '2019-12-20'
  ),
  (
    1,
    4,
    'You have serious problems',
    3,
    'Sorry I have headache.',
    '2019-12-12'
  );
set
  foreign_key_checks = 1;
COMMIT;