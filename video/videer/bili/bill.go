package bili

import (
	"encoding/json"
	"github.com/pkg/errors"
	"test/models"
	"test/utils"
)

func GetMultiPageData(html string) (*models.MultiPage, error) {
	var data models.MultiPage
	multiPageDataString := utils.MatchOneOf(
		html, `window.__INITIAL_STATE__=(.+?);\(function`,
	)

	if multiPageDataString == nil {
		return &data, errors.New("this page has no playlist")
	}
	err := json.Unmarshal([]byte(multiPageDataString[1]), &data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
