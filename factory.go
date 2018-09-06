package gomongo

import (
	"log"
	"time"

	mgo "github.com/globalsign/mgo"
)

//Factory function
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

func ConnectMongo(config *Config) (*Connection, error) {
	
	
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.Hosts},
		Timeout:  60 * time.Second,
		Database: config.Database,
		Source: config.AuthDatabase,
		Username: config.Username,
		Password: config.Password,
	}
	
	log.Println("echo : ", config, mongoDBDialInfo)

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Println("error kalp : ", err, config, mongoDBDialInfo)
		return nil, err
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	conn := new(Connection)
	conn.Session = mongoSession
	conn.Collections = make(map[string]*mgo.Collection)
	conn.Session.DB(config.Database)
	conn.Database = config.Database

	return conn, nil
}
