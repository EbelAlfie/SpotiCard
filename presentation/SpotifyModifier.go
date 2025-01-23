package presentation

import "spoti-card.com/domain/entity"

type SpoticardModel struct {
	Track     entity.TrackEntity
	IsPlaying bool
}

type CardModifier struct {
	Width  int
	Height int
	Radius int
}

type ImageModifier struct {
	Width  int
	Height int
	X      int
	Y      int
	Url    string
}

type AudioModifier struct {
	Url string
}

type EqualizerModifier struct {
	Y int
}

type TextModifier struct {
	X    int
	Y    int
	Text string
}

type ErrorModel struct {
	Error error
}