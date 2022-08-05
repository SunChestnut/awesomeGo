package filelist

import (
	"awesomeGo/errhandling/filelistingserver/customererror"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

// Error string 类型的 userError 实现了 Error 接口
func (e userError) Error() string {
	return e.Message()
}

// Message string 类型的 userError 实现了 web.go 文件内的 userError 内的 Message 接口
func (e userError) Message() string {
	return string(e)
}

// HandleFileList 展示 list 目录下的文件信息
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {

	// test 如果 URL 中的地址以 abc 开头，则返回 CustomerError 类型的错误
	if strings.Index(request.URL.Path, "/abc/") == 0 {
		return commonerror.CustomerError("wrong path 💔")
	}

	// 判断 URL 中的地址是否以 /list/ 开头
	if strings.Index(request.URL.Path, prefix) != 0 {
		//return errors.New("path must start with " + prefix)
		return userError("path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		// 在 web 服务器中，panic 受到保护，因此即使触发了 panic，服务器也不会中止运行
		//panic(err)
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}
