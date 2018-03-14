package main

import "strings"

const (
	flagPath leetcodeFlag = iota
	flagForce
	flagInteract
)

type leetcodeFlag int

var leetcodeFlags = [...]string{
	flagPath:     "path, p",
	flagForce:    "force, f",
	flagInteract: "interact, i",
}

func (l leetcodeFlag) String() string {
	return leetcodeFlags[l]
}

func (l leetcodeFlag) Flag() string {
	return flag(l.String())
}

func flag(f string) string {
	return strings.Split(f, ",")[0]
}
