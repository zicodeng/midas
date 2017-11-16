package users

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoStore struct {
	Session        *mgo.Session
	DbName   		string
	CollName 		string
}

//NewMemStore constructs and returns a new MemStore
func NewMongoStore(session *mgo.Session, dbName string, colName string) *MongoStore {
	return &MongoStore{
		Session:        session,
		DbName:   		dbName,
		CollName: 		colName,
	}
}



func (ms *MongoStore) GetByID(id bson.ObjectId) (*User, error) {
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	user := &User{}
	err := coll.FindId(id).One(user)
	return user, err
}

func (ms *MongoStore) GetAll() ([]*User, error) {
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	users := []*User{}
	err := coll.Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ms *MongoStore) GetAllID(ids []bson.ObjectId) ([]*User, error) {
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	users := []*User{}
	err := coll.Find(bson.M{"_id":bson.M{"$in": ids},}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
	//GetByEmail returns the User with the given email
func  (ms *MongoStore) GetByEmail(email string) (*User, error){
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	user := &User{}
	
	err := coll.Find(bson.M{"email": email}).One(user)
	if err != nil {		
		return nil, ErrUserNotFound
	}
	return user, nil
}

	//GetByUserName returns the User with the given Username
func  (ms *MongoStore) GetByUserName(username string) (*User, error){
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	user := &User{}
	err := coll.Find(bson.M{"username": username}).One(user)
	return user, err
}

	//Insert converts the NewUser to a User, inserts
	//it into the database, and returns it
func (ms *MongoStore) Insert(newUser *NewUser) (*User, error){
	user, err := newUser.ToUser()
	if err != nil {
		return nil, err
	}
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	err = coll.Insert(user)
	return user, err	
}

	//Update applies UserUpdates to the given user ID
func  (ms *MongoStore) Update(userID bson.ObjectId, updates *Updates) error{
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	
	dbUpdates := bson.M{"firstname": updates.FirstName, "lastname": updates.LastName}
	
	return coll.UpdateId(userID, bson.M{"$set": dbUpdates})
}

	//Delete deletes the user with the given ID
func  (ms *MongoStore) Delete(userID bson.ObjectId) error{
	coll := ms.Session.DB(ms.DbName).C(ms.CollName)
	return coll.RemoveId(userID)		
}
