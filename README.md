# Student API

## Description

A simple API that will implement the following end-points:

```bash
GET /students/ -> return all students
GET /students/{email} -> return a specific student by email
POST /students/ -> add a new student and grades
PATCH /students/{email} -> update an existing record
GET /students/average/{subject} -> return average grades for a specific subject
```

Model for the student object will be something like:

```go
email string
firstName string
lastName string
address string
dateOfBirth date
grades []Grades
```

Model for the grade object will be something like:

```go
subject string
grade int (out of 100)
```

## Docker

```bash
docker-compose up -d                # Runs the docker container in daemon mode
docker-compose down                 # To stop the docker container
docker-compose build                # To rebuild the docker image
```
