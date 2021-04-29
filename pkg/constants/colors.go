package constants

type color struct {
	Accent  string
	White   string
	Gray100 string
	Gray200 string
	Gray300 string
	Gray400 string
	Gray500 string
	Gray600 string
	Gray700 string
	Gray800 string
	Gray900 string
}

func DefaultColor() *color {
	return &color{
		Accent:  "#FD4D4D",
		White:   "#FFFFFF",
		Gray100: "#DEE3EA",
		Gray200: "#B2BDCD",
		Gray300: "#5D7290",
		Gray400: "#4F617A",
		Gray500: "#404F64",
		Gray600: "#323D4D",
		Gray700: "#242C37",
		Gray800: "#151A21",
		Gray900: "#0B0E11",
	}
}
