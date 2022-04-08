package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Person struct {
	fname, lname string
}

func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.fname, p.lname)
}

type PersonList []*Person

func (pl *PersonList) AddPerson(fname, lname string) {
	*pl = append(*pl, &Person{fname: fname, lname: lname})
}

func (pl *PersonList) Print() {
	for i := 0; i < len(*pl); i++ {
		fmt.Println((*pl)[i].GetFullName())
	}
}

func readInput() {
	var filepath string
	fmt.Println("Enter a filepath(escape space):")
	fmt.Scan(&filepath)

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(f)
	personList := new(PersonList)

	fname, err := reader.ReadString(' ')
	for ; err == nil || err == io.EOF; fname, err = reader.ReadString(' ') {
		if err == io.EOF {
			personList.AddPerson(fname, "-")
			break
		}
		lname, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			personList.AddPerson(fname[:len(fname)-1], lname[:len(lname)-1])
		}
	}

	if err != nil {
		if err == io.EOF {
			fmt.Println(' ')
		} else {
			fmt.Println(err)
		}
	}
	personList.Print()

}
