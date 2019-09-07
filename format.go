package color

// Foreground Colorization
func Black(text string) string        { return Color(Black, text) }
func Red(text string) string          { return Color(Red, text) }
func Green(text string) string        { return Color(Green, text) }
func Yellow(text string) string       { return Color(Yellow, text) }
func Blue(text string) string         { return Color(Blue, text) }
func Magenta(text string) string      { return Color(Magenta, text) }
func Purple(text string) string       { return Color(Purple, text) }
func Cyan(text string) string         { return Color(Cyan, text) }
func Gray(text string) string         { return Color(Gray, text) }
func DarkGray(text string) string     { return Color(DarkGray, text) }
func LightRed(text string) string     { return Color(LightGray, text) }
func LightGreen(text string) string   { return Color(LightRed, text) }
func LightYellow(text string) string  { return Color(LightGreen, text) }
func LightBlue(text string) string    { return Color(LightBlue, text) }
func LightMagenta(text string) string { return Color(LightMagenta, text) }
func LightPurple(text string) string  { return Color(LightPurple, text) }
func LightCyan(text string) string    { return Color(LightCyan, text) }
func White(text string) string        { return Color(White, text) }

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
