package domain

import (
	"time"

	shared "github.com/edwbaeza/coverage-api/src/shared/domain"
)

type Address struct {
	RawAddress       string `json:"raw_address,omitempty" bson:"raw_address,omitempty"`
	Neighborhood     string `json:"neighborhood,omitempty" bson:"neighborhood,omitempty"`
	Municipality     string `json:"municipality,omitempty" bson:"municipality,omitempty"`
	State            string `json:"state,omitempty" bson:"state,omitempty"`
	Country          string `json:"country,omitempty" bson:"country,omitempty"`
	ZipCode          string `json:"zip_code,omitempty" bson:"zip_code,omitempty"`
	IntNumber        string `json:"int_number,omitempty" bson:"int_number,omitempty"`
	ExtNumber        string `json:"ext_number,omitempty" bson:"ext_number,omitempty"`
	Latitude         string `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude        string `json:"longitude,omitempty" bson:"longitude,omitempty"`
	shared.BaseModel `bson:"inline"`
}

type PurchaseStatus int

const (
	Created     PurchaseStatus = 1
	Collected                  = 2
	InWarehouse                = 3
	OnRoute                    = 4
	Delivered                  = 5
	Canceled                   = 6
)

type PurchaseItem struct {
	ProductId        string      `json:"product_id,omitempty" bson:"product_id,omitempty"`
	ProductSize      ProductSize `json:"product_size,omitempty" bson:"product_size,omitempty"`
	TotalPrice       float64     `json:"total_price,omitempty" bson:"total_price,omitempty"`
	UnitPrice        float64     `json:"unit_price,omitempty" bson:"unit_price,omitempty"`
	Units            int32       `json:"units,omitempty" bson:"units,omitempty"`
	shared.BaseModel `bson:"inline"`
}

type Purchase struct {
	Status             PurchaseStatus `json:"status,omitempty" bson:"status,omitempty"`
	TotalPrice         float64        `json:"total_price,omitempty" bson:"total_price,omitempty"`
	SourceAddress      *Address       `json:"source_address,omitempty" bson:"source_address,omitempty"`
	DestinationAddress *Address       `json:"destination_address,omitempty" bson:"destination_address,omitempty"`
	shared.BaseModel   `bson:"inline"`
	Items              *[]PurchaseItem `json:"items,omitempty" bson:"items,omitempty"`
	UserId             string          `json:"user_id,omitempty" bson:"user_id,omitempty"`
}

type ProductSize int

const (
	S ProductSize = 1
	M             = 2
	L             = 3
)

func NewPurchase() Purchase {
	time_now := time.Now()
	return Purchase{
		Status: Created,
		BaseModel: shared.BaseModel{
			CreatedAt: time_now,
			UpdatedAt: time_now,
		},
	}
}

type PurchaseRepository interface {
	Create(purchase Purchase) (Purchase, error)
	Find(id string) (Purchase, error)
	List(map[string]interface{}) ([]Purchase, error)
	UpdateStatus(id string, status PurchaseStatus) (Purchase, error)
}
