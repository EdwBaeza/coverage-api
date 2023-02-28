package purchase

import (
	"net/http"

	purchaseDomain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	userDomain "github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/gin-gonic/gin"
)

func ListPurchasesHandler(respository purchaseDomain.PurchaseRepository) func(context *gin.Context) {
	return func(context *gin.Context) {
		// TODO: Add pagination
		filters := map[string]interface{}{}
		currentUser := context.MustGet("user").(userDomain.User)

		if currentUser.IsSuperAdmin() {
			filters["user_id"] = context.Query("user_id")
		} else {
			filters["user_id"] = currentUser.ID
		}

		purchases, _ := respository.List(filters)
		context.JSON(http.StatusOK, gin.H{"purchases": purchases})
	}
}
