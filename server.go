package main

import (
	setting "server/helpers/setting"
	models "server/models"
	routes "server/routes"
	//"server/helpers/logging"
)

func main() {

	// init router
	e := routes.InitRouter()

	// init DB
	models.Setup()

	// init settings
	setting.Setup()

	// init logger
	//logging.Setup();

	e.Logger.Fatal(e.Start(":1323"))
}
