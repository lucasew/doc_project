package plantuml

import "image"

type PlantUMLSVGRenderer interface {
    RenderPlantUMLToSVG(code string) (string, error)
}

type PlantUMLImageRenderer interface {
    RenderPlantUMLToImage(code string) (image.Image, error)
}

type PlantUMLTextRenderer interface {
    RenderPlantUMLToText(code string) (string, error)
}

type PlantUMLFullRenderer interface {
    PlantUMLImageRenderer
    PlantUMLSVGRenderer
    PlantUMLTextRenderer
}

