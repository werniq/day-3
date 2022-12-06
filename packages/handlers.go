package handler

import (
	"book_website"
	"encoding/json"
	"net/http"
)

// Get complete list of tasks

var GetList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Set header tp json content, otherwise data appear as plain text
	w.Header().Set("Content-Type", "application/json")

	// Connect to db and get user_id
	db, userId := book_website.OpenConnection()

	// Return all tasks (rows) as id, task, status where the user_uuid is the same as userId as I defined 2 string above
	rows, err := db.Query("SELECT id, task, status FROM tasks JOIN users ON tasks.user_id = users.user_id WHERE user_id = $1;", userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	defer db.Close()

	items := make([]Item, 0)
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.TaskNum, &item.Task, &item.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err)
		}
		items = append(items, item)
	}

	itemBytes, _ := json.MarshalIndent(items, "", "\t")
	// write to w
	_, err := w.Write(itemBytes)
	if err != nil {
		http.Error(w, err.Error, http.StatusNotFound)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
})
