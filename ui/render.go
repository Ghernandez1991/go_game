package ui

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("assets/images/One_Day_at_Horrorland.jpg")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Nothing yet.
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	scaleX := 800.0 / float64(width)
	scaleY := 600.0 / float64(height)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(img, opts)
	//screen.DrawImage(img, nil)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
