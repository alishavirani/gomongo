package gomongo

import (
	"log"
	"time"

	mgo "github.com/globalsign/mgo"
)

// Connect ...
func (n MongoDB) Connect(config *Config) (*Connection, error) {
	switch config.DbType {
	case MONGODB:
		conn, err := ConnectMongo(config)
		return conn, err
	default:
		//if type is invalid, return an error
		return nil, ErrorInvalidDBType
	}
}

// ConnectMongo ...
func ConnectMongo(config *Config) (*Connection, error) {

	var mongoDBDialInfo *mgo.DialInfo
	var mongoSession *mgo.Session
	var err error

	if len(config.Uri) > 0 {
		mongoSession, err = mgo.Dial(config.Uri)
	}else {
		mongoDBDialInfo = &mgo.DialInfo{
			Addrs:    []string{config.Hosts},
			Timeout:  60 * time.Second,
			Database: config.Database,
			Service: "mongodb",
			Mechanism: "MONGODB-CR",
			Source: config.AuthDatabase,
			Username: config.Username,
			Password: config.Password,
		}
	
		mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	}
	
	if err != nil {
		log.Println("connection error : ", err)
		return nil, err
	}

	//mongoSession.SetMode(mgo.Monotonic, true)
	// mongoSession.SetSafe(&mgo.Safe{})
	//mongoSession.SetSafe(&mgo.Safe{WMode: "majority"})

	conn := new(Connection)
	conn.Session = mongoSession
	conn.Collections = make(map[string]*mgo.Collection)
	conn.Session.DB(config.Database)
	conn.Database = config.Database

	return conn, nil
}
