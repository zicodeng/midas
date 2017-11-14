package users

import (
	"gopkg.in/mgo.v2/bson"
	"net/mail"
	"fmt"	
	"encoding/hex"	
	"crypto/md5"
	"io"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

var bcryptCost = 13

//User represents a user account in the database
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email"`
	PassHash  []byte        `json:"-"` //stored, but not encoded to clients
	UserName  string        `json:"userName"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	PhotoURL  string        `json:"photoURL"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Updates represents allowed updates to a user profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//Validate validates the new user and returns an error if
//any of the validation rules fail, or nil if its valid
func (nu *NewUser) Validate() error {
	//TODO: validate the new user according to these rules:
	//- Email field must be a valid email address (hint: see mail.ParseAddress)
	//- Password must be at least 6 characters
	//- Password and PasswordConf must match
	//- UserName must be non-zero length
	//use fmt.Errorf() to generate appropriate error messages if
	//the new user doesn't pass one of the validation rules

	_, err := mail.ParseAddress(nu.Email)
	if err != nil {
		return err
	}
	if len(nu.Password) < 6{
		return fmt.Errorf("Password must be at least 6 characters")
	}
	if nu.Password != nu.PasswordConf{
		return fmt.Errorf("Password confirmation must much password")
	}
	if len(nu.UserName) < 1{
		return fmt.Errorf("UserName must be at least 1 character")
	}
	

	return nil
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (nu *NewUser) ToUser() (*User, error) {
	//TODO: set the PhotoURL field of the new User to
	//the Gravatar PhotoURL for the user's email address.
	//see https://en.gravatar.com/site/implement/hash/
	//and https://en.gravatar.com/site/implement/images/

	//TODO: also set the ID field of the new User
	//to a new bson ObjectId
	//http://godoc.org/labix.org/v2/mgo/bson

	//TODO: also call .SetPassword() to set the PassHash
	//field of the User to a hash of the NewUser.Password
	h := md5.New()
	io.WriteString(h, strings.ToLower(strings.Trim(nu.Email," ")))
	avatar := h.Sum(nil)
	user := &User{
		ID:        bson.NewObjectId(),
		Email:     nu.Email,
		UserName:  nu.UserName,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
		PhotoURL:  gravatarBasePhotoURL + hex.EncodeToString(avatar),	
		}
	err := user.SetPassword(nu.Password)
	if err != nil {
		return nil, err
	}
		
	return user, nil
}

//FullName returns the user's full name, in the form:
// "<FirstName> <LastName>"
//If either first or last name is an empty string, no
//space is put betweeen the names
func (u *User) FullName() string {
	return strings.TrimSpace(u.FirstName + " " + u.LastName)
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	//TODO: use the bcrypt package to generate a new hash of the password
	//https://godoc.org/golang.org/x/crypto/bcrypt
	var err error
	u.PassHash, err = bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return err
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	//TODO: use the bcrypt package to compare the supplied
	//password with the stored PassHash
	//https://godoc.org/golang.org/x/crypto/bcrypt
	return bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	//TODO: set the fields of `u` to the values of the related
	//field in the `updates` struct, enforcing the following rules:
	//- the FirstName must be non-zero-length
	//- the LastName must be non-zero-length

	if len(updates.FirstName) > 0 {
		u.FirstName = updates.FirstName

	}else{
		return fmt.Errorf("FirstName updates must be at least 1 character")
	}
	if len(updates.LastName) > 0 {
		u.LastName = updates.LastName
	}else{
		return fmt.Errorf("LastName updates must be at least 1 character")
	}
	return nil
}
