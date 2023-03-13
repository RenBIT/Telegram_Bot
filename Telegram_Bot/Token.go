package main

import "io/ioutil"

//Можно удалить!! - только для разработки
func Token() string {
	Token, err := ioutil.ReadFile("../../taboo/Token.txt")
	if err != nil {
		panic(err)
	}
	return string(Token)
}
