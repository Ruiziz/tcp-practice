package main

import (
	"encoding/json"
	"git.code.oa.com/trpc-go/trpc-go/log"

	//"git.code.oa.com/trpc-go/trpc-go"
	"net/http"
	"strconv"

	_ "git.code.oa.com/trpc-go/trpc-config-tconf"
	trpc "git.code.oa.com/trpc-go/trpc-go"
	thttp "git.code.oa.com/trpc-go/trpc-go/http"
)

func main() {
	// http://127.0.0.1:8000/add?a=1&b=2
	// 加封装
	s := trpc.NewServer()
	thttp.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) error {
		_ = request.ParseForm() //解析参数
		//fmt.Println("test")
		//fmt.Println("path:", request.URL.Path)
		//拿到第一个a
		a, _ := strconv.Atoi(request.Form["a"][0])
		// 拿到第二个b
		b, _ := strconv.Atoi(request.Form["b"][0])
		writer.Header().Set("Content-Type", "application/json") //固定写法，规定为json数据协议
		// 转换成json形式以传递
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		//写回返回体
		_, _ = writer.Write(jData)
		return nil
	})
	thttp.RegisterDefaultService(s)

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}

/*
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm() //解析参数
		fmt.Println("test")
		fmt.Println("path:", request.URL.Path)
		//拿到第一个a
		a, _ := strconv.Atoi(request.Form["a"][0])
		// 拿到第二个b
		b, _ := strconv.Atoi(request.Form["b"][0])
		writer.Header().Set("Content-Type", "application/json") //固定写法，规定为json数据协议
		// 转换成json形式以传递
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		//写回返回体
		_, _ = writer.Write(jData)
	})
	// 开启监听
	http.ListenAndServe(":8000", nil) */
}
