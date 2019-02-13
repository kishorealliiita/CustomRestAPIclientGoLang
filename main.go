package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"restApi"
	"flag"
)


//parse Yaml onbarding file.
type conf struct {
	HostName         string   `yaml:"HostName"`
	EndPoint string   `yaml:"Endpoint"`
	ApiKey string   `yaml:"ApiKey"`
	Method string `yaml:"Method"`
	Headers    struct {
		Header1   string                `yaml:"Header1"`
		Header2   string 				`yaml:"Header2"`
	} `yaml:"Headers"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		fmt.Println("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()
	fmt.Println(c)

	var hostName string
	var apiKey string
    flag.StringVar(&hostName, "host_name", "http://localhost", "host name for onboarding.")
	flag.StringVar(&apiKey, "api_key", "", "Api Key to access the host.")
	
	flag.Parse()

	fmt.Println ("host name is ", hostName)
	fmt.Println ("api key is ", apiKey)

	OnboardOctopus (c, hostName, apiKey)
}


func (h* Host) getContents (c conf) (*Response, error) {
	method := c.Method
	baseURL := c.HostName
	key := c.ApiKey
	endpoint := c.EndPoint
	//Headers := make(map[string]string)
	Headers := c.Headers

	data = c.Data
	request := Request{
		Method:      method,
		BaseURL:     baseURL,
		Headers:     Headers,
		Body: data,
		Endpoint : endpoint,
	}
	req, e := BuildRequestObject(request)
	if e != nil {
		fmt.Println("Rest failed to BuildRequest. Returned error: %v", e)
	}
	if req == nil {
		fmt.Println("Failed to BuildRequest.")
	}

	//Start PrintRequest
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error : %v", err)
	}
	fmt.Println("Request : ", string(requestDump))

	customClient := &Client{&http.Client{Timeout: time.Millisecond * 10000}}
	resp, err := customClient.Send(request)
	fmt.Println ("response status code is - ", resp.StatusCode)
	code := resp.StatusCode
	if code != 201 {
		fmt.Println ("Error: Team creation failed. - ", resp.Body)
	} else {
		fmt.Println ("Team created successfully. - ", teamName)
	}

	return resp, err
}