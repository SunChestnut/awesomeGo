package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// errPanic 函数的参数➕返回值 与 appHandler 的类型相同，因此 errPanic 与 appHandler 类型相同
func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 构造一个 web server 来测试 errorWrapper 函数
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errorWrapper(tt.h)

		// httptest.NewRecorder() 返回值为 ResponseRecorder，ResponseRecorder 实现了 ResponseWriter
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://gobyexample.com/random-numbers", nil)

		// 调用 f 函数
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// 启动一个 web server 来测试 errorWrapper 函数
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		// ❓
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	// 将字节数组 b 转换成字符串，并去掉字符串末尾的换行符
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("except (%d, %s); got (%d %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
