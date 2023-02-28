package status

import (
	"net/http"

	purchaseStatusApplication "github.com/edwbaeza/coverage-api/src/purchase/application/purchase/status"
	purchaseDomain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	Status int `key:"status" json:"status,omitempty" binding:"required,oneof=1 2 3 4 5 6"`
}

func UpdatePurchaseStatusHandler(respository purchaseDomain.PurchaseRepository) func(context *gin.Context) {
	return func(context *gin.Context) {
		purchaseRequest := PurchaseRequest{}

		if err := context.ShouldBind(&purchaseRequest); err != nil {
			context.Error(err).SetType(gin.ErrorTypeBind).SetMeta(http.StatusBadRequest)
			return
		}

		currentUser := context.MustGet("user").(userDomain.User)
		updater := purchaseStatusApplication.NewPurchaseStatusUpdater(respository)
		purchase, err := updater.Update(context.Param("id"), purchaseRequest.Status, currentUser)

		if err != nil {
			context.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": purchase})
	}
}
