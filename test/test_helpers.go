package test

import (
	"io/ioutil"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// AssertContentType 比对 respose content-type 是否一致
func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Header().Get("content-type"))
	}
}

// AssertDeepEqual 使用 reflect.DeepEqual 比对两个任意类型是否相等
func AssertDeepEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}

// AssertEquals 比对两个基本类型是否相等
func AssertEquals(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("did not equal, got %#v, want %#v", got, want)
	}
}

// AssertNoError 检查如果 err 则退出测试
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
	}
}

// CreateTempFile 创建临时文件, 返回文件关闭函数句柄
func CreateTempFile(t *testing.T, initialData string, filename string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", filename)

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, err = tmpfile.Write([]byte(initialData))
	if err != nil {
		t.Fatalf("could not write temp file %v", err)
	}

	removeFile := func() {
		err := os.Remove(tmpfile.Name())
		if err != nil {
			t.Errorf("could not remove temp file %v", err)
		}
	}

	return tmpfile, removeFile
}

// Within 在指定时间 d 内, 完成 assert 函数, 否则返回测试错误
func Within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

// RetryUntil 在指定时间内, 重复执行 assert 函数, 如果成果返回 true, 否则测试失败
func RetryUntil(t *testing.T, d time.Duration, assert func() bool) bool {
	t.Helper()
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if assert() {
			return true
		}
	}
	t.Errorf("failed after %v time tying.", d)
	return false
}

// AssertWebsocketGotMsg 比对从 ws 中读取的信息是否与 want 相等
func AssertWebsocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf(`got "%s", want "%s"`, string(msg), want)
	}
}
