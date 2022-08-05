package main

import (
	"awesomeGo/errhandling/filelistingserver/customererror"
	"awesomeGo/errhandling/filelistingserver/filelist"
	"log"
	"net/http"
	"os"
)

// 定义了一个名为 appHandler，类型为 func(writer http.ResponseWriter, request *http.Request)，返回值类型为 error 的类型
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 统一的错误处理
func errorWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v\n", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {

			// 打印在控制台上的信息
			log.Printf("Error occurred handling request: %s\n", err.Error())

			// 可以展示给用户看的错误信息
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			// test
			if cusErr, ok := err.(commonerror.CustomerError); ok {
				http.Error(writer, cusErr.Msg(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	// ⚠️ 将 filelist.HandleFileList 作为参数传给 http.HandleFunc
	http.HandleFunc("/", errorWrapper(filelist.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
