package endpoint

import "net/url"

type Filter struct {
	id    string
	alias string
}

func NewFilter(idParam string, urlValues url.Values) Filter {
	var alias string

	if len(urlValues["alias"]) > 0 {
		alias = urlValues["alias"][0]
	}

	return Filter{
		id:    idParam,
		alias: alias,
	}
}

type Post struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}
