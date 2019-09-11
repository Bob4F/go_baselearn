package main
/*
  Go语言的第一个web程序
*/
import (
"fmt"
"log"
"net/http"
)

func main() {
	// HandlerFunc(路径映射，处理请求的方法)
	http.HandleFunc("/first",webSayHello)
	// :8080 web应用监听的端口 ；后面一个参数，我们现阶段为空就好了，会默认调用我们注入HandleFunc的方法 webSayHello
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		log.Fatal("ListenAndServer:",err)
	}
}

func webSayHello(w http.ResponseWriter,r *http.Request){
	// 返回给客户端的字符串
	fmt.Fprintf(w,"Hello Go_Web")
}