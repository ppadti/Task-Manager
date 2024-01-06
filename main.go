package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var tasks = make(map[int]Task)

const connectionString = "mongodb+srv://pushpapadti:Rajita12345@cluster0.mg6ymn4.mongodb.net/?retryWrites=true&w=majority"
const dbName = "tasks"
const colname = "title"

var collection *mongo.Collection

// connect with mongodb
func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongobd connection success")
	collection = client.Database(dbName).Collection(colname)

	//collecting instance
	fmt.Println("collection instance is ready")

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taskList := make([]Task, 0, len(tasks))
	// for _, task := range tasks {
	// 	taskList = append(taskList, task)
	// }
	// json.NewEncoder(w).Encode(taskList)

	//using mongo-db
	var tasks []Task

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task Task
		err := cursor.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}

func AddTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access=Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)

	//using system memory
	Id := rand.Intn(100) + 1
	newTask.ID = Id
	// tasks[newTask.ID] = newTask
	// json.NewEncoder(w).Encode(newTask)

	//using mongodb
	inserted, err := collection.InsertOne(context.Background(), newTask)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(newTask)
	fmt.Println("inserted one task with id", inserted.InsertedID)
}

func UpdateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	taskId := params["id"]

	var updatedTask Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	// id, err := strconv.Atoi(taskId)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode("invalid task id")
	// }
	// if _, ok := tasks[id]; ok {
	// 	tasks[id] = Task{ID: id, Name: updatedTask.Name}
	// 	json.NewEncoder(w).Encode(tasks[id])
	// } else {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode("Task not foud")
	// }

	//using mongodb
	// id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": taskId}
	update := bson.M{"$set": bson.M{"task": updatedTask.Task}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(updatedTask.ID)
	fmt.Println("Task updated:", updatedResult.UpsertedID)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-TYpe", "application/json")

	// params := mux.Vars(r)["id"]

	// // if _, ok := tasks[id]; ok {
	// // 	delete(tasks, id)
	// // 	w.WriteHeader((http.StatusOK))
	// // 	json.NewEncoder(w).Encode("Task deleted")
	// // } else {
	// // 	w.WriteHeader(http.StatusNotFound)
	// // 	json.NewEncoder(w).Encode("task not found")
	// // }

	//using mongodb
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]                   //get Parameter value as string
	_id, err := primitive.ObjectIDFromHex(params) // convert params to //mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to //specify language-specific rules for string comparison, such as //rules for lettercase
	res, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of //documents deleted
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks", AddTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", UpdateTasks).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	headers := handlers.AllowedHeaders([]string{"X-requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	fmt.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
