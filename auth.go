package main

import(
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var clientId = r.URL.Query().Get("clientId")  // Get clientId from URL query

		clientProfile, ok := database[clientId]  // Check if user exists in the database
		if !ok || clientId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)  // Return error if not found
			return
		}
		token := r.Header.Get("Authorization")
		if !isValidToken(clientProfile, token){
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func isValidToken(clientProfile ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer") {
		return strings.TrimPrefix(token, "Bearer") == clientProfile.Token
	}
	return false

}