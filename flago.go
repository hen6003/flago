package flago

import (
	"errors"
	"os"
)

func Exists(list []string) ([]bool, error) {
	args := os.Args[1:]

	info := make([]bool, len(list))

	for _, v := range args {
		for i, s := range list {
			switch len(s) {
			case 0:
				return info, errors.New("0 length string")

			case 1:
				s = "-" + s

			default:
				s = "--" + s
			}

			if v == s {
				info[i] = true
			}
		}
	}

	return info, nil
}
