package main

import (
	"net/http"
	"log"
	//"reflect"
	"io/ioutil"
	"strings"
	"encoding/json"
	"database/sql"
	"os"
		_ "github.com/go-sql-driver/mysql"
	"github.com/antoan-angelov/go-fuzzy"
)
type Recipe struct {
	Name string `json:"name,omitempty"`
  Type string	`json:"type,omitempty"`
	URL string		`json:"url,omitempty"`
	Keywords string	`json:"key,omitempty"`
	Cooktime int	`json:"time,omitempty"`
	Rating int	`json:"rate,omitempty"`
}

type Search struct{
	Name string
	Type string
	URL string
	Keywords string
	Cooktime string
	Rating string
}
func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/add/",addHandler)
	http.ListenAndServe(":8888",nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path[1:]
	if path==""{
		path="index.html"
	}

	 data, err := ioutil.ReadFile("docs/"+string(path))

	if err==nil{
		var contentType string
		if strings.HasSuffix(path, ".css"){
			contentType="text/css"
		} else if strings.HasSuffix(path, ".html"){
			contentType="text/html"
	  } else if strings.HasSuffix(path, ".js"){
      contentType="application/javascript"
		} else if strings.HasSuffix(path, ".png"){
      contentType="image/png"
		} else if strings.HasSuffix(path, ".svg"){
      contentType="image/svg+xml"
		} else if strings.HasSuffix(path, ".jpg"){
			contentType="image/jpeg"
		} else {
			contentType="text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else{
		w.WriteHeader(404)
		w.Write([]byte("404 this didnt work" + http.StatusText(404)))
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request){
	body, err :=ioutil.ReadAll(r.Body)
	if err !=nil{
		panic(err)
	}
	var s Search
	json.Unmarshal(body, &s)

	sqlState:="Select name, type, url, cooktime, keywords, rating from recipes"
	if (s.Type != "" || s.Rating !="" || s.Cooktime !="") {
		sqlState=sqlState+ " where"
	}
	if s.Type != "" {
		sqlState=sqlState+ " type = '" +s.Type+"'"
	}
	if s.Rating != "" {
		sqlState=sqlState + " rating = "+s.Rating
	}
	if s.Cooktime !="" {
		sqlState = sqlState + " cooktime < "+s.Cooktime
	}

	sqlState +=";"

	//log.Println(sqlState)
	sqlResult:=searchDB(sqlState)

	recipeMap:=make(map[string]Recipe)
	for sqlResult.Next(){
		var r Recipe
		//var name string
		sqlResult.Scan(&r.Name, &r.Type, &r.URL, &r.Cooktime, &r.Keywords, &r.Rating)
		recipeMap[r.Name]=r
		//log.Println("Recipe ", r.Name,  " added to map" )
	}

	if (s.Name != ""){
		recipeMap=fuzzySearch(recipeMap, "Name", s.Name)
	}
	if (s.Keywords != ""){
		keyStrip :=strings.Split(s.Keywords, ",")
		for _, value := range keyStrip{
				recipeMap=subStringSearch(recipeMap, "Keywords", value)
			}
	}

	if (len(recipeMap) ==0 ){
		var r Recipe

		r.Name = "No recipes found ....."
		r.URL, _ = os.Hostname()
		r.URL = r.URL + "/err"
		recipeMap[r.Name] = r
	}
  js, err := json.Marshal(recipeMap)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}
	w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func addHandler(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
			panic(err)
	}
	var s Search
	json.Unmarshal(body, &s)
	insert:="Insert INTO recipes (id, name, type, url, keywords, cooktime, rating) VALUES (NULL,"
	if (s.Name!=""){
		insert=insert+"'"+s.Name+"',"
	} else{
		insert=insert+"NULL,"
	}

	if (s.Type!=""){
		insert=insert+"'"+s.Type+"',"
	}	else{
		insert=insert+"NULL,"
	}

	if (s.URL!=""){
		insert=insert+"'"+s.URL+"',"
	}	else{
		insert=insert+"NULL,"
	}

	if (s.Keywords!=""){
		insert=insert+"'"+s.Keywords+"',"
	}	else{
		insert=insert+"NULL,"
	}

	if (s.Cooktime!=""){
		insert=insert+"'"+s.Cooktime+"',"
	}	else{
		insert=insert+"NULL,"
	}

	if (s.Rating!=""){
		insert=insert+"'"+s.Rating+"');"
	}	else{
		insert=insert+"NULL);"
	}
	log.Println(insert)
	insertDB(insert)
}

func insertDB(insert string){
	log.Println("insert called")
	db, err := sql.Open("mysql", "root:password@/cookbook")
	if err != nil {
    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	_, err2 := db.Query(insert)
	if err2 !=nil{
		panic(err2.Error())
	}
	defer db.Close()
}

func searchDB(sel string) *sql.Rows{
	//log.Println("select called")
	db, err := sql.Open("mysql", "root:password@/cookbook")
	if err != nil {
    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	rows, err := db.Query(sel)
	if err !=nil{
		panic(err.Error())
	}
	defer db.Close()

	return rows
}

func fuzzySearch(recipeMap map[string]Recipe, searchType string, searchTerm string) map[string]Recipe{
	fuzzySearcher := fuzzy.NewFuzzy()

	newInt :=make([]interface{}, len(recipeMap))

	index := 0
	for key :=range recipeMap{
		//log.Println(key)
		newInt[index]=recipeMap[key]
		index+=1
	}

	fuzzySearcher.Set(&newInt)
	fuzzySearcher.SetKeys([]string{searchType})

	results, err := fuzzySearcher.Search(searchTerm)

	if err != nil {
		log.Println("Search Failure")
	}

	resultMap := make(map[string]Recipe)
	for key := range results{
		log.Println(key)
		recipe, ok := results[key].(Recipe)
		if !ok{
			log.Println("Type Assertion Failure")
		}
		resultMap[recipe.Name]=recipe
		//log.Println("In the result map: ", recipe.Name)
	}

	return resultMap
}

func subStringSearch(recipeMap map[string]Recipe, searchType string, searchTerm string) map[string]Recipe{
	newMap :=make(map[string]Recipe)
	for _, value := range recipeMap{
		if (searchType == "Keywords"){
			if(strings.Contains(value.Keywords, strings.TrimSpace(searchTerm))){
				newMap[value.Name]=value
			}
		}
	}

	return newMap
}
