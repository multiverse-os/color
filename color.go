package color

const (
	escape = "\x1b"
	prefix = escape + "["
	suffix = "m"
)

// Naming uses websafe color names for consistency across platforms and since
// most developers at this point have more experience with those colors
const (
	Reset = 0
	// Primary ANSI 8 Colors
	Black    = 30
	Maroon   = 31
	Green    = 32
	Olive    = 33
	NavyBlue = 34
	Magenta  = 35
	Cyan     = 36
	Silver   = 37
	Default  = 39
	// Primary ANSI 8 Background Colors
	//BlackBg    = 40
	//MaroonBg   = 41
	//GreenBg    = 42
	//OliveBg    = 43
	//NavyBlueBg = 44
	//MagentaBg  = 45
	//CyanBg     = 46
	//SilverBg   = 47
	//DefaultBg  = 49
	// Secondary ANSI 8 Colors
	Gray    = 90
	Red     = 91
	Lime    = 92
	Yellow  = 93
	Blue    = 94
	Fuchsia = 95
	Aqua    = 96
	White   = 97
	// Secondary ANSI 8 Background Colors
	//GrayBg    = 100
	//RedBg     = 101
	//LimeBg    = 102
	//YellowBg  = 103
	//BlueBg    = 104
	//FuchsiaBg = 105
	//AquaBg    = 106
	//WhiteBg   = 107
)

// Aliasing
const (
	Normal = Reset
	Purple = Fuchsia
)

func Bg(code int) int                    { return (code + 10) }
func Sequence(code int) string           { return prefix + strconv.Itoa(code) + suffix }
func Color(code int, text string) string { return Sequence(code) + text + Sequence(Default) }

// Foreground Colorization
func Black(text string) string    { return Color(Black, text) }
func Maroon(text string) string   { return Color(Maroon, text) }
func Green(text string) string    { return Color(Green, text) }
func Olive(text string) string    { return Color(Olive, text) }
func NavyBlue(text string) string { return Color(NavyBlue, text) }
func Magenta(text string) string  { return Color(Magenta, text) }
func Cyan(text string) string     { return Color(Cyan, text) }
func Silver(text string) string   { return Color(Silver, text) }
func Gray(text string) string     { return Color(Gray, text) }
func Red(text string) string      { return Color(Red, text) }
func Lime(text string) string     { return Color(Lime, text) }
func Yellow(text string) string   { return Color(Yellow, text) }
func Blue(text string) string     { return Color(Blue, text) }
func Purple(text string) string   { return Color(Purple, text) }
func Fuchsia(text string) string  { return Color(Fuchsia, text) }
func Aqua(text string) string     { return Color(Aqua, text) }
func White(text string) string    { return Color(White, text) }

// Background Colorization
func BlackBg(text string) string    { return Color(Bg(Black), text) }
func MaroonBg(text string) string   { return Color(Bg(Maroon), text) }
func GreenBg(text string) string    { return Color(Bg(Green), text) }
func OliveBg(text string) string    { return Color(Bg(Olive), text) }
func NavyBlueBg(text string) string { return Color(Bg(NavyBlue), text) }
func MagentaBg(text string) string  { return Color(Bg(Magenta), text) }
func CyanBg(text string) string     { return Color(Bg(Cyan), text) }
func SilverBg(text string) string   { return Color(Bg(Silver), text) }
func GrayBg(text string) string     { return Color(Bg(Gray), text) }
func RedBg(text string) string      { return Color(Bg(Red), text) }
func LimeBg(text string) string     { return Color(Bg(Lime), text) }
func YellowBg(text string) string   { return Color(Bg(Yellow), text) }
func BlueBg(text string) string     { return Color(Bg(Blue), text) }
func PurpleBg(text string) string   { return Color(Bg(Purple), text) }
func FuchsiaBg(text string) string  { return Color(Bg(Fuchsia), text) }
func AquaBg(text string) string     { return Color(Bg(Aqua), text) }
func WhiteBg(text string) string    { return Color(Bg(White), text) }
