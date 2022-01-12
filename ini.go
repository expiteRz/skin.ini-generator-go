package main

import (
	"errors"
	"fmt"
	"github.com/sqweek/dialog"
	"gopkg.in/ini.v1"
	"image/color"
	"strconv"
	"strings"
)

func openFile() {
	s, err := dialog.File().Filter(".ini File (*.ini)", "ini").Load()
	if err != nil {
		return
	}

	err = readSetting(s)
	if err != nil {
		fmt.Println(err)
		errorMsg = err.Error()
		errorBox = true
		return
	}
}

func readSetting(path string) (err error) {
	i := new(ini.LoadOptions)
	i.KeyValueDelimiterOnWrite = ":"

	l, err := ini.LoadSources(*i, path)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Could not read a skin.ini\n")
	}

	sections := l.Sections()
	for i := range sections {
		switch sections[i].Name() {
		case "General":
			setting.g = readGenerals(sections[i].Keys())
		case "Colours":
			setting.c = readColours(sections[i].Keys())
		case "Fonts":
			setting.f = readFonts(sections[i].Keys())
		default:
		}
	}

	versionSelected = *getVersionLen()

	return nil
}

func readGenerals(k []*ini.Key) General {
	g := new(General)
	for _, i := range k {
		switch i.Name() {
		case "Name":
			g.Name = i.Value()
		case "Author":
			g.Author = i.Value()
		case "Version":
			g.Version = i.Value()
		case "AnimationFramerate":
			g.AnimationFramerate = parseInt32(i.Value())
		case "AllowSliderBallTint":
			g.AllowSliderBallTint = parseBool(i.Value())
		case "ComboBurstRandom":
			g.ComboBurstRandom = parseBool(i.Value())
		case "CursorCentre":
			g.CursorCentre = parseBool(i.Value())
		case "CursorExpand":
			g.CursorExpand = parseBool(i.Value())
		case "CursorRotate":
			g.CursorRotate = parseBool(i.Value())
		case "CursorTrailRotate":
			g.CursorTrailRotate = parseBool(i.Value())
		case "CustomComboBurstSounds":
			g.CustomComboBurstSounds = parseInts(i.Value())
		case "HitCircleOverlayAboveNumber":
			g.HitCircleOverlayAboveNumber = parseBool(i.Value())
		case "LayeredHitSounds":
			g.LayeredHitSounds = parseBool(i.Value())
		case "SliderBallFlip":
			g.SliderBallFlip = parseBool(i.Value())
		case "SpinnerFadePlayfield":
			g.SpinnerFadePlayfield = parseBool(i.Value())
		case "SpinnerFrequencyModulate":
			g.SpinnerFrequencyModulate = parseBool(i.Value())
		case "SpinnerNoBlink":
			g.SpinnerNoBlink = parseBool(i.Value())
		}
	}

	return *g
}

func readColours(k []*ini.Key) Colours {
	c := new(Colours)
	for _, i := range k {
		switch i.Name() {
		case "Combo1":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "Combo2":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "Combo3":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "Combo4":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "Combo5":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "Combo6":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "Combo7":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "Combo8":
			c.Combos = append(c.Combos, parseColor(i.Value()))
		case "InputOverlayText":
			c.InputOverlayText = parseColor(i.Value())
		case "MenuGlow":
			c.MenuGlow = parseColor(i.Value())
		case "SliderBall":
			c.SliderBall = parseColor(i.Value())
		case "SliderBorder":
			c.SliderBorder = parseColor(i.Value())
		case "SliderTrackOverrideToggle":
			c.SliderTrackOverrideToggle = parseBool(i.Value())
		case "SliderTrackOverride":
			c.SliderTrackOverride = parseColor(i.Value())
		case "SongSelectActiveText":
			c.SongSelectActiveText = parseColor(i.Value())
		case "SongSelectInactiveText":
			c.SongSelectInactiveText = parseColor(i.Value())
		case "SpinnerBackground":
			c.SpinnerBackground = parseColor(i.Value())
		case "StarBreakAdditive":
			c.StarBreakAdditive = parseColor(i.Value())
		}
	}

	return *c
}

func readFonts(k []*ini.Key) Fonts {
	f := new(Fonts)
	for _, i := range k {
		switch i.Name() {
		case "HitCirclePrefix":
			f.HitCirclePrefix = i.Value()
		case "HitCircleOverlap":
			f.HitCircleOverlap = parseInt32(i.Value())
		case "ScorePrefix":
			f.ScorePrefix = i.Value()
		case "ScoreOverlap":
			f.ScoreOverlap = parseInt32(i.Value())
		case "ComboPrefix":
			f.ComboPrefix = i.Value()
		case "ComboOverlap":
			f.ComboOverlap = parseInt32(i.Value())
		}
	}

	return *f
}

func parseUInt8(src string) uint8 {
	pr, _ := strconv.ParseInt(src, 10, 8)
	return uint8(pr)
}

func parseInt32(src string) int32 {
	pr, _ := strconv.ParseInt(src, 10, 32)
	return int32(pr)
}

func parseBool(src string) bool {
	b, err := strconv.ParseBool(src)
	if err != nil {
		return false
	}
	return b
}

func parseInts(src string) []int32 {
	var ia []int32
	split := strings.Split(src, ",")
	for i := range split {
		ia = append(ia, parseInt32(split[i]))
	}

	return ia
}

func parseColor(src string) color.RGBA {
	var cl color.RGBA
	split := strings.Split(src, ",")
	cl.R = parseUInt8(split[0])
	cl.G = parseUInt8(split[1])
	cl.B = parseUInt8(split[2])
	cl.A = 255

	return cl
}
