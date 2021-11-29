package db

//Updater is an interface that browse files from somewhere to another.
type Updater interface {
	Update() error //Update updates files.
}
