package formatter

import (
	"github.com/simp7/nonogram/db"
	"gopkg.in/yaml.v3"
)

type yamlFormatter struct {
	*Buffer
	*yaml.Encoder
	*yaml.Decoder
}

// Yaml returns db.Formatter that includes encoder and decoder of yaml.
func Yaml() db.Formatter {
	y := new(yamlFormatter)
	y.Buffer = Raw()
	y.Encoder = yaml.NewEncoder(y.Buffer)
	y.Decoder = yaml.NewDecoder(y.Buffer)
	return y
}

func (y *yamlFormatter) Extension() string {
	return "yaml"
}
