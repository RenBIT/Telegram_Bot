/*
Можно удалить!! -просто получает токен из файла!
*/
package main

import "io/ioutil"

func Token() string {
	Token, err := ioutil.ReadFile("../../taboo/Token.txt")
	if err != nil {
		panic(err)
	}
	return string(Token)
}
