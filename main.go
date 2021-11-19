package main

import ("fmt"
        "net/http"
        "html/template"
        "encoding/json"
        "io/ioutil"
        "log"
        "time"
        "database/sql"
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
  // Name string
  // Age uint16
//  Money int16
//  Avg_grades, Happiness float64
//  Hobbis []string
}

type Results struct {
    ResultCount  int `json:"ResultCount"`
    Results []Music `json:"Results"`
}
// func(u User) getAllInfo() string {
//   return fmt.Sprintf("User name %s. He is %  money %d", u.Name, u.Age, u.Money)
// }

// func(u *User) setNewName(newName string){
//   u.Name = newName
// }

func home_page(w http.ResponseWriter, r *http.Request){
  //bob := User{"Bob", 25}
  musics := []Music{
    {Kind: "Columbia Memorial Station", CollectionPrice: 5,},
    {Kind: "Challenger Memorial Station", CollectionPrice: 7},
    {Kind: "Cggg", CollectionPrice: 8},
}

r.ParseForm()
params := r.FormValue("foo")
switch params {
    case "download":
      workJSON()
 fmt.Println(params)

    case "remove":
  fmt.Println(params)

      case "select":
   fmt.Println(params)
 default:
     fmt.Println(params)
  }

// bob := []User{
//       {Name: "Bradbury Landing", Age: 4},
//       {Name: "Bradbury Landing2", Age: 5},
//   }


//  bob.setNewName("Alex")
//  fmt.Fprintf(w, "Test web" + bob.getAllInfo())
//fmt.Fprintf(w, <b>Main Trxt</b>")
tmpl, _ := template.ParseFiles("templates/home_page.html")
tmpl.Execute(w, musics)
}

func hadleRequest(){
http.HandleFunc("/", home_page)
//http.HandleFunc("/contacts/", contacts_page)
//http.HandleFunc("/get-time", func(w http.ResponseWriter, r *http.Request))
  fmt.Println("Connect Done")



http.ListenAndServe(":8080", nil)
}

func main(){
  //var bob User
  //bob := User{name: "Bob", age:25, money: -50, avg_grades: 4.3, happiness: 0.8}


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

    astros := Results{}

    jsonErr := json.Unmarshal(body, &astros)
	 //jsonErr := json.MarshalIndent(&astros, "", "    ")
	// fmt.Println(&astros)

    if jsonErr != nil {
        log.Fatal(jsonErr)
    }


			for i := 0; i < astros.ResultCount; i++{
	  fmt.Println(i)
		fmt.Println(astros.Results[i].Kind)
		fmt.Println(astros.Results[i].CollectionName)
		fmt.Println(astros.Results[i].TrackName)
		fmt.Println(astros.Results[i].CollectionPrice)
		fmt.Println(astros.Results[i].TrackPrice)
		fmt.Println(astros.Results[i].PrimaryGenreName)
		fmt.Println(astros.Results[i].TrackCount)
		fmt.Println(astros.Results[i].TrackNumber)
		fmt.Println(astros.Results[i].ReleaseDate)
		fmt.Println("______________")
			}
}

func WorkSQL(whattodo string){
	db, err :=sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/itunes")
	if err != nil{
	panic(err)
	}

	switch whattodo {

	    case "Insert":
	        fmt.Println("Insert")
					insert, err :=db.Query("INSERT INTO `users` (`name`, `age`) Values('Alex', 25)")
						if err != nil{
						panic(err)
						}
						defer insert.Close()

	    case "Delete":
	        fmt.Println("Delete")
					delete, err :=db.Query("DELETE FROM `users`")
						if err != nil{
						panic(err)
						}
						defer delete.Close()

	    case "Select":
	        fmt.Println("Select")
					res,  err :=db.Query("Select `name`, `age` From `users`")
						if err != nil{
						panic(err)
						}
						for res.Next() {
						  var muz Music
						  err = res.Scan(&muz.Kind, &muz.CollectionName)
						  if err != nil{
						  panic(err)
						  }
						//  fmt.Println(fmt.Sprintf("vivid", user.name, user.age))
						}

	    case "Schema":
	        fmt.Println("Schema")
					Schema, err :=db.Query("CREATE SCHEMA `test1` ;")
					if err != nil{
					panic(err)
					}
				defer Schema.Close()

				  	Schema2, err :=db.Query("CREATE TABLE `test1`.`new_table` (`id` INT NOT NULL AUTO_INCREMENT,`name` VARCHAR(45) NULL,  PRIMARY KEY (`id`))")
						if err != nil{
						panic(err)
						}
						defer Schema2.Close()

	    default:
	        fmt.Println("default")
	    }
			defer db.Close()

}
