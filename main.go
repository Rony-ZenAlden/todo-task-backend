package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task1 := Tasks{
		ID:         "1",
		TaskName:   "New Projects",
		TaskDetail: "Perfect",
		Date:       "2024-01-22",
	}
	tasks = append(tasks, task1)

	task2 := Tasks{
		ID:         "2",
		TaskName:   "Power Projects",
		TaskDetail: "Very Good",
		Date:       "2024-02-11",
	}
	tasks = append(tasks, task2)

	task3 := Tasks{
		ID:         "3",
		TaskName:   "Hello Projects",
		TaskDetail: "Nice",
		Date:       "2023-04-99",
	}
	tasks = append(tasks, task3)
	fmt.Println("Your Tasks Are: ", tasks)
}

//////////////////////////////  Routing  //////////////////////////////

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082", router))
}

///////////////////////////////////////////////////////////////////////

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Baby")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	flag := false
	for _, task := range tasks {
		if task.ID == taskId {
			json.NewEncoder(w).Encode(task)
			flag = true
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(100000000))
	currentTime := time.Now().Format("01-02-2020")
	task.Date = currentTime
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskID := mux.Vars(r)["id"]
	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			// json.NewEncoder(w).Encode(map[string]string{"status": "Task deleted successfully"})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "Error: Task not found"})
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.ID == params["id"] {
			fmt.Println("THe id is : ", item.ID)
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Tasks
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(tasks, task)
			flag = true
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func main() {
	allTasks()
	fmt.Println("Hello Flutter")
	handleRoutes()
}
