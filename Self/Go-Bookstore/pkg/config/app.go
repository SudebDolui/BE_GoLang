package config

import (
	// "context"
	// "fmt"

	"github.com/SudebDolui/Go-Bookstore/pkg/utils"
	"github.com/jinzhu/gorm"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// const connectionString = "mongodb+srv://Demo:Demo1@cluster0.6422n.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
// const dbName = "bookstore"
// const colName = "books"

// var collection *mongo.Collection

// func init() {
// 	clientOption := options.Client().ApplyURI(connectionString)
// 	client, err := mongo.Connect(context.TODO(), clientOption)

// 	utils.ErrorUsingLog(err)

// 	collection = client.Database(dbName).Collection(colName)
// 	fmt.Println("Collection instance is Ready")
// }

//IMP: github.com/jinzhu/gorm , github.com/jinzhu/dialects/mysql

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=true")
	utils.ErrorUsingPanic(err)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
