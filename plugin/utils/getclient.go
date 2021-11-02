package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
)

//APIget Call the RL API with a GET command
func APIget(conf *RLConf, path string, params map[string]string) (string, int, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://"+conf.Hostname+":"+strconv.Itoa(conf.Port)+path, nil)
	if err != nil {
		return "", 0, err
	}

	// Add any GET params
	if len(params) > 0 {
		q := request.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}
	request.SetBasicAuth(conf.User, conf.Pass)
	response, err := client.Do(request)
	if err != nil {
		return "", 0, err
	}

	bodyText, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", 0, err
	}
	s := string(bodyText)

	return s, response.StatusCode, nil

}
