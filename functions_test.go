package gomongo

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
)

//mock for testing
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Phone     string
	Salary    int
	DateTime  time.Time
}

//mock for testing
const (
	TestDBHosts  = "localhost:27017"
	TestDatabase = "golang_test"
	TestUserName = "testUser"
	TestPassword = "test1234"
)

func MockTestDB() {

	cmd := exec.Command("/bin/bash", "./mockdb.sh")
	fmt.Println("output : ", cmd)
}

func ConnectForTest() (*Connection, error) {

	//mock the database first
	MockTestDB()
	config := Config{"mongodb", TestDBHosts, TestDatabase, TestUserName, TestPassword}

	db, err := Init(MONGODB)

	if err != nil {
		log.Println("Connection error : ", err)
	}

	session, err := db.Connect(&config)

	if err != nil {
		log.Println("Connection error : ", err)
	}

	return session, err
}

func TestInit(t *testing.T) {
	db, err := Init(MONGODB)
	assert.Nil(t, err)

	config := Config{"mongodb", TestDBHosts, TestDatabase, TestUserName, TestPassword}
	_, err = db.Connect(&config)
	assert.Nil(t, err)
}

func TestInsert(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)
	tt := time.Now()
	user := new(Person)
	user.FirstName = "Amulya_T_" + tt.String()
	user.LastName = "Kasyap_T_" + tt.String()
	user.Age = 26
	user.Phone = "9559974779"
	user.Salary = 654654564
	user.DateTime = time.Now()
	conn.Collection = "users"
	insertStruct := new(InsertStruct)
	insertStruct.Data = user
	err = conn.Insert(insertStruct)
	assert.Nil(t, err)
}

