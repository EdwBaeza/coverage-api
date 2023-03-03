package purchase

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edwbaeza/coverage-api/apps/coverage/server/middlewares"
	purchaseDomain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	purchaseinfrastructure "github.com/edwbaeza/coverage-api/src/purchase/infrastructure"
	sharedDomain "github.com/edwbaeza/coverage-api/src/shared/domain"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	userinfrastructure "github.com/edwbaeza/coverage-api/src/user/infrastructure"
	"github.com/edwbaeza/coverage-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const ENDPOINT_FIND = "/api/purchases/"

func TestFindHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	mockUser := userDomain.User{
		Email:     "edwinbaeza05@gmail.com",
		Password:  "$2a$10$F4XnOU5FyTaEaW4UvNty0.l8WG6I550UX5VwBjUH38vtef6z4VU2a",
		FirstName: "Edwin",
		LastName:  "Baeza",
		Role:      1,
		BaseModel: sharedDomain.BaseModel{
			ID: "FAKE_ID",
		},
	}

	mockPurchase := purchaseDomain.Purchase{
		UserId: mockUser.ID,
		Items: &[]purchaseDomain.PurchaseItem{
			{ProductId: "PRODUCT_ID_1", UnitPrice: 10.0, Units: 1},
			{ProductId: "PRODUCT_ID_2", UnitPrice: 20.0, Units: 2},
		},
		BaseModel: sharedDomain.BaseModel{
			ID: "FAKE_PURCHASE_ID",
		},
	}

	userMockRepo := &userinfrastructure.MockRepository{}
	userMockRepo.On("Find", mock.Anything).Return(mockUser, nil)
	purchaseMockRepo := &purchaseinfrastructure.MockRepository{}
	purchaseMockRepo.On("Find", mock.Anything).Return(mockPurchase, nil)

	engine.Use(middlewares.ErrorMiddleware())
	engine.Use(middlewares.JwtAuthMiddleware(userMockRepo))
	engine.GET(ENDPOINT_FIND+":id", FindPurchaseHandler(purchaseMockRepo))
	t.Run("Returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, ENDPOINT_FIND+mockPurchase.ID, nil)
		token, err := utils.GenerateToken(mockUser.ID)
		req.Header["Authorization"] = []string{"Bearer " + token}
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		engine.ServeHTTP(rec, req)

		if err != nil {
			t.Error(err)
		}

		response := rec.Result()
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.NotEqual(t, string(body), "")
	})
}
