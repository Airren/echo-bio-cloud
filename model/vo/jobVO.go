package vo

type JobVO struct {
	Id         int64
	Algorithm  string `json:"algorithm"`
	InputFile  string `json:"inputFile"`
	OutPutFile string `json:"outPutFile"`
	Parameter  string
}
