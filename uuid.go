package main

import (
	"encoding/json"
)

func GetKeyFromUUID(uuid UUID) string {
	return GetKeyFromBytes(uuid[:])
}
func GetKeyFromMSG(msg *MSG) string {
	return GetKeyFromBytes(GetBytesFromMSG(msg))
}
func GetMSGFromKey(s string) (*MSG, error) {
	msg, err := GetMsgFromBytes(GetBytesFromKey(s))
	return msg, err
}
func GetKeyFromBytes(bytes []byte) string {
	//return base64.URLEncoding.EncodeToString(bytes)
	return string(bytes)
}
func GetBytesFromKey(s string) []byte {
	//return base64.URLEncoding.DecodeString(s)
	return []byte(s)
}
func GetBytesFromMSG(msg *MSG) []byte {
	marshal, err := json.Marshal(msg)
	if err != nil {
		println(err)
		return nil
	}
	return marshal
}
func GetMsgFromBytes(bytes []byte) (*MSG, error) {
	var msg MSG
	err := json.Unmarshal(bytes, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
