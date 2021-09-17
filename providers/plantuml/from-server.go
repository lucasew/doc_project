package plantuml

import (
	"image"
	"net/url"
)

type plantumlServerRenderer struct {
    baseURL *url.URL
}

func NewPlantUMLFromServer(baseURL string) PlantUMLFullRenderer {
    u, err := url.Parse(baseURL)
    if err != nil {
        panic(err)
    }
    if (u.Path == "/") {
        u, err = u.Parse("api/plantuml")
        if err != nil {
            panic(err)
        }
    }
    return &plantumlServerRenderer{baseURL: u}
}

func (p *plantumlServerRenderer) RenderPlantUMLToSVG(code string) (string, error) {
    panic("Missing implementation")
    // return "", nil
}

func (p *plantumlServerRenderer) RenderPlantUMLToText(code string) (string, error) {
    panic("Missing implementation")
    // return "", nil
}

func (p *plantumlServerRenderer) RenderPlantUMLToImage(code string) (image.Image, error) {
    panic("Missing implementation")
    // return nil, nil
}

// http://www.plantuml.com/plantuml/png/SyfFKj2rKt3CoKnELR1Io4ZDoSa70000
// https://www.planttext.com/api/plantuml/img/SoWkIImgAStDuNBAJrBGjLDmpCbCJbMmKiX8pSd9vt98pKi1IW80
