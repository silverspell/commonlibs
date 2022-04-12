package commonlibs

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var httpClientPool = sync.Pool{
	New: func() interface{} {
		return http.Client{}
	},
}

func HttpPostRequest(url string, data []byte, headers map[string]string) (io.ReadCloser, error) {
	client := httpClientPool.Get().(*http.Client)
	defer httpClientPool.Put(client)
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error sending %+v\n", err)
		return nil, err
	}

	return resp.Body, nil
}

func HttpGetRequest(url string, headers map[string]string) (io.ReadCloser, error) {
	client := httpClientPool.Get().(*http.Client)
	defer httpClientPool.Put(client)

	request, _ := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error sending %+v\n", err)
		return nil, err
	}
	return resp.Body, nil
}
