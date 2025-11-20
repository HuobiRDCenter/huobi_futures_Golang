package main

import (
	"bytes"
	"compress/gzip"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

func sign2(privateKeyBase64 string, method string, host string, path string, parameters string) (string, error) {
	if method == "" || host == "" || path == "" {
		return "", errors.New("failed to decode PEM block containing the private key")
	}
	// 解析 PEM 编码的私钥
	// print(len(privateKeyBase64))
	block, _ := pem.Decode([]byte(privateKeyBase64))
	// print(len(block.Bytes))
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("failed to decode PEM block containing the private key")
	}

	// 解析 x509 格式的私钥
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "nil", err
	}

	// 类型断言为 ed25519.PrivateKey
	var ok bool
	var privateKey ed25519.PrivateKey
	privateKey, ok = parsedKey.(ed25519.PrivateKey)
	if !ok {
		return "nil", errors.New("not an ed25519 private key")
	}

	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	sb.WriteString(host)
	sb.WriteString("\n")
	sb.WriteString(path)
	sb.WriteString("\n")
	sb.WriteString(parameters)

	// 拼接生成的字符串
	payload := sb.String()

	// 生成签名
	signature := ed25519.Sign(privateKey, []byte(payload))

	// 进行 Base64 编码
	return base64.StdEncoding.EncodeToString(signature), nil
}

func readLoop2(conn *websocket.Conn) {
	for conn != nil {
		msgType, buf, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Read error: %s", err)
			continue
		}
		var message string
		if msgType == websocket.BinaryMessage {
			message, err = gzipDecompress2(buf)
			if err != nil {
				fmt.Printf("UnGZip data error: %s", err)
				continue
			}
		} else if msgType == websocket.TextMessage {
			message = string(buf)
		}
		fmt.Println(message)
	}
}
func sendAuth2(conn *websocket.Conn, host string, path string, accessKey string, secretKey string) bool {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	urls := url.Values{}
	urls.Add("AccessKeyId", accessKey)
	urls.Add("SignatureMethod", "Ed25519")
	urls.Add("SignatureVersion", "2")
	urls.Add("Timestamp", timestamp)

	signature, err := sign2(secretKey, "GET", host, path, urls.Encode())
	if err != nil {

	}
	fmt.Println("signature:", signature)

	auth := fmt.Sprintf("{\"op\":\"auth\", \"type\":\"api\",\"AccessKeyId\":\"%s\", \"SignatureMethod\":\"Ed25519\", \"SignatureVersion\":\"2\", \"Timestamp\":\"%s\", \"Signature\":\"%s\"}", accessKey, timestamp, signature)
	fmt.Println("send auth:", auth)
	jdata := []byte(auth)
	conn.WriteMessage(websocket.TextMessage, jdata)
	return true
}
func gzipDecompress2(input []byte) (string, error) {
	buf := bytes.NewBuffer(input)
	reader, gzipErr := gzip.NewReader(buf)
	if gzipErr != nil {
		return "", gzipErr
	}
	defer reader.Close()

	result, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return "", readErr
	}

	return string(result), nil
}
func main() {
	host := "api.hbdm.com"
	path := "/swap-notification"
	ak := ""
	sk := ``

	url := fmt.Sprintf("wss://%s%s", host, path)
	fmt.Println("url:", url)

	//head := http.Header{"Host": {host + ":8080"}}
	//conn, _, err := websocket.DefaultDialer.Dial(url, head)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("WebSocket connected error: %s", err)
	} else {
		//fmt.Println("WebSocket connected")
		sendAuth2(conn, host, path, ak, sk)
		readLoop2(conn)
	}
}
