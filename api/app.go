package api

import (
	"fwebpanel/api/base"
	"fwebpanel/api/controllers"
	"sync"
)

func Run(group *sync.WaitGroup) {
	app := base.NewAPI(":8000")
	app.Register("/stats", new(controllers.StatsController))
	app.Run()
	group.Done()
}
