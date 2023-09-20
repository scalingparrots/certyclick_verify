package main

import (
	"fmt"

	"github.com/certyclick_verify/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Certyclick Verify")

	w.Resize(fyne.NewSize(800, 300))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	var filePath string
	output := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter previously calculated hash here")
	upload_button := widget.NewButton("Select File", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader == nil {
				return
			}
			if err != nil {
				output.SetText(fmt.Sprintf("Error reading file: %s", err.Error()))
				return
			}

			filePath = reader.URI().Path()
		}, w)
		fd.Show()
	})

	button := widget.NewButton("Verify Hash", func() {
		hashed, err := core.CalculateHash(filePath)
		if err != nil {
			output.SetText(fmt.Sprintf("Error calculating file hash: %s", err.Error()))
			return
		}

		// Compare the calculated hash with the previously stored hash
		if fmt.Sprintf("%x", hashed) == entry.Text {
			output.SetText("The hash is correct.")
		} else {
			output.SetText(fmt.Sprintf("The hash is incorrect. \n\nCalculated hash: %x\n\nProvided hash: %s", hashed, entry.Text))
		}
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("File Hash Verifier"),
		entry,
		upload_button,
		button,
		output,
	))

	w.ShowAndRun()
}
