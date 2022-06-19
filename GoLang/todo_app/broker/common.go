package broker

import (
	"log"

	uuid "github.com/google/uuid"
)

func GenUniqueId() string {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println("uuid.NewRandom() returned an error: " + err.Error())
	}
	return id.String()
}
