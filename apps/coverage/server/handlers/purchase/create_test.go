package purchase

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/edwbaeza/coverage-api/apps/coverage/server/middlewares"
	purchaseDomain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	purchaseInfraestructure "github.com/edwbaeza/coverage-api/src/purchase/infraestructure"
	sharedDomain "github.com/edwbaeza/coverage-api/src/shared/domain"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	userInfraestructure "github.com/edwbaeza/coverage-api/src/user/infraestructure"
	"github.com/edwbaeza/coverage-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const ENDPOINT_CREATE = "/api/purchases"

func TestCreateHandler(tGlobal *testing.T) {
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
		Status: 1,
		UserId: mockUser.ID,
		Items: &[]purchaseDomain.PurchaseItem{
			{ProductId: "PRODUCT_ID_1", UnitPrice: 10.0, Units: 1, ProductSize: 1},
			{ProductId: "PRODUCT_ID_2", UnitPrice: 20.0, Units: 2, ProductSize: 1},
		},
		BaseModel: sharedDomain.BaseModel{
			ID:        "FAKE_PURCHASE_ID",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}

	purchaseRequest := PurchaseRequest{
		SourceAddress: AddressRequest{
			RawAddress:   "Calle 1",
			Latitude:     "19.999999999",
			Longitude:    "19.999999999",
			Municipality: "Municipio 1",
			State:        "Estado 1",
			Country:      "Pais 1",
			ZipCode:      "12345",
			ExtNumber:    "123",
			Neighborhood: "Colonia 1",
		},
		DestinationAddress: AddressRequest{
			RawAddress:   "Calle 1",
			Latitude:     "19.999999999",
			Longitude:    "19.999999999",
			Municipality: "Municipio 1",
			State:        "Estado 1",
			Country:      "Pais 1",
			ZipCode:      "12345",
			ExtNumber:    "123",
			Neighborhood: "Colonia 1",
		},
		Items: &[]PurchaseItemRequest{
			{ProductId: "PRODUCT_ID_1", UnitPrice: 10.1, Units: 1, ProductSize: 1},
			{ProductId: "PRODUCT_ID_2", UnitPrice: 20.1, Units: 2, ProductSize: 1},
		},
	}
	userMockRepo := &userInfraestructure.MockRepository{}
	userMockRepo.On("Find", mock.Anything).Return(mockUser, nil)

	purchaseMockRepo := &purchaseInfraestructure.MockRepository{}
	purchaseMockRepo.On("Create", mock.Anything).Return(mockPurchase, nil)

	engine.Use(middlewares.ErrorMiddleware())
	engine.Use(middlewares.JwtAuthMiddleware(userMockRepo))
	engine.POST(ENDPOINT_CREATE, CreatePurchaseHandler(purchaseMockRepo))
	tGlobal.Run("Returns 201", func(t *testing.T) {
		bytesN, err := json.Marshal(purchaseRequest)

		req, err := http.NewRequest(http.MethodPost, ENDPOINT_CREATE, bytes.NewReader(bytesN))
		token, err := utils.GenerateToken(mockUser.ID)
		req.Header["Authorization"] = []string{"Bearer " + token}
		req.Header["Content-Type"] = []string{"application/json"}
		req.Header["Content-Length"] = []string{strconv.Itoa(len(bytesN))}

		require.NoError(t, err)

		rec := httptest.NewRecorder()
		engine.ServeHTTP(rec, req)

		if err != nil {
			t.Error(err)
		}
		response := rec.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusCreated, response.StatusCode)
	})
}
