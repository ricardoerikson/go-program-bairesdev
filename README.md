# Final Project - Golang Training Program (BairesDev)

This repository contains the final project developed in the Golang (Go) Training Program at BairesDev. The project consists of the backend of a simple site called QuestionsAndAnswers.com. Each question has only one answer. If someone thinks they have a better solution, they need to update the existing answer. Each question has 0 or 1 answer.

Requirements:
  - Docker
  - Git (for downloading the project)

## Features

The backend would support the following operations:
  - Get one question by its ID.
  - Get a list of all questions.
  - Get all the questions created by a given user.
  - Create a new question
  - Update an existing question (the statement and/or the answer)
  - Delete an existing question

## Technical specification

The project was developed using the following stack:

  - Docker (containers):
    - golang:1.15-alpine
    - postgres:13-alpine
  - [go-kit/kit](https://github.com/go-kit/kit)
  - [gorilla/mux](https://github.com/gorilla/mux)
  - ORM ([go-pg/pg](https://github.com/go-pg/pg)) - Golang Object Relational Mapping framework for PostgreSQL

The project uses the `go-kit` design architecture with the following layers: transport, endpoint, and service. In addition, the project uses a persistence layer, which is loosely coupled from the service layer. All the business logic goes in the service layer. The persistence layer is only concerned about storing and retrieving objects. The concrete implementation of the persistence layer can be replaced without affecting the existing code in the other layers. The persistence layer of this project contains three concrete implementations: 

  - An in-memory database (`mem`) implementation
  - A mock implementation (`mock`), which mocks the calls to the persistence layer
  - A PostgreSQL implementation (`pg`) using ORM

## Install

```
$ git clone git@github.com:ricardoerikson/go-program-bairesdev.git
$ cd go-program-bairesdev
$ docker-compose build
$ docker-compose up -d
```

## Routes

The following routes are supported:

  - `/questions` - create a new question
    - **method**: `POST`
    - **request data**: JSON object with the following fields: `user_id` and `statement`
      - example: `{"user_id": 1,"statement": "What is 3^2?"}` or an error message if the property `statement` is empty.
      - example: `{"error": "..."}`
    - **response**: question that was created
      - example: `{"id": 1, "user_id": 1,"statement": "What is 3^2?"}`
      - example: `{"error": "..."}`
    - **validation**:
      - if an `id` property is sent in the request, it is just ignored.
      - `user_id` must be greater than 0. Otherwise, an error is returned.
  - `/questions` - get all questions
    - **method**: `GET`
    - **response**: array with all questions or an empty array
      - `[{"id": 1, "user_id": 1,"statement": "What is 3^2?"}, {"id": 2, "user_id": 1,"statement": "What is 2 + 2?"}]`
      - `[]`
  - `/questions/{questionID}` - get a specific question
    - **method**: `GET`
    - **route param**: `questionID` (ID of an existing question)
    - **response**: question data for the specified `questionID`
  - `/questions/{questionID}` - update an existing question
    - **method**: `PUT`
    - **route param**: `questionID` (ID of an existing question)
    - **request data**: JSON object with the new statement for question with `questionID`
      - example: `{"statement": "new statement"}`
    - **response** JSON object containing the updated question
  - `/questions/{questionID}` - delete a question
    - **method**: `DELETE`
    - **route param**: `questionID` (ID of an existing question)
    - **response**: null (question deleted)
  - `/questions/{questionID}/answers` - update the answer of an existing question
    - **method**: `PUT`
    - **route param**: `questionID` (ID of an existing question)
    - **request data**: JSON object with the new answer
      - example: `{"answer": "new answer"}`
    - **response**: JSON containing the updated question with its answer
  - `/users/{userID}/questions` - get all the questions created by a user
    - **method**: `GET`
    - **route param**: `userID` (ID of the user)
    - **response**: JSON array containing all questions created by a user or an empty array if none were created.
      - example: - `[{"id": 1, "user_id": 1,"statement": "What is 3^2?"}, {"id": 2, "user_id": 1,"statement": "What is 2 + 2?"}]`
  
## Examples of Requests

### Create a question:
Request: 
```
$ curl -XPOST -d'{"user_id": 1,"statement": "What is 3-2?"}' localhost:8080/questions
```
Response: 
```json
{"id":2,"user_id":1,"statement":"What is 3-2?"}
```
### Get all questions:
Request:
```
$ curl -XGET localhost:8080/questions
```
Response:
```json
[
  {"id":1,"user_id":1,"statement":"What is 3^2?"},
  {"id":2,"user_id":1,"statement":"What is 3-2?"}
]
```

### Get question by ID:
Request:
```
$ curl -XGET localhost:8080/questions/1
```
Response:
```json
{"id":1,"user_id":1,"statement":"What is 3^2?"}
```

### Update a question
Request: 
```
$ curl -XPUT -d'{"statement": "What is 10*10?"}' localhost:8080/questions/1
```
Response:
```json
{"id":1,"user_id":1,"statement":"What is 10*10?"}
```

### Delete a question
Request:
```
$ curl -XDELETE localhost:8080/questions/1
```
Response: 
```
null
```

### Update answer
Request:
```
$ curl -XPUT -d'{"answer": "The answer is 1"}' localhost:8080/questions/2/answers
```
Response:
```json
{"id":2,"user_id":1,"statement":"What is 3-2?","answer":"The answer is 1"}
```

### Get all questions by a user
Request:
```
$ curl -XGET localhost:8080/users/1/questions
```

Response:
```json
[
  {"id":2,"user_id":1,"statement":"What is 3-2?","answer":"The answer is 9"}
]
```