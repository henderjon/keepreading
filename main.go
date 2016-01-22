package main

import (
	"net/http"
	"github.com/henderjon/keepreading/esvapi"
	"github.com/henderjon/keepreading/fv"
	reading "github.com/henderjon/keepreading/djrp"
	"log"
	"os"
	"html/template"
	"bytes"
	"strconv"
	"time"
)

var (
	logger *log.Logger
	expires = 720 * time.Hour
	month = "month"
	day   = "day"
)

// represent the month/day reading progress of a user
type bookmark map[string]string

// hold the byte responses from the ESV API
type page struct {
	Gospel, Nt, Ot, Wisdom, Memory bytes.Buffer
}

// our ESV responses come as html, pass HTML safe strings to the templates
type htmlPage struct {
	Gospel, Nt, Ot, Wisdom, Memory template.HTML
}

// set up a logger for debugging
func init() {
	logger = log.New(os.Stdout, ".log ", log.Lshortfile)
}

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", HandleAll())
	http.HandleFunc("/set", HandleSetBookmark())
	http.ListenAndServe(":80", nil)

}

// parse the requests cookies and display that day's reading to the user
func HandleAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		t, err := template.ParseFiles("views/reading.html")
		if err != nil {
			logger.Fatal(err)
		}

		cookie, _ := getBookmark(req)
		m, d := normalizeBookmark(cookie)

		page := page{
			esvapi.Query(reading.Plan[m][d].Gospel),
			esvapi.Query(reading.Plan[m][d].Nt),
			esvapi.Query(reading.Plan[m][d].Ot),
			esvapi.Query(reading.Plan[m][d].Wisdom),
			esvapi.Query(fv.Current(fv.Base)),
		}

		t.Execute(w, htmlPage{
			template.HTML(page.Gospel.String()),
			template.HTML(page.Nt.String()),
			template.HTML(page.Ot.String()),
			template.HTML(page.Wisdom.String()),
			template.HTML(page.Memory.String()),
		})
	}
}

// allow a user to manually set where they last read via query string
func HandleSetBookmark() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		_, form := getBookmark(req)
		m, d := normalizeBookmark(form)

		http.SetCookie(w, &http.Cookie{
			Name: month,
			Value: strconv.Itoa(m),
			Expires: time.Now().Add(expires),
		})

		http.SetCookie(w, &http.Cookie{
			Name: day,
			Value: strconv.Itoa(d),
			Expires: time.Now().Add(expires),
		})

		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	}
}

// parse both the request cookies and form (url) to suss out the provided
// month and day
func getBookmark(req *http.Request) (bookmark, bookmark) {

	cValues := make(bookmark, 0)
	fValues := make(bookmark, 0)
	keys    := []string{month, day}

	req.ParseForm()
	for _, k := range keys {
		if c, err := req.Cookie(k); err == nil {
			cValues[k] = c.Value
		}else{
			cValues[k] = ""
		}

		fValues[k] = req.FormValue(k)
	}
	return cValues, fValues
}

// normalize the given bookmark for our of range numbers, compensating for a zero index
func normalizeBookmark(b bookmark) (int, int) {

	monthInt, dayInt := 1, 1

	if len(b[month]) > 0 {
		monthInt, _ = strconv.Atoi(b[month])
	}

	for monthInt > 12  {
		monthInt -= 12
	}

	// zero indexes
	if monthInt > 0 {
		monthInt -= 1
	}

	if len(b[day]) > 0 {
		dayInt, _ = strconv.Atoi(b[day])
	}

	for dayInt > 25  {
		dayInt -= 25
	}

	// zero indexes
	if dayInt > 0 {
		dayInt -= 1
	}

	// zero indexes
	return monthInt, dayInt
}


/*
//http://www.alexedwards.net/blog/serving-static-sites-with-go

fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))

http.HandleFunc("/", serveTemplate)

log.Println("Listening...")
http.ListenAndServe(":3000", nil)
 */
