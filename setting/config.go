package setting

type Config struct {
	Color
	Text
	Language string //Language is a language of this application.
}

func (c Config) ID() string {
	return c.Language
}
