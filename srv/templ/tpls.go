package templ

import (
	"github.com/habales/gotmx-play/data/apptpl"
	"html/template"
	"log"
	"os"
)

type Tpls struct {
	DevMode bool
	tplRoot *template.Template
}

func NewTpls() *Tpls {
	tpls := &Tpls{
		tplRoot: template.New("root"),
	}

	if _, e := os.Open("data"); e == nil {
		tpls.DevMode = true
		tpls.tplRoot, e = tpls.tplRoot.ParseGlob("data/apptpl/**")
		if e != nil {
			panic(e)
		}

	} else {
		tpls.tplRoot, e = tpls.tplRoot.ParseFS(apptpl.Embedded, "*")
		if e != nil {
			panic(e)
		}

	}

	return tpls
}

func (a *Tpls) GetTemplate(name string) *template.Template {
	if a.DevMode {
		var err error
		a.tplRoot = template.New("root") //FIXME this will not include any custom functions, also might be slow when there are manny templates
		a.tplRoot, err = a.tplRoot.ParseGlob("data/apptpl/**")
		if err != nil {
			log.Printf("could not parse a template %v", err)
		}
	}
	return a.tplRoot.Lookup(name)
}
