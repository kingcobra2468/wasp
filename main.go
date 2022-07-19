package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/kingcobra2468/wasp/internal/task"
)

var (
	tokenFile   string
	credsFile   string
	taskName    string
	taskDueTime time.Duration
)

func init() {
	flag.StringVar(&tokenFile, "token-file", "token.json", "path to the token file")
	flag.StringVar(&credsFile, "creds", "", "path to the credentials file")
	flag.StringVar(&taskName, "name", "", "name of the task")
	flag.DurationVar(&taskDueTime, "time", time.Duration(21*time.Hour), "name of the task")
}

func setup() {
	if tokenFile == "" {
		if val, ok := os.LookupEnv("WASP_TOKEN_FILE"); ok {
			tokenFile = val
		} else {
			log.Fatal("token file argument was never passed")
		}
	}

	if credsFile == "" {
		if val, ok := os.LookupEnv("WASP_CREDS_FILE"); ok {
			credsFile = val
		} else {
			log.Fatal("creds file argument was never passed")
		}
	}

	if taskName == "" {
		log.Fatal("task name argument was never passed")
	}
}

func main() {
	flag.Parse()
	setup()

	client, err := task.NewClient(tokenFile, credsFile)
	if err != nil {
		fmt.Println(err)
	}

	task := task.Task{Name: taskName, Due: taskDueTime, Client: client}
	if err := task.Find(); err != nil {
		log.Fatalln(err)
	}

	late, _ := task.Late()
	done, _ := task.Done()
	if !done && late {
		fmt.Println("here")
		err := beeep.Alert("Wasp", fmt.Sprintf("%s is past due.", strings.Title(taskName)), "asset/wasp.png")
		if err != nil {
			panic(err)
		}
	}
}
