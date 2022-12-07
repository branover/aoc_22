package main

import (
	"strconv"
	// "sort"
	"strings"
	// "fmt"
)

type dir struct {
	name string
	files []file
	dirs map[string]*dir
	parent *dir
	size int
}

func (d * dir) get_size() int {
	if d.size != 0 {
		return d.size
	}
	size := 0
	for _, file := range d.files {
		// fmt.Println(file.name, file.size)
		size += file.size
	}
	for _, dir := range d.dirs {
		size += dir.get_size()
	}
	d.size = size
	return size
}

type file struct {
	name string
	size int
}

type command_parser struct {
	curr_dir *dir
	all_dirs map[string]*dir
}

func (cp *command_parser) parse(line string) {
	if line[0] == '$' {
		// Shell command
		cp.handle_shell_command(line)
	} else if line[0] == 'd' {
		// Directory
		cp.handle_dir(line)
	} else {
		// File
		cp.handle_file(line)
	}
}

func (cp *command_parser) get_sizes_less_than(size int) int {
	sum := 0
	for _, dir := range cp.all_dirs {
		dir_size := dir.get_size()
		if dir_size <= size {
			// fmt.Printf("%s %d\n", dir.name, dir_size)
			sum += dir_size
		}
	}
	return sum
}

func (cp *command_parser) handle_shell_command(line string) {
	// fmt.Println(line)
	tokens := strings.Split(line, " ")
	switch tokens[1] {
	case "cd": 
		
		dir_name := tokens[2]
		if dir_name == ".." {
			cp.curr_dir = cp.curr_dir.parent
		} else {
			if _, ok := cp.curr_dir.dirs[dir_name]; !ok {
				new_dir := dir {
					name: dir_name,
					files: []file{},
					dirs: map[string]*dir{},
					parent: cp.curr_dir,
				}	
				cp.curr_dir.dirs[dir_name] = &new_dir			
			}			
			cp.curr_dir = cp.curr_dir.dirs[dir_name]
		}
	}
	// fmt.Println(cp.curr_dir.name)
}

func (cp *command_parser) handle_dir(line string) {
	tokens := strings.Split(line, " ")
	dir_name := tokens[1]
	if _, ok := cp.curr_dir.dirs[dir_name]; ok {
		return
	}

	new_dir := dir {
		name: dir_name,
		files: []file{},
		dirs: map[string]*dir{},
		parent: cp.curr_dir,
	}
	cp.curr_dir.dirs[dir_name] = &new_dir

	if _, ok := cp.all_dirs[dir_name]; !ok {
		cp.all_dirs[new_dir.parent.name+dir_name] = &new_dir
	}	
}

func (cp *command_parser) handle_file(line string) {
	tokens := strings.Split(line, " ")
	file_name := tokens[1]
	file_size, _ := strconv.Atoi(tokens[0])
	cp.curr_dir.files = append(cp.curr_dir.files, file {
		file_name,
		file_size,
	})

}



func day7a(lines []string) int {
	var solution int
	var root_dir = dir {
		name: "/",
		files: []file{},
		dirs: map[string]*dir{},
		parent: nil,
	}
	var command_parser = command_parser {
		curr_dir: &root_dir,
		all_dirs: map[string]*dir{},
	}

	for _, line := range lines {
		// Check if line is shell command
		command_parser.parse(line)
	}

	// fmt.Println(root_dir.dirs)
	// fmt.Printf("Total size: %d\n", total_size)
	solution = command_parser.get_sizes_less_than(100000)

	return solution
}

func day7b(lines []string) int {
	var solution int
	var root_dir = dir {
		name: "/",
		files: []file{},
		dirs: map[string]*dir{},
		parent: nil,
	}
	var command_parser = command_parser {
		curr_dir: &root_dir,
		all_dirs: map[string]*dir{},
	}

	for _, line := range lines {
		// Check if line is shell command
		command_parser.parse(line)
	}
	total_size := root_dir.get_size()
	size_to_delete := total_size - (70000000 - 30000000)
	smallest_dir := total_size
	for _, dir := range command_parser.all_dirs {
		dir_size := dir.get_size()		
		if dir_size < smallest_dir && dir_size > size_to_delete{
			smallest_dir = dir_size
		}
	}
	solution = smallest_dir
	return solution
}