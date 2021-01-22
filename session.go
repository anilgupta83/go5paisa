package go5paisa

import (
	"errors"
	"net/url"
)

//GetCookie returns 5paisacookie value
func (c *Client) GetCookie() (string, error) {
	if c.connection == nil {
		return "", errors.New("Invalid connection state")
	}
	u, _ := url.ParseRequestURI(baseURL)
	cookies := c.connection.Jar.Cookies(u)
	for _, c := range cookies {
		if c.Name == "5paisacookie" {
			return c.Value, nil
		}
	}
	return "", errors.New("5paisacookie not found")
}
