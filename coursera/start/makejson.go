package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
)

type PersonMap map[string]string

const (
	keyPersonName    = "name"
	keyPersonAddress = "address"
)

func (p *PersonMap) GetName() string {
	return (*p)[keyPersonName]
}

func (p *PersonMap) GetAddress() string {
	return (*p)[keyPersonAddress]
}
func (p *PersonMap) SetName(val string) {
	(*p)[keyPersonName] = val
}

func (p *PersonMap) SetAddress(val string) {
	(*p)[keyPersonAddress] = val
}

func (p *PersonMap) StoreToJson() (string, error) {
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

	p := make(PersonMap)
	p.SetName(name[:len(name)-1])
	p.SetAddress(address[:len(address)-1])
	fileName, err := p.StoreToJson()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Your personal data stored into %s file", fileName)
}
