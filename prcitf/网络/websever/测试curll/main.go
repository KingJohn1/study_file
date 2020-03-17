// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	var n =0
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("%d  %#v",n,r)
		w.WriteHeader(404)
		w.Write([]byte("hello world"))
	})
	log.Fatal(http.ListenAndServe(":9090",nil))
}



