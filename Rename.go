package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	dir := "./imgs_textos/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		rename(dir+f.Name(), dir+normalize(f.Name()))
	}
}

func normalize(nomeArquivo string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	nomeNovo, _, err := transform.String(t, nomeArquivo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(nomeNovo))
	return string(nomeNovo)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.IsSpace(r)
}

func rename(nomeVelho string, nomeNovo string) {
	err := os.Rename(nomeVelho, nomeNovo)
	if err != nil {
		fmt.Println(err)
		return
	}
}
