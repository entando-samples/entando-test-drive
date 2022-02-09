package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	baseURL       string
	adminUserName string
	adminPassword string
}

func getToken(defaultConfig Config) string {
	tokenURL := fmt.Sprintf("%s/auth/realms/master/protocol/openid-connect/token", defaultConfig.baseURL)
	data := url.Values{}
	data.Set("username", defaultConfig.adminUserName)
	data.Set("password", defaultConfig.adminPassword)
	data.Set("client_id", "admin-cli")
	data.Set("grant_type", "password")

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Printf("err %s", err)
		return ""
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("err %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var resultMap map[string]interface{}

	unmarshalErr := json.Unmarshal(body, &resultMap)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}

	return resultMap["access_token"].(string)
}
func getRealmDeatil(defaultConfig Config, sessionToken string) string {
	realmDetailURL := fmt.Sprintf("%s/auth/admin/realms/entando/", defaultConfig.baseURL)

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", realmDetailURL, nil)
	if err != nil {
		fmt.Printf("err %s", err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", sessionToken))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("err %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func updateRealmDetails(defaultConfig Config, sessionToken string, newRealm string) int {
	realmDetailURL := fmt.Sprintf("%s/auth/admin/realms/entando/", defaultConfig.baseURL)

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("PUT", realmDetailURL, strings.NewReader(newRealm))
	if err != nil {
		fmt.Printf("err %s", err)
		return 0
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", sessionToken))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("err %s", err)
		return 0
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func main() {

	verbose := flag.Bool("v", false, "prints all the logs")

	unsecureCERT := flag.Bool("u", false, "allows unsecure certs")

	flag.Parse()

	if *unsecureCERT {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	args := flag.Args()

	defaultConfig := Config{
		baseURL:       args[0],
		adminUserName: args[1],
		adminPassword: args[2],
	}

	if *verbose {
		fmt.Printf("BASE URL: %s \n", defaultConfig.baseURL)
		fmt.Printf("adminUserName: %s \n", defaultConfig.adminUserName)
		fmt.Printf("adminPassword: %s \n", defaultConfig.adminPassword)
	}

	sessionToken := getToken(defaultConfig)
	if *verbose {
		fmt.Printf("SESSION TOKEN: %s \n", sessionToken)
	}
	currentRealmConfig := getRealmDeatil(defaultConfig, sessionToken)

	if *verbose {
		fmt.Printf("CURRENT REALM CONFIG: %s \n", currentRealmConfig)
	}

	var resultMap map[string]interface{}

	unmarshalErr := json.Unmarshal([]byte(currentRealmConfig), &resultMap)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}

	if *verbose {
		fmt.Printf("CURRENT CSP: %s \n", resultMap["browserSecurityHeaders"].(map[string]interface{})["contentSecurityPolicy"].(string))
	}

	resultMap["browserSecurityHeaders"].(map[string]interface{})["contentSecurityPolicy"] = "frame-src 'self'; object-src 'none';"
	resultMap["browserSecurityHeaders"].(map[string]interface{})["xFrameOptions"] = ""

	newRealm, err := json.Marshal(resultMap)
	if err != nil {
		return
	}

	if *verbose {
		fmt.Printf("NEW REALM: %s \n", newRealm)
	}
	retCode := updateRealmDetails(defaultConfig, sessionToken, string(newRealm))

	if *verbose {
		fmt.Printf("RETCODE: %v \n", retCode)
	}
}
