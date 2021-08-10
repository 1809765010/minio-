package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type objectAPIHandlers struct {
	ObjectAPI func() ObjectLayer
	CacheAPI  func() CacheObjectLayer
}

func (api objectAPIHandlers) sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	objectAPI := api.ObjectAPI()
	w.WriteHeader(http.StatusOK)                // 设置响应状态码为 200
	fmt.Fprintf(w, objectAPI.A1("hello world")) // 发送响应到客户端
}
func newObjectLayerFn() ObjectLayer {
	//xxx
	return globalObjectAPI
}
func newCachedObjectLayerFn() CacheObjectLayer {
	//xxx
	return globalCacheObjectAPI
}
func Router() {
	api := objectAPIHandlers{
		ObjectAPI: newObjectLayerFn,
		CacheAPI:  newCachedObjectLayerFn,
	}
	r := mux.NewRouter()
	r.HandleFunc("/hello", api.sayHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", r))
}
