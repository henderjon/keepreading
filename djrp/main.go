package djrp

import (
	"bytes"
	"github.com/henderjon/keepreading/esvapi"
	"log"
	"os"
	"text/template"
)

const (
	Jan = iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
	Gospel
	Wisdom
	Ot
	Nt
)

// renders the given genre's reading for the given month/day
func Render(g, m, d int) {
	t, _ := template.New("all").Parse("\n\n{{.}}\n\n")

	m -= 1 // zero index
	d -= 1 // zero index

	var p bytes.Buffer

	switch {
	case g == Gospel:
		p = esvapi.Query(Plan[m][d].Gospel)
	case g == Wisdom:
		p = esvapi.Query(Plan[m][d].Wisdom)
	case g == Ot:
		p = esvapi.Query(Plan[m][d].Ot)
	case g == Nt:
		p = esvapi.Query(Plan[m][d].Nt)
	}

	err := t.Execute(os.Stdout, (&p).String())

	if err != nil {
		log.Fatal(err)
	}
}

// renders the given genre's reference for the given month/day
func Ref(g, m, d int) {
	t := template.New("all")

	m -= 1 // zero index
	d -= 1 // zero index

	var ref string

	switch {
	case g == Gospel:
		ref = Plan[m][d].Gospel
		t, _ = t.Parse("\nGospel: {{.}}\n\n")
	case g == Wisdom:
		ref = Plan[m][d].Nt
		t, _ = t.Parse("\nWisdom: {{.}}\n\n")
	case g == Ot:
		ref = Plan[m][d].Ot
		t, _ = t.Parse("\nOld Testament: {{.}}\n\n")
	case g == Nt:
		ref = Plan[m][d].Wisdom
		t, _ = t.Parse("\nNew Testament: {{.}}\n\n")
	}

	err := t.Execute(os.Stdout, ref)

	if err != nil {
		log.Fatal(err)
	}
}
