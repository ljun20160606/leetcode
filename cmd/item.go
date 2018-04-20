package main

import "strconv"

type item struct {
	ID         string
	Title      string
	Topics     string
	Difficulty string
}

type itemSlice []item

func (p itemSlice) Len() int { return len(p) }
func (p itemSlice) Less(i, j int) bool {
	x, _ := strconv.Atoi(p[i].ID)
	y, _ := strconv.Atoi(p[j].ID)
	return x < y
}
func (p itemSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
