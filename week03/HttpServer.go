package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	//	"net/http/pprof"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {

	// 2、设置version
	os.Setenv("version", "V1.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	vvv := w.Header().Get("VERSION")

	fmt.Printf("os version11: %s \n", vvv)

	w.Write([]byte("<h1> This is a test page</h1>"))
	// 1、将requst中的header 设置到reponse中
	for k, v := range r.Header {

		for _, vv := range v {
			//	fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}

	// 3、记录日志并输出
	logFile, err := os.OpenFile("./logs/HttpServer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	ClientIp := getCurrentIP(r)
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Printf("Success! clientIp: %s", ClientIp)
}

func getCurrentIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// Fprintf: 来格式化并输出到 io.Writers
	fmt.Fprintf(w, "server is working")
}

func main() {
	//通过ServerMux实现一个http服务端
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start http server failed , error: %s\n", err.Error())
	}

}
