# Boolean As a Service
- Ability to get all booleans
- Ability to create, get, update and delete a boolean
- Storing data in mysql database
- Service exposing RESTful end points

## Tech Stack Used:
- Golang
- Mysql
  - Gorm as orm library
- Docker

## Configuration

### MySQL
 - Install MySQL in your machine. If you are having MacOS, you can follow [this link](https://flaviocopes.com/mysql-how-to-install/).
 - I would recommend you to create your new user in mysql and not use the root user. For eg, you can create user using the command:
 ```
CREATE USER 'boolean'@'localhost' IDENTIFIED BY 'booleanPw';
 ``` 
  - Create a database in your mysql
  ```
  CREATE DATABASE boolean;
  ``` 

- Grant all permissions for this database to the user created above
```
GRANT ALL PRIVILEGES ON boolean.* To 'boolean'@'localhost' IDENTIFIED BY 'booleanPw';
```

### Env
If you are 

## Installation
### With Docker




### Without Docker
 - Clone this repository and ```cd``` to the ```Boolean``` directory where you cloned it.
 - Install the go module
 ```
 go mod download
 ```
 - Build
 ```
 go build .
 ```
 - Run
 ```
 ./Boolean
```



