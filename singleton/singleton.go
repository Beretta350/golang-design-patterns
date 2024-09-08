package singleton

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type single struct {
	Id string
}

var once sync.Once
var instance *single

func SingleInstance() *single {
	once.Do(func() {
		instanceId, _ := uuid.NewUUID()
		instance = &single{Id: instanceId.String()}

		log.Printf("New single instance created with %v ID!\n", instanceId.String())
	})

	log.Printf("Returning instance with %v ID!\n", instance.Id)
	return instance
}
