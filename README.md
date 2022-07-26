# Introduction
    This project is for a testing management tool for testers. 
    
    If you were a tester or you ever heard about such tool like TestRail or QAtouch then it gonna be easy for you to understand this tool; 
    
    Product name is TestMaster

# 👨‍💻 Full list what has been used:
* [echo](https://github.com/labstack/echo) - Web framework
* [gorm](https://github.com/go-gorm/gorm) - Extensions to database/sql.
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
 Golang
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [jwt-go](https://github.com/dgrijalva/jwt-go) - JSON Web Tokens (JWT)
* [uuid](https://github.com/google/uuid) - UUID
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. Go
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
daemon for Go
* [Docker](https://www.docker.com/) - Docker

# Install
Make sure install Docker & make in your develop enviroment before do clone the repo:
1. Install Docker Desktop: [Download](https://www.docker.com/products/docker-desktop/)
2. Install make tool: 
- UBUNTU: `sudo apt install make`
- MAC: `brew install make`
3. Install go-migrate: [Detail](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)   
- MAC: `brew install golang-migrate`
- UBUNTU: 
```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash

sudo apt-get update

sudo apt-get install migrate
 ```

# RUN
1. Clone the repo:    
`git clone github.com/vnunicorn/testmaster.git`
2. Go to the folder:   
`cd testmaster`
3. Checkout to new branch:   
`git checkout -b mytest`
4. Start docker images:   
`make init`
5. Migrate database:   
`make up`
6. Init all library:   
`go mod tidy`


# ACCESS DATABASE USING TABLE PLUS
1. Install table Plus 
2. Access with info in Docker-compose file

# ACCESS DATABASE USING COMMAND LINE
```
docker exec -it mysql bash

mysql -u admin -p (enter)
(type password)
```

# Run
```
    cd src
    go run main.go
```
# Dev progress
## add userinfo (DOING)
    - add user management system (and Hung Lai added something)
    - init db
        - db connection
    - db crud for user
## add testing domain apis
    they're showed in doc repository