package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page! %s", r.URL.Path[1:])
}

func main() {

	p := Person{
		"Angad",
		"123456",
	}

	fmt.Println(p)

	session, err := mgo.Dial("gomongodb")
	if err != nil {
		fmt.Println("mongogb just dbname err: ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("testdb").C("people")
	_ = c.Insert(
		&Person{"Alex", "+55 89 9556 9111"},
		&Person{"John", "+55 98 8402 3256"})

	result := Person{}
	_ = c.Find(bson.M{"name": "Alex"}).One(&result)

	fmt.Println("Phone:", result.Phone)

	// start http web server
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
