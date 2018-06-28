#!/bin/bash

mongo << EOF
    use admin
    db.auth("admin", "admin1234")
    db.getUsers({"showCredentials":true})
    db.dropUser("testUser")
    use golang_test
    db.createUser({ user: "testUser", pwd: "test1234", roles: [{ role: "dbOwner", db: "golang_test" }] })
    db.auth("testUser", "test1234")
    quit()
EOF