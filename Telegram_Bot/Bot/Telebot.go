package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Bot struct {
	Token       string //Токен
	TypeMessage string //Тип сообщений
	TokenUrl    string //Ссылка вместе с токеном
	MsgText     string
	SetKeyName  []string
}

var bot Bot

const TIME_UPDATE_BOT time.Duration = 500 * time.Millisecond //Задержка обнавления запроса get

var sendMessengeBot SendMessenge
var sendAudioBot SendAudio
var sendVideoBot SendVideo
var sendPhotoBot sendPhoto
var sendDiceBot SendDice
var sendKeyBoardBot SendKeyBoard
var botMessage BotMessage

func main() {

	bot.TokenUrl = "https://api.telegram.org/bot" + bot.Token

	offset := 0
	fmt.Println("Старт:")

	for {

		updates, err := getUpdates(offset)

		if err != nil {
			log.Println("Ошибка ", err.Error())
		}

		if len(updates) > 0 { //Если сообщение пришло не пустое
			//Выведим в консоль имя  и сообщение от кого пришло
			fmt.Println("\n-", updates[0].Message.Chat.Username, ":", updates[0].Message.Text)
			bot.MsgText = updates[0].Message.Text //Запишем текст что пришол
			botMessage.ChatID = updates[0].Message.Chat.ChatID

			botMessage.Username = updates[0].Message.Chat.Username

			UpdateBotBot() //Обновляем функцию если что то пришло
			for _, update := range updates {
				offset = update.UpdateID + 1 //Получим последние сообщение
			}

		}

		time.Sleep(TIME_UPDATE_BOT) //небольшая задержка
	}

}

func getUpdates(offset int) ([]Update, error) {

	update, err := http.Get(bot.TokenUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	defer update.Body.Close()

	contents, err := io.ReadAll(update.Body)
	if err != nil {
		return nil, err
	}

	var updateResponse UpdateResponse

	err = json.Unmarshal(contents, &updateResponse)
	if err != nil {
		return nil, err
	}

	return updateResponse.Result, nil
}

func sendPhote(caption string, url string) {
	bot.TypeMessage = "sendPhoto"
	sendPhotoBot.Photo = url
	sendPhotoBot.Caption = caption
	sendPhotoBot.ParseMode = "HTML"
	sendPhotoBot.ChatID = botMessage.ChatID

	buf, err := json.Marshal(sendPhotoBot)
	if err != nil {
		return
	}

	postBot(buf)
}

func sendMessage(text string) {
	bot.TypeMessage = "sendMessage"
	sendMessengeBot.ParseMode = "HTML"
	sendMessengeBot.ChatID = botMessage.ChatID
	sendMessengeBot.Text = text

	buf, err := json.Marshal(sendMessengeBot)
	if err != nil {
		return
	}

	postBot(buf)
}

func sendAudio(url string) {

	if getSizeFile(url) < 20000000 { //#Если файл не больше 50мб

		bot.TypeMessage = "sendAudio"
		sendAudioBot.Audio = url
		//sendAudioBot.Title = "Тест"
		//sendAudioBot.Caption = "Доп инфа!!!!!!!!!!!!!!"
		sendAudioBot.ChatID = botMessage.ChatID

		buf, err := json.Marshal(sendAudioBot)
		if err != nil {
			return
		}
		//	fmt.Println("Рфзмер файла", getSizeFile(url), url)
		postBot(buf)
	} else {
		sendMessage("<a href='" + url + "'>Файл слишком большой - слушайте по ссылке..</a>")
		fmt.Println("Файл слишком большой", getSizeFile(url), url)
		return
	}
}

func sendDice(emoji string) {
	bot.TypeMessage = "sendDice"
	sendDiceBot.Emoji = emoji
	sendDiceBot.ChatID = botMessage.ChatID
	buf, err := json.Marshal(sendDiceBot)
	if err != nil {
		return
	}

	postBot(buf)

	//  Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.
	//  Dice can have values 1-6 for “🎲”, “🎯” and “🎳”,
	//  values 1-5 for “🏀” and “⚽”,
	//  and values 1-64 for “🎰”. Defaults to “🎲”
}

func sendVideo(url string) {
	bot.TypeMessage = "sendVideo"
	sendVideoBot.Vdeo = url
	//sendAudioBot.Title = "Тест"
	//sendAudioBot.Caption = "Доп инфа!!!!!!!!!!!!!!"
	sendVideoBot.ChatID = botMessage.ChatID

	buf, err := json.Marshal(sendVideoBot)
	if err != nil {
		return
	}

	postBot(buf)
}

func postBot(buff []byte) { //отправка данных на сервер телеграмм
	fmt.Println(string(buff))
	http.Post(bot.TokenUrl+"/"+bot.TypeMessage, "application/json", bytes.NewBuffer(buff))

}

func sendKeyBoard() {

	bot.TypeMessage = "sendMessage"
	sendKeyBoardBot.ParseMode = "HTML"
	sendKeyBoardBot.ChatID = botMessage.ChatID

	sendKeyBoardBot.ReplyMarkup.ResizeKeyboard = true
	sendKeyBoardBot.Text = "Рестарт : Клавиатуры и тп."

	sendKeyBoardBot.ReplyMarkup.Keyboard[0] = make([]Keyboard, len(bot.SetKeyName))

	for i := range bot.SetKeyName {

		sendKeyBoardBot.ReplyMarkup.Keyboard[0][i].Text = bot.SetKeyName[i]
		sendKeyBoardBot.ReplyMarkup.Keyboard[0][i].CallbackData = bot.SetKeyName[i] //Пока незнаю как вернуть значение нажатой кнопки , записываю текст
	}

	buf, err := json.Marshal(sendKeyBoardBot)
	if err != nil {
		return
	}

	postBot(buf)
}

func help() string {

	fContent, err := ioutil.ReadFile("file/documents/help.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
	return string(fContent)
}
