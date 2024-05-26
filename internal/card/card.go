package card

import (
	"encoding/xml"
	"fmt"
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
	w.Header().Set("Content-Type", "image/svg+xml")

	canvas := svg.New(w)
	canvas.Start(400, 600)

	frames, err := os.ReadDir("assets/fox/")
	if err != nil {
		return err
	}

	fps := 2
	frameLength := 1.0 / float64(fps)
	duration := frameLength * float64(len(frames))

	for i := range len(frames) {
		f, err := os.Open(fmt.Sprintf("assets/fox/%d.svg", i))
		if err != nil {
			return err
		}
		s, err := parseSVG(f)
		if err != nil {
			return err
		}

		frameStart := float64(i) / float64(len(frames))
		frameEnd := float64(i+1) / float64(len(frames))

		canvas.Group(fmt.Sprintf(`id="frame-%d"`, i), `transform="scale(10)"`)
		canvas.Writer.Write(s.Doc)
		canvas.Writer.Write([]byte(fmt.Sprintf(`<animate attributeName="opacity" dur="%.2fs" values="0; 0; 1; 1; 0; 0;"
			keyTimes="0; %.2f; %.2f; %.2f; %.2f; 1" repeatCount="indefinite" />`, duration, frameStart, frameStart+0.0001, frameEnd-0.0001, frameEnd)))
		canvas.Gend()
	}

	canvas.End()
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
