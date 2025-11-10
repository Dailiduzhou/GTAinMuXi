// Passport : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiT01FTiIsImlhdCI6MTc2Mjc3MjU5MiwibmJmIjoxNzYyNzcyNTkyfQ.btvPcam3oegSr8rCcKHZQmY6O2fNHh7ZBtK-tqI3v_M
package main

import (
	"fmt"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)

	// write your code below
	// ...
	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	resp.ShowHeader()

}
