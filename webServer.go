package main

import ("net/http"
	"log"
	"io/ioutil"
	"strings"
	"encoding/json"
	"strconv"
	"database/sql"
		_ "github.com/go-sql-driver/mysql"
)
type Recipe struct {
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
	Cooktime int
	Rating int
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
	var s Search
	json.Unmarshal(body, &s)

	log.Println(s.Rating)

	sqlState:="Select * from recipes where "
	noArgs:=true

	if (s.Name!=""){
		sqlState=sqlState+"name='"+s.Name+"' and "
		noArgs=false
	}

	if (s.Type!=""){
		sqlState=sqlState+"type='"+s.Type+"' and "
		noArgs=false
	}

	if (s.Keywords!=""){
		sqlState=sqlState+"keywords='"+s.Keywords+"' and "
		noArgs=false
	}

	if (strconv.Itoa(s.Rating)!="0"){
		sqlState=sqlState+" rating='"+strconv.Itoa(s.Rating)+"' and "
		noArgs=false
	}

	if(noArgs){
		sqlState=sqlState[:len(sqlState)-7]+";"
	} else{
		sqlState=sqlState[:len(sqlState)-5]+";"
	}


	recipes:=searchDB(sqlState)

  js, err := json.Marshal(recipes)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}
	w.Header().Set("Content-Type", "application/json")
  w.Write(js)
	log.Println(s.Name)
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

	if (strconv.Itoa(s.Cooktime)!="0"){
		insert=insert+"'"+strconv.Itoa(s.Cooktime)+"',"
	}	else{
		insert=insert+"NULL,"
	}

	if (strconv.Itoa(s.Rating)!="0"){
		insert=insert+"'"+strconv.Itoa(s.Rating)+"');"
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
	log.Println("select called")
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
