package main

import (
	"encoding/json"
	"io/ioutil"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func getTodos() (todos []Todo) {
	fileBytes, err := ioutil.ReadFile("todos.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &todos)
	if err != nil {
		panic(err)
	}

	return todos
}

func saveTodos(todos []Todo) {
	fileBytes, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("todos.json", fileBytes, 0644)
	if err != nil {
		panic(err)
	}
}
