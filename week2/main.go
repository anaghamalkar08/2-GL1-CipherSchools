package main

import (
	"log"
	"os"
)

/*
import (
	"fmt"
	"log"

	"github.com/anaghamalkar08/2-GL1-CipherSchools/database"
	"github.com/anaghamalkar08/2-GL1-CipherSchools/routers"
)

func init() {
	database.Setup()
}
func init() {
	//database.Setup()
	fmt.Print(1)
}*/

type User struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}

type something struct {
	Name    string    `json:"name`
	Married bool      `json:"married"`
	Age     float64   `json:"age`
	Address []Address `json:"address"`
	Marks   []int     `json:"marks"`
}

type Address struct {
	Street int    `json:"street"`
	City   string `json:"city"`
}

func main() {
	// database.Setup()           //establish database
	// engine := routers.Router() //customised engine
	// err := engine.Run("127.0.0.1:8080")
	// if err != nil {
	// log.Fatal(err)
	// }
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	//	jsonByteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	// var something Something
	// json.Unmarshal(jsonByteValues, &something)
	// log.Println(something)
	// fmt.Println(string(jsonByteValues))
	// newJsonByteValues, err := json.Marshal(something)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//os.WriteFile("some.json",newJsonByteValues)
	//consuming rest api in go code
	//returned value?= json

}
