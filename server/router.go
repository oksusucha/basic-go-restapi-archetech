package server

import customtask "basic_archetech/controller/v1/custom_task"

type APIEvents struct {
	CT *customtask.CustomTaskAPIEvent
	// Please modify to your api path.
}

func newApiEvents() *APIEvents {
	db := makeRDBMSConnection()

	return &APIEvents{
		// Please modify to your api path.
		CT: &customtask.CustomTaskAPIEvent{DB: db},
	}
}

func initRouter() {
	apiEvents := newApiEvents()

	// r.Use()

	api := r.Group("api")

	v1 := api.Group("v1")

	// Please modify to your api path.
	customTask := v1.Group("ct")
	{
		customTask.GET("/ct", apiEvents.CT.PleaseChangeToYourAPIServiceName)
	}
}
