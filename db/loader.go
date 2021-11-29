package db

//Loader is an interface that loads data to the program from somewhere.
type Loader interface {
	Load(fileName string, target interface{}) error //Load loads data from Loader to argument. argument should be address of wanted object.
	List() ([]string, error)
}
