package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
	filename        string
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
		if count >= int32(len(version)) {
			count = 0
			break
		}
		count++
	}

	return &count
}

func getCursorImage() string {
	var def = func() string {
		ex, err := os.Executable()
		if err != nil {
			log.Fatalf("Error: Cursor image not found -> %v\n", err)
			return ""
		}
		return filepath.Dir(ex) + "\\cursor.png"
	}
	if filename == "" {
		return def()
	}

	if _, err := os.Stat(strings.TrimSuffix(filename, "skin.ini") + "cursor.png"); err != nil {
		if _, err := os.Stat(strings.TrimSuffix(filename, "skin.ini") + "cursor@2x.png"); err != nil {
			return def()
		}
		return strings.TrimSuffix(filename, "skin.ini") + "cursor@2x.png"
	}

	return strings.TrimSuffix(filename, "skin.ini") + "cursor.png"
}

// false = width, true = height
func getImageSize(m bool) float32 {
	errHndlr := func(err error) {
		fmt.Println(err)
		errorMsg = "Could not get image size"
		errorBox = true
	}

	f, err := os.Open(getCursorImage())
	if err != nil {
		errHndlr(err)
		return 0
	}

	i, _, err := image.Decode(f)
	if err != nil {
		errHndlr(err)
		return 0
	}

	if m {
		return float32(i.Bounds().Dy())
	}
	return float32(i.Bounds().Dx())
}

func setImage() {
	img, _ := g.LoadImage(getCursorImage())
	g.NewTextureFromRgba(img, func(t *g.Texture) {
		texture = t
	})
}

func getImageSetPoint(m bool) image.Point {
	img, _ := os.Open(getCursorImage())
	decode, _, _ := image.Decode(img)

	if decode.Bounds().Dx() <= 100 {
		decX, decY := decode.Bounds().Dx(), decode.Bounds().Dy()
		spX, spY := 100-decX, 100-decY

		if !setting.g.CursorCentre {
			if m {
				return image.Pt(530-spX, 380-spY)
			}
			return image.Pt(410, 260)
		}

		if m {
			return image.Pt(470-(spX/2), 320-(spY/2))
		}
		return image.Pt(350+(spX/2), 200+(spY/2))
	}

	if !setting.g.CursorCentre {
		if m {
			return image.Pt(530, 380)
		}
		return image.Pt(410, 260)
	}

	if m {
		return image.Pt(470, 320)
	}
	return image.Pt(350, 200)
}
