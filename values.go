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
	g := General{
		Name:                        "",
		Author:                      "",
		Version:                     "latest",
		AnimationFramerate:          -1,
		AllowSliderBallTint:         false,
		ComboBurstRandom:            false,
		CursorCentre:                true,
		CursorExpand:                true,
		CursorRotate:                true,
		CursorTrailRotate:           true,
		CustomComboBurstSounds:      []int32{},
		HitCircleOverlayAboveNumber: true,
		LayeredHitSounds:            true,
		SliderBallFlip:              true,
		SpinnerFadePlayfield:        false,
		SpinnerFrequencyModulate:    true,
		SpinnerNoBlink:              false,
	}

	c := Colours{
		Combos: []color.RGBA{
			{255, 192, 0, 255},
			{0, 202, 0, 255},
			{18, 124, 255, 255},
			{242, 24, 57, 255},
		},
		InputOverlayText:       color.RGBA{A: 255},
		MenuGlow:               color.RGBA{G: 78, B: 255, A: 255},
		SliderBall:             color.RGBA{R: 2, G: 170, B: 255, A: 255},
		SliderBorder:           color.RGBA{R: 255, G: 255, B: 255, A: 255},
		SongSelectActiveText:   color.RGBA{A: 255},
		SongSelectInactiveText: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		SpinnerBackground:      color.RGBA{R: 100, G: 100, B: 100, A: 255},
		StarBreakAdditive:      color.RGBA{R: 255, G: 182, B: 193, A: 255},
	}

	f := Fonts{
		HitCirclePrefix:  "default",
		HitCircleOverlap: -2,
		ScorePrefix:      "score",
		ScoreOverlap:     0,
		ComboPrefix:      "score",
		ComboOverlap:     0,
	}

	setting = Setting{g, c, f}
	versionSelected = 0
}

func getVersionLen() *int32 {
	var count int32 = 0
	for _, s := range version {
		if s == setting.g.Version {
			break
		}
		count++
	}

	return &count
}
