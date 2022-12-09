# Exam project

## Technologies used in this service
### Technologies


### Programming language:


golang programming language


### Database:


PostgresSql

### Technologies:


gRPC for connecting microservices


kafka for connecting microservice


### Connection Database:


sqlx Conn to connect to PostgreSql


### Libraries used:


     "google.golang.org/grpc"
	 "google.golang.org/grpc/reflection"
     "github.com/spf13/cast"
     "github.com/jmoiron/sqlx"
     "google.golang.org/grpc/codes"
     "google.golang.org/grpc/status"


## Get Clone

ssh: 
```
git@github.com:kabulovdev/pro.git 

```
https:
```
https://github.com/kabulovdev/pro.git

```

## Note!!!

This project's database is built on AWS

If you want to restore the database
you should cd services and 
up:
```
make migrate_down

```
down:
```
make migrate_up
```

## Run 

Steps


first
```
sudo docker compose build
```
second
```
sudo odcker compose up
```
third


You should 
create topic in localhost:8080; Name topic customer.customer

Swagger doc:
```
http://localhost:9079/v1/swagger/index.html#/
```



## After run and creating topic in terminal

```
docker stop post_service
```

```
docker start post_service
```

## Admin & Moder
Admin
name: 
```
abduazim 
```
password: 
```
sdy12197
``` 

Moder
name: 
```
kabulov 
```
password: 
```
sdy12197
``` 