package basic

type VoiceInterface interface {
	Start() error
	Stop()
	SpeedSub() (string, float64, error)
	SpeedAdd() (string, float64, error)
	SpeekLoop() error
	Requset(string, string) (string, error)
	Say() error
	ChSpeakMe(string)
	Exit()
	GetPause() bool
	SetPause()
	TanslateOrNot() bool
	InvertTranslate()
}
