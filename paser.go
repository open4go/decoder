package decoder

import "net/url"

func PaserUrl(query string) url.Values {
	m, err := url.ParseQuery(query)
	if err != nil {
		return nil
	}
	return m
}
