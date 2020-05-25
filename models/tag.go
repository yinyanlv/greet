package models

type Tag struct {
	Common
	Code     string
	Text     string
	Sort     uint32
}
