package users

import (
	"gopkg.in/mgo.v2/bson"
)


type MemStore struct {
	entries []*User
}

//NewMemStore constructs and returns a new MemStore
func NewMemStore() *MemStore {
	return &MemStore{
		entries: []*User{},
	}
}

func  (mus *MemStore) GetByID(id bson.ObjectId) (*User, error){
	for _, user := range mus.entries {
		if user.ID == id {
			return user, nil
		}
	}
	return nil,ErrUserNotFound
}

	//GetByEmail returns the User with the given email
func  (mus *MemStore) GetByEmail(email string) (*User, error){
	for _, user := range mus.entries {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

	//GetByUserName returns the User with the given Username
func  (mus *MemStore) GetByUserName(username string) (*User, error){
	for _, user := range mus.entries {
		if user.UserName == username {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (mus *MemStore) GetAll() ([]*User, error) {
	users := make([]*User, len(mus.entries))
	copy(users, mus.entries)
	return users, nil
}

func (mus *MemStore) GetAllID(ids []bson.ObjectId) ([]*User, error) {
	users := []*User{}
	for	_, userID := range ids{
		user, err := mus.GetByID(userID)
		if err != nil {
			return nil,err
		}
		users = append(users, user)	
	}
	return users, nil
}
	//Insert converts the NewUser to a User, inserts
	//it into the database, and returns it
func (mus *MemStore) Insert(newUser *NewUser) (*User, error){
	user, err := newUser.ToUser()
	if err != nil {
		return nil, err
	}
	mus.entries = append(mus.entries, user)	
	return user, nil
}

	//Update applies UserUpdates to the given user ID
func  (mus *MemStore) Update(userID bson.ObjectId, updates *Updates) error{
	user, err := mus.GetByID(userID)
	if err != nil {
		return err
	}
	err = user.ApplyUpdates(updates)
	return err
}

	//Delete deletes the user with the given ID
func  (mus *MemStore) Delete(userID bson.ObjectId) error{
	for i, user := range mus.entries {
		if user.ID == userID {
			mus.entries = append(mus.entries[:i], mus.entries[i+1:]...)
			return nil
		}
	}
	return ErrUserNotFound		
}
