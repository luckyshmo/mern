package service

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/luckyshmo/api-example/models/keep"
	"github.com/pkg/errors"
)

type KeepService struct {
	//? some endpoint
}

func NewKeepService() *KeepService {
	return &KeepService{}
}

func (ks *KeepService) GetAll() (keep.Note, error) {

	// curl -X GET --data {"email":"mishka2017@gmail.com","token":"iopnilguhgigbbht","name":"English words"} http://localhost:5001/get_words

	params := url.Values{}
	// params.Add(`{"email":"mishka2017@gmail.com","token":"iopnilguhgigbbht","name":"English words"}`, "")
	params.Add("email", "mishka2017@gmail.com")
	params.Add("token", "iopnilguhgigbbht")
	params.Add("name", "English words")
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "http://localhost:5001/get_words", body)
	if err != nil {
		return keep.Note{}, errors.Wrap(err, "error creating Request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return keep.Note{}, errors.Wrap(err, "error while making request")
	}
	defer resp.Body.Close()

	var I map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&I)
	if err != nil {
		panic(err)
	}

	arrayOfI := I["value"].([]interface{})

	strArr := make([]string, len(arrayOfI))

	for _, v := range arrayOfI {
		strArr = append(strArr, strings.Trim(v.(string), "☐ "))
	}

	return keep.Note{
		Data: strArr,
	}, nil
}
