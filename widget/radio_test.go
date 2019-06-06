package widget

import (
	"fmt"
	"testing"

	"fyne.io/fyne"
	_ "fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"github.com/stretchr/testify/assert"
)

func TestRadio_MinSize(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, nil)
	min := radio.MinSize()

	assert.True(t, min.Width > theme.Padding()*2)
	assert.True(t, min.Height > theme.Padding()*2)

	radio2 := NewRadio([]string{"Hi", "He"}, nil)
	min2 := radio2.MinSize()

	assert.Equal(t, min.Width, min2.Width)
	assert.True(t, min2.Height > min.Height)
}

func TestRadio_BackgroundStyle(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, nil)
	bg := Renderer(radio).BackgroundColor()

	assert.Equal(t, bg, theme.BackgroundColor())
}

func TestRadio_MinSize_Horizontal(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, nil)
	min := radio.MinSize()

	radio2 := NewRadio([]string{"Hi", "He"}, nil)
	radio2.Horizontal = true
	min2 := radio2.MinSize()

	assert.True(t, min2.Width > min.Width)
	assert.Equal(t, min.Height, min2.Height)
}

func TestRadio_Selected(t *testing.T) {
	selected := ""
	radio := NewRadio([]string{"Hi"}, func(sel string) {
		selected = sel
	})
	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(theme.Padding(), theme.Padding())})

	assert.Equal(t, "Hi", selected)
}

func TestRadio_Unselected(t *testing.T) {
	selected := "Hi"
	radio := NewRadio([]string{"Hi"}, func(sel string) {
		selected = sel
	})
	radio.Selected = selected
	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(theme.Padding(), theme.Padding())})

	assert.Equal(t, "", selected)
}

func TestRadio_DisableWhenSelected(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, nil)
	radio.SetSelected("Hi")
	render := Renderer(radio).(*radioRenderer)
	resName := render.items[0].icon.Resource.Name()

	assert.Equal(t, resName, theme.RadioButtonCheckedIcon().Name())

	radio.Disable()
	resName = render.items[0].icon.Resource.Name()
	assert.Equal(t, resName, fmt.Sprintf("disabled_%v", theme.RadioButtonCheckedIcon().Name()))
}

func TestRadio_DisableWhenNotSelected(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, nil)
	render := Renderer(radio).(*radioRenderer)
	resName := render.items[0].icon.Resource.Name()

	assert.Equal(t, resName, theme.RadioButtonIcon().Name())

	radio.Disable()
	resName = render.items[0].icon.Resource.Name()
	assert.Equal(t, resName, fmt.Sprintf("disabled_%v", theme.RadioButtonIcon().Name()))
}

func TestRadio_SelectedOther(t *testing.T) {
	selected := "Hi"
	radio := NewRadio([]string{"Hi", "Hi2"}, func(sel string) {
		selected = sel
	})
	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(theme.Padding(), radio.MinSize().Height-theme.Padding())})

	assert.Equal(t, "Hi2", selected)
}

func TestRadio_SelectedHorizontal(t *testing.T) {
	selected := "Hi"
	radio := NewRadio([]string{"Hi", "Hi2"}, func(sel string) {
		selected = sel
	})
	radio.Horizontal = true
	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(radio.MinSize().Width-theme.Padding(), theme.Padding())})

	assert.Equal(t, "Hi2", selected)
}

func TestRadio_SelectedNone(t *testing.T) {
	selected := ""
	radio := NewRadio([]string{"Hi"}, func(sel string) {
		selected = sel
	})

	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(0, -2)})
	assert.Equal(t, "", selected)

	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(0, 25)})
	assert.Equal(t, "", selected)
}

func TestRadio_Append(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, nil)

	assert.Equal(t, 1, len(radio.Options))
	assert.Equal(t, 1, len(Renderer(radio).(*radioRenderer).items))

	radio.Options = append(radio.Options, "Another")
	Refresh(radio)

	assert.Equal(t, 2, len(radio.Options))
	assert.Equal(t, 2, len(Renderer(radio).(*radioRenderer).items))
}

func TestRadio_Remove(t *testing.T) {
	radio := NewRadio([]string{"Hi", "Another"}, nil)

	assert.Equal(t, 2, len(radio.Options))
	assert.Equal(t, 2, len(Renderer(radio).(*radioRenderer).items))

	radio.Options = radio.Options[:1]
	Refresh(radio)

	assert.Equal(t, 1, len(radio.Options))
	assert.Equal(t, 1, len(Renderer(radio).(*radioRenderer).items))
}

func TestRadio_SetSelected(t *testing.T) {
	radio := NewRadio([]string{"Hi", "Another"}, nil)

	radio.SetSelected("Another")

	assert.Equal(t, "Another", radio.Selected)
}

func TestRadio_SetSelectedWithSameOption(t *testing.T) {
	radio := NewRadio([]string{"Hi", "Another"}, nil)

	radio.Selected = "Another"
	Refresh(radio)

	radio.SetSelected("Another")

	assert.Equal(t, "Another", radio.Selected)
}

func TestRadio_SetSelectedEmpty(t *testing.T) {
	radio := NewRadio([]string{"Hi", "Another"}, nil)

	radio.Selected = "Another"
	Refresh(radio)

	radio.SetSelected("")

	assert.Equal(t, "", radio.Selected)
}

func TestRadio_DuplicatedOptions(t *testing.T) {
	radio := NewRadio([]string{"Hi", "Hi", "Hi", "Another", "Another"}, nil)

	assert.Equal(t, 2, len(radio.Options))
	assert.Equal(t, 2, len(Renderer(radio).(*radioRenderer).items))
}

func TestRadio_AppendDuplicate(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, nil)

	radio.Append("Hi")

	assert.Equal(t, 1, len(radio.Options))
	assert.Equal(t, 1, len(Renderer(radio).(*radioRenderer).items))
}

func TestRadio_Disable(t *testing.T) {
	selected := ""
	radio := NewRadio([]string{"Hi"}, func(sel string) {
		selected = sel
	})

	radio.Disable()
	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(theme.Padding(), theme.Padding())})

	assert.Equal(t, "", selected, "Radio should have been disabled.")
}

func TestRadio_Enable(t *testing.T) {
	selected := ""
	radio := NewRadio([]string{"Hi"}, func(sel string) {
		selected = sel
	})

	radio.Disable()
	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(theme.Padding(), theme.Padding())})
	assert.Equal(t, "", selected, "Radio should have been disabled.")

	radio.Enable()
	radio.Tapped(&fyne.PointEvent{Position: fyne.NewPos(theme.Padding(), theme.Padding())})
	assert.Equal(t, "Hi", selected, "Radio should have been re-enabled.")
}

func TestRadio_Disabled(t *testing.T) {
	radio := NewRadio([]string{"Hi"}, func(string) {})
	assert.False(t, radio.Disabled())
	radio.Disable()
	assert.True(t, radio.Disabled())
	radio.Enable()
	assert.False(t, radio.Disabled())
}
