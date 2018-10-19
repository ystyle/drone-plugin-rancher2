package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	//跳过证书验证
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{
		Transport: tr,
	} // http client
	api       string // rancher api地址
	accessKey string // rancher access key
	secretKey string // rancher secret key
)

// DATA 字段元素
type container struct {
	Name  string `json:"name"`  // 名称
	Image string `json:"image"` // 镜像
}

// DATA 字段
type containers []container

// 取容器名称对应的image
func (cs containers) get(name string) string {
	for i := 0; i < len(cs); i++ {
		c := cs[i]
		if c.Name == name {
			return c.Image
		}
	}
	return ""
}

// 构建请求并返回结果
func buildRequest(method string, data string) []byte {
	req, err := http.NewRequest(method, api, strings.NewReader(data))
	if err != nil {
		log.Printf("An error occurred during build request: %s", err.Error())
		os.Exit(1)
	}
	req.SetBasicAuth(accessKey, secretKey)
	resp, err := client.Do(req)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		log.Printf("An error occurred during get rancher workload metadata: %s", err.Error())
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("An error occurred during reading rancher workload metadata: %s", err.Error())
		os.Exit(1)
	}
	return body
}

// 格式化json
func formatJson(data []byte) string {
	var out bytes.Buffer
	json.Indent(&out, data, "", "  ")
	return string(out.Bytes())
}

func init() {
	// 初始化数据
	api = os.Getenv("PLUGIN_API")
	accessKey = os.Getenv("PLUGIN_ACCESS_KEY")
	secretKey = os.Getenv("PLUGIN_SECRET_KEY")
}

func main() {
	data := os.Getenv("PLUGIN_DATA")
	// 打印参数
	log.Printf("\n API %s \n ACCESS_KEY: %s\n SECRET_KEY: %s\n DATA: %s", api, accessKey, secretKey, data)

	// 解析参数
	cs := &containers{}
	err := json.Unmarshal([]byte(data), &cs)
	if err != nil {
		log.Printf("An error occurred during pasre params: data,  %s", err.Error())
		os.Exit(1)
	}
	// 请求数据
	body := buildRequest("GET", "")
	log.Printf("receive data,  %s", formatJson(body))
	// 解析containers
	var apiData map[string]interface{}
	json.Unmarshal([]byte(body), &apiData)
	containersData, ok := apiData["containers"]
	if !ok {
		log.Printf("An error occurred during parse rancher workload metadata: %s", err.Error())
		os.Exit(1)
	}
	// 查找参数对应名称的容器并把修改为参数的镜像
	containersArrays := containersData.([]interface{})
	for _, csItemData := range containersArrays {
		csItem := csItemData.(map[string]interface{})
		if nameData, ok := csItem["name"]; ok {
			name := nameData.(string)
			image := cs.get(name)
			if image != "" {
				csItem["image"] = image
			}
		}
	}
	// 调用请求更新rancher workload
	target, err := json.Marshal(apiData)
	log.Printf("POST DATA: %s\n", formatJson(target))
	body = buildRequest("PUT", string(target))
	log.Printf("result: %s\n", formatJson(body))
	log.Println("update successfull")
}
