package domain

import "strings"

type Gender struct {
	Id     int
	Gender string
}

type Genders []Gender

func (genders Genders) GenderToId(gender string) (id int) {
	for _, gen := range genders {
		if strings.ToLower(gen.Gender) == strings.ToLower(gender) {
			id = gen.Id
			return
		}
	}
	return
}
