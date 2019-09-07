package color

const (
	escape = "\x1b"
	prefix = escape + "["
	suffix = "m"
)

// Naming uses websafe color names for consistency across platforms and since
// most developers at this point have more experience with those colors
const (
	Reset        = 0
	Black        = 30
	Maroon       = 31
	Green        = 32
	Olive        = 33
	NavyBlue     = 34
	Magenta      = 35
	Cyan         = 36
	Silver       = 37
	DefaultColor = 39
	//BlackBackground         = 40
	//RedBackground           = 41
	//GreenBackground         = 42
	//YellowBackground        = 43
	//BlueBackground          = 44
	//MagentaBackground       = 45
	//CyanBackground          = 46
	//GrayBackground          = 47
	//BackgroundOff           = 49
	Gray    = 90
	Red     = 91
	Lime    = 92
	Yellow  = 93
	Blue    = 94
	Fuchsia = 95
	Aqua    = 96
	White   = 97
	//DarkGrayBackground      = 100
	//BrightRedBackground     = 101
	//BrightGreenBackground   = 102
	//BrightYellowBackground  = 103
	//BrightBlueBackground    = 104
	//BrightMagentaBackground = 105
	//BrightCyanBackground    = 106
	//WhiteBackground         = 107
)

// Aliasing
const (
	Normal = Reset
	Purple = Fuchsia
)

func BgCode(code int) int { return (code + 10) }

func Sequence(code int) string { return escape + prefix + strconv.Itoa(code) + suffix }

func Reset(text string) string          { return Reset.Sequence() + text }
func DefaultColor(text string) string   { return Sequence(DefaultColor.Int()) + text }
func DefaultBgColor(text string) string { return Sequence(BgCode(DefaultColor.Int())) + text }

func Fg(code int, text string) string {
	return Sequence(code) + text + Sequence(DefaultColor.Int())
}
func Bg(code int, text string) string {
	return Sequence(BgCode(code)) + text + Sequence(BgCode(DefaultColor.Int()))
}

func Color(color int, text string) string {
	switch code := c.Int(); {
	case Reset.Int():
		return Reset(text)
	case (code >= 30 && code <= 37), (code >= 90 && code <= 97):
		return Fg(c, text)
	case DefaultColor.Int():
		return DefaultColor(text)
	case (code >= 40 && code <= 47), (code >= 100 && code <= 107):
		return Bg(c, text)
	default:
		return text
	}
}

// Foreground Colorization
func Black(text string) string    { return Fg(Black, text) }
func Maroon(text string) string   { return Fg(Maroon, text) }
func Green(text string) string    { return Fg(Green, text) }
func Olive(text string) string    { return Fg(Olive, text) }
func NavyBlue(text string) string { return Fg(NavyBlue, text) }
func Magenta(text string) string  { return Fg(Magenta, text) }
func Cyan(text string) string     { return Fg(Cyan, text) }
func Silver(text string) string   { return Fg(Silver, text) }
func Gray(text string) string     { return Fg(Gray, text) }
func Red(text string) string      { return Fg(Red, text) }
func Lime(text string) string     { return Fg(Lime, text) }
func Yellow(text string) string   { return Fg(Yellow, text) }
func Blue(text string) string     { return Fg(Blue, text) }
func Purple(text string) string   { return Fg(Purple, text) }
func Fuchsia(text string) string  { return Fg(Fuchsia, text) }
func Aqua(text string) string     { return Fg(Aqua, text) }
func White(text string) string    { return Fg(White, text) }

// Background Colorization
func BlackBg(text string) string    { return Bg(Black, text) }
func MaroonBg(text string) string   { return Bg(Maroon, text) }
func GreenBg(text string) string    { return Bg(Green, text) }
func OliveBg(text string) string    { return Bg(Olive, text) }
func NavyBlueBg(text string) string { return Bg(NavyBlue, text) }
func MagentaBg(text string) string  { return Bg(Magenta, text) }
func CyanBg(text string) string     { return Bg(Cyan, text) }
func SilverBg(text string) string   { return Bg(Silver, text) }
func GrayBg(text string) string     { return Bg(Gray, text) }
func RedBg(text string) string      { return Bg(Red, text) }
func LimeBg(text string) string     { return Bg(Lime, text) }
func YellowBg(text string) string   { return Bg(Yellow, text) }
func BlueBg(text string) string     { return Bg(Blue, text) }
func PurpleBg(text string) string   { return Bg(Purple, text) }
func FuchsiaBg(text string) string  { return Bg(Fuchsia, text) }
func AquaBg(text string) string     { return Bg(Aqua, text) }
func WhiteBg(text string) string    { return Bg(White, text) }
