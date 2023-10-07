package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

type File struct {
	name     string
	size     int
	subfiles []*File
	parent   *File
}

func (file File) Size() int {
	total := file.size
	if total == 0 {
		for _, subfile := range file.subfiles {
			total += subfile.Size()
		}
		// file.size = total
	}
	//TODO
	// if (total >= root - 40000000) {
	// 	fmt.Println(total, file.name)
	// }
	return total
}

func (file File) Total() int {
	total := 0
	if file.Size() <= 100000 && file.subfiles != nil {
		total = file.Size()
	}
	for _, subfile := range file.subfiles {
		total += subfile.Total()
	}
	return total
}

func main() {
	input := utils.ReadFile(2022, 7, "\n")
	fmt.Println("2022 Day 07")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) int {
	currentFile := &File{
		name:     "/",
		size:     0,
		subfiles: []*File{},
		parent:   nil,
	}
	rootFile := currentFile
	for lineIndex := 1; lineIndex < len(input); lineIndex++ {
		line := strings.Split(input[lineIndex], " ")
		if input[lineIndex] == "$ cd .." {
			currentFile = (*currentFile).parent
		} else if line[0] == "$" && line[1] == "cd" {
			for _, subfile := range (*currentFile).subfiles {
				if subfile.name == line[2] {
					currentFile = subfile
					break
				}
			}
		} else if line[0] == "dir" {
			(*currentFile).subfiles = append((*currentFile).subfiles, &File{
				name:     line[1],
				size:     0,
				subfiles: []*File{},
				parent:   currentFile,
			})
		} else if line[0] != "$" {
			(*currentFile).subfiles = append((*currentFile).subfiles, &File{
				name:     line[1],
				size:     utils.StringToInt(line[0]),
				subfiles: nil,
				parent:   currentFile,
			})
		}
	}
	return (*rootFile).Total()
}

func part2(input []string) int {
	currentFile := &File{
		name:     "/",
		size:     0,
		subfiles: []*File{},
		parent:   nil,
	}
	rootFile := currentFile
	for lineIndex := 1; lineIndex < len(input); lineIndex++ {
		line := strings.Split(input[lineIndex], " ")
		if input[lineIndex] == "$ cd .." {
			currentFile = (*currentFile).parent
		} else if line[0] == "$" && line[1] == "cd" {
			for _, subfile := range (*currentFile).subfiles {
				if subfile.name == line[2] {
					currentFile = subfile
					break
				}
			}
		} else if line[0] == "dir" {
			(*currentFile).subfiles = append((*currentFile).subfiles, &File{
				name:     line[1],
				size:     0,
				subfiles: []*File{},
				parent:   currentFile,
			})
		} else if line[0] != "$" {
			(*currentFile).subfiles = append((*currentFile).subfiles, &File{
				name:     line[1],
				size:     utils.StringToInt(line[0]),
				subfiles: nil,
				parent:   currentFile,
			})
		}
	}
	return (*rootFile).Size()
}
