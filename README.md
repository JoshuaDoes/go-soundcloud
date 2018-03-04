# go-soundcloud
Single file library for accessing SoundCloud metadata and audio URLs in Golang without extra dependencies

# Installing
`go get github.com/JoshuaDoes/go-soundcloud`

# Example
```go
package main

import "fmt"
import "github.com/JoshuaDoes/go-soundcloud"

func main() {
	//Initialize a new client
	c := &soundcloud.Client{ClientID:"your client id here", AppVersion:"your app version here"}

	//Get metadata and an audio URL
	res, err := c.GetTrackInfo("https://soundcloud.com/artist/title")

	if err != nil {
		panic(err)
	}
	
	fmt.Println("Track ID: " + res.Track)
	fmt.Println("Title: " + res.Title)
	fmt.Println("Artist: " + res.Artist)
	fmt.Println("Description: " + res.Description)
	fmt.Println("Download URL: " + res.DownloadURL)
	fmt.Println("Artwork URL: " + res.ArtURL)
}
```
### Output

```
> go run main.go

< Track ID
< Title
< Artist
< Description
< Download URL
< Artwork URL
```

## License
The source code for go-soundcloud is released under the MIT License. See LICENSE for more details.

## Donations
All donations are appreciated and helps me stay awake at night to work on this more. Even if it's not much, it helps a lot in the long run!
You can find the donation link here: [Donation Link](https://paypal.me/JoshuaDoes)
