package lib

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"web/ascii"
)

var counter int

func RunServer() {

	//file server for image used for logo
	fs := http.FileServer(http.Dir("img"))
	http.HandleFunc("/img", images)
	// We're binding the handler to the `/images` route, here.
	http.Handle("/img/", http.StripPrefix("/img/", fs))

	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	js := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	font := http.FileServer(http.Dir("files"))
	http.Handle("/files/", http.StripPrefix("/files/", font))

	handler := http.HandlerFunc(handleRequest)
	http.Handle("/output.txt", handler)

	//start server and handle /ascii-art
	http.HandleFunc("/", HandlePage)
	err := http.ListenAndServe("0.0.0.0:8080", nil)

	//log if error
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("counter")
	strcounter := cookie.Value
	fileBytes, err := os.ReadFile("files/output" + strcounter + ".txt")
	if err != nil {
		panic(err)
	}
	//convert filsezie to string
	number := len(fileBytes)
	converted := strconv.Itoa(number)

	//set header
	w.Header().Set("Content-Type", " text/html; charset=UTF-8")
	w.Header().Set("Content-Length", converted)
	w.Header().Set("Content-Disposition", "attachment; filename=output"+strcounter+".txt")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}

// function to handle images
func images(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/link.html")

	if err != nil {
		log.Panic("There's an error with the template:", err)
	}

	t.ExecuteTemplate(w, "link", nil)
}

type TodoPageData struct {
	PageTitle string
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/ascii-art.html")
	if err != nil {
		log.Fatalln("There's an error with the template:", err)
	}

	//t.Execute(w, nil)

	//Error handling
	if r.URL.Path != "/ascii-art" && r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.URL.Path != "/ascii-art" && r.URL.Path != "/" {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/ascii-art" && r.URL.Path != "/" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	//switch statement to handle GET and POST
	switch r.Method {
	case "GET":
		fmt.Println("GET")
		t.Execute(w, nil)
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		//variable used to pass to the ascii-art program
		text := r.FormValue("inputtext")

		banner := sanitizebanner(r.FormValue("banner"))
		if banner == "ok" {
			http.Redirect(w, r, "files/error.html", http.StatusSeeOther)
			banner = "standard.txt"
		}
		align := sanitizealign(r.FormValue("align"))
		if align == "ok" {
			http.Redirect(w, r, "files/error.html", http.StatusSeeOther)
			align = "left"
		}

		var counter int
		counter++
		counter2 := strconv.Itoa(counter)
		cookie := &http.Cookie{
			Name:   "counter",
			Value:  counter2,
			MaxAge: 600,
		}
		http.SetCookie(w, cookie)
		cookiecounter := strconv.Itoa(counter)

		//struct for info
		data := TodoPageData{
			PageTitle: ascii.PrintAscii(text, banner, cookiecounter),
		}

		//switch cases for simple alignment without the need for program change.
		switch align {
		case "left":
			data = TodoPageData{
				PageTitle: ascii.PrintAscii(text, banner, cookiecounter),
			}
			t.Execute(w, data)
		case "right":
			data = TodoPageData{
				PageTitle: ascii.PrintAscii(text, banner, cookiecounter),
			}
			t.Execute(w, data)
		case "center":
			data = TodoPageData{
				PageTitle: ascii.PrintAscii(text, banner, cookiecounter),
			}
			t.Execute(w, data)
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func sanitizebanner(s string) string {
	if s == "" {
		return "standard.txt"
	} else if s != "standard.txt" && s != "shadow.txt" && s != "thinkertoy.txt" {
		fmt.Printf("Invalid banner, using standard.txt was: %s\n", s)
		return "ok"
	} else {
		return s
	}
}

func sanitizealign(s string) string {
	fmt.Println("")
	if s == "" {
		return "left"
	} else if s != "left" && s != "right" && s != "center" {
		fmt.Printf("Invalid alignment, using left was: %s\n", s)
		return "ok"
	} else {
		return s
	}
}
