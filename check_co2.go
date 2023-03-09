package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getCO2() string {
	out, err := exec.Command("sudo", "python3", "/usr/local/bin/get_co2.py").Output()
	if err != nil {
		log.Fatal(err)
	}
	CO2Value := string(out)
	return CO2Value
}

func main() {
	var (
		w = flag.Int("w", 0, "int flag")
		c = flag.Int("c", 0, "int flag")
	)
	flag.Parse()

	co2String := getCO2()

	s := strings.ReplaceAll(co2String, "\n", "")
	co2Value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	msg := "CO2 Concentration: %vppm\n"

	fmt.Printf(msg, co2Value)
	if co2Value >= *c {
		os.Exit(2)
	} else if co2Value >= *w {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
