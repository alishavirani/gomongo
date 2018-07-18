package gomongo

import (
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// BulkInsert : Function inserts the data in bulk to the collection
// Input Parameters :
// 		*BulkInsertStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
// Output Parameters
// 		record(*mgo.BulkResult) : return count of matched and modified results
// 		error : if it was error then return error else nil
func (conn *Connection) BulkInsert(bulkInsertStruct *BulkInsertStruct) (*mgo.BulkResult, error) {
	var info *mgo.BulkResult
	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	bulk := collection.Bulk()
	bulk.Unordered()
	bulk.Insert(bulkInsertStruct.Data...)
	info, err := bulk.Run()
	if err != nil {
		log.Println(err)
		info = nil
	}
	return info, err
}

// Insert : Function inserts the data object into the collection
// Input Parameters :
// 		*InsertStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
// Output Parameters
// 		error : if it was error then return error else nil
func (conn *Connection) Insert(insertStruct *InsertStruct) error {
	var err error

	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	err = collection.Insert(&insertStruct.Data)
	if err != nil {
		log.Println(err)
	}
	return err
}

// InsertAsync : Function inserts the data object into the collection
// Input Parameters :
// 		*InsertStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
//			callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback(data, error) : sends data, error back to channel
func (conn *Connection) InsertAsync(insertStruct *InsertStruct, callback chan *Callback) {
	var err error
	err = conn.Insert(insertStruct)
	cb := new(Callback)
	cb.Data = nil
	cb.Error = err
	callback <- cb
}

// Update : Function Updates the record into the collection
// Input Parameters :
// 		*UpdateStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
// 			Id(string) : The record id whose details have to be updated
// Output Parameters
// 		error : if it was error then return error else nil

func (conn *Connection) Update(updateStruct *UpdateStruct) error {
	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	err := collection.UpdateId(bson.ObjectIdHex(updateStruct.Id), updateStruct.Data)
	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
		return err
	}
	return err
}

// UpdateAsync : Function Updates the record into the collection
// Input Parameters
// 		*UpdateStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
// 			Id(string) : The record id whose details have to be updated
//			callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback(data, error) : sends if error back to channel
func (conn *Connection) UpdateAsync(updateStruct *UpdateStruct, callback chan *Callback) {
	var err error
	err = conn.Update(updateStruct)
	cb := new(Callback)
	cb.Data = nil
	cb.Error = err
	callback <- cb
}

// Upsert : Function Updates the record if found else inserts as a new record into the collection
// Input Parameters :
// 		*UpsertStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
// 			Id(string) : The record id whose details have to be updated
// Output Parameters
// 		info(*mgo.ChangeInfo) : returns updated, removed, matched, record details count
// 		error : if it was error then return error else nil

func (conn *Connection) Upsert(upsertStruct *UpsertStruct) (*mgo.ChangeInfo, error) {
	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	info, err := collection.UpsertId(bson.ObjectIdHex(upsertStruct.Id), upsertStruct.Data)
	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
	}
	return info, err
}

// UpsertAsync : Function Updates the record if found else inserts as a new record into the collection
// Input Parameters :
// 		*UpsertStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
// 			Id(string) : The record id whose details have to be updated
// Output Parameters
// 		callback(data, error) : sends updated, matched, modified counts back to channel

func (conn *Connection) UpsertAsync(upsertStruct *UpsertStruct, callback chan *Callback) {

	info, err := conn.Upsert(upsertStruct)
	cb := new(Callback)
	cb.Data = info
	cb.Error = err
	callback <- cb
}

// UpdateAll : Function Updates all the record into the collection
// Input Parameters
// 		*UpdateAllStruct (Struct) :
//	 		Data([] interfaces{}]) : the object which has to be inserted
// 			Query(bson Object) : Criteria as per the update should execute
// Output Parameters
// 		records(*mgo.ChangeInfo) : returns updated, matched, modified counts
// 		error : if it was error then return error else nil

