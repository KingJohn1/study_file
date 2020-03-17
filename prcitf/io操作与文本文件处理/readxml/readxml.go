package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"writeServer"`
	Description string   `xml:",innerxml"`
	TestAttr    string   `xml:"serverName,testAttr"`
	TestNoMatch string //`xml:"-"`
	TestMatchAny string `xml:",any"`
	TestAnnotate string `xml:",comments"`
	TestDeep string `xml:"default3>default4>deep"`

}

type server struct {
	XMLName    xml.Name `xml:"writeServer"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("io操作与文本文件处理/readxml/data/servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%#v", string(data))
	fmt.Println()
	fmt.Println()
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%#v", v)
}
