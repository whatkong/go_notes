package main

import (
	"./filelist"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		err := handler(w, req)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}

			http.Error(w, http.StatusText(code), code)
		}

	}
}

func main() {
	//req,err := http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)
	//if err != nil{
	//	panic(" network error")
	//}
	//req.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//
	//resp, _ := http.DefaultClient.Do(req)
	//defer resp.Body.Close()
	//content,_ := httputil.DumpResponse(resp,true)
	//fmt.Printf("%s\n",content)

	http.HandleFunc("/list/", errWrapper(filelist.FileListingHandler))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
