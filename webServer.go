package main

import ("net/http"
	"log"
	"io/ioutil"
	"strings"
	"encoding/json"
	//"database/sql"
	 //_ "github.com/go-sql-driver/mysql"
)
type Recipe struct {
  Type string	`json:"type,omitempty"`
	URL string		`json:"url,omitempty"`
	Keywords string	`json:"key,omitempty"`
	Cooktime int	`json:"time,omitempty"`
	Rating int	`json:"rate,omitempty"`
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

	log.Println(path+ " handle func")
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
	log.Println(body)
	recipes:=make(map[string]Recipe)

	recipes["buffalo-chicken-stuffed-shells"]=Recipe{"Dinner","dummyurl","chicken hot", 40, 3}

  js, err := json.Marshal(recipes)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}
	w.Header().Set("Content-Type", "application/json")
  w.Write(js)
	log.Println(recipes["buffalo-chicken-stuffed-shells"])
	log.Println("no error")
}

func addHandler(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
			panic(err)
	}
	log.Println(string(body))
	var t Recipe
	err = json.Unmarshal(body, &t)
	if err != nil {
			panic(err)
	}
	log.Println(t.Type)
}
