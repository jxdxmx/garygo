package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("mongodb://xingcuntian:123456@localhost:27017/xingcuntian")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("xingcuntian").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)
}
