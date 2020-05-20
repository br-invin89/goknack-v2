package main

import (
	setting "github.com/br-invin89/goknack-v2/server/helpers/setting"
	models "github.com/br-invin89/goknack-v2/server/models"
	routes "github.com/br-invin89/goknack-v2/server/routes"
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
