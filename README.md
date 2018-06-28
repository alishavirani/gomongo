# MONGODB WRAPPER FOR GOLANG | mongodb, golang

> Basic functionalities for data management like insert, remove, find, findbyid, update, upsert, etc
>
> All functions are available in sync as well as async flavours like insert, insertAsync
>
> Just Append `Async` after the function name like insert, insertAsync

``` bash
# clone the repo
git clone https://gitlab.com/amulyakashyap09/gomongo.git
```

``` bash
import "github.com/amulyakashyap09/gomongo"
```

### Instantiate the connection
``` bash
    config := Config{"mongodb", MongoDBHosts, AuthDatabase, AuthUserName, AuthPassword}
    db, err := Init(MONGODB)
    
```

### Connect to database
``` bash
    sess, err := db.Connect(&config)
```

### Insert 
``` bash

# Insert sample
#   user := new(User)
#   user.FirstName = "AmulyaXIVXIV"
#   user.LastName = "KashyapXIVXIV"
#   user.Age = 26
#   user.Phone = "9559974779"
#   user.Salary = 654654564
#   user.DateTime = time.Now()

#   insertStr := new(InsertStruct)
#   insertStr.Data = user

#   sess.Collection = "users"

#   err = sess.Insert(insertStr)

#   if err!=nil {
#       fmt.Println("error in inserting : ", err)    
#   }else{
#       fmt.Println("record inserted successfully : ")    
#   }
```

### Update 
``` bash

# Update sample
#   updateStruct := new(UpdateStruct)
#   conn.Collection = "users"
#   updateStruct.Id = "5b28da94a34bd180f5ab0f5a"
#   updateStruct.Data = bson.M{"$set": bson.M{"firstname": "AmulyaXXX", "lastname": "Kashyap", "age": 26, "phone": "9559974779", "salary": "7854693210", "datetime": time.Now()}}
#   err = conn.Update(updateStruct)

#   if err!=nil {
#       fmt.Println("error in update : ", err)    
#   }else{
#       fmt.Println("record updated successfully : ")    
#   }
```


### Find 
``` bash

# Find sample
#   findStr := new(FindStruct)
#   findStr.Fields = bson.M{"firstname": 1}
#   findStr.Options = make(map[string]int)
#   findStr.Options["skip"] = 0
#   findStr.Options["limit"] = 10
#   findStr.Query = bson.M{"_id": bson.ObjectIdHex("5b28da94a34bd180f5ab0f5a")}
#   findData, err := sess.Find(findStr)

#   if err!=nil {
#       fmt.Println("error in find : ", err)    
#   }else{
#       fmt.Println("record found successfully : ", findData)    
#   }
```


### Remove 
``` bash

# Remove sample
#   removeStr := new(RemoveStruct)
#   removeStr.Query = bson.M{"_id": bson.ObjectIdHex("5b28da94a34bd180f5ab0f5a")}
#   err = sess.Remove(removeStr)

#   if err!=nil {
#       fmt.Println("error in find : ", err)    
#   }else{
#       fmt.Println("record found successfully : ", findData)    
#   }
```

## Project Details

### Author
```bash
    Amulya Kasyap
```

## Maintainer
```bash
    Amulya Ratan
        email : amulyakashyap09@gmail.com
        contact : +91-9559974779
```

### Version
```bash
    1.0.0
```

### License

```bash
    This project is licensed under the MIT License
```




