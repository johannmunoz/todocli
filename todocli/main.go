package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getAll := getCmd.Bool("all", false, "Get all todos")
	getDone := getCmd.Bool("done", false, "Get done todos")
	getDue := getCmd.Bool("due", false, "Get due todos")
	getId := getCmd.String("id", "", "Get todo by ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addTitle := addCmd.String("title", "", "Todo title")

	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)

	doneId := doneCmd.String("id", "", "Mark as done todo ID")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get', 'add' or 'done' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getDone, getDue, getId)
	case "add":
		HandleAdd(addCmd, addTitle)
	case "done":
		HandleDone(doneCmd, doneId)
	case "default":

	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, done *bool, due *bool, id *string) {

	getCmd.Parse(os.Args[2:])

	if !*all && !*done && !*due && *id == "" {
		fmt.Print("required subcomands")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	todos := getTodos()

	if *all {

		fmt.Printf("ID \t Title \t Done \n")

		for _, todo := range todos {
			fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Title, todo.Done)
		}
		return
	}

	if *done {
		fmt.Printf("ID \t Title \t Done \n")

		for _, todo := range todos {
			if todo.Done {
				fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Title, todo.Done)
			}
		}
		return
	}

	if *due {
		fmt.Printf("ID \t Title \t Done \n")

		for _, todo := range todos {
			if !todo.Done {
				fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Title, todo.Done)
			}
		}
		return
	}

	if *id != "" {
		fmt.Printf("ID \t Title \t Done \n")

		for _, todo := range todos {
			if strconv.Itoa(todo.Id) == *id {
				fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Title, todo.Done)
			}
		}
		return
	}
}

func HandleAdd(addCmd *flag.FlagSet, title *string) {
	addCmd.Parse(os.Args[2:])
	if *title == "" {
		fmt.Print("required title subcommand")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	todos := getTodos()

	todo := Todo{
		Id:    len(todos) + 1,
		Title: *title,
		Done:  false,
	}

	todos = append(todos, todo)

	saveTodos(todos)
}

func HandleDone(doneCmd *flag.FlagSet, id *string) {
	doneCmd.Parse(os.Args[2:])
	if *id == "" {
		fmt.Print("required id subcommand")
		doneCmd.PrintDefaults()
		os.Exit(1)
	}

	todos := getTodos()

	for index, todo := range todos {
		if strconv.Itoa(todo.Id) == *id {
			todos[index].Done = true
			saveTodos(todos)
		}
	}

}
