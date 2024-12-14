package main

import (
	"fmt"
	"net/http"
)

var taskList = []string{
	"read a book",
	"go for run",
	"watch golang cours",
}

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":5000", nil)

}

func helloUser(writer http.ResponseWriter, reponse *http.Request) {
	var greeting = "Hello User. Welcom to our Todolist App!"
	fmt.Fprintf(writer, greeting)
}

func showTasks(writer http.ResponseWriter, respone *http.Request) {
	fmt.Fprintln(writer, taskList)
}

func printTask(taskList []string) {

	fmt.Println("list of my todos")
	for index, tasks := range taskList {
		fmt.Printf("%d: %s\n", index+1, tasks)
		fmt.Printf("%d: %s\n", index+1, tasks)
	}
}
