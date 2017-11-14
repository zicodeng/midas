package users
import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
	"encoding/hex"	
	"crypto/md5"
	"io"
)

//TODO: add tests for the various functions in user.go, as described in the assignment.
//use `go test -cover` to ensure that you are covering all or nearly all of your code paths.


// Test the (nu *NewUser) Validate() function to ensure it catches all possible validation errors, 
// and returns no error when the new user is valid.
func TestValidateUser(t *testing.T) {
	cases := []struct {
		name          string
		hint          string
		nu    		  *NewUser
		expectError   bool
	}{
		{
			"Valid New User",
			"Idk it should be valid rn",
			&NewUser{
				Email: "uw@uw.uw",
				Password: "password",
				PasswordConf: "password",
				UserName: "userName",
				FirstName: "firstName",
				LastName: "LastName",
			},
			false,
		},
		{
			"Bad Password conf",
			"Check you password and passwordconf comparision",
			&NewUser{
				Email: "uw@uw.uw",
				Password: "password",
				PasswordConf: "passwords",
				UserName: "userName",
			},
			true,
		},
		{
			"Bad Email",
			"Check your email validation",
			&NewUser{
				Email: "uw.uw.uw",
				Password: "password",
				PasswordConf: "password",
				UserName: "userName",
			},
			true,
		},
		{
			"Short UserName",
			"Check UserName Validation",
			&NewUser{
				Email: "uw@uw.uw",
				Password: "password",
				PasswordConf: "password",
				UserName: "",
			},
			true,
		},
		{
			"Short password",
			"Check UserName Validation",
			&NewUser{
				Email: "uw@uw.uw",
				Password: "pass",
				PasswordConf: "pass",
				UserName: "",
			},
			true,
		},
	}
	for _, c := range cases {
		err := c.nu.Validate()
		if err != nil && !c.expectError {
			t.Errorf("case %s: unexpected error validating NewUser: %v\nHINT: %s", c.name, err, c.hint)
		}
		if err == nil && c.expectError {
			t.Errorf("case %s: unexpected error validating NewUser: %v\nHINT: %s", c.name, err, c.hint)
		}
		
	}
}
// Test the (nu *NewUser) ToUser() function to ensure it calculates the PhotoURL field correctly,
// even when the email address has upper case letters or spaces, and sets the PassHash field to 
// the password hash. Since bcrypt hashes are 
// salted with a random value, you can't anticipate what the hash should be, but you can verify the 
// generated hash by comparing it to the original password using the bcrypt package functions.

func TestToUser(t *testing.T) {
	
	newUser := &NewUser{
		Email: "MyEmailAddress@example.com ",
		Password: "password",
		PasswordConf: "password",
		UserName: "userName",
		FirstName: "firstName",
		LastName: "LastName",
	}
	user, err := newUser.ToUser()

	if err !=nil {
		t.Errorf("unexpected error making User: %v", err)
	}

	
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(newUser.Password)); err !=nil {
		t.Errorf("unexpected error hashing password: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte("passwords")); err ==nil {
		t.Errorf("expected error hashing incorrect password: %v", err)
	}
	if strings.HasPrefix(user.PhotoURL,gravatarBasePhotoURL){

		hash := strings.Replace(user.PhotoURL, gravatarBasePhotoURL, "", 1)
		h := md5.New()
		io.WriteString(h , "myemailaddress@example.com")
		avatar := h.Sum(nil)
		if hash != hex.EncodeToString(avatar){
			t.Errorf("Invalid processing of email for photoUrl")					
		}
	}else{
		t.Errorf("Invalid photoUrl base")		
	}
}

// Test the (u *User) FullName() function to verify that it returns the correct results given the 
// various possible inputs (no FirstName, no LastName, neither field set, both fields set).
func TestFullName(t *testing.T) {
	newUser := &NewUser{
		Email: "MyEmailAddress@example.com ",
		Password: "password",
		PasswordConf: "password",
		UserName: "userName",
		FirstName: "firstName",
		LastName: "LastName",
	}
	user, err := newUser.ToUser()

	if err !=nil {
		t.Errorf("unexpected error making User: %v", err)
	}
	if user.FullName() != "firstName LastName"{
		t.Errorf("Error in making fullname with first and last name set")	
	}

	user.FirstName=""
	if user.FullName() != "LastName"{
		t.Errorf("Error in making fullname with just last name set")	
	}
	user.FirstName="firstName"
	user.LastName=""
	if user.FullName() != "firstName"{
		t.Errorf("Error in making fullname with just firstName set")	
	}
	user.FirstName=""
	user.LastName=""
	if user.FullName() != ""{
		t.Errorf("Error in making fullname with no name set")	
	}
}

// Test the (u *User) Authenticate() function to verify that authentication happens correctly for 
// the various possible inputs (incorrect password, correct password).
func TestAuthenticate(t *testing.T) {
	newUser := &NewUser{
		Email: "MyEmailAddress@example.com ",
		Password: "password",
		PasswordConf: "password",
		UserName: "userName",
		FirstName: "firstName",
		LastName: "LastName",
	}
	user, err := newUser.ToUser()
	if err !=nil {
		t.Errorf("unexpected error making User: %v", err)
	}
	if err := user.Authenticate("password" ); err!= nil {
		t.Errorf("unexpected error hashing password: %v", err)
	}
	
	if err := user.Authenticate("passwords" ); err== nil {
		t.Errorf("expected error hashing incorrect password")
	}
	
}

// Test the (u *User) ApplyUpdates() function to ensure the user's fields are updated properly given 
// an Updates struct.
func TestApplyUpdates(t *testing.T) {
	newUser := &NewUser{
		Email: "MyEmailAddress@example.com ",
		Password: "password",
		PasswordConf: "password",
		UserName: "userName",
		FirstName: "firstName",
		LastName: "LastName",
	}
	user, err := newUser.ToUser()

	if err !=nil {
		t.Errorf("unexpected error making User: %v", err)
	}
	updates := &Updates{
		FirstName: "joy",
		LastName:"jaeger",
	}
	if err := user.ApplyUpdates(updates); err!= nil {
		t.Errorf("unexpected error updating firstname and lastname: %v", err)
	}

	updates = &Updates{
		FirstName: "",
		LastName:"jaeger",
	}
	if err := user.ApplyUpdates(updates); err== nil {
		t.Errorf("expected error updating with invalid firstname")
	}
	updates = &Updates{
		FirstName: "joy",
		LastName:"",
	}
	if err := user.ApplyUpdates(updates); err== nil {
		t.Errorf("expected error updating with invalid LastName")
	}

}