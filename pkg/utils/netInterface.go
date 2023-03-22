package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const (
//VERSION = "0.0.0.1"

// STATUS string = "status"
// OK     string = "OK"
// FAIL   string = "FAIL"
// POST   string = "POST"
// GET_REQ    string = "GET_REQ"
// ERROR  string = "hasError"
)

func GetReq(url string, token string) (*http.Response, *GetError) {
	urlGet := "https://app.rocketbot.pro/api/mobile/" + url
	var jsonStr = []byte("{}")
	req, err := http.NewRequest("GET", urlGet, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-agent", "RocketBot PoS Service/Go_"+VERSION)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		_ = req.Body.Close()
		client.CloseIdleConnections()
		return nil, &GetError{message: err.Error(), statusCode: http.StatusInternalServerError}
	}

	if resp.StatusCode != 200 {
		//ReportMessage(fmt.Sprintf("response Status: %d", resp.StatusCode))
		//body, _ := io.ReadAll(resp.Body)
		//ReportMessage(fmt.Sprintf("response body:%s", string(body)))
		//ReportMessage(fmt.Sprintf("url:%s", url))
		_ = req.Body.Close()
		_ = resp.Body.Close()
		client.CloseIdleConnections()
		return nil, &GetError{message: resp.Status, statusCode: http.StatusUnauthorized}
	}
	_ = req.Body.Close()
	client.CloseIdleConnections()
	return resp, nil
}

func GETAny(url string) (*http.Response, *GetError) {
	urlGet := url
	var jsonStr = []byte("{}")
	req, err := http.NewRequest("GET", urlGet, bytes.NewBuffer(jsonStr))
	req.Header.Set("User-agent", "RocketBot PoS Service/Go_"+VERSION)
	req.Header.Set("Content-Type", "plain/text")
	req.Header.Set("Connection", "close")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		_ = req.Body.Close()
		client.CloseIdleConnections()
		return nil, &GetError{message: err.Error(), statusCode: http.StatusInternalServerError}
	}

	//fmt.Println("response Status:", resp.StatusCode)
	//fmt.Println("response Headers:", resp.Header)
	//fmt.Println("response body:", resp.Body)
	if resp.StatusCode != 200 {
		_ = req.Body.Close()
		_ = resp.Body.Close()
		client.CloseIdleConnections()
		return nil, &GetError{message: resp.Status, statusCode: http.StatusUnauthorized}
	}
	_ = req.Body.Close()
	client.CloseIdleConnections()
	return resp, nil
}

//func ContactServer(url string, endpoint string, json []byte) (*http.Response, *GetError) {
//	urlGet := "http://" + strings.TrimSpace(url) + ":7465" + endpoint
//	//fmt.Println(urlGet)
//	var jsonStr = json
//	req, err := http.NewRequest("POST", urlGet, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		return nil, &GetError{message: err.Error(), statusCode: http.StatusInternalServerError}
//	}
//	req.Header.Set("RPCUser-agent", "RocketBot PoS Service/Go_"+VERSION)
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Connection", "close")
//
//	client := &http.Client{}
//	resp, errD := client.Do(req)
//	if errD != nil {
//		_ = req.Body.Close()
//		client.CloseIdleConnections()
//		return nil, &GetError{message: errD.Error(), statusCode: http.StatusInternalServerError}
//	}
//
//	if resp.StatusCode != 200 {
//		fmt.Println("response Status:", resp.StatusCode)
//		fmt.Println("response Headers:", resp.Header)
//		_ = req.Body.Close()
//		_ = resp.Body.Close()
//		client.CloseIdleConnections()
//		return nil, &GetError{message: resp.Status, statusCode: http.StatusUnauthorized}
//	}
//	_ = req.Body.Close()
//	client.CloseIdleConnections()
//	return resp, nil
//}
//
//func ContactServerRetry(url string, endpoint string, json []byte) error {
//	urlGet := "http://" + strings.TrimSpace(url) + ":7465" + endpoint
//	var jsonStr = json
//	//fmt.Println(urlGet)
//	req, err := http.NewRequest("POST", urlGet, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		_ = req.Body.Close()
//		return errors.New(err.Error())
//		//return &GetError{message: err.Error(), statusCode: http.StatusInternalServerError}
//	}
//	req.Header.Set("RPCUser-agent", "RocketBot PoS Service/Go_"+VERSION)
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Connection", "close")
//
//	client := &http.Client{}
//	resp, errD := client.Do(req)
//	if errD != nil {
//		_ = req.Body.Close()
//		return errors.New(fmt.Sprintf("%d", resp.StatusCode))
//		//return nil, &GetError{message: errD.Error(), statusCode: http.StatusInternalServerError}
//	}
//
//	if resp.StatusCode != 200 {
//		_ = req.Body.Close()
//		_ = resp.Body.Close()
//		client.CloseIdleConnections()
//		return errors.New(fmt.Sprintf("%d", resp.StatusCode))
//		//return nil, &GetError{message: resp.Status, statusCode: http.StatusUnauthorized}
//	}
//	_ = req.Body.Close()
//	_ = resp.Body.Close()
//	client.CloseIdleConnections()
//	return nil
//}
//
//func ContactServerEncrypt(url string, endpoint string, json []byte, urlSc string) (*http.Response, error) {
//	urlGet := "http://" + strings.TrimSpace(url) + ":7465" + endpoint
//	var jsonStr = json
//	//fmt.Println(urlGet)
//	req, err := http.NewRequest("POST", urlGet, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		return nil, errors.New(err.Error())
//		//return &GetError{message: err.Error(), statusCode: http.StatusInternalServerError}
//	}
//	req.Header.Set("RPCUser-agent", "RocketBot PoS Service/Go_"+VERSION)
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("url", urlSc)
//
//	client := &http.Client{}
//	resp, errD := client.Do(req)
//	if errD != nil {
//		_ = req.Body.Close()
//		client.CloseIdleConnections()
//		return nil, errors.New("error client")
//		//return nil, &GetError{message: errD.Error(), statusCode: http.StatusInternalServerError}
//	}
//
//	if resp.StatusCode != 200 {
//		_ = req.Body.Close()
//		_ = resp.Body.Close()
//		client.CloseIdleConnections()
//		return nil, errors.New(fmt.Sprintf("%d", resp.StatusCode))
//		//return nil, &GetError{message: resp.Status, statusCode: http.StatusUnauthorized}
//	}
//	_ = req.Body.Close()
//	client.CloseIdleConnections()
//	return resp, nil
//}
//
//func ContactClient(url string, endpoint string, token string, json []byte, params ...string) (*http.Response, *GetError) {
//	urlGet := "http://" + strings.TrimSpace(url) + ":7466" + endpoint
//	//fmt.Println(urlGet)
//	var method = http.MethodPost
//	var jsonStr = json
//	if len(params) != 0 {
//		method = http.MethodGet
//	}
//	req, err := http.NewRequest(method, urlGet, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		_ = req.Body.Close()
//		return nil, &GetError{message: err.Error(), statusCode: http.StatusInternalServerError}
//	}
//	req.Header.Set("RPCUser-agent", "RocketBot PoS Service/Go_"+VERSION)
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
//
//	client := &http.Client{}
//	resp, errD := client.Do(req)
//	if errD != nil {
//		_ = req.Body.Close()
//		client.CloseIdleConnections()
//		return nil, &GetError{message: errD.Error(), statusCode: http.StatusInternalServerError}
//	}
//
//	if resp.StatusCode != 200 {
//		fmt.Println("response Status:", resp.StatusCode)
//		fmt.Println("response Headers:", resp.Header)
//		_ = req.Body.Close()
//		_ = resp.Body.Close()
//		client.CloseIdleConnections()
//		//WrapErrorLog(err.Error())
//		return nil, &GetError{message: resp.Status, statusCode: http.StatusUnauthorized}
//	}
//	_ = req.Body.Close()
//	client.CloseIdleConnections()
//	return resp, nil
//}

