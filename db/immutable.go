package db

type Immutable interface {
	Updater
	Loader
}
