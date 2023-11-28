package payload

import (
	"app/model"
	"encoding/json"
	"log"
	"time"
)

type Student struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age"`
	Grade        float32   `json:"grade"s`
	ClassName    string    `json:"class_name"`
	EntranceDate time.Time `json:"entrance_date"`
}

func (c *Student) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
func (c *Student) ToModel() model.Student {
	return model.Student{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Age:          c.Age,
		Grade:        c.Grade,
		ClassName:    c.ClassName,
		EntranceDate: c.EntranceDate,
	}
}
