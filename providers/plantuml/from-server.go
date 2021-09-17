package plantuml

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type plantumlServerRenderer struct {
    baseURL *url.URL
}

func NewPlantUMLFromServer(baseURL string) PlantUMLFullRenderer {
    u, err := url.Parse(baseURL)
    if err != nil {
        panic(err)
    }
    println(u.Path)
    if (u.Path == "/") {
        u, err = u.Parse("api/plantuml")
        if err != nil {
            panic(err)
        }
    }
    println(u.String())
    return &plantumlServerRenderer{baseURL: u}
}

func (p *plantumlServerRenderer) requestAPI(uri string, v ...interface{}) ([]byte, error) {
    extraURL := fmt.Sprintf(uri, v...)
    joined, err := url.Parse(path.Join(p.baseURL.String(), extraURL))
    if err != nil {
        return nil, err
    }
    finalURL := strings.Replace(joined.String(), "///", "//", 1)
    println(finalURL)
    res, err := http.Get(finalURL)
    if err != nil {
        return nil, err
    }
    return ioutil.ReadAll(res.Body)
}

func (p *plantumlServerRenderer) RenderPlantUMLToSVG(code string) (string, error) {
    encoded, err := EncodeStatement(code)
    if err != nil {
        return "", err
    }
    data, err := p.requestAPI("./svg/%s", encoded)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (p *plantumlServerRenderer) RenderPlantUMLToText(code string) (string, error) {
    encoded, err := EncodeStatement(code)
    if err != nil {
        return "", err
    }
    data, err := p.requestAPI("./txt/%s", encoded)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (p *plantumlServerRenderer) RenderPlantUMLToImage(code string) (image.Image, error) {
    encoded, err := EncodeStatement(code)
    if err != nil {
        return nil, err
    }
    data, err := p.requestAPI("./png/%s", encoded)
    if err != nil {
        return nil, err
    }
    buf := bytes.NewBuffer(data)
    img, _, err := image.Decode(buf)
    return img, err
}


// http://www.plantuml.com/plantuml/png/SyfFKj2rKt3CoKnELR1Io4ZDoSa70000
// https://www.planttext.com/api/plantuml/img/SoWkIImgAStDuNBAJrBGjLDmpCbCJbMmKiX8pSd9vt98pKi1IW80
