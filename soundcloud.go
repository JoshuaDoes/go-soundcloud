package soundcloud

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Client struct {
	ClientID string
}

type Track struct {
	Artist string
	ArtURL string
	Description string
	DownloadURL string
	Title string
	Track string
}

func (c *Client) GetTrackInfo(url string) (*Track, error) {
	request, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(request.Body)
	html := buffer.String()

	regex := regexp.MustCompile("soundcloud:tracks:([0-9]+)")
	track := regex.FindStringSubmatch(html)
	if len(track) <= 0 {
		return nil, errors.New("Error finding track ID")
	}
	regex = regexp.MustCompile("[\"']secret_token[\"'] *: *[\"'](.-)[\"']")
	secret := regex.FindStringSubmatch(html)
	//regex = regexp.MustCompile("[\"']title[\"'] *: *\"(.-[^\\])\"")
	regex = regexp.MustCompile("(?s)\"title\":\"(.*?)\"")
	title := regex.FindStringSubmatch(html)
	if len(title) <= 0 {
		return nil, errors.New("Error finding title")
	}
	regex = regexp.MustCompile("(?s)\"description\":\"(.*?)\"")
	description := regex.FindStringSubmatch(html)
	if len(description) <= 0 {
		return nil, errors.New("Error finding description")
	}
	regex = regexp.MustCompile("(?s)\"username\":\"(.*?)\"")
	artist := regex.FindStringSubmatch(html)
	if len(artist) <= 0 {
		return nil, errors.New("Error finding artist")
	}
	regex = regexp.MustCompile("(?s)\"artwork_url\":\"(.*?)\"")
	artURL := regex.FindStringSubmatch(html)
	if len(artURL) <= 0 {
		return nil, errors.New("Error finding artwork URL")
	}

	regex = regexp.MustCompile("https://api-v2\\.soundcloud\\.com/media/soundcloud:tracks:(?:\\d*)/([[:alnum:]-]*)/stream/progressive")
	identifier := regex.FindStringSubmatch(html)
	if len(identifier) <= 0 {
		return nil, errors.New("Error finding identifier")
	}

	url = fmt.Sprintf("https://api-v2.soundcloud.com/media/soundcloud:tracks:%s/%s/stream/progressive?client_id=%s", track[1], identifier[1], c.ClientID)
	if len(secret) > 0 {
		url += "&secret_token=" + secret[1]
	}

	data := &Track{}
	result, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	buffer = new(bytes.Buffer)
	buffer.ReadFrom(result.Body)
	downloadData := buffer.String()

	regex = regexp.MustCompile("(?s)\"url\":\"(.*?)\"")
	mp3128URL := regex.FindStringSubmatch(downloadData)
	if len(mp3128URL) <= 0 {
		return nil, errors.New("Error finding download URL")
	}
	downloadURL := strings.Replace(mp3128URL[1], "\\u0026", "&", -1)

	data.Artist = artist[1]
	data.ArtURL = artURL[1]
	data.Description = description[1]
	data.DownloadURL = downloadURL
	data.Title = title[1]
	data.Track = track[1]

	return data, err
}
