package color

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
