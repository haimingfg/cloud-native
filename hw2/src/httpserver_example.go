package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// 1.接收客户端 request，并将 request 中带的 header 写入 response header
// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func hello(w http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.Header)
	for k, v := range req.Header {
		w.Header().Add(k, strings.Join(v, ";"))
	}

	// VERSION
	w.Header().Add("GO VERSION", runtime.Version())
	w.Write([]byte("Ok ba"))

	for k, v := range w.Header() {
		w.Write([]byte(k + ":" + strings.Join(v, " ") + "\n"))
	}
}

// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func accesssFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIp := ReadUserIP(r)
		fmt.Printf("clientIp: %s\n, http status :%d\n", clientIp, http.StatusOK)
		h.ServeHTTP(w, r)
	})
}

// 4. 当访问 localhost/healthz 时，应返回200
func healthz(w http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("healthz, http status :%d\n", http.StatusOK)
	w.Write([]byte(s))
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", healthz)
	fmt.Printf("start...\n")
	http.ListenAndServe(":8080", accesssFunc(http.DefaultServeMux))
}
