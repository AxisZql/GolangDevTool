package main

import (
	"fmt"
	"net/http"
)
func index(write http.ResponseWriter,req *http.Request){
	fmt.Fprintf(write,"This is homePage")
}
func main(){
	fmt.Printf("Hello world")
	//CreateSsl()
	mux := http.NewServeMux()//创建多路复用器
	mux.HandleFunc("/",index)
	server :=http.Server{
		Addr: "127.0.0.1:9000",
		Handler: mux,
	}
	server.ListenAndServeTLS("cert.pem","key.pem")//必须是使用https协议才能访问
}
