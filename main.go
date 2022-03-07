package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type Message struct {
	Content string `json:"content"`
}

func getInterfaces() string {
	var msg string = ""

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	for _, i := range ifaces {

		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, a := range addrs {
			var isUP bool = false
			if strings.Contains(i.Flags.String(), "up") {
				isUP = true
			}
			msg += "Index: " + strconv.FormatInt(int64(i.Index), 32) + "\nName: " + string(i.Name) + "\nMAC: " + string(a.String()) + "\nStatus: " + strconv.FormatBool(isUP) + "\n\n"
		}

	}
	return msg
}

func main() {
	var interfaces string = getInterfaces()

	client := http.Client{}

	bd := Message{
		Content: interfaces,
	}

	body, err := json.Marshal(bd)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	var url string = ""
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Errorf(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println(resp)
}
