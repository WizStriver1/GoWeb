package controller

import (
	"fmt"
	"log"
	"net/http"

	"app/common/passhash"
	"app/common/session"
	"app/model"
)

// RegisterPOST handles the registration form submission
func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Prevent brute force login attempts by not hitting MySQL and pretending like it was invalid :-)
	// if sess.Values["register_attempt"] != nil && sess.Values["register_attempt"].(int) >= 5 {
	// 	log.Println("Brute force register prevented")
	// 	http.Redirect(w, r, "/register", http.StatusFound)
	// 	return
	// }
	// type User struct {
	// 	ObjectID   bson.ObjectId `bson:"_id"`
	// 	ID         uint32        `db:"id" bson:"id,omitempty"`     // Don't use Id, use UserID() instead for consistency with MongoDB
	// 	UserName   string        `db:"user_name" bson:"user_name"` // 用户名
	// 	Email      string        `db:"email" bson:"email"`
	// 	Phone      string        `db:"phone" bson:"phone"`
	// 	Password   string        `db:"password" bson:"password"`
	// 	CreateTime time.Time     `db:"create_time" bson:"create_time"`
	// 	UpdateTime time.Time     `db:"update_time" bson:"update_time"`
	// 	StatusID   uint8         `db:"status_id" bson:"status_id"`
	// 	RoleID     uint8         `db:"role_id" bson:"role_id"`
	// 	Deleted    uint8         `db:"deleted" bson:"deleted"` // deleted contains data delete status (0 deleted/1 stored)
	// }
	// Validate with required fields
	if validate, missingField := Validate(r, []string{"username", "email", "password"}); !validate {
		fmt.Fprint(w, "Field missing: "+missingField)
		sess.Save(r, w)
		return
	}

	// Get form values
	userName := r.FormValue("username")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	password, errp := passhash.HashString(r.FormValue("password"))

	// If password hashing failed
	if errp != nil {
		log.Println(errp)
		fmt.Fprint(w, "An error occurred on the server. Please try again later.")
		sess.Save(r, w)
		return
	}

	// Get database result
	_, err := model.UserByEmail(email)

	if err == model.ErrNoResult { // If success (no user exists with that email)
		ex := model.UserCreate(userName, email, phone, password)
		// Will only error if there is a problem with the query
		if ex != nil {
			log.Println(ex)
			fmt.Fprint(w, "An error occurred on the server. Please try again later.")
			sess.Save(r, w)
		} else {
			fmt.Fprint(w, "Account created successfully for: "+email)
			sess.Save(r, w)
			return
		}
	} else if err != nil { // Catch all other errors
		log.Println(err)
		fmt.Fprint(w, "Account created successfully for: "+email)
		sess.Save(r, w)
	} else { // Else the user already exists
		fmt.Fprint(w, "Account already exists for: "+email)
		sess.Save(r, w)
	}

	return
}
