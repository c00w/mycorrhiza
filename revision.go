package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gomarkdown/markdown"
)

// Revision represents a revision, duh.
// A revision is a version of a hypha at a point in time.
type Revision struct {
	Id         int      `json:"-"`
	FullName   string   `json:"-"`
	Tags       []string `json:"tags"`
	ShortName  string   `json:"name"`
	Comment    string   `json:"comment"`
	Author     string   `json:"author"`
	Time       int      `json:"time"`
	TextMime   string   `json:"text_mime"`
	BinaryMime string   `json:"binary_mime"`
	TextPath   string   `json:"-"`
	BinaryPath string   `json:"-"`
}

// IdAsStr returns revision's id as a string.
func (rev *Revision) IdAsStr() string {
	return strconv.Itoa(rev.Id)
}

// hasBinaryData returns true if the revision has any binary data associated.
// During initialisation, it is guaranteed that r.BinaryMime is set to "" if the revision has no binary data.
func (rev *Revision) hasBinaryData() bool {
	return rev.BinaryMime != ""
}

// AsHtml returns HTML representation of the revision.
// If there is an error, it will be told about it in `w`.
// In any case, some http data is written to `w`.
func (rev *Revision) AsHtml(w http.ResponseWriter) (ret string, err error) {
	ret += `<article class="page">
	<h1 class="page__title">` + rev.FullName + `</h1>
`
	// TODO: support things other than images
	if rev.hasBinaryData() {
		ret += fmt.Sprintf(`<img src="/%s?action=getBinary&rev=%d" class="page__amnt"/>`, rev.FullName, rev.Id)
	}

	contents, err := ioutil.ReadFile(rev.TextPath)
	if err != nil {
		log.Println("Failed to render", rev.FullName)
		w.WriteHeader(http.StatusInternalServerError)
		return "", err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// TODO: support more markups.
	// TODO: support mycorrhiza extensions like transclusion.
	switch rev.TextMime {
	case "text/markdown":
		html := markdown.ToHTML(contents, nil, nil)
		ret += string(html)
	default:
		ret += fmt.Sprintf(`<pre>%s</pre>`, contents)
	}

	ret += `
</article>`
	return ret, nil
}

// ActionGetBinary is used with `?action=getBinary`.
// It writes binary data of the revision. It also sets the MIME-type.
func (rev *Revision) ActionGetBinary(w http.ResponseWriter) {
	fileContents, err := ioutil.ReadFile(rev.BinaryPath)
	if err != nil {
		log.Println("Failed to load binary data of", rev.FullName, rev.Id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", rev.BinaryMime)
	w.WriteHeader(http.StatusOK)
	w.Write(fileContents)
	log.Println("Serving binary data of", rev.FullName, rev.Id)
}

// ActionRaw is used with `?action=raw`.
// It writes text content of the revision without any parsing or rendering.
func (rev *Revision) ActionRaw(w http.ResponseWriter) {
	fileContents, err := ioutil.ReadFile(rev.TextPath)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", rev.TextMime)
	w.WriteHeader(http.StatusOK)
	w.Write(fileContents)
	log.Println("Serving text data of", rev.FullName, rev.Id)
}

// ActionZen is used with `?action=zen`.
// It renders the revision without any layout or navigation.
func (rev *Revision) ActionZen(w http.ResponseWriter) {
	html, err := rev.AsHtml(w)
	if err == nil {
		fmt.Fprint(w, html)
		log.Println("Rendering", rev.FullName, "in zen mode")
	}
}

// ActionView is used with `?action=view` or without any action.
// It renders the revision with layout and navigation.
func (rev *Revision) ActionView(w http.ResponseWriter, layoutFun func(Revision, string) string) {
	html, err := rev.AsHtml(w)
	if err == nil {
		fmt.Fprint(w, layoutFun(*rev, html))
		log.Println("Rendering", rev.FullName)
	}
}
