package main

const UUID_LEN = 16

type UUID [UUID_LEN]byte
type CNT int64
type TimeStamp int64
type MSG struct {
	Src       UUID      `json:"src"`
	Content   []byte    `json:"content"`
	TimeStamp TimeStamp `json:"timeStamp"`
}
type Backend interface {
	Len(uuid UUID) (CNT, error)
	Get(uuid UUID, cnt CNT) ([]MSG, error)
	Write(dst UUID, msg *MSG) error
	Skip(uuid UUID, cnt CNT) (CNT, error)
}
