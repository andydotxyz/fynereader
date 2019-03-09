package fynereader

import (
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/mmcdole/gofeed"
)

type reader struct {
	feedList, itemList *widget.Box
	itemGroup          *widget.Group
	itemContent        *widget.Label
	itemLink           *widget.Hyperlink
}

func (r *reader) load(feedURL string) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)

	if err != nil {
		log.Println("Unable to load feed!")
		return
	}

	r.itemGroup.Text = feed.Title
	r.itemList.Children = nil
	widget.Refresh(r.itemGroup)

	for i := range feed.Items {
		item := feed.Items[i] // keep a reference to the slices
		r.itemList.Append(widget.NewButton(item.Title, func() {
			r.itemContent.SetText(item.Description)
			r.itemLink.SetURLFromString(fmt.Sprintf("%s#about", item.Link))
		}))
	}
}

// Show loads the main reader window for the specified app context
func Show(app fyne.App) {
	read := &reader{}
	read.feedList = widget.NewVBox(widget.NewButton("All Feeds", func() {}))
	read.itemList = widget.NewVBox()

	w := app.NewWindow("FyneReader")
	feeds := widget.NewGroup("Feeds", read.feedList)
	tools := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}))
	feedPanel := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, tools, nil, nil),
		feeds, tools)

	defaultURL := "http://fyne.io/feed.xml"
	read.itemGroup = widget.NewGroupWithScroller("Items", read.itemList)
	read.itemContent = widget.NewLabel("TODO add content here")

	read.itemLink = widget.NewHyperlink("Read more...", nil)
	body := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, read.itemLink, nil, nil),
		fyne.NewContainerWithLayout(layout.NewGridLayout(1),
			read.itemGroup, widget.NewGroupWithScroller("Content", read.itemContent)),
		read.itemLink)
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, feedPanel, nil),
		feedPanel, body))

	feeds.Append(widget.NewButton("Fyne", func() {
		read.load(defaultURL)
	}))
	go read.load(defaultURL)

	w.Resize(fyne.NewSize(1024, 576))
	w.Show()
}
