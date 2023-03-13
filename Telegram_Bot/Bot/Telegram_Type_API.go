package Bot

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type BotMessage struct {
	ChatID   int    `json:"id"`
	Username string `json:"first_name"`
	Text     string `json:"text"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatID   int    `json:"id"`
	Username string `json:"first_name"`
}

type UpdateResponse struct {
	Result []Update `json:"result"`
}

type SendMessenge struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type sendPhoto struct {
	ChatID    int    `json:"chat_id"`
	Photo     string `json:"photo"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}
type SendDice struct {
	ChatID int    `json:"chat_id"`
	Emoji  string `json:"emoji"`
}

type SendAudio struct {
	ChatID int    `json:"chat_id"`
	Audio  string `json:"audio"`
	//Caption string `json:"caption"`
	//Title   string `json:"title"`
}

type SendVideo struct {
	ChatID int    `json:"chat_id"`
	Vdeo   string `json:"video"`
	//Caption string `json:"caption"`
	//Title   string `json:"title"`
}

type InputFile struct {
	InputFile string `json:"-"`
}

// type SendKeyBoard struct {
// 	ChatID      int    `json:"chat_id"`
// 	ParseMode   string `json:"parse_mode"`
// 	Text        string `json:"text"`
// 	ReplyMarkup struct {
// 		ResizeKeyboard bool `json:"resize_keyboard"`
// 		Keyboard       [1][4]struct {
// 			Text         string `json:"text"`
// 			CallbackData string `json:"callback_data"`
// 		} `json:"keyboard"`
// 	} `json:"reply_markup"`
// }

type SendKeyBoard struct {
	ChatID      int         `json:"chat_id"`
	ParseMode   string      `json:"parse_mode"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}
type Keyboard struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}
type ReplyMarkup struct {
	ResizeKeyboard bool          `json:"resize_keyboard"`
	Keyboard       [1][]Keyboard `json:"keyboard"`
}
