package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"app/common/passhash"
	"app/common/session"
	"app/model"

	"github.com/gorilla/sessions"
)

const (
	// Name of the session variable that tracks login attempts
	sessLoginAttempt = "login_attempt"
)

// loginAttempt increments the number of login attempts in sessions variable
func loginAttempt(sess *sessions.Session) {
	// Log the attempt
	if sess.Values[sessLoginAttempt] == nil {
		sess.Values[sessLoginAttempt] = 1
	} else {
		sess.Values[sessLoginAttempt] = sess.Values[sessLoginAttempt].(int) + 1
	}
}

// LoginPOST handles the login form submission
func LoginPOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)
	// Prevent brute force login attempts by not hitting MySQL and pretending like it was invalid :-)
	// if sess.Values[sessLoginAttempt] != nil && sess.Values[sessLoginAttempt].(int) >= 5 {
	// 	log.Println("Brute force login prevented")
	// 	fmt.Fprint(w, "Sorry, no brute force - Attempt: "+fmt.Sprintf("%v", sess.Values[sessLoginAttempt]))
	// 	sess.Save(r, w)
	// 	return
	// }

	// Form values
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Get database result
	result, err := model.UserByEmail(email)

	// Determine if user exists
	if err == model.ErrNoResult {
		loginAttempt(sess)
		sess.Save(r, w)
		fmt.Fprint(w, "Username does not exist :-)")
	} else if err != nil {
		// Display error message
		log.Println(err)
		sess.Save(r, w)
		fmt.Fprint(w, "There was an error. Please try again later :-)")
	} else if passhash.MatchString(result.Password, password) {
		if result.StatusID != 1 {
			// User inactive and display inactive message
			sess.Save(r, w)
			fmt.Fprint(w, "Account is inactive so login is disabled :-)")
		} else {
			// Login successfully
			session.Empty(sess)
			log.Println("Login successful!")
			sess.Values["id"] = result.UserID()
			sess.Values["email"] = email
			sess.Values["username"] = result.UserName
			json.NewEncoder(w).Encode(result)
			sess.Save(r, w)
		}
	} else {
		loginAttempt(sess)
		fmt.Fprint(w, "Password is incorrect :-)")
		sess.Save(r, w)
	}
	return
}
