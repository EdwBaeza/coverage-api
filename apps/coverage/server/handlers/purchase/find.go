package purchase

import (
	"net/http"

	purchaseApplication "github.com/edwbaeza/coverage-api/src/purchase/application/purchase"
	purchaseDomain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/gin-gonic/gin"
)

func FindPurchaseHandler(respository purchaseDomain.PurchaseRepository) func(context *gin.Context) {
	return func(context *gin.Context) {
		currentUser := context.MustGet("user").(userDomain.User)
		finder := purchaseApplication.NewPurchaseFinder(respository)
		purchase, err := finder.Find(context.Param("id"), currentUser)

		if err != nil {
			context.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusOK, gin.H{"purchase": purchase})
	}
}
