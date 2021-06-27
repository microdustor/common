package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (error, string) {
	resp, err := http.Get(url)
	if err != nil {
		return err, ""
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Print(err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	return nil, string(body)
}
