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
}

func handler(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "text/html")
	res.Header().Set("X-Frame-Options", "DENY")
	res.Header().Set("X-XSS-Protection", "1; mode=block")
	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.Header().Set("Content-Security-Policy", "default-src 'self'")

	var status = 404
	const html = `<doctype html>
<html>
	<head>
		<script src="/js/title.js"></script>
	</head>
	<body>
	</body>
</html>`

	// 最後に出力するコンテンツ
	var content string

	var pathSplit = strings.Split(req.URL.Path[1:], "/")
	switch(len(pathSplit)){
	case 1:
		switch(pathSplit[0]){
		case "settings":
		case "login":
		case "logout":
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
		var resourcePath = fmt.Sprintf("./resources%s", req.URL.Path)
		var data, error = ioutil.ReadFile(resourcePath)
		if(error != nil){
			content = html
		}else{
			var index = strings.LastIndex(req.URL.Path, ".")
			if(index != -1){
				var extension = req.URL.Path[index:]
				var mimeType = mime.TypeByExtension(extension)
				res.Header().Set("Content-Type", mimeType)
			}else{
				var mimeType = mime.TypeByExtension(".txt")
				res.Header().Set("Content-Type", mimeType)
			}

			status = 200
			content = string(data)
		}
	}

	res.WriteHeader(status)
	fmt.Fprintf(res, "%s", content)
}

func main(){
	var configData, error = ioutil.ReadFile("./config.json")
	if(error != nil){
		fmt.Println(error)
	}

	var config Config
	error = json.Unmarshal(configData, &config)
	if(error != nil){
		fmt.Println(error)
	}

	var mux = http.NewServeMux()
	mux.HandleFunc("/", handler)
	fmt.Println("Akitter API Server KIDOU!")

	var port = ":" + strconv.Itoa(config.Port)
	http.ListenAndServe(port, mux)
}