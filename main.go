package main

import (
	"bufio"
	"flag"
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
	profile := flag.Bool("profile", false, "enable CPU profiler")
	filename := flag.String("filename", "", "measurements filename")
	flag.Parse()

	if *profile {
		cleanupProfiler := setupProfiler()
		defer cleanupProfiler()
	}

	fmt.Println("CPU profiling enabled: ", *profile)
	fmt.Println("Mesurements file:", *filename)

	fmt.Println("\nLet's process a billion rows!\n")

	brc(*filename)
}

func brc(filename string) {
	cities := map[string]CityData{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
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
