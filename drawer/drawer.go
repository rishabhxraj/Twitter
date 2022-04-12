package drawer

import (
	"fmt"
	"image/color"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"github.com/rishabhxraj/twitter-bot/quotable"
)

type Image struct {
	Canvas *gg.Context
	Quote  *quotable.Quote
	Footer string
}

func (img *Image) Create() error {
	if err := img.createBackground(); err != nil {
		return errors.Wrap(err, "failed to draw image")
	}
	img.drawOverlay()
	if err := img.drawLogo(); err != nil {
		return errors.Wrap(err, "failed to draw image")
	}
	if err := img.writeFooter(); err != nil {
		return errors.Wrap(err, "failed to draw image")
	}
	if err := img.writeQuote(); err != nil {
		return errors.Wrap(err, "failed to draw image")
	}
	outfilePath := filepath.Join("static", "img", "out.png")
	if err := img.Canvas.SavePNG(outfilePath); err != nil {
		return errors.Wrap(err, "save png")
	}
	fmt.Println("Sucessfully created image")
	return nil
}

func (img *Image) createBackground() error {
	bgfilePath := filepath.Join("static", "img", "bg.png")
	backgroundImage, err := gg.LoadImage(bgfilePath)
	if err != nil {
		return errors.Wrap(err, "load background image")
	}
	img.Canvas.DrawImage(backgroundImage, 0, 0)
	backgroundImage = imaging.Fill(backgroundImage, img.Canvas.Width(), img.Canvas.Height(), imaging.Center, imaging.Lanczos)
	return nil
}

func (img *Image) drawOverlay() {
	margin := 20.0
	x := margin
	y := margin
	w := float64(img.Canvas.Width()) - (2.0 * margin)
	h := float64(img.Canvas.Height()) - (2.0 * margin)
	img.Canvas.SetColor(color.RGBA{0, 0, 0, 204})
	img.Canvas.DrawRectangle(x, y, w, h)
	img.Canvas.Fill()
}

func (img *Image) drawLogo() error {
	fontPath := filepath.Join("static", "font", "Roboto-Bold.ttf")
	if err := img.Canvas.LoadFontFace(fontPath, 25); err != nil {
		return errors.Wrap(err, "load font")
	}
	img.Canvas.SetColor(color.White)
	marginX := 50.0
	marginY := 25.0
	textWidth, textHeight := img.Canvas.MeasureString(img.Quote.Author)
	x := float64(img.Canvas.Width()) - textWidth - marginX
	y := float64(img.Canvas.Height()) - textHeight - marginY
	img.Canvas.DrawString(img.Quote.Author, x, y)
	return nil
}

func (img *Image) writeFooter() error {
	textColor := color.White
	fontPath := filepath.Join("static", "font", "Roboto-Light.ttf")
	if err := img.Canvas.LoadFontFace(fontPath, 20); err != nil {
		return errors.Wrap(err, "failed to load font")
	}
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	img.Canvas.SetColor(mutedColor)
	marginY := 30.0
	_, textHeight := img.Canvas.MeasureString(img.Footer)
	x := 70.0
	y := float64(img.Canvas.Height()) - textHeight - marginY
	img.Canvas.DrawString(img.Footer, x, y)
	return nil
}

func (img *Image) writeQuote() error {
	textShadowColor := color.Black
	textColor := color.White
	fontPath := filepath.Join("static", "font", "GreatVibes-Regular.ttf")
	if err := img.Canvas.LoadFontFace(fontPath, 90); err != nil {
		return errors.Wrap(err, "load Playfair_Display")
	}
	textRightMargin := 60.0
	textTopMargin := 90.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(img.Canvas.Width()) - textRightMargin - textRightMargin
	img.Canvas.SetColor(textShadowColor)
	img.Canvas.DrawStringWrapped(img.Quote.Content, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	img.Canvas.SetColor(textColor)
	img.Canvas.DrawStringWrapped(img.Quote.Content, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return nil
}
