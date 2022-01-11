package main

import "image/color"

type Setting struct {
	g General
	c Colours
	f Fonts
}

type General struct {
	Name                        string
	Author                      string
	Version                     string
	AnimationFramerate          int32
	AllowSliderBallTint         bool
	ComboBurstRandom            bool
	CursorCentre                bool
	CursorExpand                bool
	CursorRotate                bool
	CursorTrailRotate           bool
	CustomComboBurstSounds      []int32
	HitCircleOverlayAboveNumber bool
	LayeredHitSounds            bool
	SliderBallFlip              bool
	SpinnerFadePlayfield        bool
	SpinnerFrequencyModulate    bool
	SpinnerNoBlink              bool
}

type Colours struct {
	Combos                    []color.RGBA
	InputOverlayText          color.RGBA
	MenuGlow                  color.RGBA
	SliderBall                color.RGBA
	SliderBorder              color.RGBA
	SliderTrackOverrideToggle bool
	SliderTrackOverride       color.RGBA
	SongSelectActiveText      color.RGBA
	SongSelectInactiveText    color.RGBA
	SpinnerBackground         color.RGBA
	StarBreakAdditive         color.RGBA
}

type Fonts struct {
	HitCirclePrefix  string
	HitCircleOverlap int32
	ScorePrefix      string
	ScoreOverlap     int32
	ComboPrefix      string
	ComboOverlap     int32
}

//type color struct {
//	r int32
//	g int32
//	b int32
//}

var (
	setting         Setting
	versionSelected int32
)

func initSetting() {
	setting.g.AnimationFramerate = -1
	setting.g.CursorCentre = true
	setting.g.CursorExpand = true
	setting.g.CursorRotate = true
	setting.g.CursorTrailRotate = true
	setting.g.HitCircleOverlayAboveNumber = true
	setting.g.LayeredHitSounds = true
	setting.g.SliderBallFlip = true
	setting.g.SpinnerFrequencyModulate = true

	setting.c.Combos = []color.RGBA{
		{255, 192, 0, 255},
		{0, 202, 0, 255},
		{18, 124, 255, 255},
		{242, 24, 57, 255},
	}
	setting.c.InputOverlayText = color.RGBA{A: 255}
	setting.c.MenuGlow = color.RGBA{G: 78, B: 255, A: 255}
	setting.c.SliderBall = color.RGBA{R: 2, G: 170, B: 255, A: 255}
	setting.c.SliderBorder = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	setting.c.SongSelectActiveText = color.RGBA{A: 255}
	setting.c.SongSelectInactiveText = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	setting.c.SpinnerBackground = color.RGBA{R: 100, G: 100, B: 100, A: 255}
	setting.c.StarBreakAdditive = color.RGBA{R: 255, G: 182, B: 193, A: 255}

	setting.f.HitCirclePrefix = "default"
	setting.f.HitCircleOverlap = -2
	setting.f.ScorePrefix = "score"
	setting.f.ScoreOverlap = 0
	setting.f.ComboPrefix = "score"
	setting.f.ComboOverlap = 0
}
