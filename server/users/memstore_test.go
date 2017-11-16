package users

import (
	"testing"
)

/*
TestMemStore tests the MemStore object

Since a Store is like a database, you can't really test methods like Get()
or Delete() without also calling (and therefore testing) methods like Save(),
so instead of testing individual methods in isolation, this test runs through
a full CRUD cycle, ensuring the correct behavior occurs at each point in that
cycle. You should use a similar approach when testing your RedisStore implementation.
*/
func TestMemStore(t *testing.T) {
	newUser := &NewUser{
		Email: "MyEmailAddress@example.com ",
		Password: "password",
		PasswordConf: "password",
		UserName: "userName",
		FirstName: "firstName",
		LastName: "LastName",
	}

	store := NewMemStore()
	user, err := store.Insert(newUser)

	if err != nil {
		t.Errorf("Error inserting newuser: %v",err)
	}
	
	if _, err := store.GetByID(user.ID); err != nil {
		t.Errorf("Error getting user by id: %v",err)
	}
	

	
	if _, err := store.GetByEmail(user.Email); err != nil {
		t.Errorf("Error getting user by email: %v",err)
	}
	
	if _, err := store.GetByUserName(user.UserName); err != nil {
		t.Errorf("Error getting user by UserName: %v",err)
	}
	
	updates := &Updates{
		FirstName: "joy",
		LastName:  "Jaeger",
	}
	if err := store.Update(user.ID, updates); err != nil {
		t.Errorf("Error updating user: %v",err)
	}

	userID, err := store.GetByID(user.ID)
	if err != nil {
		t.Errorf("Error getting user by id after update: %v",err)
	}
	if userID.FirstName != "joy" {
		t.Errorf("FirstName not updated")
	}
	if userID.LastName != "Jaeger" {
		t.Errorf("LastName not updated")
	}

	if err := store.Delete(user.ID); err != nil {
		t.Errorf("Error deleting user: %v",err)
	}
	if err := store.Delete(user.ID); err == nil {
		t.Errorf("Expected error for deleting user")
	}

	if _, err := store.GetByID(user.ID); err == nil {
		t.Errorf("Expected error getting user by id")
	}
	

	if _, err := store.GetByEmail(user.Email); err == nil {
		t.Errorf("Expected error getting user by email")
	}
	
	if _, err := store.GetByUserName(user.UserName); err == nil {
		t.Errorf("Expected error getting user by UserName")
	}
	
	if err := store.Update(user.ID, updates); err == nil {
		t.Errorf("Expected error updating user")
	}

	
	
}
