package flago

import (
	"errors"
	"os"
)

// CheckFlags checks the flags against what flags you want
func CheckFlags(list []string) ([]bool, error) {
	args := os.Args[1:]

	info := make([]bool, len(list))

	for _, v := range args {
		for i, s := range list {
			switch len(s) {
			case 0:
				err := errors.New("0 width string")
				return info, err

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

// NonFlags returns all args that are not flags
func NonFlags() []string {
	args := os.Args[1:]
	var nonFlags []string

	for _, v := range args {
		if string(v[0]) != "-" {
			nonFlags = append(nonFlags, v)
		} else if v == "-" {
			nonFlags = append(nonFlags, v)
		}
	}

	return nonFlags
}
