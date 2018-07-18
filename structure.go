package gomongo

import (
	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Callback struct {
	Data  interface{}
	Error error
}

type DB interface {
	Connect(*Config) (*Connection, error)
}

type Operations interface {
	BulkInsertStruct(*BulkInsertStruct)
	Insert(*InsertStruct) error
	InsertAsync(*InsertStruct, *Callback)
	Update(*UpdateStruct) error
	UpdateAsync(*UpdateStruct, *Callback)
	Upsert(*UpsertStruct) (*mgo.ChangeInfo, error)
	UpsertAsync(*UpsertStruct, *Callback)
	UpdateAll(*UpdateAllStruct) (*mgo.ChangeInfo, error)
	UpdateAllAsync(*UpdateAllStruct, *Callback)
	UpsertAll(*UpsertAllStruct) (*mgo.ChangeInfo, error)
	UpsertAllAsync(*UpsertAllStruct, *Callback)
	FindByID(*FindByIDStruct) (interface{}, error)
	FindByIDAsync(*FindByIDStruct, *Callback)
	Find(*FindStruct) ([]interface{}, error)
	FindAsync(*FindStruct, *Callback)
	FindAll(*FindAllStruct) ([]interface{}, error)
	FindAllAsync(*FindAllStruct, *Callback)
	Remove(*RemoveStruct) error
	RemoveAsync(*RemoveStruct, *Callback)
	RemoveAll(*RemoveAllStruct) (*mgo.ChangeInfo, error)
	RemoveAllAsync(*RemoveAllStruct, *Callback)
}

type Connection struct {
	Database    string                     //database name
	DialInfo    *mgo.DialInfo              // connection info
	Session     *mgo.Session               //session info
	Collections map[string]*mgo.Collection //all collections
	Collection  string                     //collection name
}

type Config struct {
	DbType   string
	Hosts    string //connection url i.e, localhost:27017
	Database string //database name
	Username string //username
	Password string //password
}

type BulkInsertStruct struct {
	Config *Config
	Data   []interface{}
}

type InsertStruct struct {
	Data interface{}
}

type UpdateStruct struct {
	Id   string
	Data interface{}
}

type UpsertStruct struct {
	Id   string
	Data interface{}
}

type UpdateAllStruct struct {
	Query bson.M
	Data  interface{}
}

type UpsertAllStruct struct {
	Query bson.M
	Data  interface{}
}

type FindByIDStruct struct {
	Id     string
	Fields bson.M
}

type FindStruct struct {
	Query   bson.M
	Options map[string]int
	Fields  bson.M
}

type FindAllStruct struct {
	Fields bson.M
}

type RemoveStruct struct {
	Query bson.M
}

type RemoveAllStruct struct {
}

type MongoDB struct{}
