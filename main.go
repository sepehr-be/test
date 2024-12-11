package main

import (
	"fmt"
	"net/http"

)


firstTask := "read a book"
secondTask := "go for run"
thirdTask := "watch golang cours"

var taskList  = []string { firstTask, secondTask, thirdTask}


func main() {

	http.HandleFunc("/",helloUser)
	http.HandleFunc("/show-tasks",showTasks)

	http.ListenAndServe(":5000",nil)

	// addTask(taskList,"watching tv")
}

func helloUser(http.ResponseWriter, *http.Request)  {
	var greeting = "Hello User. Welcom to our Todolist App!"
	fmt.Fprintf(writer, greeting)
}

func showTasks(http.ResponseWriter, *http.Request)  {
	fmt.Fprintf(writer, printTask)
}


func printTask(taskList []string)  {
	
	fmt.Println("list of my todos")
	for index, tasks := range taskList {
		fmt.Printf("%d: %s\n",index+1 ,tasks)
	}
}

// func addTask(taskList []string, newTask string) []string {
// 	var updatedList = append(taskList, newTask)
// 	return updatedList
	
// }