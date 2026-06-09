package main

import (
	"fmt"
)
type Icon struct{
	domain string
	image  string
}

func (i *Icon) Display() {
	fmt.Printf("Displaying icon for %s\n", i.domain)
}
func NewIcon(domain string)*Icon{

	return &Icon{
		domain: domain,
		image: domain+".ico",
	}

}
type IconFactory struct{
	cache map[string]*Icon
}

func (i *IconFactory) GetIcon(domain string) *Icon{
	icon ,ok := i.cache[domain]

	if ok {
		return icon
	}
	fmt.Printf("Cache MISS -> %s\n", domain)

	icon = NewIcon(domain)

	i.cache[domain]=icon


	return icon

}

func NewIconFactory() *IconFactory{
	return &IconFactory{
		cache: make(map[string]*Icon),
	}
}


type BookMark struct{
	Name string
	Icon *Icon
	Url string
}

func NewBookmark(
	title string,
	url string,
	icon *Icon,
) *BookMark {

	return &BookMark{
		Name: title,
		Url:   url,
		Icon:  icon,
	}
}

func (b *BookMark) Show() {
	fmt.Printf(
		"Bookmark: %-15s URL: %-25s Icon Address: %p\n",
		b.Name,
		b.Url,
		b.Icon,
	)
}

func main() {

	factory := NewIconFactory()

	bookmarks := []*BookMark{
		NewBookmark(
			"Google Search",
			"https://google.com/search",
			factory.GetIcon("google.com"),
		),

		NewBookmark(
			"Google Maps",
			"https://maps.google.com",
			factory.GetIcon("google.com"),
		),

		NewBookmark(
			"Google Docs",
			"https://docs.google.com",
			factory.GetIcon("google.com"),
		),

		NewBookmark(
			"GitHub",
			"https://github.com",
			factory.GetIcon("github.com"),
		),

		NewBookmark(
			"GitHub Issues",
			"https://github.com/issues",
			factory.GetIcon("github.com"),
		),
	}


	for _, bookmark := range bookmarks {
		bookmark.Show()
	}
}

