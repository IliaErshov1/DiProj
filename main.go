package main

import ("fmt"
        "net/http"
        "html/template"
        "encoding/json"
        "io/ioutil"
        "log"
        "time"
        "strconv"
        "strings"
        "database/sql"
        "os"
        "bufio"
      _ "github.com/go-sql-driver/mysql")

type Music struct{
  Kind string `json:"Kind"`
  CollectionName string `json:"CollectionName"`
  TrackName string `json:"TrackName"`
  CollectionPrice float32 `json:"CollectionPrice"`
  TrackPrice float32 `json:"TrackPrice"`
  PrimaryGenreName string `json:"PrimaryGenreName"`
  TrackCount int `json:"TrackCount"`
  TrackNumber int `json:"TrackNumber"`
  ReleaseDate string `json:"ReleaseDate"`

}

type Results struct {
    ResultCount  int `json:"ResultCount"`
    Results []Music `json:"Results"`
}

var muzSh Music
var Bks = make([]*Music, 0)
var BksClean = make([]*Music, 0)
var Pbaseserver="127.0.0.1"
var Pportbase="3306"
var Pbase="itunes"
var Plogin="root"
var Ppass="root"
var Pportweb="8080"

func home_page(w http.ResponseWriter, r *http.Request){
r.ParseForm()
params := r.FormValue("foo")
switch params {
    case "download":
      workJSON()
 fmt.Println(params)

    case "remove":
  fmt.Println(params)
 WorkSQL("Delete",  "s")

      case "select":
   fmt.Println(params)
WorkSQL("Select", "tmpl")
  //tmpl.Execute(w, muz)

  case "cleanlist":
fmt.Println(params)
     Bks=BksClean

  case "schema":
    fmt.Println(params)
    WorkSQL("Schema", "")
    WorkSQL("Table", "")

 default:
     fmt.Println(params)
  }

tmpl, _ := template.ParseFiles("templates/home_page.html")
tmpl.Execute(w, Bks)
}

func hadleRequest(){
http.HandleFunc("/", home_page)
//http.HandleFunc("/contacts/", contacts_page)
//http.HandleFunc("/get-time", func(w http.ResponseWriter, r *http.Request))
  fmt.Println("Connect Done")



http.ListenAndServe(":"+Pportweb, nil)
}

func main(){
  FileRead()
  WorkSQL("Schema", "")
  WorkSQL("Table", "")
  hadleRequest()
}

func testb(){
  fmt.Println("Connect Done")
}

func workJSON() {
    url := "https://itunes.apple.com/search?term=The+Beatles"

    var netClient = http.Client{
        Timeout: time.Second * 10,
    }

    res, err := netClient.Get(url)

    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)

  //  fmt.Println(body)

    if err != nil {
        log.Fatal(err)
    }

    as := Results{}

    jsonErr := json.Unmarshal(body, &as)

    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

			for i := 0; i < as.ResultCount; i++{
	  fmt.Println(i)
		fmt.Println(as.Results[i].Kind)
		fmt.Println(as.Results[i].CollectionName)
		fmt.Println(as.Results[i].TrackName)
		fmt.Println(as.Results[i].CollectionPrice)
		fmt.Println(as.Results[i].TrackPrice)
		fmt.Println(as.Results[i].PrimaryGenreName)
		fmt.Println(as.Results[i].TrackCount)
		fmt.Println(as.Results[i].TrackNumber)
		fmt.Println(as.Results[i].ReleaseDate)
		fmt.Println("______________"+"','"+ as.Results[i].Kind +"','")

    s := ("REPLACE INTO `music` (`Kind`, `CollectionName`, `TrackName`, `CollectionPrice`, `TrackPrice`, `PrimaryGenreName`, `TrackCount`, `TrackNumber`, `ReleaseDate`) Values('"+ as.Results[i].Kind +"', '"+ as.Results[i].CollectionName  +"', '"+ strings.Replace(as.Results[i].TrackName, "'", string([]rune{'\u005c', '\u0027'}) , 1) +"', '"+  fmt.Sprintf("%.2f", as.Results[i].CollectionPrice)  +"', '"+  fmt.Sprintf("%.2f", as.Results[i].TrackPrice)  +"', '"+  as.Results[i].PrimaryGenreName  +"', '"+  strconv.Itoa(as.Results[i].TrackCount)  +"', '"+  strconv.Itoa(as.Results[i].TrackNumber)  +"', '"+  strings.NewReplacer("T", " ", "Z", "").Replace(as.Results[i].ReleaseDate) +"')")
fmt.Println(s)
  WorkSQL("Insert",  s)
      }
}

