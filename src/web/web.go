package web

import (
	"log"
	"net/http"
	"github.com/gobuffalo/packr"
)

type Server struct {
}

func New() *Server {
	return new(Server)
}

func (this *Server) Run(listen string) error {

	box := packr.NewBox("../../assets/public")
	http.HandleFunc("/",this.handlerShowHeaders)
	http.Handle("/static/",http.StripPrefix("/static",http.FileServer(box)))

	log.Printf("info: Start Listening at http://%s/", listen)
	if err := http.ListenAndServe(listen, nil); err != nil {
		return err
	}
	return nil
}

