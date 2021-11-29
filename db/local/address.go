package local

import (
	"errors"
	"os"
	"path"
)

var (
	homePathErr = errors.New("HOME does not exist")
)

type address struct {
	sub []string
}

func addressTo(sub ...string) address {
	return address{sub}
}

func rootAddress() address {
	return addressTo()
}

func mapAddress() address {
	return addressTo("maps")
}

func languageAddress() address {
	return addressTo("language")
}

func (a address) Real() (string, error) {
	base, err := rootDir()
	return path.Join(append(base, a.sub...)...), err
}

func (a address) Skel() string {
	return path.Join(append([]string{"skel"}, a.sub...)...)
}

func (a address) Goto(dirName string) address {
	return address{append(a.sub, dirName)}
}

func homeEnv() (string, error) {
	if root, ok := os.LookupEnv("HOME"); ok {
		return root, nil
	}
	return "", homePathErr
}

func rootDir() ([]string, error) {
	home, err := homeEnv()
	if err != nil {
		return nil, err
	}
	return []string{home, "nonogram"}, nil
}
