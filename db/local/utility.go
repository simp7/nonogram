package local

import (
	"io"
	"io/fs"
	"os"
	"path"
)

func getRealPath(addr address, fileName string) (realPath string, err error) {

	if realPath, err = addr.Real(); err != nil {
		return
	}

	realPath = path.Join(realPath, fileName)
	return

}

func getSkelPath(addr address, fileName string) string {
	return path.Join(addr.Skel(), fileName)
}

// Reason of using os.IsNotExist: https://gist.github.com/mattes/d13e273314c3b3ade33f
func mkDir(addr address) (err error) {

	target := ""
	if target, err = addr.Real(); err == nil {
		err = os.Mkdir(target, 0755)
	}

	if !os.IsNotExist(err) {
		err = nil
	}
	return

}

func readRealFile(addr address, fileName string) ([]byte, error) {

	realPath, err := getRealPath(addr, fileName)
	if err != nil {
		return nil, err
	}

	realFile, err := os.Open(realPath)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(realFile)

}

func readSkelFile(addr address, fileName string) (data []byte, err error) {
	if skelFile, err := f.Open(getSkelPath(addr, fileName)); err == nil {
		data, err = io.ReadAll(skelFile)
	}
	return
}

func readSkelDir(addr address) (result []fs.DirEntry, err error) {
	return f.ReadDir(addr.Skel())
}

func readRealDir(addr address) ([]os.DirEntry, error) {

	realPath, err := addr.Real()
	if err != nil {
		return nil, err
	}

	return os.ReadDir(realPath)

}

func writeFile(addr address, fileName string, data []byte) (err error) {
	if realPath, err := getRealPath(addr, fileName); err == nil {
		err = os.WriteFile(realPath, data, 0644)
	}
	return
}

func copyFile(addr address, fileName string) (err error) {
	if data, err := readSkelFile(addr, fileName); err == nil {
		err = writeFile(addr, fileName, data)
	}
	return
}

//func isInitial() (bool, error) {
//
//	path, err := rootAddress()
//	if err != nil {
//		return false, err
//	}
//
//	_, err1 := readDir(path)
//	_, err2 := readFile(path)
//
//	return err1 != nil && err2 != nil, err
//
//}
