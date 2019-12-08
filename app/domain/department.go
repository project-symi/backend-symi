package domain

import "strings"

type Department struct {
	Id   int
	Name string
}

type Departments []Department

func (departments Departments) DepartmentToId(name string) (id int) {
	for _, department := range departments {
		if strings.ToLower(department.Name) == strings.ToLower(name) {
			id = department.Id
			return
		}
	}
	return
}