func WorkSQL(whattodo string, queryin string) {
  pconnect:= Plogin+":"+Ppass+"@tcp("+Pbaseserver+":"+Pportbase+")/"+Pbase
  fmt.Println(pconnect)
	db, err :=sql.Open("mysql", pconnect)
	if err != nil{
	panic(err)
	}

	switch whattodo {

	    case "Insert":
	        fmt.Println("Insert")
					insert, err :=db.Query(queryin)
						if err != nil{
						panic(err)
						}
						defer insert.Close()

	    case "Delete":
	        fmt.Println("Delete")
					delete, err :=db.Query("DELETE FROM `music`")
						if err != nil{
						panic(err)
						}
						defer delete.Close()

	    case "Select":
	        fmt.Println("Select")
					res,  err :=db.Query("Select `CollectionName`, `ReleaseDate` From `music` ORDER BY `ReleaseDate` DESC")
						if err != nil{
						panic(err)
						}

						for res.Next() {
              bk := new(Music)
						  err = res.Scan(&bk.CollectionName, &bk.ReleaseDate)
              Bks = append(Bks, bk)
						  if err != nil{
						  panic(err)
						  }
						}

	    case "Schema":
        pconnectsh:= Plogin+":"+Ppass+"@tcp("+Pbaseserver+":"+Pportbase+")/"
        fmt.Println(pconnectsh)
      	db, err :=sql.Open("mysql", pconnectsh)
      	if err != nil{
      	panic(err)
      }
	        fmt.Println("Schema")
					Schema, err :=db.Query("CREATE SCHEMA `itunes`;")
					if err != nil{
            fmt.Println(err)
            return
					//panic(err)
					}
				defer Schema.Close()

      case "Table":
        pconnectsh:= Plogin+":"+Ppass+"@tcp("+Pbaseserver+":"+Pportbase+")/"+Pbase
        fmt.Println(pconnectsh)
        db, err :=sql.Open("mysql", pconnectsh)
        if err != nil{
        panic(err)
      }
				  	Schema2, err :=db.Query("CREATE TABLE `itunes`.`music` (`id` INT NOT NULL AUTO_INCREMENT, `Kind` VARCHAR(45) NULL, `CollectionName` VARCHAR(300) NULL, `TrackName` VARCHAR(300) NULL, `CollectionPrice` DOUBLE NULL, `TrackPrice` DOUBLE NULL, `PrimaryGenreName` VARCHAR(300) NULL, `TrackCount` INT NULL, `TrackNumber` INT NULL, `ReleaseDate` DATETIME NULL, PRIMARY KEY (`id`));")
						if err != nil{
            fmt.Println(err)
            return
            //panic(err)
						}
						defer Schema2.Close()

	    default:
	        fmt.Println("default")
	    }
			defer db.Close()

}

func FileRead(){
  file, err := os.Open("settings.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      fmt.Println(scanner.Text())
      words := strings.Split(scanner.Text(), ":")
      settings :=words[0]
      switch settings {

    	    case "connect":
            Pbaseserver= words[1]

          case "portbase":
            Pportbase= words[1]

          case "base":
              Pbase= words[1]

          case "login":
              Plogin = words[1]

          case "pass":
              Ppass= words[1]

          case "portweb":
              Pportweb= words[1]

  }
}
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}
