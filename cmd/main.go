package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sleepytaka/graph_tutorial_go_sample/internal/apps"
)

func main() {
	fmt.Println("Go Graph Tutorial")
	appId := os.Getenv("APP_ID")
	scopes := strings.Split(os.Getenv("SCOPES"), ",")

	app := apps.New(appId, scopes)
	user, err := app.GetMe()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s!\n", user.DisplayName)

	choice := -1
	for choice != 0 {
		fmt.Println("Please choose one of the following options:")
		fmt.Println("0. Exit")
		fmt.Println("1. Display access token")
		fmt.Println("2. View this week's calendar")
		fmt.Println("3. Add an event")

		choice = -1
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			s := scanner.Text()
			i, err := strconv.Atoi(s)
			if err == nil {
				choice = i
			}
		}

		switch choice {
		case 0:
			fmt.Println("Goodbye...")
		case 1:
			// Display access token
			fmt.Println("Access token is " + app.GetAccessToken())
		case 2:
			// List the calendar
			err := app.ListCalendarEvents()
			if err != nil {
				fmt.Println(err.Error())
			}
		case 3:
			// Create a new event
			subject := getUserInput("subject")
			fmt.Printf("> %v\n", subject)
			start := getUserInput("start ex) 2021-01-02 15:04:05 ")
			t1, err := time.Parse("2006-01-02 15:04:05", start)
			if err != nil {
				t1 = time.Now().UTC()
			}
			fmt.Printf("> %v\n", t1)

			end := getUserInput("end ex) 2021-01-02 15:04:05")
			t2, err := time.Parse("2006-01-02 15:04:05", end)
			if err != nil {
				t2 = time.Now().UTC().AddDate(0, 0, 1)
			}
			fmt.Printf("> %v\n", t2)

			content := getUserInput("content")
			fmt.Printf("> %v\n", content)

			err = app.CreateEvent(subject, t1, t2, content)
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}

func getUserInput(fieldName string) string {
	fmt.Printf("%s:", fieldName)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
