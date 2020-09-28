package main

import (
	"github.com/mauriciobergallo/school-app-students-api/pkg/adding"
	"github.com/mauriciobergallo/school-app-students-api/pkg/deleting"
	"github.com/mauriciobergallo/school-app-students-api/pkg/handlers"
	"github.com/mauriciobergallo/school-app-students-api/pkg/listing"
	"github.com/mauriciobergallo/school-app-students-api/pkg/logging"
	"github.com/mauriciobergallo/school-app-students-api/pkg/storage/memory"
	"github.com/mauriciobergallo/school-app-students-api/pkg/updating"
)

func main() {
	// Creating Logging Service
	l := logging.NewStdoutLogging("DEBUG")

	// Creating Repo
	sm := new(memory.Storage)
	err := sm.RunSeeds()
	if err != nil {
		panic(err)
	}

	// Creating Services with DI
	as := adding.NewService(sm, l)
	ls := listing.NewService(sm, l)
	ds := deleting.NewService(sm, l)
	us := updating.NewService(sm, l)

	handlers.NewRestService(as, ls, ds, us)
}