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

func Color(color int, text string) string { return Sequence(color) + text + Sequence(ColorOff) }
func Background(color int, text string) string {
	return Sequence(color+10) + text + Sequence(ColorOff+10)
}
