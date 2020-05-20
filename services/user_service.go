package services

import (
	"time"

	models "github.com/br-invin89/goknack-v2/server/models"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// GetByEmail func queries for and returns a single user object by email address
func GetByEmail(c echo.Context) (models.User, error) {
	email := c.Param("email")

	user, err := models.GetUserByEmail(email)

	if err != nil {
		return user, err
	}

	return user, nil
}

// GetOrCreate func
func GetOrCreate(c echo.Context) (models.User, error) {

	// fetch params
	email := c.FormValue("email")
	first := c.FormValue("first")
	last := c.FormValue("last")
	provider := c.FormValue("provider")
	img := c.FormValue("image")

	user, err := models.GetUserByEmail(email)

	if err != nil {

		if err.Error() == "not found" {
			user, err = models.CreateUser(models.User{ID: bson.NewObjectId(), Email: email, FirstName: first, LastName: last, Image: img, FullName: first + " " + last, LastLogin: time.Now(), Created: time.Now()})
		} else {
			return models.User{}, err
		}
	}

	// write auth logs
	switch provider {
	case "google":
		models.WriteGoogleLog(user.ID, c)
		break
	case "facebook":
		models.WriteFacebookLog(user.ID, c)
		break
	default:
		break
	}

	return user, nil
}

// All func
func All() ([]models.User, error) {

	return models.GetUsers()

}
