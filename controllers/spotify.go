package controllers

import (
	"net/url"
	"fmt"
	"github.com/boratanrikulu/s-lyrics/models"
	"html/template"
	"net/http"
)

// Public Methods

func WrongGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/wrong.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		panic("Something is wrong")
	}
}

func SpotifyGet(w http.ResponseWriter, r *http.Request) {
	// Creates a spotify model with it's secrets.
	spotify := new(models.Spotify)
	spotify.InitSecrets()

	// Gets result for RefreshAndAccessTokes request.
	spotify.Authorization.Response.Code = r.URL.Query().Get("code")
	// TODO: Check if there is a code value. (that means user is login.)
	// Result is in spotify.ResponseRefreshAndAccessTokens
	err := spotify.GetRefreshAndAccessTokensResponse()
	if err != nil {
		http.Redirect(w, r, "/wrong", http.StatusMovedPermanently)
	}

	// Gets current song.
	artistName, songName, err := spotify.GetCurrentlyPlaying()
	if err != nil {
		http.Redirect(w, r, "/wrong", http.StatusMovedPermanently)
	}

	// Redirects to lyrics page.
	u, _ := url.Parse("/lyric")
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("artistName", artistName)
	q.Add("songName", songName)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, fmt.Sprint(u), http.StatusFound)
}
