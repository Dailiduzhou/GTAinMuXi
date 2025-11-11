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

const filePath = "C:\\Users\\OMEN\\Desktop\\GTAinMuXi\\GTAinMuXi\\checkpoint3\\test.jpg"
const filePath1 = "C:\\Users\\OMEN\\Desktop\\GTAinMuXi\\combination.go"

func main() {
	// req, err := httptool.NewRequest(
	// 	httptool.GETMETHOD,
	// 	"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
	// 	"",
	// 	httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req1, err := httptool.NewRequest(
	// 	httptool.GETMETHOD,
	// 	"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key",
	// 	"",
	// 	httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(req)

	// write your code below
	// ...
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiT01FTiIsImlhdCI6MTc2Mjc3MjYyNSwibmJmIjoxNzYyNzcyNjI1fQ.vbdUhVyXAwJ35N5uWQVYmJ--MEbDDEgGdnTgRPcBI28"

	// req1.Req.Header.Set("passport", passport)
	// resp, err := req1.SendRequest()
	// if err != nil {
	// 	fmt.Printf("error sending request: %q", err)
	// }
	// resp.ShowHeader()
	// resp.ShowBody()
	// secret := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	// errCode, err := encrypt.Base64Decode(secret)
	// if err != nil {
	// 	fmt.Printf("error decoding: %q", err)
	// }
	// fmt.Println(errCode)
	// secret_key := []byte("MuxiStudio203304")
	// error_code := []byte("for {go func(){time.Sleep(1*time.Hour)}()}")
	// encrption, err := encrypt.AESEncryptOutInBase64(error_code, secret_key)
	// if err != nil {
	// 	fmt.Printf("error encrtpting: %q", err)
	// }
	// req2, err := httptool.NewRequest(
	// 	httptool.PUTMETHOD,
	// 	"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate",
	// 	string(encrption),
	// 	httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req2.Req.Header.Set("passport", passport)
	// resp1, err := req2.SendRequest()
	// if err != nil {
	// 	fmt.Printf("error sending request: %q", err)
	// }
	// resp1.ShowBody()

	// req3, err := httptool.NewRequest(
	// 	httptool.GETMETHOD,
	// 	"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
	// 	"",
	// 	httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req3.Req.Header.Set("passport", passport)
	// resp2, err := req3.SendRequest()
	// if err != nil {
	// 	fmt.Printf("error sending request: %q", err)
	// }
	// resp2.ShowHeader()
	// resp2.ShowBody()

	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Printf("Failed to get working directory: %v", err)
	// }

	// req4, err := httptool.NewRequest(
	// 	httptool.GETMETHOD,
	// 	"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
	// 	filePath,
	// 	httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	// )
	// if err != nil {
	// 	fmt.Printf("error new: %q", err)
	// }
	// req4.Req.Header.Set("passport", passport)
	// resp3, err := req4.SendRequest()
	// if err != nil {
	// 	fmt.Printf("error sending request: %q", err)
	// }
	// // resp3.ShowHeader()
	// // resp3.ShowBody()
	// err = resp3.Save(filePath)
	// if err != nil {
	// 	fmt.Printf("error saving: %q", err)
	// }

	// req5, err := httptool.NewRequest(
	// 	httptool.POSTMETHOD,
	// 	"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
	// 	filePath,
	// 	httptool.FILE,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req5.Req.Header.Set("passport", passport)
	// resp4, err := req5.SendRequest()
	// if err != nil {
	// 	fmt.Printf("error sending request: %q", err)
	// }
	// resp4.ShowHeader()
	// resp4.ShowBody()

	// req6, err := httptool.NewRequest(
	// 	httptool.GETMETHOD,
	// 	"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
	// 	"",
	// 	httptool.DEFAULT,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req6.Req.Header.Set("passport", passport)
	// resp5, err := req6.SendRequest()
	// if err != nil {
	// 	fmt.Printf("error sending request: %q", err)
	// }
	// resp5.ShowHeader()
	// resp5.ShowBody()

	req7, err := httptool.NewRequest(
		httptool.POSTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
		filePath1,
		httptool.FILE,
	)
	if err != nil {
		fmt.Println(err)
	}
	req7.Req.Header.Set("passport", passport)
	resp6, err := req7.SendRequest()
	if err != nil {
		fmt.Printf("error sending request: %q", err)
	}
	resp6.ShowHeader()
	resp6.ShowBody()
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
