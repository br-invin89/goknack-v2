package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"database/sql"
	"crypto/tls"
	"net"
	"time"
)

var (
	db *sql.DB

	// Session var
	Session *mgo.Session
)

// Setup func
func Setup() {
	var err error

	log.Println("Initializing MongoDB connection...")

    defer func() {
        if r := recover(); r != nil {
            log.Fatalf("Mongo panic detected: %v", err)
            var ok bool
            err, ok := r.(error)
            if !ok {
                fmt.Printf("pkg:  %v,  error: %s", r, err)
            }
        }
    }()

	
	z := &mgo.DialInfo{}
	z.Addrs = []string{"goknack1-shard-00-00-8zw76.mongodb.net:27017", "goknack1-shard-00-01-8zw76.mongodb.net:27017", "goknack1-shard-00-02-8zw76.mongodb.net:27017"}
	z.Username = "goknack_admin"
	z.Password = "T2MDLZ6Ci2AVUeg"
	z.Database = "goknack"
	z.ReplicaSetName = "Goknack1-shard-0"
	z.Timeout = 60 * time.Second
	z.Source = "admin"

	tlsConfig := &tls.Config{}
    z.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	Session, err = mgo.DialWithInfo(z)

    if err!=nil{
		log.Println("MongoDB Init Error: %v", err)
        // panic(err)
    }
    log.Println("Successfully connected to MongoDB - goknack1-shard-00-00-8zw76.mongodb.net")
	Session.SetMode(mgo.Monotonic, true)

}

// CloseDB func
func CloseDB() {
	defer db.Close()
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}