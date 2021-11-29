package local

import (
	"embed"
)

//go:embed skel
var f embed.FS

type updater struct {
	target address
}

func (u *updater) Update() error {
	return u.updateDir(u.target)
}

func (u *updater) updateDir(addr address) (err error) {

	if err = mkDir(addr); err != nil {
		return
	}

	files, err := readSkelDir(addr)
	if err != nil {
		return
	}

	for _, v := range files {

		fileName := v.Name()

		if v.IsDir() {
			err = u.updateDir(addr.Goto(fileName))
		} else {
			err = copyFile(addr, v.Name())
		}

		if err != nil {
			return
		}

	}

	return

}
