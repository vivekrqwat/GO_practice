package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivekrqwat/monogo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://vivekrawaturfsam:Golang123@golangapi.bxlu1zh.mongodb.net/?retryWrites=true&w=majority&appName=GolangApi"
const dbname = "netflix"
const colname = "watchlist"

var collection *mongo.Collection
var cd *mongo.Collection

// init()
func init() {
	clientOption := options.Client().ApplyURI(connectionstring)
	client, er := mongo.Connect(context.TODO(), clientOption)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println("connected")
	collection = client.Database(dbname).Collection(colname)
	fmt.Println("collection is genrated")
}

// inserting data in mongodb
func insertone(movie model.Netflix) {
	in, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		// log.Println("Insert error:", err)
		fmt.Println("erorr in insert")
		return
	}
	fmt.Println("insertd", in.InsertedID)

}

// update
func updateone(movie string) {
	id, _ := primitive.ObjectIDFromHex(movie)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Insert error:", err)
	}
	fmt.Println("updates", res.MatchedCount)

}

// delete
func delete(movieid string) {
	id, _ := primitive.ObjectIDFromHex(movieid)
	filter := bson.M{"_id": id}
	_, er := collection.DeleteOne(context.Background(), filter)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println("deleter")
}

// delete all
func deleteAll() {
	filter := bson.D{{}}
	res, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println("Delete error:", err)
		return
	}
	fmt.Println("Deleted count:", res.DeletedCount)
}

func find(moviid string) model.Netflix {
	id, _ := primitive.ObjectIDFromHex(moviid)
	i := collection.FindOne(context.Background(), bson.M{"_id": id})
	var movive model.Netflix
	err := i.Decode(&movive)
	if err != nil {
		log.Println("find error:", err)

	}
	fmt.Println("movie", movive)
	return movive

}

// func get all
func findall() []bson.M {
	cur, er := collection.Find(context.Background(), bson.M{})
	if er != nil {
		log.Fatal(er)
	}
	var movies []bson.M
	for cur.Next(context.Background()) {
		var mov bson.M
		err := cur.Decode(&mov)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, mov)
	}
	defer cur.Close(context.Background())
	return movies
}

// actual controller
func Getall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMov := findall()
	json.NewEncoder(w).Encode(allMov)
}

func GetByid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	allMov := find(params["id"])
	json.NewEncoder(w).Encode(allMov)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "PUT")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	insertone(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkasWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateone(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleByid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	delete(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "DELETE")

	deleteAll()
	json.NewEncoder(w).Encode("all docs deleted")

}
