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

func RootHTTPHandler(w http.ResponseWriter, r *http.Request) {
	message := ""
	message += "有傻逼来访问啦哈哈哈哈\n"
	message += fmt.Sprintln("r.URL.Path:", r.URL.Path)
	message += fmt.Sprintln("r.Method：", r.Method)
	message += fmt.Sprintln("r.Host:", r.Host)
	message += fmt.Sprintln("r.URL.Query():", r.URL.Query())
	message += fmt.Sprintln("r.URL.RawQuery", r.URL.RawQuery)
	message += fmt.Sprintln("r.Body:", r.Body)
	message += fmt.Sprintln("r.RequestURI", r.RequestURI)
	message += fmt.Sprintln("a:", r.URL.Query().Get("a"))
	message += "\n"
	message += "\n"
	message += "\n"
	message += fmt.Sprintln(r)
	fmt.Fprintf(w, message)
}

func main() {
	// 测试使用对应的结构体来作为handler
	//http.Handle("/", &Server{})
	// 测试使用func 来作为handler
	http.HandleFunc("/", RootHTTPHandler)
	err := http.ListenAndServe(":8100", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
