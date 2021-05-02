package main

import(
	"fmt"
	"io/ioutil"
	"mime"
	"encoding/json"
	"strings"
	"net/http"
	"strconv"
)

type Config struct {
	Port int `json:"port"`
	API string `json:API`
}
var config Config

func handler(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type", "text/html; charset=UTF-8")
	response.Header().Set("X-Frame-Options", "DENY")
	response.Header().Set("X-XSS-Protection", "1; mode=block")
	response.Header().Set("X-Content-Type-Options", "nosniff")
	response.Header().Set("Content-Security-Policy", "default-src 'self'")

	var status = 404
	const html = `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="/css/reset.css">
		<link rel="stylesheet" href="/css/main.css">
		<script src="/js/title.js"></script>
		<script src="/js/templete.js"></script>
	</head>
	<body>
	</body>
</html>`

	// 最後に出力するコンテンツ
	var content string

	// 最初のスラッシュ削除しつつスライス？配列？に分割
	var pathSplit = strings.Split(request.URL.Path[1:], "/")

	switch(len(pathSplit)){
	case 1:
		switch(pathSplit[0]){
		case "": fallthrough
		case "settings": fallthrough
		case "login": fallthrough
		case "logout": fallthrough
		case "signup":
			status = 200
			content = html
			break
		}
		break

	case 2:
		switch(pathSplit[0]){
		case "tweet":
			status = 200
			content = html
			break

		case "profile":
			status = 200
			content = html
			break
		}
		break

	default:
		content = html
		break
	}

	if(content == ""){
		var resourcePath = fmt.Sprintf("./resources%s", request.URL.Path)
		var bytes, error = ioutil.ReadFile(resourcePath)
		if(error != nil){
			content = html
		}else{
			var index = strings.LastIndex(request.URL.Path, ".")
			if(index != -1){
				var extension = request.URL.Path[index:]
				var mimeType = mime.TypeByExtension(extension) + "; charset=UTF-8"
				response.Header().Set("Content-Type", mimeType)
			}else{
				var mimeType = mime.TypeByExtension(".txt") + "; charset=UTF-8"
				response.Header().Set("Content-Type", mimeType)
			}

			status = 200
			content = string(bytes)
		}
	}

	response.WriteHeader(status)
	fmt.Fprintf(response, "%s", content)
}

func main(){
	var bytes, err = ioutil.ReadFile("./config.json")
	if(err != nil){
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(bytes, &config)
	if(err != nil){
		fmt.Println(err)
		return
	}

	var port = ":" + strconv.Itoa(config.Port)

	var mux = http.NewServeMux()
	mux.HandleFunc("/", handler)
	fmt.Println("Akitter Server KIDOU!")
	http.ListenAndServe(port, mux)
}