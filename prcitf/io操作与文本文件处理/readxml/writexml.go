package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name      `xml:"servers"`
	Version string        `xml:"version,attr"`
	Svs     []writeServer `xml:"writeServer"`
}

type writeServer struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP>serverIp1>serverIp2"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, writeServer{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, writeServer{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " - ", "  --  ")
	//output,err=xml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}