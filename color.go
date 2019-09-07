package color

const (
	escape = "\x1b"
	prefix = "["
	suffix = "m"
)

const (
	Normal        = 0
	Reset         = 0
	ForegroundOff = 39
	BackgroundOff = 49
)

const (
	Black         = 30
	Red           = 31
	Green         = 32
	Yellow        = 33
	Blue          = 34
	Magenta       = 35
	Cyan          = 36
	Gray          = 37
	DarkGray      = 90
	BrightRed     = 91
	BrightGreen   = 92
	BrightYellow  = 93
	BrightBlue    = 94
	BrightMagenta = 95
	BrightCyan    = 96
	White         = 97
)

func Sequence(code int) string { return escape + prefix + strconv.Itoa(code) + suffix }

func Foreground(color int, text string) string { return ForegroundOn(color) + text + ForegroundOff() }
func Background(color int, text string) string { return BackgroundOn(color) + text + BackgroundOff() }

func ForegroundOn(code int) string { return Sequence(code) }
func BackgroundOn(code int) string { return Sequence(BackgroundCode(code + 10)) }

func ForegroundOff() string { return Sequence(ForegroundOff) }
func BackgroundOff() string { return Sequence(BackgroundOff) }
