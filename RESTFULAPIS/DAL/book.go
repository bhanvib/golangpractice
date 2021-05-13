package DAL

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Id     int
	Name   string `json:"Book Name,"`
	Isbn   string
	Author string `json:"-"`  //ignore field
	Pages  int `json:"myName,omitempty"` 
}

func InsertBook(book Book) {
	session := Connect()
	collection := session.Database("webinardb").Collection("books") //
	result, err := collection.InsertOne(context.TODO(), book)
	log.Println(result)
	log.Println(err)
}

func GetBookByIsbn(isbn string) Book {
	var book Book
	session := Connect()
	collection := session.Database("webinardb").Collection("books") // from a given db access books collection(Table)
	filter := bson.M{"isbn": isbn}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		log.Println(err)
	}
	return book
}

func GetBooks() []Book {

	var books []Book

	// //var book = Book{Id:10}

	// books = append(books, Book{Id: 10, Name: "JS"})
	// books = append(books, Book{Id: 11, Name: "GoLang"})
	// books = append(books, Book{Id: 12, Name: "Node.js"})

	// return books

	session := Connect()

	collection := session.Database("webinardb").Collection("books") // from a given db access books collection(Table)

	cur, _ := collection.Find(context.TODO(), bson.M{}) // emtpy filter (select * from books)

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) { // Iterating through records

		var book Book
		//COnvert mongo DB record into book object
		err := cur.Decode(&book) // convert book document(record) to Book struct object

		if err != nil {
			log.Fatal(err)
		}

		//Add book object to Slice
		books = append(books, book)
	}
	return books

}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected...")
	return client
}
