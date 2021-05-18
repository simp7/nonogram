package local

import (
	"github.com/simp7/nonogram/file"
	"sync"
)

var instance file.System
var once sync.Once

//System returns struct that implements file.System by local storage
//Returned struct by System is standard option for using file package.
func System() (file.System, error) {

	var err error

	once.Do(func() {
		if isInitial() {
			var u file.Updater
			u, err = allUpdater()
			u.Update()
		}
		instance = new(system)
	})

	return instance, err

}

type system struct {
}

func (s *system) Map(name string, formatter file.Formatter) (file.Storage, error) {
	return mapStorage(name, formatter)
}

func (s *system) Setting(formatter file.Formatter) (file.Storage, error) {
	return settingStorage(formatter)
}

func (s *system) LanguageOf(language string, formatter file.Formatter) (file.Loader, error) {
	return languageLoader(language, formatter)
}

func (s *system) Language() (file.Updater, error) {
	return languageUpdater()
}

func (s *system) Maps(unit int) file.MapList {
	return newMapList(unit)
}
