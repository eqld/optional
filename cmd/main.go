package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	pkg := os.Getenv("GOPACKAGE")
	dir := os.Getenv("PWD")
	file := os.Getenv("GOFILE")
	line := os.Getenv("GOLINE")

	tt := os.Args[1:]

	if tn, ok := parse(dir, file, line); ok {
		tt = append(tt, tn)
	}

	for _, t := range tt {
		generate(dir, pkg, t)
	}
}

func parse(dir, file, line string) (string, bool) {
	f, err := os.Open(dir + "/" + file)
	if err != nil {
		panic(err)
	}

	ln, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	for i := 0; i <= ln; i++ {
		s.Scan()
	}

	l := strings.SplitN(strings.TrimSpace(s.Text()), " ", 3)

	if l[0] == "type" {
		return l[1], true
	}

	return "", false
}

func generate(dir, pkg, t string) {
	tl := strings.ToLower(t)
	to := strings.ToUpper(t[:1]) + t[1:]

	if pkg != "optional" {
		to = "Optional" + to
	}

	filename := "optional_" + tl + ".generated.go"

	f, err := os.Create(dir + "/" + filename)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(strings.NewReplacer("_P_", pkg, "_TO_", to, "_T_", t).Replace(template))
	if err != nil {
		panic(err)
	}
}
