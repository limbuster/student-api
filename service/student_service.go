package service

interface StudentService {
	AddStudent(student Student) 
	GetStudent(email string) Grade 
}

type Grade struct {
	subject string
	grade int
}

type Student struct {
	email string
	firstName string
	lastName string
	dateOfBirth string
	grades []Grade
}