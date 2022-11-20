package main

import (
	"17dir/database"
	_ "17dir/database/migration"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type PhoneInfo struct {
	UserName  string
	UserPhone string
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok { // Это означает ok = false
		http.Error(w, "GET parameter 'id' is required", http.StatusBadRequest)
		return
	}

	user_ID := ids[0]
	user_Ids, err := strconv.ParseInt(user_ID, 10, 64)
	if err != nil {
		http.Error(w, "GET parameter 'id' is not valid", http.StatusBadRequest)
		return
	}

	query := `SELECT user_name, phone FROM phone WHERE id=$1;`
	row := database.DB.QueryRow(query, user_Ids)

	var UserName, UserPhone string
	err = row.Scan(&UserName, &UserPhone)
	switch err {
	case sql.ErrNoRows:
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	case nil:
	default:
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	ans := fmt.Sprintf("%s,%s", UserName, UserPhone)
	_, err = fmt.Fprintf(w, ans)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

}

func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := `SELECT user_name, phone FROM phone;`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	var phoneInfos []PhoneInfo
	for rows.Next() {
		var UserName, UserPhone string
		rows.Scan(&UserName, &UserPhone)
		if err != nil {
			log.Println(err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
		phoneInfos = append(phoneInfos, PhoneInfo{
			UserName:  UserName,
			UserPhone: UserPhone,
		})
	}

	ans := ""
	for i := range phoneInfos {
		ans += fmt.Sprintf("%s,%s\n", phoneInfos[i].UserName, phoneInfos[i].UserPhone)
	}

	_, err = fmt.Fprintf(w, ans)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok { // Это означает ok = false
		http.Error(w, "GET parameter 'id' is required", http.StatusBadRequest)
		return
	}

	user_ID := ids[0]
	user_Ids, err := strconv.ParseInt(user_ID, 10, 64)
	if err != nil {
		http.Error(w, "GET parameter 'id' is not valid", http.StatusBadRequest)
		return
	}

	UserPhone, ok := r.URL.Query()["phone"]
	if !ok || len(UserPhone) == 0 { // Это означает ok = false
		http.Error(w, "GET parameter 'Phone' is required", http.StatusBadRequest)
		return
	}
	UserPhones := UserPhone[0]

	query := `UPDATE phone SET phone=$1 WHERE id=$2;`
	_, err = database.DB.Exec(query, UserPhones, user_Ids)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	ans := fmt.Sprintf("ok")
	_, err = fmt.Fprintf(w, ans)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	UserName, ok := r.URL.Query()["user_name"]
	if !ok || len(UserName) == 0 { // Это означает ok = false
		http.Error(w, "GET parameter 'User Name' is required", http.StatusBadRequest)
		return
	}
	UserPhone, ok := r.URL.Query()["phone"]
	if !ok || len(UserPhone) == 0 { // Это означает ok = false
		http.Error(w, "GET parameter 'Phone' is required", http.StatusBadRequest)
		return
	}

	UserNames := UserName[0]
	UserPhones := UserPhone[0]
	query := `INSERT INTO phone (user_name, phone) VALUES ($1, $2);`
	_, err := database.DB.Exec(query, UserNames, UserPhones)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	ans := fmt.Sprintf("ok")
	_, err = fmt.Fprintf(w, ans)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {

	dbSettings := database.Settings{
		User:   "postgres",
		Pass:   "12345",
		Name:   "phone",
		Host:   "localhost",
		Port:   "5432",
		Reload: false,
	}
	err := database.Connect(dbSettings)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/get", get)
	http.HandleFunc("/all", getAll)
	http.HandleFunc("/update", update)
	http.HandleFunc("/add", add)

	fmt.Println("Запуск сервера\n")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
