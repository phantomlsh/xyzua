package utils

type SimpleResponse struct {
	Success bool
	Message interface{}
}

type ChallengeInfo struct {
	Username  string
	Random    string
	TimeStamp uint
}
