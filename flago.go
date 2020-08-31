package flago

import (
	"errors"
	"os"
	"strings"
)

func AllFlags() []string {
	var flags []string

	for _, v := range os.Args[1:] {
		if v == "-" {
		} else if string(v[0]) == "-" {
			v = strings.Replace(v, "-", "", -1)
			flags = append(flags, v)
		}
	}

	return flags
}

// CheckFlags checks the flags against what flags you want
func CheckFlags(list []string) ([]bool, error) {
	args := AllFlags()

	info := make([]bool, len(list))

	for _, v := range args {
		for i, s := range list {
			if len(s) == 0 {
				err := errors.New("0 length string")
				return info, err
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
