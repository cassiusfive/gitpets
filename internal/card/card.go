package card

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/cassiusfive/gitpets/internal/pet"
)

type CardStyles struct {
	text       string
	background string
}

func Generate(w http.ResponseWriter, pet pet.Pet, styles CardStyles) error {
	f, err := os.Open("assets/fox/1.svg")
	if err != nil {
		return err
	}
	s, err := parseSVG(f)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	canvas := svg.New(w)
	canvas.Start(400, 600)
	placeSvg(canvas, s)

	return nil
}

// SVG contains the parsed attributes and xml from the given file.
type SVG struct {
	// Width and Height are attributes of the <svg> tag
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
	// Doc is all all of the contents within the <svg> tags, specified by the
	// `innerxml` struct tag
	Doc []byte `xml:",innerxml"`
}

func parseSVG(src io.Reader) (SVG, error) {
	var s SVG
	data, err := io.ReadAll(src)
	if err != nil {
		return SVG{}, err
	}
	err = xml.Unmarshal(data, &s)
	if err != nil {
		return SVG{}, err
	}
	return s, nil
}

func placeSvg(canvas *svg.SVG, s SVG) {
	// create clip path of svg size
	// canvas.Group(`clip-path="url(#embed)"`)
	// canvas.ClipPath(`id="embed"`)
	// canvas.Rect(0, 0, s.Width, s.Height)
	// canvas.ClipEnd()
	// // append embedded svg
	canvas.Group(`id="frame-1"`, `transform="scale(10)"`)
	canvas.Writer.Write(s.Doc)
	canvas.Writer.Write([]byte(`<animate attributeName="opacity" dur="1s" values="1; 1; 0; 0;" keyTimes="0; 0.5; 0.51; 1" repeatCount="indefinite" />`))
	canvas.Gend()
	canvas.Group(`id="frame-2"`, `transform="scale(10)"`, `opacity="0%"`)
	canvas.Writer.Write(s.Doc)
	canvas.Gend()
	canvas.End()
}
