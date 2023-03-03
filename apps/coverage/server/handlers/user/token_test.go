package user

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/edwbaeza/coverage-api/apps/coverage/server/middlewares"
	sharedDomain "github.com/edwbaeza/coverage-api/src/shared/domain"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	userinfrastructure "github.com/edwbaeza/coverage-api/src/user/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const ENDPOINT = "/api/users/tokens"

func TestTokenHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	data := map[string]string{
		"username": "edwinbaeza05@gmail.com",
		"password": "devtest01",
	}
	mockRepo := &userinfrastructure.MockRepository{}
	user := userDomain.User{
		Email:    data["email"],
		Password: "$2a$10$F4XnOU5FyTaEaW4UvNty0.l8WG6I550UX5VwBjUH38vtef6z4VU2a",
		BaseModel: sharedDomain.BaseModel{
			ID: "FAKE_ID",
		},
	}
	mockRepo.On("FindByEmail", data["username"]).Return(user, nil)
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(user, nil)

	engine.Use(middlewares.ErrorMiddleware())
	engine.POST(ENDPOINT, TokenHandler(mockRepo))

	t.Run("Returns 200", func(t *testing.T) {
		bytesN, err := json.Marshal(data)

		req, err := http.NewRequest(http.MethodPost, ENDPOINT, bytes.NewReader(bytesN))
		require.NoError(t, err)
		req.Header["Content-Type"] = []string{"application/json"}
		req.Header["Content-Length"] = []string{strconv.Itoa(len(bytesN))}

		rec := httptest.NewRecorder()
		engine.ServeHTTP(rec, req)

		if err != nil {
			t.Error(err)
		}
		response := rec.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusOK, response.StatusCode)
		body, _ := ioutil.ReadAll(response.Body)
		assert.NotEqual(t, string(body), "")
	})
}
