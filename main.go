package main

import (
	"fmt"
	"game-of-life/game"
	"html/template"
	"net/http"
)

func main() {

	templates, err := template.ParseFiles("./templates/layout.html", "./templates/gameBoard.html")
	styles := http.FileServer(http.Dir("./stylesheets"))

	if err != nil {
		panic(err)
	}

	fmt.Println("Templates loaded!")

	g := game.Init(100)
	g.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "layout", g.State)
	})

	http.Handle("/styles/", http.StripPrefix("/styles", styles))

	http.ListenAndServe(":3000", nil)
}
