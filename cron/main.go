package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/robfig/cron"
)

func task() {
	fmt.Println("I am runnning task:", time.Now())
	// Platform specific code
	cmd := exec.Command("bash", "migrate.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	c := cron.New()
	c.AddFunc("@every 10s", task) // run at every second
	c.Start()
	//select {} // Keep-alive loop
	// Funcs are invoked in their own goroutine, asynchronously.
	// ... your code ...
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
