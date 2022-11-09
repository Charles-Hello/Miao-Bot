package telegram_bot

type Ret struct {
	Code    int `default:"200"`
	OpSlice []string
}

var Return = new(Ret)
