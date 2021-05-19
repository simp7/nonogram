package nonogram

import (
	"github.com/simp7/nonogram/file"
	"github.com/simp7/nonogram/setting"
	"github.com/simp7/nonogram/unit"
)

type Core interface {
	LoadSetting() (result setting.Config, err error)
	SaveSetting(target setting.Config) error
	InitMap(bitmap [][]bool) unit.Map
	LoadMap(name string) (result unit.Map, err error)
	SaveMap(name string, nonomap unit.Map) error
	MapList(unit int) file.MapList
}
