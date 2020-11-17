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
  - go-kit - [github.com/go-kit/kit](https://github.com/go-kit/kit)
  - mux - [github.com/gorilla/mux](https://github.com/gorilla/mux)


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
  - `/questions/{questionID}/answers` - update the answer of an existing question
    - **method**: `PUT`
    - **route param**: `questionID` (ID of an existing question)
    - **request data**: JSON object with the new answer
      - example: `{"answer": "new answer"}`
    - **response**: JSON containing the updated question with its answer
  - `/questions/{questionID}` - delete a question
    - **method**: `DELETE`
    - **route param**: `questionID` (ID of an existing question)
    - **response**: null (question deleted)
  - `/users/{userID}/questions` - get all the questions created by a user
    - **method**: `GET`
    - **route param**: `userID` (ID of the user)
    - **response**: JSON array containing all questions created by a user or an empty array if none were created.
      - example: - `[{"id": 1, "user_id": 1,"statement": "What is 3^2?"}, {"id": 2, "user_id": 1,"statement": "What is 2 + 2?"}]`
  

