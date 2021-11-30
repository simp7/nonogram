package nonogram

import (
	"github.com/simp7/nonogram/setting"
	"github.com/simp7/nonogram/unit"
)

const LatestLanguageVersion = "1.0"
const SettingFile = "setting"

type Core struct {
	fs        System
	prototype unit.Map
}

func New(fs System, prototype unit.Map) *Core {
	return &Core{fs: fs, prototype: prototype}
}

func (c *Core) LoadSetting() (result setting.Config, err error) {

	loader := c.fs.Setting()
	if loader.Load(SettingFile, &result) != nil {
		return
	}

	text, err := c.initLanguage(result.Language)
	result.SetText(text)
	return

}

func (c *Core) initLanguage(language string) (result setting.Text, err error) {

	immutable := c.fs.Language()
	if immutable.Load(language, &result) != nil {
		return
	}

	if result.FileVersion != LatestLanguageVersion {
		if c.updateLanguage() != nil {
			return
		}
		err = immutable.Load(language, &result)
	}

	return

}

func (c *Core) SaveSetting(target setting.Config) error {

	willSaved := target
	willSaved.SetText(setting.Text{})

	saver := c.fs.Setting()
	return saver.Save("setting", willSaved)

}

func (c *Core) updateLanguage() error {
	updater := c.fs.Language()
	return updater.Update()
}

func (c *Core) InitMap(bitmap [][]bool) unit.Map {
	return c.prototype.Init(bitmap)
}

func (c *Core) LoadMap(name string) (result unit.Map, err error) {
	result = c.prototype
	loader := c.fs.Map()
	err = loader.Load(name, result)
	return
}

func (c *Core) SaveMap(name string, nonomap unit.Map) error {
	saver := c.fs.Map()
	return saver.Save(name, nonomap)
}

func (c *Core) Maps() ([]string, error) {
	return c.fs.Maps()
}
