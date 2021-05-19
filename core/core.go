package core

import (
	"github.com/simp7/nonogram/file"
	"github.com/simp7/nonogram/setting"
	"github.com/simp7/nonogram/unit"
)

const LatestLanguageVersion = "1.0"

type core struct {
	fs                file.System
	prototype         unit.Map
	settingFormatter  file.Formatter
	languageFormatter file.Formatter
}

func New(fs file.System, prototype unit.Map, settingFormat file.Formatter, languageFormat file.Formatter) *core {
	return &core{fs, prototype, settingFormat, languageFormat}
}

func (c *core) LoadSetting() (result setting.Config, err error) {

	loader, err := c.fs.Setting(c.settingFormatter)
	if err != nil {
		return
	}

	err = loader.Load(&result)
	if err != nil {
		return
	}

	result.Text, err = c.initLanguage(result.Language)
	return

}

func (c *core) initLanguage(language string) (result setting.Text, err error) {

	loader, err := c.fs.LanguageOf(language, c.languageFormatter)
	if err != nil {
		return
	}

	err = loader.Load(&result)

	if result.FileVersion != LatestLanguageVersion {

		if c.updateLanguage() != nil {
			return
		}

		loader, err = c.fs.LanguageOf(language, c.languageFormatter)
		if err != nil {
			return
		}

		err = loader.Load(&result)

	}

	return

}

func (c *core) SaveSetting(target setting.Config) (err error) {

	willSaved := target
	willSaved.Text = setting.Text{}

	saver, err := c.fs.Setting(c.settingFormatter)
	if err != nil {
		return
	}

	err = saver.Save(willSaved)
	return

}

func (c *core) updateLanguage() error {
	updater, err := c.fs.Language()
	if err == nil {
		updater.Update()
	}
	return err
}

func (c *core) InitMap(bitmap [][]bool) unit.Map {
	return c.prototype.Init(bitmap)
}

func (c *core) LoadMap(name string) (result unit.Map, err error) {

	loader, err := c.fs.Map(name, c.prototype.GetFormatter())
	if err != nil {
		return
	}

	err = loader.Load(&result)
	return

}

func (c *core) SaveMap(name string, nonomap unit.Map) error {

	saver, err := c.fs.Map(name, c.prototype.GetFormatter())
	if err == nil {
		err = saver.Save(&nonomap)
	}

	return err

}

func (c *core) MapList(unit int) file.MapList {
	return c.fs.Maps(unit)
}
