package main

import (
	"apiTest/verification"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

// تعریف ساختار برای ذخیره داده‌ها
type Reservation struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	NationalID  string `json:"national_id"`
	TicketCount int    `json:"ticket_count"`
}

// استفاده از یک slice برای ذخیره اطلاعات رزرو
var reservations = make([]Reservation, 0)

// تابعی برای ذخیره رزرو در حافظه
func saveReservation(res Reservation) {
	reservations = append(reservations, res)
}

// هندلر برای دریافت درخواست رزرو
func reservationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var res Reservation
		
		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		fmt.Println(res, reflect.TypeOf(res), reflect.TypeOf(res.FirstName), reflect.TypeOf(res.LastName), reflect.TypeOf(res.Email), reflect.TypeOf(res.NationalID), reflect.TypeOf(res.TicketCount))

		
		errors := verification.VerificationData(res.FirstName, res.LastName, res.Email, res.NationalID, res.TicketCount)
		if len(errors) > 0 {
			fmt.Println(errors)
			http.Error(w, fmt.Sprintf("Validation errors: %v", errors), http.StatusBadRequest)
			return
		}

		saveReservation(res)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reservation successful"))

	case http.MethodGet:
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reservations)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please send data to '/reserve' for reservation.")
}

func main() {


	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/reserve", reservationHandler) 

	// راه‌اندازی سرور HTTP
	fmt.Println("Server is running on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
