package card

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/cassiusfive/gitpets/internal/pet"
)

type CardStyles struct {
	Text       string
	Background string
}

func Generate(w http.ResponseWriter, p pet.Pet, styles CardStyles) error {
	w.Header().Set("Content-Type", "image/svg+xml")
	if styles.Text == "" {
		styles.Text = "black"
	}
	canvas := svg.New(w)
	canvas.Start(32*5, 32*5)

	width := 18
	xpProgress := float32(p.Xp) / float32(pet.ExperienceToLevel(p.Level))
	xpStr := fmt.Sprintf("xp: %s%s %.0f%%", strings.Repeat("▰", int(xpProgress*10)), strings.Repeat("▱", 10-int(xpProgress*10)), xpProgress*100)
	moodStr := fmt.Sprintf("mood: %-*s", width-6, p.Mood)

	canvas.Writer.Write([]byte(fmt.Sprintf(`<text x="50%%" y="25" dominant-baseline="middle" text-anchor="middle" fill="%s" style="font-family:'Courier New',monospace;font-size:0.8rem;font-weight:bold;white-space:pre">%s Lv%d</text>`, styles.Text, p.Name, p.Level)))
	canvas.Writer.Write([]byte(fmt.Sprintf(`<text x="50%%" y="130" dominant-baseline="middle" text-anchor="middle" fill="%s" style="font-family:'Courier New',monospace;font-size:0.8rem;white-space:pre">%s</text>`, styles.Text, xpStr)))
	canvas.Writer.Write([]byte(fmt.Sprintf(`<text x="50%%" y="150" dominant-baseline="middle" text-anchor="middle" fill="%s" style="font-family:'Courier New',monospace;font-size:0.8rem;white-space:pre">%s</text>`, styles.Text, moodStr)))

	speciesDir := path.Join("assets", p.Species)
	fmt.Printf("speciesDir: %vn", speciesDir)
	frames, err := os.ReadDir(speciesDir)
	if err != nil {
		return err
	}

	fps := 3
	frameLength := 1.0 / float64(fps)
	duration := frameLength * float64(len(frames))

	canvas.Translate(0, -40)
	for i := range len(frames) {
		f, err := os.Open(path.Join(speciesDir, strconv.Itoa(i)+".svg"))
		if err != nil {
			fmt.Println(err)
			return err
		}
		s, err := parseSVG(f)
		if err != nil {
			return err
		}

		frameStart := float64(i) / float64(len(frames))
		frameEnd := float64(i+1) / float64(len(frames))

		canvas.Group(fmt.Sprintf(`id="frame-%d"`, i), `transform="scale(5)"`)
		canvas.Writer.Write(s.Doc)
		canvas.Writer.Write([]byte(fmt.Sprintf(`<animate attributeName="opacity" dur="%.2fs" values="0; 0; 1; 1; 0; 0;"
			keyTimes="0; %.2f; %.2f; %.2f; %.2f; 1" repeatCount="indefinite" />`, duration, frameStart, frameStart+0.0001, frameEnd-0.0001, frameEnd)))
		canvas.Gend()
	}
	canvas.Gend()

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
