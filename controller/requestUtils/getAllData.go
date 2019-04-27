package requestUtils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"plane-ticket-search-web/dto"
)

func GetAllTicketData() ([]dto.TicketPriceDTO, error) {
	res, err := http.Get(getAllTicketUrl())

	if err != nil {
		return nil, err
	}

	getData, _ := ioutil.ReadAll(res.Body)
	var parsedJsonData []dto.TicketPriceDTO
	err = json.Unmarshal(getData, &parsedJsonData)
	if err != nil {
		return nil, err
	}

	return parsedJsonData, nil
}

