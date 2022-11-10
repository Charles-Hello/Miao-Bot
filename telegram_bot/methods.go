package telegram_bot

import "strconv"

func TG_reply(a string) {
	Return.OpSlice = append(Return.OpSlice, "reply"+a)
}

func TG_sleep(a int) {
	Return.OpSlice = append(Return.OpSlice, "sleep"+strconv.Itoa(a))
}
func TG_edit(a string) {
	Return.OpSlice = append(Return.OpSlice, "edit"+a)
}
func TG_delete() {
	Return.OpSlice = append(Return.OpSlice, "delete")
}

func TG_send(a string) {
	Return.OpSlice = append(Return.OpSlice, "send"+a)
}
