package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"mongoAPI/model"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "netflix"
const colName = "watchlist"

var COllection *mongo.Collection

// connect
func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	connectionString := os.Getenv("MONGO_URI")
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	checkNil(err)
	fmt.Println("success->connected")

	COllection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance is ready")
}

//MongoDB helpers

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := COllection.InsertOne(context.Background(), movie)
	checkNil(err)
	fmt.Println(inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	checkNil(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := COllection.UpdateOne(context.Background(), filter, update)
	checkNil(err)

	fmt.Println(result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	checkNil(err)

	filter := bson.M{"_id": id}
	result, err := COllection.DeleteOne(context.Background(), filter)
	checkNil(err)

	fmt.Println(result.DeletedCount)
}

// delete all
func deleteAllMovie() int64 {

	result, err := COllection.DeleteMany(context.Background(), bson.D{{}}, nil)
	checkNil(err)

	fmt.Println(result.DeletedCount)
	return result.DeletedCount
}

// get all collections
func getAllCollections() []primitive.M {
	cur, err := COllection.Find(context.Background(), bson.D{{}})
	checkNil(err)

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M

		err := cur.Decode(&movie)
		checkNil(err)

		movies = append(movies, movie)

	}
	defer cur.Close(context.Background())
	return movies
}

// actual controllers
// get all movies from DB
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllCollections()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMyOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("deleted: " + params["id"])
}

func DeleteMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode("deleted: " + strconv.Itoa(int(count)) + " movies")
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
