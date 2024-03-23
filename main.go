package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

type CityData struct {
	min   float64
	max   float64
	sum   float64
	count int
}

const filename = "measurements-1b.txt"

func setupProfiler() func() {
	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		panic(err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}

func main() {
	cleanupProfiler := setupProfiler()
	defer cleanupProfiler()

	fmt.Println("Let's process a billion rows!")

	cities := map[string]CityData{}

	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line string
	var parts []string
	var city string
	var temp float64

	for scanner.Scan() {
		line = scanner.Text()
		parts = strings.Split(line, ";")
		city = parts[0]
		temp, _ = strconv.ParseFloat(parts[1], 64)

		data, present := cities[city]

		if !present {
			cities[city] = CityData{min: temp, max: temp, sum: temp, count: 1}
		} else {
			data.count++
			data.sum += temp

			if temp < data.min {
				data.min = temp
			}
			if temp > data.max {
				data.max = temp
			}

			cities[city] = data
		}
	}

	for city, data := range cities {
		mean := data.sum / float64(data.count)
		fmt.Printf("%s=%.1f/%.1f/%.1f\n", city, data.min, mean, data.max)
	}
}
