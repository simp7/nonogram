package formatter

import (
	"encoding/json"
	"github.com/simp7/nonogram/db"
)

type jsonFormatter struct {
	*Buffer
	*json.Encoder
	*json.Decoder
}

// Json returns db.Formatter that includes encoder and decoder of json.
func Json() db.Formatter {
	f := new(jsonFormatter)
	f.Buffer = Raw()
	f.Encoder = json.NewEncoder(f.Buffer)
	f.Decoder = json.NewDecoder(f.Buffer)
	return f
}

func (j *jsonFormatter) Extension() string {
	return "json"
}
