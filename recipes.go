package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antoan-angelov/go-fuzzy"
)

type Recipe struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	URL      string `json:"url,omitempty"`
	Keywords string `json:"key,omitempty"`
	Cooktime int    `json:"time,omitempty"`
	Rating   int    `json:"rate,omitempty"`
}

type Search struct {
	Name     string
	Type     string
	URL      string
	Keywords string
	Cooktime string
	Rating   string
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var s Search
	json.Unmarshal(body, &s)

	sqlState := "Select name, type, url, cooktime, keywords, rating from recipes"
	if s.Type != "" || s.Rating != "" || s.Cooktime != "" {
		sqlState = sqlState + " where"
	}
	if s.Type != "" {
		sqlState = sqlState + " type = '" + s.Type + "'"
	}
	if s.Rating != "" {
		sqlState = sqlState + " rating = " + s.Rating
	}
	if s.Cooktime != "" {
		sqlState = sqlState + " cooktime < " + s.Cooktime
	}

	sqlState += ";"

	//log.Println(sqlState)
	sqlResult := SearchDB(sqlState)

	recipeMap := make(map[string]Recipe)
	for sqlResult.Next() {
		var r Recipe
		sqlResult.Scan(&r.Name, &r.Type, &r.URL, &r.Cooktime, &r.Keywords, &r.Rating)
		recipeMap[r.Name] = r
		//log.Println("Recipe ", r.Name,  " added to map" )
	}

	if s.Name != "" {
		recipeMap = fuzzySearch(recipeMap, "Name", s.Name)
	}
	if s.Keywords != "" {
		keyStrip := strings.Split(s.Keywords, ",")
		for _, value := range keyStrip {
			recipeMap = subStringSearch(recipeMap, "Keywords", value)
		}
	}

	if len(recipeMap) == 0 {
		var r Recipe

		r.Name = "No recipes found ....."
		r.URL, _ = os.Hostname()
		r.URL = r.URL + "/err"
		recipeMap[r.Name] = r
	}
	js, err := json.Marshal(recipeMap)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var s Search
	json.Unmarshal(body, &s)
	insert := "Insert INTO recipes (id, name, type, url, keywords, cooktime, rating) VALUES (NULL,"
	if s.Name != "" {
		insert = insert + "'" + s.Name + "',"
	} else {
		insert = insert + "NULL,"
	}

	if s.Type != "" {
		insert = insert + "'" + s.Type + "',"
	} else {
		insert = insert + "NULL,"
	}

	if s.URL != "" {
		insert = insert + "'" + s.URL + "',"
	} else {
		insert = insert + "NULL,"
	}

	if s.Keywords != "" {
		insert = insert + "'" + s.Keywords + "',"
	} else {
		insert = insert + "NULL,"
	}

	if s.Cooktime != "" {
		insert = insert + "'" + s.Cooktime + "',"
	} else {
		insert = insert + "NULL,"
	}

	if s.Rating != "" {
		insert = insert + "'" + s.Rating + "');"
	} else {
		insert = insert + "NULL);"
	}
	log.Println(insert)
	InsertDB(insert)
}

func fuzzySearch(recipeMap map[string]Recipe, searchType string, searchTerm string) map[string]Recipe {
	fuzzySearcher := fuzzy.NewFuzzy()

	newInt := make([]interface{}, len(recipeMap))

	index := 0
	for key := range recipeMap {
		//log.Println(key)
		newInt[index] = recipeMap[key]
		index += 1
	}

	fuzzySearcher.Set(&newInt)
	fuzzySearcher.SetKeys([]string{searchType})

	results, err := fuzzySearcher.Search(searchTerm)

	if err != nil {
		log.Println("Search Failure")
	}

	resultMap := make(map[string]Recipe)
	for key := range results {
		log.Println(key)
		recipe, ok := results[key].(Recipe)
		if !ok {
			log.Println("Type Assertion Failure")
		}
		resultMap[recipe.Name] = recipe
		//log.Println("In the result map: ", recipe.Name)
	}

	return resultMap
}

func subStringSearch(recipeMap map[string]Recipe, searchType string, searchTerm string) map[string]Recipe {
	newMap := make(map[string]Recipe)
	for _, value := range recipeMap {
		if searchType == "Keywords" {
			if strings.Contains(value.Keywords, strings.TrimSpace(searchTerm)) {
				newMap[value.Name] = value
			}
		}
	}

	return newMap
}
