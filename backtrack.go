package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Tile struct{
	wall	bool
	visited bool
}

var filename = "template2.txt"
var size = 7

func main() {
	arr := createArray(size)
	loadTemplate(filename, arr)
	arr[11].visited = true
	fmt.Println(recur(arr, 11, 1, size*size))
}

func loadTemplate(name string, arr map[int]*Tile) {
	f, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if number, err := strconv.Atoi(scanner.Text()); err == nil{
			arr[number].wall = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createArray(size int) map[int]*Tile {
	arr := map[int]*Tile{}
	for x := 1; x <= size; x++ {
		for y := 1; y <= size; y++{
			arr[x*10 + y] = &Tile{wall: false, visited: false}
		}
	}
	return arr
}

func isPossible(arr map[int]*Tile, position int) (possible bool){
	if _, ok := arr[position]; ok {
		if !arr[position].visited {
			if !arr[position].wall {
				return !possible
			}
		}
	}
	return
} //controls whether the movement is possible

func recur(arr map[int]*Tile, position int, distance int, minDist int) int{
	//fmt.Println(position)
	if position == 11*size {
		return int(math.Min(float64(distance),float64(minDist)))
	}

	arr[position].visited = true


	if isPossible(arr, position+1){
		minDist = recur(arr, position+1, distance+1, minDist)
	}

	if isPossible(arr, position-1){
		minDist = recur(arr, position-1, distance+1, minDist)
	}
	if isPossible(arr, position+10){
		minDist = recur(arr, position+10, distance+1, minDist)
	}

	if isPossible(arr, position-10){
		minDist = recur(arr, position-10, distance+1, minDist)
	}

	arr[position].visited = false
	return minDist
}