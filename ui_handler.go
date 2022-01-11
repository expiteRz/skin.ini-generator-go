package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"image/color"
)

func NewRGB(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 255}
}

func NewRGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

func callLayout() []g.Widget {
	contents := []g.Widget{g.PrepareMsgbox()}
	contents = append(contents, callMenu()...)
	contents = append(contents, callContents()...)

	return []g.Widget{
		g.Style().
			SetColor(g.StyleColorWindowBg, NewRGB(42, 31, 73)).
			SetColor(g.StyleColorMenuBarBg, NewRGB(44, 36, 68)).
			SetColor(g.StyleColorHeaderHovered, NewRGB(82, 61, 143)).
			SetColor(g.StyleColorChildBg, NewRGBA(53, 51, 59, 87)).
			SetColor(g.StyleColorFrameBg, NewRGBA(76, 69, 94, 90)).
			SetColor(g.StyleColorFrameBgHovered, NewRGBA(76, 69, 94, 134)).
			SetColor(g.StyleColorButton, NewRGB(82, 61, 143)).
			SetColor(g.StyleColorButtonHovered, NewRGB(107, 78, 191)).
			SetColor(g.StyleColorButtonActive, NewRGB(72, 46, 147)).
			SetColor(g.StyleColorCheckMark, NewRGB(138, 98, 255)).
			SetColor(g.StyleColorBorder, NewRGBA(76, 69, 94, 168)).
			SetColor(g.StyleColorTab, NewRGB(53, 51, 59)).
			SetColor(g.StyleColorTabActive, NewRGB(76, 69, 94)).
			SetColor(g.StyleColorTabHovered, NewRGB(107, 78, 191)).To(contents...)}
}

func callContents() []g.Widget {
	return []g.Widget{
		g.TabBar().TabItems(
			g.TabItem("General").Layout(callGeneral()...),
			g.TabItem("Colors").Layout(callColours()...),
			g.TabItem("Fonts").Layout(callFonts()...),
		),
	}
}

func callMenu() []g.Widget {
	return []g.Widget{
		g.MenuBar().Layout(
			g.Menu("File").Layout(
				g.MenuItem("New"),
				g.MenuItem("Open"),
				g.Separator(),
				g.MenuItem("Save"),
				g.MenuItem("Save as new"),
				g.Separator(),
				g.MenuItem("Quit").OnClick(close),
			),
			g.Menu("Help").Layout(
				g.MenuItem("About").OnClick(openAbout),
				g.PopupModal("About").Layout(
					g.Label("skin.ini Generator v1.0")),
				g.Separator(),
				g.MenuItem("Check for update")),
		),
	}
}

func callGeneral() []g.Widget {
	return []g.Widget{
		g.InputText(&setting.g.Name).Label("Name"),
		g.InputText(&setting.g.Author).Label("Author"),
		g.Combo("Version", version[versionSelected], version, &versionSelected),
		g.InputInt(&setting.g.AnimationFramerate).Label("AnimationFramerate").OnChange(func() {
			if setting.g.AnimationFramerate <= -1 {
				setting.g.AnimationFramerate = -1
			}
		}),
		g.Checkbox("AllowSliderBallTint", &setting.g.AllowSliderBallTint),
		g.Checkbox("ComboBurstRandom", &setting.g.ComboBurstRandom),
		g.Row(
			g.Column(
				g.Checkbox("CursorCentre", &setting.g.CursorCentre),
				g.Checkbox("CursorExpand", &setting.g.CursorExpand),
				g.Checkbox("CursorRotate", &setting.g.CursorRotate),
				g.Checkbox("CursorTrailRotate", &setting.g.CursorTrailRotate),
			),
			// Display cursor image
		),
		g.Row(
			g.Label("CustomComboBurstSounds"),
			g.Button("Add").OnClick(func() {
				setting.g.CustomComboBurstSounds = append(setting.g.CustomComboBurstSounds, 0)
			}), g.Button("Remove").OnClick(func() {
				if len(setting.g.CustomComboBurstSounds) <= 0 {
					return
				}
				setting.g.CustomComboBurstSounds = setting.g.CustomComboBurstSounds[:len(setting.g.CustomComboBurstSounds)-1]
			}),
		),
		g.Table().FastMode(true).Size(450, 150).Rows(callComboBurstRows()...),
		g.Checkbox("HitCircleOverlayAboveNumber", &setting.g.HitCircleOverlayAboveNumber),
		g.Checkbox("LayeredHitSounds", &setting.g.LayeredHitSounds),
		g.Checkbox("SliderBallFlip", &setting.g.SliderBallFlip),
		g.Checkbox("SpinnerFadePlayfield", &setting.g.SpinnerFadePlayfield),
		g.Checkbox("SpinnerFrequencyModulate", &setting.g.SpinnerFrequencyModulate),
		g.Checkbox("SpinnerNoBlank", &setting.g.SpinnerNoBlink),
	}
}

func callColours() []g.Widget {
	return []g.Widget{
		g.Row(
			g.Label("Combo Colors"),
			g.Button("Add").OnClick(func() {
				if len(setting.c.Combos) >= 8 {
					return
				}
				setting.c.Combos = append(setting.c.Combos, color.RGBA{A: 255})
			}), g.Button("Remove").OnClick(func() {
				if len(setting.c.Combos) <= 1 {
					return
				}
				setting.c.Combos = setting.c.Combos[:len(setting.c.Combos)-1]
			}),
		),
		g.Table().FastMode(true).Size(450, 210).Rows(callComboColorRows()...),
		g.ColorEdit("InputOverlayText", &setting.c.InputOverlayText),
		g.ColorEdit("MenuGlow", &setting.c.MenuGlow),
		g.ColorEdit("SliderBall", &setting.c.SliderBall),
		g.ColorEdit("SliderBorder", &setting.c.SliderBorder),
		g.ColorEdit("SongSelectActiveText", &setting.c.SongSelectActiveText),
		g.ColorEdit("SongSelectInactiveText", &setting.c.SongSelectInactiveText),
		g.ColorEdit("SpinnerBackground", &setting.c.SpinnerBackground),
		g.ColorEdit("StarBreakAdditive", &setting.c.StarBreakAdditive),
	}
}

func callFonts() []g.Widget {
	return []g.Widget{
		g.InputText(&setting.f.HitCirclePrefix).Label("HitCirclePrefix"),
		g.InputText(&setting.f.ScorePrefix).Label("ScorePrefix"),
		g.InputText(&setting.f.ComboPrefix).Label("ComboPrefix"),
		g.InputInt(&setting.f.HitCircleOverlap).Label("HitCircleOverlap"),
		g.InputInt(&setting.f.ScoreOverlap).Label("ScoreOverlap"),
		g.InputInt(&setting.f.ComboOverlap).Label("ComboOverlap"),
	}
}

func callComboBurstRows() []*g.TableRowWidget {
	var rows []*g.TableRowWidget
	for i := range setting.g.CustomComboBurstSounds {
		rows = append(rows, g.TableRow(g.InputInt(&setting.g.CustomComboBurstSounds[i]).OnChange(func() {
			if setting.g.CustomComboBurstSounds[i] <= 0 {
				setting.g.CustomComboBurstSounds[i] = 0
			}
		})))
	}

	return rows
}

func callComboColorRows() []*g.TableRowWidget {
	var rows []*g.TableRowWidget
	for i := range setting.c.Combos {
		rows = append(rows, g.TableRow(g.ColorEdit(fmt.Sprintf("Combo %d", i+1), &setting.c.Combos[i])))
	}

	return rows
}

func openAbout() { g.OpenPopup("About") }
