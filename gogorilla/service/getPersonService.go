package service

import "gogorilla/model"

func GetPersonsService() ([]model.Person, error) {
	var persons []model.Person

	var person model.Person

	person.Id = 1
	person.FirstName = "Issac"
	person.LastName = "N"
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "E"
	persons = append(persons, person)
	return persons, nil
}
