package session

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

// generateRandomBytes generates a specified number of random bytes.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// generateSessionToken generates a random session token using base64 encoding.
func generateSessionToken() (string, error) {
	bytes, err := generateRandomBytes(32) // Adjust token length as per your requirements
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// createSessionCookie creates a new session cookie.
func createSessionCookie(response http.ResponseWriter, sessionToken string) {
	var cookie = http.Cookie{
		Name:     "my_little_session_cookie",
		Value:    sessionToken, // Set doing runtime
		Path:     "/",
		Expires:  time.Time{},
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(response, &cookie)
}

// deleteSessionCookie deletes the session cookie by setting its expiration to the past.
func deleteSessionCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "my_little_session_cookie",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
}

// Active retrieves the session token from the request cookie.
func Active(response http.ResponseWriter, request *http.Request) (string, error) {
	_, err := request.Cookie("my_little_session_cookie")
	if err != nil {
		return "inactive", err
	}
	return "active", nil
}

func Create(response http.ResponseWriter, r *http.Request, userID string) (any, error) {
	// Generate session token
	sessionToken, err := generateSessionToken()
	if err != nil {
		return nil, err
	}

	// Create session cookie
	createSessionCookie(response, sessionToken)
	return "Started session", nil
}

func Close(w http.ResponseWriter, r *http.Request) (any, error) {
	// Delete session cookie
	deleteSessionCookie(w)
	return "Closed session", nil
}
