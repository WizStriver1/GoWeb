package model

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"app/common/database"

	"github.com/boltdb/bolt"
	"gopkg.in/mgo.v2/bson"
)

// *****************************************************************************
// User
// *****************************************************************************

// User table contains the information for each user
type User struct {
	ObjectID   bson.ObjectId `bson:"_id"`
	ID         uint32        `db:"id" bson:"id,omitempty"`     // Don't use Id, use UserID() instead for consistency with MongoDB
	UserName   string        `db:"user_name" bson:"user_name"` // 用户名
	Email      string        `db:"email" bson:"email"`
	Phone      string        `db:"phone" bson:"phone"`
	Password   string        `db:"password" bson:"password"`
	CreateTime time.Time     `db:"create_time" bson:"create_time"`
	UpdateTime time.Time     `db:"update_time" bson:"update_time"`
	StatusID   uint8         `db:"status_id" bson:"status_id"`
	RoleID     uint8         `db:"role_id" bson:"role_id"`
	Deleted    uint8         `db:"deleted" bson:"deleted"` // deleted contains data delete status (0 deleted/1 stored)
}

// UserStatus table contains every possible user status (active/inactive)
type UserStatus struct {
	ID         uint8     `db:"id" bson:"id"`
	Status     string    `db:"status" bson:"status"`
	CreateTime time.Time `db:"create_time" bson:"create_time"`
	UpdateTime time.Time `db:"update_time" bson:"update_time"`
	Deleted    uint8     `db:"deleted" bson:"deleted"`
}

// UserRoles table contains every possible user status (administrator/standard)
type UserRoles struct {
	ID         uint8     `db:"id" bson:"id"`
	Role       string    `db:"role" bson:"role"`
	CreateTime time.Time `db:"create_time" bson:"create_time"`
	UpdateTime time.Time `db:"update_time" bson:"update_time"`
	Deleted    uint8     `db:"deleted" bson:"deleted"`
}

// Passwords table contains the information for each user who set a new password
type Passwords struct {
	ObjectID   bson.ObjectId `bson:"_id"`
	ID         uint32        `db:"id" bson:"id,omitempty"` // Don't use Id, use PasswordID() instead for consistency with MongoDB
	WebsiteID  bson.ObjectId `bson:"website_id"`
	WID        uint32        `db:"website_id" bson:"websiteid,omitempty"`
	UserName   string        `db:"user_name" bson:"user_name"`
	Password   string        `db:"password" bson:"password"`
	CreateTime time.Time     `db:"create_time" bson:"create_time"`
	UpdateTime time.Time     `db:"update_time" bson:"update_time"`
	Deleted    uint8         `db:"deleted" bson:"deleted"`
}

// UserID returns the user id
func (u *User) UserID() string {
	return u.ObjectID.Hex()
}

// PasswordID returns the password id
func (u *Passwords) PasswordID() string {
	return u.ObjectID.Hex()
}

// UserByEmail gets user information from email
func UserByEmail(email string) (User, error) {
	var err error

	result := User{}

	err = database.View("user", email, &result)
	if err != nil {
		err = ErrNoResult
	}

	return result, standardizeError(err)
}

// UserCreate creates user
func UserCreate(userName, email, phone, password string) error {
	var err error

	now := time.Now()

	user := &User{
		ObjectID:   bson.NewObjectId(),
		UserName:   userName,
		Email:      email,
		Phone:      phone,
		Password:   password,
		StatusID:   1,
		RoleID:     2,
		CreateTime: now,
		UpdateTime: now,
		Deleted:    0,
	}

	err = database.Update("user", user.Email, &user)

	return standardizeError(err)
}

// PasswordByID gets password by ID
func PasswordByID(websiteID string, passwordID string) (Passwords, error) {
	var err error

	result := Passwords{}

	err = database.View("passwords", websiteID+passwordID, &result)
	if err != nil {
		err = ErrNoResult
	}
	if result.WebsiteID != bson.ObjectIdHex(websiteID) {
		result = Passwords{}
		err = ErrUnauthorized
	}

	return result, standardizeError(err)
}

// PasswordsByWebsiteID gets all password for a website
func PasswordsByWebsiteID(websiteID string) ([]Passwords, error) {
	var err error

	var result []Passwords

	// View retrieves a record set in Bolt
	err = database.BoltDB.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte("passwords"))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Get the iterator
		c := b.Cursor()

		prefix := []byte(websiteID)
		for k, v := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = c.Next() {
			var single Passwords

			// Decode the record
			err := json.Unmarshal(v, &single)
			if err != nil {
				log.Println(err)
				continue
			}

			result = append(result, single)
		}

		return nil
	})

	return result, standardizeError(err)
}

// PasswordCreate creates a note
func PasswordCreate(userName, password, websiteID string) error {
	var err error

	now := time.Now()

	passwords := &Passwords{
		ObjectID:   bson.NewObjectId(),
		WebsiteID:  bson.ObjectIdHex(websiteID),
		UserName:   userName,
		Password:   password,
		CreateTime: now,
		UpdateTime: now,
		Deleted:    0,
	}

	err = database.Update("passwords", websiteID+passwords.ObjectID.Hex(), &passwords)

	return standardizeError(err)
}

// PasswordUpdate updates a note
func PasswordUpdate(password string, websiteID string, passwordID string) error {
	var err error

	now := time.Now()

	var passwords Passwords
	passwords, err = PasswordByID(websiteID, passwordID)
	if err == nil {
		// Confirm the owner is attempting to modify the note
		if passwords.WebsiteID.Hex() == websiteID {
			passwords.UpdateTime = now
			passwords.Password = password
			err = database.Update("passwords", websiteID+passwords.ObjectID.Hex(), &passwords)
		} else {
			err = ErrUnauthorized
		}
	}

	return standardizeError(err)
}
