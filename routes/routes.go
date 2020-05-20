package routes

import (
	"net/http"
	"github.com/labstack/echo"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})



	return e;
}


//TODO OLD ROUTES THAT NEED TO BE REFACTORED INTO ECHO
// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

// r.GET("/auth/check-token", api.CheckAuth)
// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// r.POST("/upload", api.UploadImage)

// protected routes
// apiv1 := r.Group("/api/v1")

// IF you want the route secured
// apiv1.Use(jwt.JWT())
// apiv1.Use()
// {

	// user := apiv1.Group("/user")
// 	user.Use()
// 	{
// 		// user routes
// 		user.GET("/all", v1.GetUsers)
// 		user.POST("/find", v1.GetOrCreateUser)
// 	}

// 	content := apiv1.Group("/content")
// 	content.Use(jwt.JWT())
// 	{
// 		content.POST("/add", v1.AddContent)
// 	}


// }




func main() {}