package Infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	Infrastructure "api/task_manager/Infrastructure"
	domain "api/task_manager/Domain"
)

func TestJWTMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(Infrastructure.JWTMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	user := domain.User{
		ID:       1,
		Username: "testuser",
		Role:     "user",
	}

	token, _ := Infrastructure.GenerateJWT(user)

	tests := []struct {
		name       string
		token      string
		expectedCode int
	}{
		{
			name:       "Valid token",
			token:      "Bearer " + token,
			expectedCode: http.StatusOK,
		},
		{
			name:       "Missing token",
			token:      "",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:       "Invalid token",
			token:      "Bearer invalidtoken",
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
			req.Header.Set("Authorization", tt.token)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestAdminOnly(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("role", "user")
	})
	r.Use(Infrastructure.AdminOnly())
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "admin access"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/admin", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
