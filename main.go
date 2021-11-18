package main

import ("fmt"
        "net/http"
        "html/template")

type User struct{
  Name string
  Age uint16
//  Money int16
//  Avg_grades, Happiness float64
//  Hobbis []string
}

// func(u User) getAllInfo() string {
//   return fmt.Sprintf("User name %s. He is %  money %d", u.Name, u.Age, u.Money)
// }

// func(u *User) setNewName(newName string){
//   u.Name = newName
// }

func home_page(w http.ResponseWriter, r *http.Request){
  //bob := User{"Bob", 25}
  bob := []User{
    {Name: "Columbia Memorial Station", Age: 5,},
    {Name: "Challenger Memorial Station", Age: 7},
    {Name: "Cggg", Age: 8},
}

r.ParseForm()
    // они все тут
//  params := r.Form
params := r.FormValue("foo")
switch params {
    case "download":
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
tmpl.Execute(w, bob)
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
