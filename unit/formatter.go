package unit

//Formatter is just a duplication of file.Formatter.
type Formatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
	Raw(from []byte) error
	Content() []byte
	Extension() string
}
