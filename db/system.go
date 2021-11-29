package db

type system struct {
	maps     Storage
	setting  Storage
	language Immutable
}

func New(mapStorage Storage, settingStorage Storage, langImmutable Immutable) *system {
	sys := new(system)
	sys.maps = mapStorage
	sys.setting = settingStorage
	sys.language = langImmutable
	return sys
}

func (s *system) Map() Storage {
	return s.maps
}

func (s *system) Setting() Storage {
	return s.setting
}

func (s *system) Language() Immutable {
	return s.language
}

func (s *system) Maps() ([]string, error) {
	return s.maps.List()
}
