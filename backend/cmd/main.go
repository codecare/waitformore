package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
	"waitformore/internal/waitformore"
)

func main() {

	connect := waitformore.MakeDbConnect()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connect.Hostname, connect.Port, connect.User, connect.Password, connect.Database)
	db, err := sql.Open("postgres", psqlInfo)


	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Second * 5)

	defer db.Close()

	// repository := waitformore.RamQueueRepository{}

	repository := waitformore.PostgresqlQueueRepository{DatabaseConnectionPool: db}

	router := mux.NewRouter()
	router.HandleFunc("/api/user/{userKey}", MakeUserViewHandler(&repository)).Methods("GET")
	router.HandleFunc("/api/user/{userKey}/pull", MakeUserPullHandler(&repository)).Methods("GET")
	router.HandleFunc("/api/queue", MakeNewQueueHandler(&repository)).Methods("POST")
	router.HandleFunc("/api/queue/{adminKey}", MakeAdminViewHandler(&repository)).Methods("GET")
	router.HandleFunc("/api/queue/{adminKey}/call", MakeAdminCallHandler(&repository)).Methods("GET")
	router.HandleFunc("/api/startus", MakeStatusHandler(&repository)).Methods("GET")

	if waitformore.ReadBoolFromEnv("Activate_CORS", false) {
		httpErr := http.ListenAndServe(":8081", handlers.CORS()(router))
		if httpErr != nil {
			panic(httpErr)
		}
	} else {
		httpErr := http.ListenAndServe(":8081", router)
		if httpErr != nil {
			panic(httpErr)
		}
	}
}

type QueueTitleRequest struct {
	QueueTitle string
}

func MakeNewQueueHandler(repository waitformore.QueueRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var queueTitleRequest QueueTitleRequest

		if err := decoder.Decode(&queueTitleRequest); err != nil {
			fmt.Printf("failed to parse request %s\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newQueue, repositoryError := repository.CreateNewQueue(queueTitleRequest.QueueTitle)

		if repositoryError != nil {
			w.WriteHeader(repositoryError.HttpStatus)
			return
		}

		fmt.Printf("new queue: %v\n", newQueue)
		jsonData, _ := json.Marshal(newQueue)
		_, err := w.Write(jsonData)
		if err != nil {
			fmt.Printf("failed to write response %sn", err)
		}
	}
}

func MakeAdminViewHandler(repository waitformore.QueueRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		adminKey := vars["adminKey"]
		if len(adminKey) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		queueDefinition, repositoryError := repository.GetQueueDefinition(adminKey)

		if repositoryError != nil {
			w.WriteHeader(repositoryError.HttpStatus)
			return
		}

		fmt.Printf("admin queue: %v\n", queueDefinition)
		jsonData, err := json.Marshal(queueDefinition)
		if err != nil {
			panic(err) // golang should be able to marshal to json!
		}

		_, err = w.Write(jsonData)
		if err != nil {
			fmt.Printf("failed to write response %sn", err)
		}
	}
}

func MakeUserViewHandler(repository waitformore.QueueRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		userKey := vars["userKey"]
		if len(userKey) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userView, repositoryError := repository.GetUserView(userKey)

		if repositoryError != nil {
			w.WriteHeader(repositoryError.HttpStatus)
			return
		}

		fmt.Printf("userview: %v\n", userView)
		jsonData, _ := json.Marshal(userView)
		_, err := w.Write(jsonData)
		if err != nil {
			fmt.Printf("failed to write response %sn", err)
		}
	}
}

func MakeUserPullHandler(repository waitformore.QueueRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		userKey := vars["userKey"]
		if len(userKey) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		queueDefinition, repositoryError := repository.PullByUserKey(userKey)

		if repositoryError != nil {
			w.WriteHeader(repositoryError.HttpStatus)
			return
		}

		userView := waitformore.UserView{
			UserKey:      queueDefinition.UserKey,
			QueueTitle:   queueDefinition.QueueTitle,
			QueueState:   queueDefinition.QueueState,
			NumberPulled: queueDefinition.QueueState.LastPulled,
		}

		fmt.Printf("user pull: %v\n", userView)
		jsonData, _ := json.Marshal(userView)
		_, err := w.Write(jsonData)
		if err != nil {
			fmt.Printf("failed to write response %sn", err)
		}
	}
}

func MakeAdminCallHandler(repository waitformore.QueueRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		adminKey := vars["adminKey"]
		if len(adminKey) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		queueDefinition, repositoryError := repository.CallByAdminKey(adminKey)

		if repositoryError != nil {
			w.WriteHeader(repositoryError.HttpStatus)
		}

		fmt.Printf("admin call %v\n", queueDefinition)
		jsonData, _ := json.Marshal(queueDefinition)
		_, err := w.Write(jsonData)
		if err != nil {
			fmt.Printf("failed to write response %sn", err)
		}
	}
}


func MakeStatusHandler(repository waitformore.QueueRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		statusOverview, repositoryError := repository.GetStatus()

		if repositoryError != nil {
			w.WriteHeader(repositoryError.HttpStatus)
		}

		fmt.Printf("status call %v\n", statusOverview)
		jsonData, _ := json.Marshal(statusOverview)
		_, err := w.Write(jsonData)
		if err != nil {
			fmt.Printf("failed to write response %sn", err)
		}
	}
}
