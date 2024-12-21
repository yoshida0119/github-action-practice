package main

var version string // ビルド時にldflag経由で使う

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even change"
	} else {
		return "odd"
	}
}