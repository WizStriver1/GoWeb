package controller

import "net/http"

// Validate returns true if all the required form values are passed
func Validate(req *http.Request, required []string) (bool, string) {
	for _, v := range required {
		if req.FormValue(v) == "" {
			return false, v
		}
	}

	return true, ""
}
