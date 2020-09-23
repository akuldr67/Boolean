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
If you are running project `without docker`, and have given different user name/password or database name in mysql configuration above,  than change the `.env` file variables.

## Installation
### With Docker

- Build
```
docker build -t boolean .
```

- Run (change variable values accordingly)
```
docker run -p 8080:8080 -e DOCKER=true -e DB_USER='boolean' -e DB_PASS='booleanPw' -e DB_NAME='boolean' boolean
```

### Without Docker
 - Clone this repository and `cd` to the `Boolean` directory where you cloned it.
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


## API
### Base URL
```
http://localhost:8080
```

### Get All Booleans
- Use `GET /` to get all booleans
- Response: 
```
[
    {
        "id": "1adc9cfd-ff45-428b-b4eb-96fced361ac3",
        "key": "a boolean",
        "value": false
    },
    {
        "id": "27259a61-189c-48b0-9db6-d3fe6f4078a6",
        "key": "name",
        "value": false
    }
]
```

### Create a Boolean
 - Use `POST /` to create a boolean
 - Request:
    ```
    {
      "value":true,
      "key": "name" // this is optional
    }
    ```
    - Value should be either `true` or `false` (`boolean`, not a string)
    - Key should be a string
 - Response:
    ```
    {
      "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
      "value": true,
      "key": "name"
    } 
    ```

### Get a Boolean
- Use `GET /:id` to get a particular boolean
- Response:
  ```
  {
    "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
    "value": true,
    "key": "name"
  }
  ```

### Update a Boolean
- Use `PATCH /:id` to update a particular boolean
- Request:
  ```
  {
    "value":false,
    "key": "new name" // this is optional
  }
  ```
- Response:
  ```
  {
    "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
    "value": false,
    "key": "new name"
  }
  ```

### Delete a Boolean
- Use `DELETE /:id` to delete a particular boolean
- Response:
  ```
  HTTP 204 No Content
  ```