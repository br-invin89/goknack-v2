package models

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
	//"fmt"
)

// User struct
type User struct {
	ID			bson.ObjectId	`json:"id" bson:"_id"`
	FullName  	string    		`json:"fullName" bson:"fullName"`
	FirstName 	string    		`json:"firstName" bson:"firstName"`
	LastName  	string    		`json:"lastName" bson:"lastName"`
	Image     	string   		`json:"image" bson:"image"`
	GImage    	string    		`json:"gImage" bson:"gImage"`
	FImage	  	string    		`json:"fImage" bson:"fImage"`
	Email     	string    		`json:"email" bson:"email"`
	IsActive  	bool      		`json:"isActive" bson:"isActive"`
	Updated   	time.Time 		`json:"updated" bson:"updated"`
	Created   	time.Time 		`json:"created" bson:"created"`
	LastLogin 	time.Time 		`json:"lastLogin" bson:"lastLogin"`
}

// FacebookLog struct
type FacebookLog struct {
	UserID		string 			`json:"userID" bson:"userID"`
	FacebookID	string			`json:"fbID" bson:"facebookID"`
	Timestamp	time.Time		`json:"timestamp" bson:"timestamp"`
}

// GoogleLog struct
type GoogleLog struct {
	UserID			string 		`json:"userID" bson:"userID"`
	GoogleID		string		`json:"googleID" bson:"googleID"`
	AccessToken		string		`json:"accessToken" bson:"accessToken"`
	Expires 		string 		`json:"expires" bson:"expires"`
	Issued 			string 		`json:"issued" bson:"issued"`
	Jwt 			string		`json:"jwt" bson:"jwt"`
}

// update the last login timestamp on the user record
func updateLastLogin(email string) {

	mongo := Session.Copy()
	defer mongo.Close()

	mongo.DB("goknack").C("user").Update(bson.M{"email": email}, bson.M{"$set": bson.M{"lastLogin": time.Now()}})
	mongo.Close()
}

// WriteGoogleLog func
func WriteGoogleLog(id bson.ObjectId, c echo.Context) {

	mongo := Session.Copy()
	defer mongo.Close()

	col := mongo.DB("goknack").C("auth_google")
	
	g := GoogleLog{}
	g.UserID = bson.ObjectId(id).Hex()
	g.GoogleID = c.FormValue("googleID")
	g.AccessToken = c.FormValue("accessToken")
	g.Expires = c.FormValue("expires")
	g.Issued = c.FormValue("issued")
	g.Jwt = c.FormValue("idToken")

   	err := col.Insert(&g)

	if err != nil {
		log.Println("INSERT ERROR: Google Log -", err)
	}
}

// WriteFacebookLog func
func WriteFacebookLog(id bson.ObjectId, c echo.Context) {
	mongo := Session.Copy()
	defer mongo.Close()

	col := mongo.DB("goknack").C("auth_fb")
	
	f := FacebookLog{}
	f.UserID = bson.ObjectId(id).Hex()
	f.FacebookID = c.FormValue("fbID")
	f.Timestamp = time.Now()

	err := col.Insert(&f)

 if err != nil {
	 log.Println("INSERT ERROR: Facebook Log -", err)
 }
}


// GetUsers func - return all users
func GetUsers() ([]User, error) {
	mongo := Session.Copy()
	defer mongo.Close()
	users := []User{}
	err := mongo.DB("goknack").C("user").Find(nil).All(&users)

	if err != nil {
		return users, err
	}
	
	return users, nil
}

// GetUserByEmail func
func GetUserByEmail (email string) (User, error){
	mongo := Session.Copy()
	defer mongo.Close()
	user := User{}
	err := mongo.DB("goknack").C("user").Find(bson.M{"email": email}).One(&user)

	if err != nil {
		return user, err
	}

	updateLastLogin(email)
	return user, nil
}

// CreateUser func
func CreateUser (user User) (User, error){
	mongo := Session.Copy()
	defer mongo.Close()

	c := mongo.DB("goknack").C("user")

	// insert new user record
	user.ID = bson.NewObjectId()
	err := c.Insert(&user)

	if err != nil {
		return user, err
	}

	// fetch new record
	err = c.Find(bson.M{"_id": user.ID}).One(&user)

	if err != nil {
		return user, err
	}

	return user, nil

}

// FindByID func
func FindByID(id string) {

}

