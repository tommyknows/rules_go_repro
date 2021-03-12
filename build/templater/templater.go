package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		outputFile   = flag.String("out", "", "the output file that will be written")
		templateFile = flag.String("template", "", "the template that will be...templated")
	)
	flag.Parse()

	fmt.Println("Writing to output file", *outputFile)

	out, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("could not open output file: %v", err)
	}
	defer func() { _ = out.Close() }()

	b, err := ioutil.ReadFile(*templateFile)
	if err != nil {
		log.Fatalf("could not read template file: %v", err)
	}

	b = append(b, []byte("// hello!")...)
	_, err = out.Write(b)
	if err != nil {
		log.Fatalf("could not write to file: %v", err)
	}
}
