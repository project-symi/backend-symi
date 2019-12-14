START TRANSACTION;
set
  foreign_key_checks = 0;
DELETE FROM news;
INSERT INTO news (title, description, photo_link)
VALUES
  (
    'Celebrate Friday the 13th!',
    'Go home early and watch a scary movie! Company provides popcord.\nCheck with your manager for details!',
    'https://media1.giphy.com/media/PVDE7QM5tfokg/giphy.gif'
  ),
  (
    'Bring Your Pup to Work!',
    'You can now bring your puppy to work on Wednesdays! That''s something to celebrate.',
    'https://media.giphy.com/media/mRB9PmJFOjAw8/giphy.gif'
  ),
  (
    'Kentucky Christmas Party!',
    'Join the Kentucky Christmas party 12/23! We look forward to having you.',
    'https://media.giphy.com/media/in4t9IzuZKhqg/giphy.gif'
  ),
  (
    'CC-X Epic Graduation',
    'Come and see the magnificent CC 10 graduate projects on 12/26 (19:00~)',
    'https://slack-imgs.com/?c=1&o1=ro&url=https%3A%2F%2Fsecure.meetupstatic.com%2Fphotos%2Fevent%2F1%2Fe%2F2%2F1%2F600_487087713.jpeg'
  );
set
  foreign_key_checks = 1;
COMMIT;