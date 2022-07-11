package middleware

import (
	"net/http"
	"regexp"
	"strings"
)

//CheckAccountValid ...
func CheckAccountValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		address := strings.ToLower(r.Form.Get("address"))
		re := regexp.MustCompile("^0x[0-9a-f]{40}$")
		if !re.MatchString(address) {
			http.Error(w, "Account incorect", http.StatusBadRequest)
			return
		}
		r.Form["address"] = []string{address}
		next.ServeHTTP(w, r)
	})
}
