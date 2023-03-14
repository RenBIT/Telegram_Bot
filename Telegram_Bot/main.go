package main

import (
	"fmt"
	"time"
)

func init() {

	bot.Token = Token() // Тут нужно вписать ваш токен
	//	ParseSiteMp3()            //Подключим парсер сайта с mp3

	//Выводим кнопки
	sendButton("Привет!", "Как дела?", "Подскажи время?", "Мой ник?")

	mass := [...]int{1, 299, 3, 4, 76, 9}
	max := mass[0]
	for i := range mass {
		if max < mass[i] {
			max = mass[i]
		}

	}
	fmt.Println("Максимальное число в массиве:", max)
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
	// отправить аудио ссылка URL
	//sendAudio(url_mp3.File)

	isCommands() //Проверить на текст команды или нажата кнопка
}

// Проверка на команды и нажатий кнопок
func isCommands() {
	// Если пришла команда или нажали кнопку
	switch bot.MsgText {
	case "/start":
		sendKeyBoard() //Рестарт кнопки
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
	case "Мой ник?":
		sendMessage("<b>Твой ник: </b>" + botMessage.Username)
		return
	default:
		//	sendMessage(Command)
		return
	}

}
