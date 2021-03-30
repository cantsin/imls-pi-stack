package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gsa.gov/18f/session-counter/config"
)

func postJSON(cfg *config.Config, tok *config.AuthConfig, uri string, data map[string]string) (int, error) {
	log.Println("storing JSON to", uri)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	var reqBody []byte
	var err error
	// FIXME
	// Directus takes posted data directly.
	// ReVal is currently looking for it to be "wrapped" in a struct.
	// We should modify ReVal so that it takes the exact same POSTed data
	// as Directus, so that we cannot tell the difference from the client-side.
	// switch svr.Name {
	// case "directus":
	// 	reqBody, err = json.Marshal(data)
	// case "reval":
	// 	source := map[string][]map[string]string{"source": {data}}
	// 	reqBody, err = json.Marshal(source)
	// }

	reqBody, err = json.Marshal(data)

	if err != nil {
		return -1, errors.New("api: unable to marshal post of data to JSON")
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(reqBody))
	if err != nil {
		return -1, errors.New("api: unable to construct request for data POST")
	}

	req.Header.Set("Content-type", "application/json")
	if tok != nil {
		log.Printf("Using access token: %v\n", tok.Umbrella.Token)
		req.Header.Set("Authorization", fmt.Sprintf("X-Api-Key %s", tok.Umbrella.Token))
	} else {
		log.Printf("api: failed to set headers for authorization.")
	}

	// The first thing we do is post an event. This will return a "magic index"
	// or a foreign key, that we will use in our post of the data. This associates
	// every piece of data entered with a session, and indexes the post in that session.
	// That way, we can say "this set of data was entry 293 of session ABC."
	magic_index := -1

	log.Printf("req:\n%v\n", req)
	resp, err := client.Do(req)
	log.Printf("resp: %v\n", resp)
	if err != nil {
		log.Printf("err resp: %v\n", resp)
		return -1, fmt.Errorf("api: failure in client attempt to POST to %v", uri)
	} else {
		// If we get things back, the errors will be encoded within the JSON.
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			log.Printf("api: bad status on POST to: %v\n", uri)
			log.Printf("api: bad status on POST response: [ %v ]\n", resp.Status)
		} else {
			var dat map[string]interface{}
			body, _ := ioutil.ReadAll(resp.Body)
			err := json.Unmarshal(body, &dat)
			if err != nil {
				return -1, fmt.Errorf("api: could not unmarshal response body")
			}
			// 2021/03/26 14:00:18 resp.Body {"data":{"magic_index":12,"device_uuid":"1000000089bbf88b","lib_user":"10x@gsa.gov","session_id":"effc67d0068b4e7f","localtime":"2021-03-26T18:00:17Z","servertime":"2021-03-26T18:00:17Z","tag":"nil","info":"{}"}}
			log.Println("resp.Body", string(body))
			if _, ok := dat["data"]; ok {
				inner := dat["data"].(map[string]interface{})
				if _, ok := inner["magic_index"]; ok {
					magic_index = int(inner["magic_index"].(float64))
				}
			}
		}
	}
	// Close the body at function exit.
	defer resp.Body.Close()

	return magic_index, nil
}