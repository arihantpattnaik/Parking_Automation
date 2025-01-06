package control

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/parking_automation/modelstruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB HELPERs - generally in separate File
const connectionString = "mongodb+srv://pattnaikarihant:pattnaikarihant@arihantcluster.wgbiv.mongodb.net/?retryWrites=true&w=majority&appName=Arihantcluster"
const dbName = "Parking Automation"
const colName = "parkinglot"

var collection *mongo.Collection // Corrected the type

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	} else {
		fmt.Println("COnnected to the Database")
	}
	// Assign the collection
	collection = client.Database(dbName).Collection(colName) // Corrected method names
}

func addnewslot(slot model.Parking) {
	inserted, err := collection.InsertOne(context.Background(), slot)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The ID of the inserted data is ", inserted.InsertedID)
}

// controler
func CreateNewSlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var slot model.Parking

	json.NewDecoder(r.Body).Decode(&slot)
	addnewslot(slot)
	json.NewEncoder(w).Encode(slot)
}

// DB CONTROLLER
func updateslot(parkid string) {
	id, _ := primitive.ObjectIDFromHex(parkid)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"ParkPrice": 0}}

	resp, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Update Slot print", resp.ModifiedCount)
}

// controller
func MarkUnavail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateslot(params["parkid"])

	json.NewEncoder(w).Encode(params["parkid"])
}

func deleteone(parkid string) {
	id, _ := primitive.ObjectIDFromHex(parkid)
	filter := bson.M{"_id": id}

	dlt, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The 1 parkingid that got deleted is", dlt)
}

func DeleteOneSlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	// Get parkid from the URL parameters
	params := mux.Vars(r)
	parkid := params["id"]

	// Call the deleteone function to delete the parking slot by id
	deleteone(parkid)

	// Return success message
	json.NewEncoder(w).Encode(map[string]string{"message": "Parking slot deleted successfully", "parkid": parkid})
}

func deleteall() {
	// Delete all documents in the collection
	dlt, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of parking slots deleted:", dlt.DeletedCount)
}

func DeleteAllSlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	// Call the function to delete all parking slots
	deleteall()

	// Send success response back to client
	json.NewEncoder(w).Encode(map[string]string{"message": "All parking slots deleted successfully"})
}

func getallcollection() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var parking []primitive.M

	for cur.Next(context.Background()) {
		var ans bson.M
		err := cur.Decode(&ans)
		if err != nil {
			log.Fatal(err)
		}

		parking = append(parking, ans)

	}

	defer cur.Close(context.Background())
	return parking
}

// Actual Controller
func GetAllParking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlcode")
	allpark := getallcollection()

	json.NewEncoder(w).Encode(allpark)
}