func TestInsertAsync(t *testing.T) {

	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	outputCh := make(chan *Callback)

	tt := time.Now()
	user := new(Person)
	user.FirstName = "Amulya_T_" + tt.String()
	user.LastName = "Kasyap_T_" + tt.String()
	user.Age = 26
	user.Phone = "9559974779"
	user.Salary = 654654564
	user.DateTime = time.Now()

	insertStruct := new(InsertStruct)
	insertStruct.Data = user
	conn.Collection = "users"

	go conn.InsertAsync(insertStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}

func TestFindAll(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	findAllStruct := new(FindAllStruct)
	conn.Collection = "users"
	findAllStruct.Fields = bson.M{"firstname": 1, "lastname": 1}
	_, err = conn.FindAll(findAllStruct)
	assert.Nil(t, err)
}

func TestFindAllAsync(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	outputCh := make(chan *Callback)

	findAllStruct := new(FindAllStruct)
	conn.Collection = "users"
	findAllStruct.Fields = bson.M{"firstname": 1, "lastname": 1}
	go conn.FindAllAsync(findAllStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}

func TestFindById(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	findByIDStruct := new(FindByIDStruct)
	conn.Collection = "users"
	findByIDStruct.Id = "5b28da94a34bd180f5ab0f5a"
	findByIDStruct.Fields = bson.M{"firstname": 1, "lastname": 1}
	_, err = conn.FindByID(findByIDStruct)
	assert.Nil(t, err)
}

func TestFindByIdAsync(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	outputCh := make(chan *Callback)
	findByIDStruct := new(FindByIDStruct)
	conn.Collection = "users"
	findByIDStruct.Id = "5b28da94a34bd180f5ab0f5a"
	findByIDStruct.Fields = bson.M{"firstname": 1, "lastname": 1}
	go conn.FindByIDAsync(findByIDStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}

func TestFind(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	findStruct := new(FindStruct)
	conn.Collection = "users"
	findStruct.Query = bson.M{"_id": bson.ObjectIdHex("5b2a0045693a62e597564ddd")}
	findStruct.Fields = bson.M{"firstname": 1, "lastname": 1}
	_, err = conn.Find(findStruct)
	assert.Nil(t, err)
}

func TestFindAsync(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	outputCh := make(chan *Callback)
	findStruct := new(FindStruct)
	conn.Collection = "users"
	findStruct.Query = bson.M{"_id": bson.ObjectIdHex("5b28da94a34bd180f5ab0f5a")}
	findStruct.Fields = bson.M{"firstname": 1, "lastname": 1}
	go conn.FindAsync(findStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}

func TestBulkInsert(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	bulkData := make([]interface{}, 1000)

	for i := 1; i <= 1000; i++ {
		tt := time.Now()
		user := new(Person)
		user.FirstName = "Amulya_T_" + tt.String()
		user.LastName = "Kasyap_T_" + tt.String()
		user.Age = i + 26
		user.Phone = "9559974779"
		user.Salary = 1000 * i
		user.DateTime = time.Now()
		bulkData[i-1] = user
	}

	bulkInsertStruct := new(BulkInsertStruct)
	conn.Collection = "users"
	bulkInsertStruct.Data = bulkData

	_, err = conn.BulkInsert(bulkInsertStruct)
	assert.Nil(t, err)
}

func TestUpsert(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	objectId := bson.NewObjectId()

	upsertStruct := new(UpsertStruct)
	conn.Collection = "users"
	upsertStruct.Id = objectId.Hex()
	upsertStruct.Data = bson.M{"$set": bson.M{"firstname": "AmulyaUpsert1", "lastname": "KashyapUpsert1", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}

	_, err = conn.Upsert(upsertStruct)
	assert.Nil(t, err)
}

func TestUpsertAsync(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	outputCh := make(chan *Callback)

	upsertStruct := new(UpsertStruct)
	conn.Collection = "users"
	upsertStruct.Id = "5b28da94a34bd180f5ab0f5a"
	upsertStruct.Data = bson.M{"$set": bson.M{"firstname": "AmulyaXXX", "lastname": "Kashyap", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}

	go conn.UpsertAsync(upsertStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}

func TestUpdate(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	updateStruct := new(UpdateStruct)
	conn.Collection = "users"
	updateStruct.Id = "5b28da94a34bd180f5ab0f5a"
	updateStruct.Data = bson.M{"$set": bson.M{"firstname": "AmulyaXXX", "lastname": "Kashyap", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}

	err = conn.Update(updateStruct)
	assert.Nil(t, err)
}

func TestUpdateOne(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	updateStruct := new(UpdateOneStruct)
	conn.Collection = "users"
	updateStruct.Query = bson.M{"firstname": "Amulya", "lastname": "Kashyap"}
	updateStruct.Data = bson.M{"$set": bson.M{"firstname": "AmulyaXXX", "lastname": "KashyapXXX", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}

	err = conn.UpdateOne(updateStruct)
	assert.Nil(t, err)
}

func TestUpdateAsync(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	outputCh := make(chan *Callback)

	updateStruct := new(UpdateStruct)
	conn.Collection = "users"
	updateStruct.Id = "5b28da94a34bd180f5ab0f5a"
	updateStruct.Data = bson.M{"$set": bson.M{"firstname": "AmulyaXXX", "lastname": "Kashyap", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}

	go conn.UpdateAsync(updateStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}

func TestUpdateAll(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)

	assert.Nil(t, err)

	var updateStructAll UpdateAllStruct
	// updateStructAll := new(UpdateAllStruct)
	conn.Collection = "users"
	updateStructAll.Query = bson.M{"_id": bson.ObjectIdHex("5b28da94a34bd180f5ab0f5a")}
	updateStructAll.Data = bson.M{"$set": bson.M{"firstname": "AmulyaXXX", "lastname": "Kashyap", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}

	_, err = conn.UpdateAll(updateStructAll)
	assert.Nil(t, err)

}

func TestUpdateAllAsync(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)
	outputCh := make(chan *Callback)
	var updateStructAll UpdateAllStruct
	updateStructAll := new(UpdateAllStruct)
	conn.Collection = "users"
	updateStructAll.Query = bson.M{"_id": bson.ObjectIdHex("5b28da94a34bd180f5ab0f5a")}
	updateStructAll.Data = bson.M{"$set": bson.M{"firstname": "AmulyaXXX", "lastname": "Kashyap", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}
	go conn.UpdateAllAsync(updateStructAll, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}

func TestUpsertAll(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)

	assert.Nil(t, err)

	upsertStructAll := new(UpsertAllStruct)
	conn.Collection = "users"
	upsertStructAll.Query = bson.M{"_id": bson.NewObjectId()}
	upsertStructAll.Data = bson.M{"$set": bson.M{"firstname": "TIMESTAMP_AMULYA", "lastname": "Kashyap", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}

	_, err = conn.UpsertAll(upsertStructAll)
	assert.Nil(t, err)

}

func TestRemove(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	removeStruct := new(RemoveStruct)
	conn.Collection = "users"
	removeStruct.Query = bson.M{"_id": bson.ObjectIdHex("5b28da94a34bd180f5ab0f67")}

	err = conn.Remove(removeStruct)
	assert.Nil(t, err)
}

func TestRemoveAsync(t *testing.T) {
	conn, err := ConnectForTest()
	defer Close(conn)
	assert.Nil(t, err)

	outputCh := make(chan *Callback)

	removeStruct := new(RemoveStruct)
	conn.Collection = "users"
	removeStruct.Query = bson.M{"_id": bson.ObjectIdHex("5b28da94a34bd180f5ab0f63")}

	go conn.RemoveAsync(removeStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
	assert.True(t, true, output.Data)
}

func TestRemoveAll(t *testing.T) {

	conn, err := ConnectForTest()
	assert.Nil(t, err)
	removeAllStruct := new(RemoveAllStruct)
	conn.Collection = "users"
	_, err = conn.RemoveAll(removeAllStruct)
	assert.Nil(t, err)
}

func TestRemoveAllAsync(t *testing.T) {

	conn, err := ConnectForTest()
	assert.Nil(t, err)

	outputCh := make(chan *Callback)

	removeAllStruct := new(RemoveAllStruct)
	conn.Collection = "users"
	go conn.RemoveAllAsync(removeAllStruct, outputCh)
	output := <-outputCh
	assert.Nil(t, output.Error)
}
