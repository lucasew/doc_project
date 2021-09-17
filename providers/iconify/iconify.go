package iconify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
    ErrInvalidIconId = errors.New("Invalid icon id, should be pack:icon")
    ErrIconNotFound = errors.New("Icon not found")
)

type iconifyAPIResult struct {
    Prefix string `json:"prefix"`
    Icons map[string]iconifyIcon `json:"icons"`
    Width int `json:"width"`
    Height int `json:"height"`
    NotFound *[]string `json:"not_found"`
}

type iconifyIcon struct {
    Body string `json:"body"`
}

func GetIconSVG(id string) (string, error) {
    parts := strings.Split(id, ":")
    if (len(parts) != 2) {
        return "", ErrInvalidIconId
    }
    pack := parts[0]
    icon := parts[1]
    fetchedIcon, err := fetchIcon(pack, icon) // TODO: rasterize icon
    if err != nil {
        return "", err
    }
    if fetchedIcon.NotFound != nil {
        return "", ErrIconNotFound
    }
    return fetchedIcon.Icons[icon].Body, nil // TODO: raster image

}

func fetchIcon(pack string, icon string) (*iconifyAPIResult, error) {
    res, err := http.Get(fmt.Sprintf("https://api.iconify.design/%s.json?icons=%s", pack, icon))
    if err != nil {
        return nil, err
    }
    var ret iconifyAPIResult
    err = json.NewDecoder(res.Body).Decode(&ret)
    if err != nil {
        return nil, err
    }
    return &ret, nil
}
