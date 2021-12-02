package db

//Storage is an interface that implements Saver and Loader.
type Storage interface {
	Saver
	Loader
	Updater
}
