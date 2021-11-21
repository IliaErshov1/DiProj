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
  //bob := User{"Bob", 25}
// musics := Music
//    musics := []Music{
//     {Kind: "Columbia Memorial Station", CollectionPrice: 5,},
//     {Kind: "Challenger Memorial Station", CollectionPrice: 7},
//     {Kind: "Cggg", CollectionPrice: 8},
// }

//musics := Music{Kind: "Columbia Memorial Station", CollectionPrice: 5,}


//musics := Music{}
// for i := 0; i < 10; i++{
//   musics{Kind: "Test Item 1", CollectionPrice: 5}
// }

// as := Results{}
// for i := 0; i < 10; i++{
// as.Results[i].Kind="oo"
// as.Results[i].CollectionPrice=0.2
//  }
// fmt.Println(as.Results[0].Kind)


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

  case "clean":
fmt.Println(params)
     Bks=BksClean

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

    as := Results{}

    jsonErr := json.Unmarshal(body, &as)
	 //jsonErr := json.MarshalIndent(&astros, "", "    ")
	// fmt.Println(&astros)

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
// pp :=as.Results[i].CollectionPrice
// ppp :=as.Results[i].TrackPrice
// ppCollectionPrice := fmt.Sprintf("%.6f",as.Results[i].CollectionPrice)
// ppTrackPrice := fmt.Sprintf("%.6f",ppp)
//replacer :=strings.NewReplacer("T", " ", "Z", "")
//replacer := strings.ReplaceAll(as.Results[i].ReleaseDate, "T", " ")
// res1 := strings.Replace(str1, "e", "E", 3)
//outData := replacer.Replace(as.Results[i].ReleaseDate)
//strings.NewReplacer("T", " ", "Z", "").Replace(as.Results[i].ReleaseDate)
// str1 := as.Results[i].ReleaseDate
// replacer := strings.NewReplacer("T", " ", "Z", "")
// 	out := replacer.Replace(str1)
//аа := ("INSERT INTO `music` (`Kind`) Values('"+ Don't Let Me Down +"')")
//INSERT INTO `music` (`Kind`, `CollectionName`, `TrackName`, `CollectionPrice`, `TrackPrice`, `PrimaryGenreName`, `TrackCount`, `TrackNumber`, `ReleaseDate`) Values('song', 'The Beatles (The White Album)', 'Rocky Raccoon', '12.99', '1.29', 'Rock', '17', '13', '1968-11-22 12:00:00')
//INSERT INTO `music` (`Kind`, `CollectionName`, `TrackName`, `CollectionPrice`, `TrackPrice`, `PrimaryGenreName`, `TrackCount`, `TrackNumber`, `ReleaseDate`) Values('song', 'The Beatles 1967-1970 (The Blue Album)', 'Don't Let Me Down', '12.99', '1.29', 'Rock', '14', '5', '1969-04-11 12:00:00')

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
					res,  err :=db.Query("Select `Kind`, `CollectionPrice` From `music` ORDER BY `ReleaseDate` ASC")
						if err != nil{
						panic(err)
						}

          //  bks := make([]*Music, 0)
						for res.Next() {
						 // var muz Music
              //musics := []Music{
              bk := new(Music)
						  err = res.Scan(&bk.Kind, &bk.CollectionPrice)
              Bks = append(Bks, bk)
						  if err != nil{
						  panic(err)
						  }


              	// for rows.Next() {
              	// 	bk := new(Table_view)
              	// 	rows.Scan(&bk.id, &bk.fam, &bk.name)
              	// 	bks = append(bks, bk)
                //as := Results{}


             // muzZn{
             //       {Kind: muzSh.Kind, CollectionPrice: muzSh.CollectionPrice},
             //     }
						 // fmt.Println(fmt.Sprintf("vivid", muzSh.Kind, muzSh.CollectionPrice))
             // fmt.Println(muzSh)
             // tmpl, _ := template.ParseFiles("templates/home_page.html")
            //  tmpl.Execute(w, ms)
						}
        //     for i := 0; i < 10; i++{
        //   fmt.Println(i)
        //   fmt.Println(Bks[i].Kind)
        // //	fmt.Println(bks[i].CollectionName)
        // //	fmt.Println(bks[i].TrackName)
        //   fmt.Println(Bks[i].CollectionPrice)
        //	fmt.Println(bks[i].TrackPrice)
        //	fmt.Println(bks[i].PrimaryGenreName)
        //	fmt.Println(bks[i].TrackCount)
        //	fmt.Println(bks[i].TrackNumber)
        //	fmt.Println(bks[i].ReleaseDate)
        //	fmt.Println("______________"+"','"+ bks[i].Kind +"','")
    //    }



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
      // fmt.Println(words[0])
      // fmt.Println(words[1])
      settings :=words[0]
      switch settings {

    	    case "connect":
            Pbaseserver= words[1]

          case "baset":
            Pportbase= words[1]

          case "portbase":
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
