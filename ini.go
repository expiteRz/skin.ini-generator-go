package main

import (
	"errors"
	"fmt"
	"github.com/sqweek/dialog"
	"gopkg.in/ini.v1"
	"image/color"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func openFile() {
	s, err := dialog.File().Filter(".ini File (*.ini)", "ini").Load()
	if err != nil {
		return
	}

	filename = s
	err = readSetting(s)
	if err != nil {
		fmt.Println(err)
		errorMsg = err.Error()
		errorBox = true
		return
	}
}

func getFontPrefix(r string) string {
	sur := strings.TrimSuffix(filename, "skin.ini")
	d, err := dialog.File().SetStartDir(sur).Filter(".png file (*.png)", "png").Load()
	if err != nil {
		return r
	}

	s := d[len(sur):]
	s = strings.TrimSuffix(s, ".png")
	s = strings.TrimSuffix(s, "@2x")
	if strings.HasSuffix(s, "-comma") {
		return strings.TrimSuffix(s, "-comma")
	}
	if strings.HasSuffix(s, "-dot") {
		return strings.TrimSuffix(s, "-dot")
	}
	if strings.HasSuffix(s, "-percent") {
		return strings.TrimSuffix(s, "-percent")
	}
	s = strings.ReplaceAll(s, `\`, "/")

	return s[:len(s)-2]
}

func readSetting(path string) (err error) {
	skin, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	println(string(skin))

	i := ini.LoadOptions{
		KeyValueDelimiterOnWrite:    ":",
		SkipUnrecognizableLines:     true,
		UnescapeValueCommentSymbols: true,
		UnescapeValueDoubleQuotes:   true,
	}

	l, err := ini.LoadSources(i, path)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Could not read a skin.ini\n")
	}

	initSetting()
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

//parseUInt8 parses uint8 from string
func parseUInt8(src string) uint8 {
	pr, _ := strconv.ParseInt(src, 10, 16)
	return uint8(pr)
}

//parseInt32 parses int32 from string
func parseInt32(src string) int32 {
	pr, _ := strconv.ParseInt(src, 10, 32)
	return int32(pr)
}

// parseBool parses boolean from string,
// 0 -> false, 1 -> true
func parseBool(src string) bool {
	b, err := strconv.ParseBool(src)
	if err != nil {
		return false
	}
	return b
}

// parseInts parses and stores int32 array from source string
func parseInts(src string) []int32 {
	var ia []int32
	split := strings.Split(src, ",")
	for i := range split {
		ia = append(ia, parseInt32(split[i]))
	}

	return ia
}

// parseColor parses color.RGBA from string arrays
func parseColor(src string) color.RGBA {
	var cl color.RGBA
	split := strings.Split(src, ",")
	for i := range split {
		split[i] = removeUnNum(split[i])
	}

	cl.R = parseUInt8(split[0])
	cl.G = parseUInt8(split[1])
	cl.B = parseUInt8(split[2])
	cl.A = 255

	return cl
}

// removeUnNum only stores numbers from source string
func removeUnNum(str string) string {
	s := regexp.MustCompile("[\\d\\-]+").FindAllString(str, -1)
	return strings.Join(s, "")
}
