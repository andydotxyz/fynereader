package fynereader

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func (r *reader) inputNewFeed(win fyne.Window) {
	input := widget.NewEntry()
	input.SetPlaceHolder("https://                                           ")

	dialog.ShowCustomConfirm("Add URL to feed", "Add", "Cancel", input,
		func(ok bool) {
			if ok {
				r.add(input.Text)
			}
		}, win)
}

func (r *reader) confirmRemove(feed *Feed, win fyne.Window) {
	if feed == nil {
		dialog.ShowInformation("Remove feed", "Cannot remove \"All Feeds\"", win)
		return
	}

	message := fmt.Sprintf("Are you sure you want to remove \"%s\"?", feed.Title)
	dialog.ShowConfirm("Remove feed", message,
		func(ok bool) {
			if ok {
				r.remove(feed)
			}
		}, win)
}