func (conn *Connection) UpdateAll(updateAllStruct *UpdateAllStruct) (*mgo.ChangeInfo, error) {
	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	records, err := collection.UpdateAll(updateAllStruct.Query, updateAllStruct.Data)
	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
		return nil, err
	}
	return records, nil
}

// UpdateAllAsync : Function Updates all the record into the collection
// Input Parameters
// 		*UpdateAllStruct (Struct) :
// 			Query(bson Object) : Criteria as per the update should execute
//	 		Data([] interfaces{}]) : the object which has to be inserted
//			callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback(data, error) : sends updated, matched, modified counts back to channel

func (conn *Connection) UpdateAllAsync(updateAllStruct *UpdateAllStruct, callback chan *Callback) {
	records, err := conn.UpdateAll(updateAllStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// UpsertAll : Function Upserts record into the collection if not found else updates
// Input Parameters
//	*UpsertAllStruct (Struct) :
//		Data([] interfaces{}]) : the object which has to be inserted
//		Query(bson Object) : Criteria as per the update should execute
// Output Parameters
//	records(*mgo.ChangeInfo) : returns updated, matched, modified counts
//	error : if it was error then return error else nil

func (conn *Connection) UpsertAll(upsertAllStruct *UpsertAllStruct) (*mgo.ChangeInfo, error) {
	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	records, err := collection.Upsert(upsertAllStruct.Query, upsertAllStruct.Data)
	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
		return nil, err
	}
	return records, nil
}

// UpsertAllAsync : Function Upserts record into the collection if not found else updates
// Input Parameters
//	*UpsertAllStruct (Struct) :
//		Data([] interfaces{}]) : the object which has to be inserted
//		Query(bson Object) : Criteria as per the update should execute
// Output Parameters
//	callback(data, error) : sends data, error to channel

func (conn *Connection) UpsertAllAsync(upsertAllStruct *UpsertAllStruct, callback chan *Callback) {
	records, err := conn.UpsertAll(upsertAllStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

//	FindByID : Function FindByID finds and returns record by Hexadecimal ID
//	Input Parameters :
// 		*FindByIDStruct (Struct) :
// 			Id(string) : The record id whose details have to be updated
//	Output :
//		record(interface{}) : Returns the mongo Object
//		error : Return error object if found
func (conn *Connection) FindByID(findByIDStruct *FindByIDStruct) (interface{}, error) {
	var record interface{}
	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	err := collection.FindId(bson.ObjectIdHex(findByIDStruct.Id)).Select(findByIDStruct.Fields).One(&record)
	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
	}
	return record, err
}

//	FindByIDAsync : Function FindByIDAsync finds and returns record by Hexadecimal ID
//	Input Parameters :
// 		*FindByIDStruct (Struct) :
// 			Id(string) : The record id whose details have to be updated
//			callback (channel) : which returns data to goroutine
//	Output :
//		callback(data, error) : sends data, error to channel
func (conn *Connection) FindByIDAsync(findByIDStruct *FindByIDStruct, callback chan *Callback) {
	records, err := conn.FindByID(findByIDStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// Find : Function finds the record into the collection according to the query/criteria
// Input Parameters
//		*FindStruct (Struct) :
// 			Query(bson Object) : Criteria as per the update should execute
//			Options(map[string]int) : optional things like, limit, skip, etc
// Output Parameters
// 		records([]interface{]}) : Return the result mapped as interface
// 		error(error) : if it was error then return error else nil
func (conn *Connection) Find(findStruct *FindStruct) ([]interface{}, error) {

	var records []interface{}
	var err error

	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(conn.Database).C(conn.Collection)

	limit, isLimit := findStruct.Options["limit"]
	skip, isSkip := findStruct.Options["isSkip"]

	if isLimit && isSkip {
		err = collection.Find(findStruct.Query).Select(findStruct.Fields).Skip(skip).Limit(limit).All(&records)
	} else if isLimit {
		err = collection.Find(findStruct.Query).Select(findStruct.Fields).Limit(limit).All(&records)
	} else if isSkip {
		err = collection.Find(findStruct.Query).Select(findStruct.Fields).Skip(skip).All(&records)
	} else {
		err = collection.Find(findStruct.Query).Select(findStruct.Fields).All(&records)
	}

	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
	}

	return records, err
}

// FindAsync : Function finds the record into the collection according to the query/criteria
// Input Parameters
//		*FindStruct (Struct) :
// 			Query(bson Object) : Criteria as per the update should execute
//			Options(map[string]int) : optional things like, limit, skip, etc
//			callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback(data, error) : returns data, error to channel
func (conn *Connection) FindAsync(findStruct *FindStruct, callback chan *Callback) {
	records, err := conn.Find(findStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// FindAll : Function finds all the records into the collection
// Input Parameters
//		*FindAllStruct (Struct) :
// Output Parameters
// 		records([]interface{}) : all the records present in database
// 		error : if it was error then return error else nil
func (conn *Connection) FindAll(findAllStruct *FindAllStruct) ([]interface{}, error) {

	var records []interface{}

	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(conn.Database).C(conn.Collection)

	err := collection.Find(nil).Select(findAllStruct.Fields).All(&records)
	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
	}

	return records, err
}

// FindAllAsync : Function finds all the records into the collection
// Input Parameters
//		*FindAllStruct (Struct) :
//			callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback(data, error) : returns data, error to channel
func (conn *Connection) FindAllAsync(findAllStruct *FindAllStruct, callback chan *Callback) {
	records, err := conn.FindAll(findAllStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// Remove : Function removes the record from the collection as per criteria/query
// Input Parameters
//		*RemoveStruct (Struct) :
// 			Query(bson Object) : Criteria as per the update should execute
// Output Parameters
// 		records(boolean) : returns true / false depending on output of operation
// 		error : if it was error then return error else nil
func (conn *Connection) Remove(removeStruct *RemoveStruct) error {
	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	err := collection.Remove(removeStruct.Query)
	if err != nil {
		log.Println(err)
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
	}
	return err
}

// removeAsync : Function removes the record from the collection as per criteria/query
// Input Parameters
//		*RemoveStruct (Struct) :
// 			Query(bson Object) : Criteria as per the update should execute
//			callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback(data, error) : returns data, error to channel
func (conn *Connection) RemoveAsync(removeStruct *RemoveStruct, callback chan *Callback) {
	err := conn.Remove(removeStruct)
	cb := new(Callback)
	cb.Data = nil
	cb.Error = err
	callback <- cb
}

// RemoveAll : Function removes all the record from the collection
// Input Parameters
//		*RemoveAllStruct (Struct) :
// Output Parameters
// 		records(*mgo.ChangeInfo) : returns matched modified removed count
// 		error : if it was error then return error else nil
func (conn *Connection) RemoveAll(removeAllStruct *RemoveAllStruct) (*mgo.ChangeInfo, error) {

	sessionCopy := conn.Session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(conn.Database).C(conn.Collection)
	records, err := collection.RemoveAll(nil)
	if err != nil {
		log.Println(err)
		records = nil
		if err.Error() == MongoErrorNotFound.Error() {
			err = nil
		}
	}

	return records, err
}

// RemoveAllAsync : Function removes all the record from the collection asynchronously
// Input Parameters
//		*RemoveAllStruct (Struct) :
//		callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback(data, error) : returns data, error to channel
func (conn *Connection) RemoveAllAsync(removeAllStruct *RemoveAllStruct, callback chan *Callback) {
	records, err := conn.RemoveAll(removeAllStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

func Close(conn *Connection) error {
	conn.Session.Close()
	return nil
}

//Database factory
//Return a DB for general purpose
func Init(driver string) (DB, error) {
	switch driver {
	case MONGODB:
		return new(MongoDB), nil
	default:
		return nil, ErrorInvalidDriver
	}
}
