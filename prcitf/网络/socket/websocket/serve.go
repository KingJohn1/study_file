// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func pushMessage(ws *websocket.Conn){
	var ch1 chan int=make(chan int)
	defer close(ch1)

	go func(){

		for {
			time.Sleep(1 * time.Second)
			ch1 <- 1
		}

	}()

	go func(){

		for {
			select{
			case <-ch1:
				websocket.Message.Send(ws,"主动发消息"+time.Now().String())
			}
		}


	}()
}

func Echo(ws *websocket.Conn) {
	var err error
	log.Println("receive msg")


	//pushMessage()

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive",err)
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}


	}
}

func view(w http.ResponseWriter, req *http.Request){

	readFile, err := ioutil.ReadFile("socket/websocket/client_html.html")
	if err!=nil{
		http.Error(w,"open html fail",500)
	}
	fmt.Println("收到消息")
	w.Write(readFile)
}

func main(){

	http.Handle("/",websocket.Handler(Echo))
	http.HandleFunc("/test/",view)
	if err:=http.ListenAndServe(":1234",nil);err!=nil{
		log.Fatal(" err:",err)
	}
}
