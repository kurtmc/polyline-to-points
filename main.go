package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/scanner"
	"strconv"
)

func main() {
	fmt.Println("test")
	s, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(s))
	parsePolyline(string(s))
}

func expect(s scanner.Scanner, expected string) {
	if s.TokenText() != expected {
		log.Fatal("At position", s.Pos(), ", was '", s.TokenText(), "' but expected '", expected, "'")
	}
}

func getFloat64(s scanner.Scanner) float64 {
	f, err := strconv.ParseFloat(s.TokenText(), 64)
	if err != nil {
		log.Fatal("At position", s.Pos(), ", was '", s.TokenText(), "' but expected float string")
	}
	return f
}

func parsePolyline(input string) {
	var polylines = make([][]interface{}, 0)
	var polyline []interface{}
	var x, y, z float64
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	var tok rune
	for tok != scanner.EOF {
		if s.TokenText() == "Open" {
			polyline = make([]interface{}, 0)
			polylines = append(polylines, polyline)
		}

		if s.TokenText() == "at" {
			tok = s.Scan()
			expect(s, "point")

			tok = s.Scan()
			expect(s, "X")
			tok = s.Scan()
			expect(s, "=")
			tok = s.Scan()
			x = getFloat64(s)

			tok = s.Scan()
			expect(s, "Y")
			tok = s.Scan()
			expect(s, "=")
			tok = s.Scan()
			y = getFloat64(s)

			tok = s.Scan()
			expect(s, "Z")
			tok = s.Scan()
			expect(s, "=")
			tok = s.Scan()
			z = getFloat64(s)

			p := Point{X: x, Y: y, Z: z}

			polyline = append(polyline, p)
		}

		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
	}

	fmt.Printf("polyline: %v\n", polyline)
}
