package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edwbaeza/coverage-api/apps/coverage/server/middlewares"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	userInfraestructure "github.com/edwbaeza/coverage-api/src/user/infraestructure"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const ENDPOINT_REGISTRATION = "/api/users"

func TestRegistrationHandler(tGlobal *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	data := map[string]string{
		"email":                 "edwinbaeza05@gmail.com",
		"password":              "devtest01",
		"password_confirmation": "devtest01",
		"first_name":            "Edwin",
		"last_name":             "Baeza",
	}
	mockRepo := &userInfraestructure.MockRepository{}
	mockUser := userDomain.User{
		Email:     data["email"],
		Password:  "$2a$10$F4XnOU5FyTaEaW4UvNty0.l8WG6I550UX5VwBjUH38vtef6z4VU2a",
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Role:      2,
	}
	mockRepo.On("Create", mock.Anything).Return(mockUser, nil)
	mockRepo.On("FindByEmail", mock.Anything).Return(mockUser, errors.New("No documents in result"))
	engine.Use(middlewares.ErrorMiddleware())
	engine.POST(ENDPOINT_REGISTRATION, RegistrationHandler(mockRepo))
	tGlobal.Run("Returns 201", func(t *testing.T) {
		var dataReader bytes.Buffer
		json.NewEncoder(&dataReader).Encode(data)

		req, err := http.NewRequest(http.MethodPost, ENDPOINT_REGISTRATION, &dataReader)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		engine.ServeHTTP(rec, req)

		if err != nil {
			t.Error(err)
		}

		response := rec.Result()
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, http.StatusCreated, response.StatusCode)
		assert.NotEqual(t, string(body), "")
	})
}
