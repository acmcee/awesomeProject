package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	message string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.message = ""
	s.message += "有傻逼来访问啦哈哈哈哈\n"
	s.message += fmt.Sprintln("r.URL.Path:", r.URL.Path)
	s.message += fmt.Sprintln("r.Method：", r.Method)
	s.message += fmt.Sprintln("r.Host:", r.Host)
	s.message += fmt.Sprintln("r.URL.Query():", r.URL.Query())
	s.message += fmt.Sprintln("r.URL.RawQuery", r.URL.RawQuery)
	s.message += fmt.Sprintln("r.Body:", r.Body)
	s.message += fmt.Sprintln("r.RequestURI", r.RequestURI)
	s.message += fmt.Sprintln("a:", r.URL.Query().Get("a"))
	s.message += "\n"
	s.message += "\n"
	s.message += "\n"
	s.message += fmt.Sprintln(r)
	fmt.Fprintf(w, s.message)
}

func main() {
	http.Handle("/", &Server{})
	http.HandleFunc()
	err := http.ListenAndServe(":8100", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
