package main

import "fmt"

const SizePole = 40

type Square struct {
	Name  string
	Cell  [][] int
	Width int
	Pole  *[SizePole][SizePole] int
}

type Rectangle struct {
	Name   string
	Cell   [][] int
	Width  int
	Height int
	Pole   *[SizePole][SizePole] int
}

type Snake struct {
	Name   string
	Cell   [][] int
	Length int
	Pole   *[SizePole][SizePole] int
}

type Game interface {
}

func main() {

	var pole [SizePole][SizePole] int //Поле

	for _, v := range pole {
		fmt.Println(v)
	}

}
