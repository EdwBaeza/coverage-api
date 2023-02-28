package purchase

import (
	"net/http"

	purchaseApplication "github.com/edwbaeza/coverage-api/src/purchase/application/purchase"
	"github.com/edwbaeza/coverage-api/src/purchase/domain"
	purchaseDomain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/gin-gonic/gin"
)

type PurchaseItemRequest struct {
	ProductId   string  `binding:"required" key:"product_id" json:"product_id,omitempty" binding:"required"`
	ProductSize int32   `binding:"required" key:"product_size" json:"product_size,omitempty" binding:"required,oneof=1 2 3"`
	Units       int32   `binding:"required" key:"units" json:"units,omitempty" binding:"required"`
	UnitPrice   float64 `binding:"required" key:"unit_price" json:"unit_price,omitempty" binding:"required"`
}

type AddressRequest struct {
	RawAddress   string `binding:"required" json:"raw_address,omitempty" key:"raw_address,omitempty"`
	Neighborhood string `binding:"required" json:"neighborhood,omitempty" key:"neighborhood,omitempty"`
	Municipality string `binding:"required" json:"municipality,omitempty" key:"municipality,omitempty"`
	State        string `binding:"required" json:"state,omitempty" key:"state,omitempty"`
	Country      string `binding:"required" json:"country,omitempty" key:"country,omitempty"`
	ZipCode      string `binding:"required" json:"zip_code,omitempty" key:"zip_code,omitempty"`
	IntNumber    string `json:"int_number,omitempty" key:"int_number,omitempty"`
	ExtNumber    string `binding:"required" json:"ext_number,omitempty" key:"ext_number,omitempty"`
	Latitude     string `binding:"required" json:"latitude,omitempty" key:"latitude,omitempty"`
	Longitude    string `binding:"required" json:"longitude,omitempty" key:"longitude,omitempty"`
}

type PurchaseRequest struct {
	UserId             string                 `key:"user_id" json:"user_id,omitempty" binding:"omitempty"`
	SourceAddress      AddressRequest         `key:"source_address" json:"source_address,omitempty" binding:"required"`
	DestinationAddress AddressRequest         `key:"destination_address" json:"destination_address,omitempty" binding:"required"`
	Items              *[]PurchaseItemRequest `key:"items" json:"items,omitempty" binding:"required"`
}

func CreatePurchaseHandler(respository purchaseDomain.PurchaseRepository) func(context *gin.Context) {
	return func(context *gin.Context) {
		requestPurchaseNewTest := &PurchaseRequest{}
		err := context.ShouldBind(requestPurchaseNewTest)
		if err != nil {
			context.Error(err).SetType(gin.ErrorTypeBind).SetMeta(http.StatusBadRequest)
			return
		}
		// Maybe is better personification
		user := context.MustGet("user").(userDomain.User)

		if !user.IsSuperAdmin() {
			requestPurchaseNewTest.UserId = user.ID
		}

		creator := purchaseApplication.NewPurchaseCreator(respository)
		purchaseCreated, err := creator.Create(
			domain.Purchase{
				UserId: requestPurchaseNewTest.UserId,
				SourceAddress: &domain.Address{
					RawAddress:   requestPurchaseNewTest.SourceAddress.RawAddress,
					Neighborhood: requestPurchaseNewTest.SourceAddress.Neighborhood,
					Municipality: requestPurchaseNewTest.SourceAddress.Municipality,
					State:        requestPurchaseNewTest.SourceAddress.State,
					Country:      requestPurchaseNewTest.SourceAddress.Country,
					ZipCode:      requestPurchaseNewTest.SourceAddress.ZipCode,
					IntNumber:    requestPurchaseNewTest.SourceAddress.IntNumber,
					ExtNumber:    requestPurchaseNewTest.SourceAddress.ExtNumber,
					Latitude:     requestPurchaseNewTest.SourceAddress.Latitude,
					Longitude:    requestPurchaseNewTest.SourceAddress.Longitude,
				},
				DestinationAddress: &domain.Address{
					RawAddress:   requestPurchaseNewTest.DestinationAddress.RawAddress,
					Neighborhood: requestPurchaseNewTest.DestinationAddress.Neighborhood,
					Municipality: requestPurchaseNewTest.DestinationAddress.Municipality,
					State:        requestPurchaseNewTest.DestinationAddress.State,
					Country:      requestPurchaseNewTest.DestinationAddress.Country,
					ZipCode:      requestPurchaseNewTest.DestinationAddress.ZipCode,
					IntNumber:    requestPurchaseNewTest.DestinationAddress.IntNumber,
					ExtNumber:    requestPurchaseNewTest.DestinationAddress.ExtNumber,
					Latitude:     requestPurchaseNewTest.DestinationAddress.Latitude,
					Longitude:    requestPurchaseNewTest.DestinationAddress.Longitude,
				},
				Items: mapToPurchaseItems(requestPurchaseNewTest.Items),
			},
		)

		if err != nil {
			context.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusCreated, gin.H{"data": purchaseCreated})
	}
}

func mapToPurchaseItems(items *[]PurchaseItemRequest) *[]domain.PurchaseItem {
	purchaseItems := []domain.PurchaseItem{}
	for _, item := range *items {
		purchaseItems = append(purchaseItems, domain.PurchaseItem{
			ProductId:   item.ProductId,
			ProductSize: purchaseDomain.ProductSize(item.ProductSize),
			UnitPrice:   item.UnitPrice,
			Units:       item.Units,
		})
	}
	return &purchaseItems
}
