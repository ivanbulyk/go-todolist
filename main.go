package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	username = "myuser"
	password = "password"
	portnum = 3306
	dbname = "todolist"
	host = "localhost"
)

type Todo struct {
	Id    int
	Title string
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}


func main() {



	log.Print("Starting the service...")


	r := mux.NewRouter()



	DSN := username + ":" + password + "@(" + host +":" + strconv.Itoa(portnum) + ")/" + dbname + "?parseTime=true"

	//fmt.Println(DSN)

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	log.Printf("Connected to DB %s successfully\n", dbname)

	query := `CREATE TABLE IF NOT EXISTS todos (
                        id int(11) unsigned NOT NULL AUTO_INCREMENT,
                        title VARCHAR(128) NOT NULL DEFAULT "",
                        PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`

	_, err = db.Exec(query)
	if err != nil {
		log.Printf("Error %s when creating todos table", err)
		return
	}

	homeTempl := template.Must(template.ParseFiles("./templates/index.html"))



	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		todos, err := db.Query("select * from todos")

		if err != nil {
			log.Printf("Error %s when selecting todos", err)
			//return
		}
		log.Println("Successfully selecting todos")

		defer todos.Close()


		var todo []Todo

		for todos.Next() {
			var t Todo
			err := todos.Scan(&t.Id, &t.Title)

			if err != nil {
				log.Printf("Error %s", err)
				return
			}
			log.Println("Successfully selecting todo", t)
			todo = append(todo, t)

		}



		data := TodoPageData{
			PageTitle: "To do list",
			Todos:     todo,
		}



		fmt.Println(data)
		er := homeTempl.Execute(w, data)

		if er != nil {
			log.Println(er)
		}
	})


	r.HandleFunc("/remove-todo/{id}", func(w http.ResponseWriter, r *http.Request){
		id := mux.Vars(r)["id"]
		_, err := db.Exec(`Delete from todos where id = ?`, id)

		if err != nil{
			log.Printf("Error %s", err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Printf("Successfully deleted todo with ID: %s\n", id)

	})



	r.HandleFunc("/add-todo", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost {
			title := r.FormValue("todotitle")
			_, err = db.Exec("insert into todos(title) values(?)", title)

			if err != nil{
				log.Printf("Error %s", err)
			}
			log.Println("no errors")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})


	r.
		PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("."+"/assets/"))))


	log.Fatal(http.ListenAndServe(":8080", r))
}
