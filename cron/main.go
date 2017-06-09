package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/robfig/cron"
)

func task() {
	fmt.Println("I am runnning task:", time.Now())
	//
	// Platform specific code
	cmd := exec.Command("bash", "migrate.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", task) // run at every second
	c.Start()
	select {} // Keep-alive loop
	// Funcs are invoked in their own goroutine, asynchronously.
	// ... your code ...
}
