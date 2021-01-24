package nonBlockingCache

import (
	"io/ioutil"
	"net/http"
)

func  HttpGetBody(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close();

	// ioutil.ReadAll returns a []byte and an error which matches the return type
	// of HttpGetBody
	return ioutil.ReadAll(response.Body)
}


