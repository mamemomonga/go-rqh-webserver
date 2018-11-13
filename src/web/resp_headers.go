package web

import (
	"net/http"
	"sort"
	"fmt"
	"log"
	"html/template"
	"github.com/gobuffalo/packr"
//	"github.com/davecgh/go-spew/spew"
)


type KV struct {
	K string
	V interface{}
}

type tmplRH struct {
	Req     interface{}
	Headers []KV
}

func (this *Server) rh2tmpl(r *http.Request) tmplRH {
	t := tmplRH{}
	t.Req = r

	tb := []KV{}
	for k,v := range r.Header {
		tb = append(tb, KV{K: k, V: v})
	}
	sort.Slice(tb, func(i,j int) bool {
		return tb[i].K < tb[j].K
	})
	t.Headers = tb

	return t
}


func (this *Server) handlerShowHeaders(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATH: %s", r.URL.Path)

	box := packr.NewBox("../../assets/templates")
	s, err := box.FindString("headers.tpl")
	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}

	t,err := template.New("T").Parse(s)
	// log.Print( "debug: \n" + spew.Sdump(tp) )

	err = t.Execute(w, this.rh2tmpl(r))

	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}


/*

	type Data struct {
		Key    string
		Values []string
	}
	type Table []Data
	var table Table

	for k,v := range r.Header {
		table = append(table,Data{Key: k, Values: v})
	}
	sort.Slice(table,func(i,j int) bool {
		return table[i].Key < table[j].Key
	})

	log.Print( "debug: \n" + spew.Sdump(table) )

	box := packr.NewBox("../../assets/templates")

	s, err := box.FindString("headers.tpl")
	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}

	t,err := template.New("T").Parse(s)

	err = t.Execute(w, struct {
		Title      string
		TblReqHdr  Table
	}{
		Title:     "Headers",
		TblReqHdr:  table,
	})

	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}
*/

}
