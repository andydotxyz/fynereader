package fynereader

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type reader struct {
	feedList    *widget.Radio
	itemList    *widget.Box
	itemGroup   *widget.Group
	itemContent *widget.Label
	itemLink    *widget.Hyperlink

	feeds   []*Feed
	current *Feed
}

func (r *reader) add(feedURL string) {
	feed, err := newFeed(feedURL)

	if err != nil {
		log.Println("Unable to load feed!")
		return
	}

	r.feeds = append(r.feeds, feed)
	r.feedList.Options = append(r.feedList.Options, feed.Title)
	r.feedList.Selected = feed.Title
	r.current = feed
	widget.Refresh(r.feedList)

	r.load(feedURL)
}

func (r *reader) load(feedURL string) {
	feed, err := newFeed(feedURL)

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
			r.itemContent.SetText(textWrap(stripTags(item.Description)))
			r.itemLink.SetURLFromString(item.Link)
		}))
	}
}

func (r *reader) remove(feed *Feed) {
	// TODO
	log.Println("TODO remove from the UI and feed list")
}

// Show loads the main reader window for the specified app context
func Show(app fyne.App) {
	read := &reader{}
	read.feedList = widget.NewRadio([]string{"All Feeds"}, func(title string) {
		read.current = nil
		for _, item := range read.feeds {
			if item.Title == title {
				read.current = item
				read.load(item.Link)
			}
		}
	})
	read.itemList = widget.NewVBox()

	w := app.NewWindow("FyneReader")
	feeds := widget.NewGroup("Feeds", read.feedList)
	tools := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			read.inputNewFeed(w)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			read.confirmRemove(read.current, w)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}))
	feedPanel := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, tools, nil, nil),
		feeds, tools)

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

	go read.add("http://fyne.io/feed.xml")

	w.Resize(fyne.NewSize(1024, 576))
	w.Show()
}
