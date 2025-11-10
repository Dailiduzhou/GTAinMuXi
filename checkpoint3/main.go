package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWTPayload struct {
	Code string `json:"code"`
	Iat  int64  `json:"iat"`
	Nbf  int64  `json:"nbf"`
}

type MergedJWT struct {
	Header  JWTHeader  `json:"header"`
	Payload JWTPayload `json:"payload"`
}

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
	req1, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)

	// write your code below
	// ...
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiT01FTiIsImlhdCI6MTc2Mjc3MjYyNSwibmJmIjoxNzYyNzcyNjI1fQ.vbdUhVyXAwJ35N5uWQVYmJ--MEbDDEgGdnTgRPcBI28"
	merged, err := decodeJWTStructured(passport)
	if err != nil {
		fmt.Printf("error decoding: %q", err)
	}
	mergedJSON, err := json.Marshal(merged)
	if err != nil {
		fmt.Printf("error marshaling: %q", err)
	}
	req1.Req.Header.Set("passport", string(mergedJSON))
	resp, err := req1.SendRequest()
	if err != nil {
		fmt.Printf("error sending request: %q", err)
	}
	resp.ShowBody()
}

func decodeJWTStructured(jwt string) (*MergedJWT, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("无效的 JWT 格式")
	}

	// 解码头部
	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}

	// 解码载荷
	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var header JWTHeader
	var payload JWTPayload

	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return nil, err
	}

	return &MergedJWT{
		Header:  header,
		Payload: payload,
	}, nil
}
