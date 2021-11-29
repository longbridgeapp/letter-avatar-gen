package main

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/disintegration/letteravatar"
)

var names = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "#",
}

func hexColor(hexColor string) color.Color {
	b, err := hex.DecodeString(hexColor)
	if err != nil {
		log.Fatal(err)
	}

	color := color.RGBA{b[0], b[1], b[2], 255}
	return color
}

func colorPalettes(hexColors []string) []color.Color {
	colors := make([]color.Color, len(hexColors))
	for i, c := range hexColors {
		colors[i] = hexColor(c)
	}
	return colors
}

func main() {
	for _, name := range names {
		firstLetter, _ := utf8.DecodeRuneInString(name)

		img, err := letteravatar.Draw(500, firstLetter, &letteravatar.Options{
			Palette: colorPalettes([]string{
				"00B8D4",
				"00BFA5",
				"00C853",
				"64DD17",
				"D50000",
				"C51162",
				"AA00FF",
				"6200EA",
				"304FFE",
				"2962FF",
				"0091EA",
				"AEEA00",
				"D32F2F",
				"4A148C",
				"1A237E",
				"311B92",
				"0D47A1",
				"FFD600",
				"FFAB00",
				"FF6D00",
				"DD2C00",
				"BF360C",
			}),
			PaletteKey: fmt.Sprintf("%c", firstLetter),
		})
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create("dist/" + strings.ToLower(name) + ".png")
		if err != nil {
			log.Fatal(err)
		}

		err = png.Encode(file, img)
		if err != nil {
			log.Fatal(err)
		}
	}
}
