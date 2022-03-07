package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
)

type Person map[string]string

const (
	keyPersonName    = "name"
	keyPersonAddress = "address"
)

func (p *Person) GetName() string {
	return (*p)[keyPersonName]
}

func (p *Person) GetAddress() string {
	return (*p)[keyPersonAddress]
}
func (p *Person) SetName(val string) {
	(*p)[keyPersonName] = val
}

func (p *Person) SetAddress(val string) {
	(*p)[keyPersonAddress] = val
}

func (p *Person) StoreToJson() (string, error) {
	json, err := json.Marshal(*p)

	if err != nil {
		return "", err
	}
	fmt.Println(string(json))
	fileName := fmt.Sprintf("./person_%x.json", md5.Sum([]byte(p.GetName())))

	f, err := os.Create(fileName)
	defer f.Close()

	if err != nil {
		return fileName, err
	}

	_, err = f.Write(json)

	return fileName, err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter your address: ")
	address, _ := reader.ReadString('\n')

	p := make(Person)
	p.SetName(name[:len(name)-1])
	p.SetAddress(address[:len(address)-1])
	fileName, err := p.StoreToJson()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Your personal data stored into %s file", fileName)
}
