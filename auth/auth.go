package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"samla-admin/clerk"
	"strings"
	"time"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	clerkjwt "github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/joho/godotenv"
)

// OrganizationIDKey is the context key for storing organization ID
type OrganizationIDKey struct{}

// GetOrganizationID retrieves the organization ID from the request context
func GetOrganizationID(r *http.Request) (string, bool) {
	organizationID, ok := r.Context().Value(OrganizationIDKey{}).(string)
	return organizationID, ok
}

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}

// authorizedOrgId is the organization ID that is authorized to access the API
var authorizedOrgId string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	authorizedOrgId = os.Getenv("AUTHORIZED_ORG_ID")
	if authorizedOrgId == "" {
		log.Fatalf("[AUTH] ERROR: AUTHORIZED_ORG_ID is not set")
	}
}

// VerifyingMiddleware is the general middleware that verifies the passed JWT Token from clerk and extracts the user ID and organization ID to pass it to the next handler
func VerifyingMiddleware(next http.Handler) http.Handler {
	return clerkhttp.RequireHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[AUTH] Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		startTime := time.Now()

		// Log authorization header presence (without revealing the token)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Printf("[AUTH] ERROR: Missing authorization header for %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("[AUTH] Authorization header present for %s %s", r.Method, r.URL.Path)

		userID, err := extractUserIDFromAuthHeader(r)
		if err != nil {
			log.Printf("[AUTH] ERROR: Failed to extract user ID for %s %s: %v", r.Method, r.URL.Path, err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("[AUTH] Successfully extracted user ID: %s for %s %s", userID, r.Method, r.URL.Path)

		organizationID, err := clerk.GetUserOrganizationId(userID)
		if err != nil {
			log.Printf("[AUTH] ERROR: Failed to get organization ID for user %s on %s %s: %v", userID, r.Method, r.URL.Path, err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("[AUTH] Successfully retrieved organization ID: %s for user %s on %s %s", organizationID, userID, r.Method, r.URL.Path)

		if organizationID == "" || organizationID != authorizedOrgId {
			log.Printf("[AUTH] ERROR: Organization ID is not authorized for user %s on %s %s", userID, r.Method, r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Add organization ID to request context
		ctx := context.WithValue(r.Context(), OrganizationIDKey{}, organizationID)
		r = r.WithContext(ctx)

		// Wrap the response writer to capture the status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("[AUTH] Response: %s %s -> STATUS: %d completed in %v (User: %s, Org: %s)", r.Method, r.URL.Path, rw.statusCode, time.Since(startTime), userID, organizationID)
	}))
}

// extractUserIDFromAuthHeader extracts the user ID from the Authorization header
func extractUserIDFromAuthHeader(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing authorization header")
	}

	// Check if it's a Bearer token
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header format")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Verify the JWT token and extract claims
	claims, err := clerkjwt.Verify(context.Background(), &clerkjwt.VerifyParams{
		Token: token,
	})
	if err != nil {
		return "", fmt.Errorf("failed to verify token: %v", err)
	}

	// Extract user ID from the subject claim
	userID := claims.RegisteredClaims.Subject
	if userID == "" {
		return "", fmt.Errorf("no user ID found in token")
	}

	return userID, nil
}
