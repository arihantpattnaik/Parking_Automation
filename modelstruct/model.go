package model

import (
	"fmt"
 
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Parking struct {
	ParkId        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ParkFloor     int                `json:"parkfloor,omitempty"`
	ParkAvailable bool               `json:"parkavailable"`
	ParkPrice     float32            `json:"parkprice"`
	Owner         Owner              `json:"owner"` // Add Owner field here
}
type Owner struct {
	OwnerName   string `json:"ownername"`
	OwnerNumber string `json:"ownernumber"`
}

func Ghi() {
	fmt.Println("I AM MODEL FUNCTION")
}
