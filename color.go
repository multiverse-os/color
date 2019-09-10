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
	reset = 0
	// Primary ANSI 8 Colors
	black   = 30
	maroon  = 31
	green   = 32
	olive   = 33
	blue    = 34
	magenta = 35
	cyan    = 36
	silver  = 37
	off     = 39
	// Primary ANSI 8 Background Colors
	//blackBg    = 40
	//maroonBg   = 41
	//greenBg    = 42
	//oliveBg    = 43
	//blueBg     = 44
	//magentaBg  = 45
	//cyanBg     = 46
	//silverBg   = 47
	//bgOff      = 49
	// Secondary ANSI 8 Colors
	gray    = 90
	red     = 91
	lime    = 92
	yellow  = 93
	skyBlue = 94
	fuchsia = 95
	aqua    = 96
	white   = 97
	// Secondary ANSI 8 Background Colors
	//grayBg    = 100
	//redBg     = 101
	//limeBg    = 102
	//yellowBg  = 103
	//blueBg    = 104
	//fuchsiaBg = 105
	//aquaBg    = 106
	//whiteBg   = 107
)

func Bg(code int) int          { return (code + 10) }
func Sequence(code int) string { return prefix + strconv.Itoa(code) + suffix }
func Reset() string            { return Sequence(reset) }

func Open(code int) string { return Sequence(code) }
func Close(code int) string {
	switch {
	case (code >= 40 && code < 49), (code >= 100 && code < 107):
		return Sequence(Bg(off))
	default:
		return Sequence(reset)
	}
}

func Text(code int, text string) string {
	switch {
	case (code >= 30 && code <= 39), (code >= 90 && code <= 97):
		return Open(code) + text + Close(OFF)
	case (code >= 40 && code <= 49), (code >= 100 && code <= 107):
		return Sequence(code) + text + Sequence(Bg(OFF))
	default:
		return text
	}
}

// Foreground Colorization
func Black(text string) string   { return Text(black, text) }
func Maroon(text string) string  { return Text(maroon, text) }
func Green(text string) string   { return Text(green, text) }
func Olive(text string) string   { return Text(olive, text) }
func Brown(text string) string   { return Text(brown, text) }
func Blue(text string) string    { return Text(blue, text) }
func Magenta(text string) string { return Text(magenta, text) }
func Cyan(text string) string    { return Text(cyan, text) }
func Silver(text string) string  { return Text(silver, text) }
func Gray(text string) string    { return Text(gray, text) }
func Red(text string) string     { return Text(red, text) }
func Lime(text string) string    { return Text(lime, text) }
func Yellow(text string) string  { return Text(yellow, text) }
func SkyBlue(text string) string { return Text(skyBlue, text) }
func Purple(text string) string  { return Text(purple, text) }
func Fuchsia(text string) string { return Text(fuchsia, text) }
func Aqua(text string) string    { return Text(aqua, text) }
func White(text string) string   { return Text(white, text) }

// Background Colorization
func BlackBg(text string) string   { return Text(Bg(black), text) }
func MaroonBg(text string) string  { return Text(Bg(maroon), text) }
func GreenBg(text string) string   { return Text(Bg(green), text) }
func OliveBg(text string) string   { return Text(Bg(olive), text) }
func BrownBg(text string) string   { return Text(Bg(brown), text) }
func BlueBg(text string) string    { return Text(Bg(blue), text) }
func MagentaBg(text string) string { return Text(Bg(magenta), text) }
func CyanBg(text string) string    { return Text(Bg(cyan), text) }
func SilverBg(text string) string  { return Text(Bg(silver), text) }
func GrayBg(text string) string    { return Text(Bg(gray), text) }
func RedBg(text string) string     { return Text(Bg(red), text) }
func LimeBg(text string) string    { return Text(Bg(lime), text) }
func YellowBg(text string) string  { return Text(Bg(yellow), text) }
func SkyBlueBg(text string) string { return Text(Bg(skyBlue), text) }
func PurpleBg(text string) string  { return Text(Bg(purple), text) }
func FuchsiaBg(text string) string { return Text(Bg(fuchsia), text) }
func AquaBg(text string) string    { return Text(Bg(aqua), text) }
func WhiteBg(text string) string   { return Text(Bg(white), text) }
