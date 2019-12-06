To run the docker container: 
  docker run -it -p 3000:8080 symi-backend

To drop migrations from aws:
  docker run -v /${PWD}/app/migration:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://symidb:symi-database@tcp(symi-db.chzk1jthuvd3.ap-northeast-1.rds.amazonaws.com:3306)/symi_test" drop