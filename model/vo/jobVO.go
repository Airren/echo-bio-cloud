package vo

type JobVO struct {
	Id         uint64 `json:"id"`
	Algorithm  string `json:"algorithm"`
	InputFile  string `json:"inputFile"`
	OutPutFile string `json:"outPutFile"`
	Parameter  string
}
