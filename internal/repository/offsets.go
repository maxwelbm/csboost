package repository

import (
	"encoding/json"
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"io"
	"net/http"
)

const OffsetsURL string = "https://raw.githubusercontent.com/frk1/hazedumper/master/csgo.json"

func GetOffset() (*model.OffSet, error) {
	resp, err := http.Get(OffsetsURL)
	if err != nil {
		return nil, fmt.Errorf("fail to get offset. Error - %v. Using default offsets. Cheat may not work", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	strBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to parse returned offset json. Error - %v. Using default offsets. Cheat may not work", err)
	}

	var offSet model.OffSet
	err = json.Unmarshal(strBytes, &offSet)
	if err != nil {
		return nil, fmt.Errorf("unable to parse returned offset json. Error - %v. Using default offsets. Cheat may not work", err)
	}
	return &offSet, nil
}
