package color

const (
	escape = "\x1b"
	prefix = "["
	suffix = "m"
)

const (
	Reset    = 0
	Black    = 30
	Red      = 31
	Green    = 32
	Yellow   = 33
	Blue     = 34
	Magenta  = 35
	Cyan     = 36
	Gray     = 37
	ColorOff = 39
	//BlackBackground         = 40
	//RedBackground           = 41
	//GreenBackground         = 42
	//YellowBackground        = 43
	//BlueBackground          = 44
	//MagentaBackground       = 45
	//CyanBackground          = 46
	//GrayBackground          = 47
	//BackgroundOff           = 49
	DarkGray      = 90
	BrightRed     = 91
	BrightGreen   = 92
	BrightYellow  = 93
	BrightBlue    = 94
	BrightMagenta = 95
	BrightCyan    = 96
	White         = 97
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
	Purple = Magenta
)

func Sequence(code int) string { return escape + prefix + strconv.Itoa(code) + suffix }

func Reset(text string) string         { return Reset.Sequence() + text }
func ForegroundOff(text string) string { return Sequence(ForegroundOff.Int()) + text }
func BackgroundOff(text string) string { return Sequence(BackgroundOff.Int()) + text }

func Foreground(code int, text string) string {
	return Sequence(code) + text + Sequence(ForegroundOff.Int())
}
func Background(code int, text string) string {
	return Sequence(code+10) + text + Sequence(BackgroundOff.Int()+10)
}

func BackgroundCode(code int) int { return (code + 10) }

func Color(color int, text string) string {
	switch code := c.Int(); {
	case Reset.Int():
		return Reset(text)
	case (code >= 30 && code <= 37), (code >= 90 && code <= 97):
		return Color(c, text)
	case ForegroundOff.Int():
		return ForegroundOff(text)
	case BackgroundOff.Int():
		return BackgroundOff(text)
	case (code >= 40 && code <= 47), (code >= 100 && code <= 107):
		return Background(c, text)
	default:
		return text
	}
}

// Foreground Colorization
func Black(text string) string        { return Foreground(Black, text) }
func Red(text string) string          { return Foreground(Red, text) }
func Green(text string) string        { return Foreground(Green, text) }
func Yellow(text string) string       { return Foreground(Yellow, text) }
func Blue(text string) string         { return Foreground(Blue, text) }
func Magenta(text string) string      { return Foreground(Magenta, text) }
func Purple(text string) string       { return Foreground(Purple, text) }
func Cyan(text string) string         { return Foreground(Cyan, text) }
func Gray(text string) string         { return Foreground(Gray, text) }
func DarkGray(text string) string     { return Foreground(DarkGray, text) }
func LightRed(text string) string     { return Foreground(LightGray, text) }
func LightGreen(text string) string   { return Foreground(LightRed, text) }
func LightYellow(text string) string  { return Foreground(LightGreen, text) }
func LightBlue(text string) string    { return Foreground(LightBlue, text) }
func LightMagenta(text string) string { return Foreground(LightMagenta, text) }
func LightPurple(text string) string  { return Foreground(LightPurple, text) }
func LightCyan(text string) string    { return Foreground(LightCyan, text) }
func White(text string) string        { return Foreground(White, text) }

// Background Colorization
func BlackBackground(text string) string        { return Background(Black, text) }
func RedBackground(text string) string          { return Background(Red, text) }
func GreenBackground(text string) string        { return Background(Green, text) }
func YellowBackground(text string) string       { return Background(Yellow, text) }
func BlueBackground(text string) string         { return Background(Blue, text) }
func MagentaBackground(text string) string      { return Background(Magenta, text) }
func PurpleBackground(text string) string       { return Background(Purple, text) }
func CyanBackground(text string) string         { return Background(Cyan, text) }
func GrayBackground(text string) string         { return Background(Gray, text) }
func DarkGrayBackground(text string) string     { return Background(DarkGray, text) }
func LightRedBackground(text string) string     { return Background(LightGray, text) }
func LightGreenBackground(text string) string   { return Background(LightRed, text) }
func LightYellowBackground(text string) string  { return Background(LightGreen, text) }
func LightBlueBackground(text string) string    { return Background(LightBlue, text) }
func LightMagentaBackground(text string) string { return Background(LightMagenta, text) }
func LightPurpleBackground(text string) string  { return Background(LightPurple, text) }
func LightCyanBackground(text string) string    { return Background(LightCyan, text) }
func WhiteBackground(text string) string        { return Background(White, text) }
