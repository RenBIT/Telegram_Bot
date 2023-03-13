package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func init() {

	//Можно удалить!! - только для разработки
	Token, err := ioutil.ReadFile("../../taboo/Token.txt")
	if err != nil {
		panic(err)
	}
	//////////////////////////////////////////////////

	bot.Token = string(Token) // Тут нужно вписать ваш токен

	//	ParseSiteMp3()            //Подключим парсер сайта с mp3

	//Выводим кнопки
	bot.SetKeyName = []string{"Привет!", "Как дела?", "Подскажи время?", "время?", "Все ?"}
}

/*
Эта функция сработает при входящем сообщение!
*/
func UpdateBot() {
	defer fmt.Println("- Все сообщения отправлены!")
	//проиграть зыук при поступление нового сообщения
	PlaySoundA_Windows("file/audio/AudioMessenge.wav")

	// Отправить картинку
	//sendPhote("<b>Текст картинки</b>", "https://interesnoznat.com/wp-content/uploads/big_1473653108_image.jpg")
	// отправить сообщение
	//sendMessage("Привет мир!!! ")
	// отправить анимацию, мини игра
	//sendDice("⚽")
	// отправить видео ссылка
	//sendVideo("https://static.videezy.com/system/resources/previews/000/037/501/original/Hi-Tech_HUD__global_warming_concept_00352.mp4")
	// отправить аудио ссылка
	//sendAudio(url_mp3.File)

	isCommands()
}

// Проверка на команды и нажатий кнопок
func isCommands() {
	// Если пришла команда или нажали кнопку
	switch bot.MsgText {
	case "/start":
		sendKeyBoard() //кнопки
		return
	case "/help":
		sendMessage(help())
	case "/image":
		sendPhote("<b>Текст картинки</b>", "https://interesnoznat.com/wp-content/uploads/big_1473653108_image.jpg")
	case "Привет!":
		sendMessage("Привет " + botMessage.Username)
		return
	case "Как дела?":
		sendMessage("Все отлично")
		return
	case "Подскажи время?":
		sendMessage("<b>Текущие время: </b>" + time.Now().Format(" 15:04"))
		return
	case "время?":
		sendMessage("Нажата время?")
		return
	default:
		//	sendMessage(Command)
		return
	}

}
