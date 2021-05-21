package unit

//Formatter is just an duplication of file.GetFormatter.
type Formatter interface {
	Encode(interface{}) error
	Decode(interface{}) error
	GetRaw(from []byte) error
	Content() []byte
}