package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Note struct {
	ID        string
	Message   string
	RecipeID  string
	Timestamp float64
}

func AddNotesHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var note Note
	json.Unmarshal(body, &note)
	sqlStatement := fmt.Sprintf("INSERT INTO notes (id, msg, recipeId, timestamp) VALUES (NULL, %s, %s, %d)", note.Message, note.RecipeID, time.Now().Unix())
	InsertDB(sqlStatement)
}

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var recipeID string
	json.Unmarshal(body, &recipeID)
	sqlStatement := fmt.Sprintf("SELECT ID, Message, RecipeID from notes where RecipeID == %s ORDER BY Timestamp DESC", recipeID)
	sqlResult := SearchDB(sqlStatement)
	notesMap := make(map[string]Note)
	for sqlResult.Next() {
		var n Note
		sqlResult.Scan(&n.ID, &n.Message, &n.RecipeID, &n.Timestamp)
		notesMap[n.ID] = n
		//log.Println("Note ", n.ID,  " added to map" )
	}

	js, err := json.Marshal(notesMap)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {

}
