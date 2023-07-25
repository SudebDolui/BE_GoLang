package config

import (
	"context"
	"encoding/json"
	"fmt"

	misc "github.com/SudebDolui/CollegeApi/Packages/Misc"
	models "github.com/SudebDolui/CollegeApi/Packages/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Demo:Demo1@cluster0.6422n.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "college"
const colName = "students"

var collection *mongo.Collection

func Connect() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	misc.ErrorUsingFatal(err)
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection Instance is Ready")
}

func HelperRegisterStudent(student models.Students) {
	register, err := collection.InsertOne(context.Background(), student)
	misc.ErrorUsingFatal(err)
	fmt.Println("Registered Student with Roll Number:", register.InsertedID)
}

func HelperGetStudents() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	misc.ErrorUsingFatal(err)
	var students []primitive.M

	for cur.Next(context.Background()) {
		var student bson.M
		err := cur.Decode(&student)
		misc.ErrorUsingFatal(err)
		students = append(students, student)
	}

	defer cur.Close(context.Background())
	return students
}

func HelperGetStudent(rollnumber string) models.Students {
	id, _ := primitive.ObjectIDFromHex(rollnumber)
	filter := bson.M{"_id": id}
	var result models.Students
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	misc.ErrorUsingFatal(err)
	fmt.Println("Found Document", result)
	return result
}

func HelperUpdateStudent(student *models.Students, rollnumber string) []byte {
	id, _ := primitive.ObjectIDFromHex(rollnumber)
	filter := bson.M{"_id": id}
	update := &student
	result, err := collection.UpdateByID(context.Background(), filter, *update)
	misc.ErrorUsingFatal(err)
	res, err := json.Marshal(*update)
	misc.ErrorUsingFatal(err)
	fmt.Println("Update done on:", result.UpsertedID)
	return res
}

func HelperDeleteStudents() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	misc.ErrorUsingFatal(err)
	return result.DeletedCount
}

func HelperDeleteStudent(rollnumber string) int64 {
	id, _ := primitive.ObjectIDFromHex(rollnumber)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	misc.ErrorUsingFatal(err)
	return result.DeletedCount
}
