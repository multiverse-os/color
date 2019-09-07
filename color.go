package color

import (
	"strconv"
)

const (
	escape = "\x1b"
	prefix = escape + "["
	suffix = "m"
)

// Naming uses websafe color names for consistency across platforms and since
// most developers at this point have more experience with those colors
// NOTE: These are commented out to provide developers context, but they are not
// required to be uncommented because we provide functions that will provide
// these values and use less memory.
const (
	RESET = 0
	// Primary ANSI 8 Colors
	BLACK   = 30
	MAROON  = 31
	GREEN   = 32
	OLIVE   = 33
	BLUE    = 34
	MAGENTA = 35
	CYAN    = 36
	SILVER  = 37
	OFF     = 39
	// Primary ANSI 8 Background Colors
	//BLACK_BG    = 40
	//MAROON_BG   = 41
	//GREEN_BG    = 42
	//OLIVE_BG    = 43
	//BLUE_BG     = 44
	//MAGENTA_BG  = 45
	//CYAN_BG     = 46
	//SILVER_BG   = 47
	//BG_OFF      = 49
	// Secondary ANSI 8 Colors
	GRAY     = 90
	RED      = 91
	LIME     = 92
	YELLOW   = 93
	SKY_BLUE = 94
	FUCHSIA  = 95
	AQUA     = 96
	WHITE    = 97
	// Secondary ANSI 8 Background Colors
	//GRAY_BG    = 100
	//RED_BG     = 101
	//LIME_BG    = 102
	//YELLOW_BG  = 103
	//BLUE_BG    = 104
	//FUCHSIA_BG = 105
	//AQUA_BG    = 106
	//WHITE_BG   = 107
)

// Aliasing
const (
	PURPLE = FUCHSIA
	BROWN  = OLIVE
)

func Bg(code int) int          { return (code + 10) }
func Sequence(code int) string { return prefix + strconv.Itoa(code) + suffix }
func Reset() string            { return Sequence(reset) }

func Open(code int) string { return Sequence(code) }
func Close(code int) string {
	switch {
	case (code >= 40 && code <= 49), (code >= 100 && code <= 107):
		return Sequence(Bg(off))
	default:
		return ""
	}
}

func Text(code int, text string) string {
	switch {
	case (code >= 30 && code <= 39), (code >= 90 && code <= 97):
		return Open(code) + text + Close(off)
	case (code >= 40 && code <= 49), (code >= 100 && code <= 107):
		return Sequence(code) + text + Sequence(Bg(off))
	default:
		return text
	}
}

// Foreground Colorization
func Black(text string) string   { return Text(BLACK, text) }
func Maroon(text string) string  { return Text(MAROON, text) }
func Green(text string) string   { return Text(GREEN, text) }
func Olive(text string) string   { return Text(OLIVE, text) }
func Brown(text string) string   { return Text(BROWN, text) }
func Blue(text string) string    { return Text(BLUE, text) }
func Magenta(text string) string { return Text(MAGENTA, text) }
func Cyan(text string) string    { return Text(CYAN, text) }
func Silver(text string) string  { return Text(SILVER, text) }
func Gray(text string) string    { return Text(GRAY, text) }
func Red(text string) string     { return Text(RED, text) }
func Lime(text string) string    { return Text(LIME, text) }
func Yellow(text string) string  { return Text(YELLOW, text) }
func SkyBlue(text string) string { return Text(SKY_BLUE, text) }
func Purple(text string) string  { return Text(PURPLE, text) }
func Fuchsia(text string) string { return Text(FUCHSIA, text) }
func Aqua(text string) string    { return Text(AQUA, text) }
func White(text string) string   { return Text(WHITE, text) }

// Background Colorization
func BlackBg(text string) string   { return Text(Bg(BLACK), text) }
func MaroonBg(text string) string  { return Text(Bg(MAROON), text) }
func GreenBg(text string) string   { return Text(Bg(GREEN), text) }
func OliveBg(text string) string   { return Text(Bg(OLIVE), text) }
func BrownBg(text string) string   { return Text(Bg(BROWN), text) }
func BlueBg(text string) string    { return Text(Bg(BLUE), text) }
func MagentaBg(text string) string { return Text(Bg(MAGENTA), text) }
func CyanBg(text string) string    { return Text(Bg(CYAN), text) }
func SilverBg(text string) string  { return Text(Bg(SILVER), text) }
func GrayBg(text string) string    { return Text(Bg(GRAY), text) }
func RedBg(text string) string     { return Text(Bg(RED), text) }
func LimeBg(text string) string    { return Text(Bg(LIME), text) }
func YellowBg(text string) string  { return Text(Bg(YELLOW), text) }
func SkyBlueBg(text string) string { return Text(Bg(SKY_BLUE), text) }
func PurpleBg(text string) string  { return Text(Bg(PURPLE), text) }
func FuchsiaBg(text string) string { return Text(Bg(FUCHSIA), text) }
func AquaBg(text string) string    { return Text(Bg(AQUA), text) }
func WhiteBg(text string) string   { return Text(Bg(WHITE), text) }
