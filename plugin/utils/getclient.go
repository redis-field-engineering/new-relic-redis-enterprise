package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
)

//APIget Call the RL API with a GET command
func APIget(conf *RLConf, path string) (string, int, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://"+conf.Hostname+":"+strconv.Itoa(conf.Port)+path, nil)
	request.SetBasicAuth(conf.User, conf.Pass)
	response, err := client.Do(request)
	if err != nil {
		return "", 0, err
	}

	bodyText, err := ioutil.ReadAll(response.Body)
	s := string(bodyText)

	return s, response.StatusCode, nil

}
