package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	g "github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

// get 操作
func testGet(c *g.Context) {

	url := "http://161.117.0.57:10082/api/v1/sql/server"

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

}

/// post 操作
func testPost(c *g.Context) {

	requestBody := fmt.Sprintf(`
	{
    "config":{
            "channelName":"assetpublish",
            "chainCodeName":"weiFinanceTest2Up",
            "funcName":"query"
    },
    "data": ["%s"],
    "user":{
        "dataTime":"2019-06-27",
        "uuid":"dl-01-dp-04-1f7c1c15-12c6-4053-b92e-xxx1",
        "orgName":"瑞泰格公司",
        "orgId":"dc25f534-1e1b-4903-9694-90516b09732f",
        "userName":"张飞飞",
        "userId":"6a2149a7-a487-4fca-985a-de54aa3d99d4",
        "peer":"peer0",
        "anchor":"and",
        "affiliationId":"DL-XE-AS-04"
    }
}
`, "523137470488705")
	var jsonStr = []byte(requestBody)

	url := "http://161.117.0.57:10082/api/v1/wei/query"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

}