func Retry(attempts int, sleep time.Duration, f func() error) (err error) {
	errorChan := make(chan error)
	go func() {
		for i := 0; i < attempts; i++ {
			errorChan <- f()
			if err != nil {
				fmt.Println("This was attempt number", i+1)
				fmt.Println(colorRed, fmt.Sprintf("error occured after attempt %d from %d: %s", i+1, attempts, err.Error()))
				//fmt.Println(string(colorRed), errorovia[i])
				fmt.Print(colorReset, "")
				//utils.WrapErrorLogF("error occured after attempt number %d: %s", i+1, err.Error())
				fmt.Println("sleeping for: ", sleep.String())
				time.Sleep(sleep)
				//sleep *= 2
				continue
			}
			break
		}
	}()

	select {
	case err := <-errorChan:
		return err
	}
}

//func RetryServer(attempts int, sleep time.Duration, serverURL string, endpoint string, payload []byte) (err error) {
//	for i := 0; i < attempts; i++ {
//		err = ContactServerRetry(serverURL, endpoint, payload)
//		if err != nil && err.Error() != "409" {
//			fmt.Println("This was attempt number", i+1)
//			fmt.Println(colorRed, fmt.Sprintf("error occured after attempt %d from %d: %s", i+1, attempts, err.Error()))
//			//fmt.Println(string(colorRed), errorovia[i])
//			fmt.Print(colorReset, "")
//			//utils.WrapErrorLogF("error occured after attempt number %d: %s", i+1, err.Error())
//			fmt.Println("sleeping for: ", sleep.String())
//			time.Sleep(sleep)
//			//sleep *= 2
//			continue
//		}
//		break
//	}
//	if err == nil {
//		return nil
//	} else if err.Error() == "409" {
//		return nil
//	} else {
//		return err
//	}
//}

//func RetryServerPayload(attempts int, sleep time.Duration, serverURL string, endpoint string, payload []byte, url string) (response *http.Response, err error) {
//	var resp *http.Response
//	for i := 0; i < attempts; i++ {
//		resp, err = ContactServerEncrypt(serverURL, endpoint, payload, url)
//		if err != nil && err.Error() != "409" {
//			fmt.Println("This was attempt number", i+1)
//			fmt.Println(colorRed, fmt.Sprintf("error occured after attempt %d from %d: %s", i+1, attempts, err.Error()))
//			//fmt.Println(string(colorRed), errorovia[i])
//			fmt.Print(colorReset, "")
//			//utils.WrapErrorLogF("error occured after attempt number %d: %s", i+1, err.Error())
//			fmt.Println("sleeping for: ", sleep.String())
//			time.Sleep(sleep)
//			//sleep *= 2
//			continue
//		}
//		break
//	}
//
//	if err == nil {
//		return resp, nil
//	} else {
//		return nil, err
//	}
//}

//func DecryptRequest(r *http.Request, url string) ([]byte, string, error) {
//	var urlNode string
//	cryptKey, errCrypt := database.ReadRow[string]("SELECT encryptKey FROM masternodes WHERE url = ?", urlNode, url)
//	if errCrypt != nil {
//		WrapErrorLog(errCrypt.Error())
//		return nil, "", errCrypt
//	}
//
//	body, err := io.ReadAll(r.Body)
//	bodyString := string(body)
//
//	message, err := DecryptMessage([]byte(cryptKey), bodyString)
//	if err != nil {
//		WrapErrorLog(err.Error())
//		return nil, "", err
//	}
//
//	return []byte(message), cryptKey, nil
//}
