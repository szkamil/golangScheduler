// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/carlescere/scheduler"
// )

// func main() {
// 	job := func() {
// 		fmt.Println("Time's up!")
// 	}
// 	scheduler.Every(5).Seconds().Run(job)
// }

// func sched() {

// 	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(responseData))
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/carlescere/scheduler"
)

func main() {
	job := func() {
		t := time.Now()
		fmt.Println("Time's up! @", t.UTC())
		test()
		getConfig()
	}
	// Run every 2 seconds but not now.
	scheduler.Every(2).Seconds().NotImmediately().Run(job)

	// Run now and every X.
	scheduler.Every(5).Minutes().Run(job)
	scheduler.Every().Day().Run(job)
	scheduler.Every().Monday().At("08:30").Run(job)

	// Keep the program from not exiting.
	runtime.Goexit()
}
func test() {
	fmt.Println("calling test function")
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	fmt.Println(string("worked"))
}

func getConfig() {
	type Configuration struct {
		Users  []string
		Groups []string
	}

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Users) // output: [UserA, UserB]
	fmt.Println(configuration.Groups)

}
