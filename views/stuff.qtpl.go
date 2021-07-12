// Code generated by qtc from "stuff.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/stuff.qtpl:1
package views

//line views/stuff.qtpl:1
import "path/filepath"

//line views/stuff.qtpl:2
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/stuff.qtpl:3
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/stuff.qtpl:4
import "github.com/bouncepaw/mycorrhiza/user"

//line views/stuff.qtpl:5
import "github.com/bouncepaw/mycorrhiza/util"

//line views/stuff.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/stuff.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/stuff.qtpl:7
func StreamBaseHTML(qw422016 *qt422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:7
	qw422016.N().S(`
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>`)
//line views/stuff.qtpl:13
	qw422016.E().S(title)
//line views/stuff.qtpl:13
	qw422016.N().S(`</title>
		<link rel="shortcut icon" href="/static/favicon.ico">
		<link rel="stylesheet" href="/static/style.css">
		<script src="/static/shortcuts.js"></script>
		`)
//line views/stuff.qtpl:17
	for _, el := range headElements {
//line views/stuff.qtpl:17
		qw422016.N().S(el)
//line views/stuff.qtpl:17
	}
//line views/stuff.qtpl:17
	qw422016.N().S(`
	</head>
	<body>
		<header>
			<nav class="header-links main-width">
				<ul class="header-links__list">
`)
//line views/stuff.qtpl:23
	for _, link := range cfg.HeaderLinks {
//line views/stuff.qtpl:23
		qw422016.N().S(`					<li class="header-links__entry"><a class="header-links__link" href="`)
//line views/stuff.qtpl:24
		qw422016.E().S(link.Href)
//line views/stuff.qtpl:24
		qw422016.N().S(`">`)
//line views/stuff.qtpl:24
		qw422016.E().S(link.Display)
//line views/stuff.qtpl:24
		qw422016.N().S(`</a></li>
`)
//line views/stuff.qtpl:25
	}
//line views/stuff.qtpl:25
	qw422016.N().S(`					<li class="header-links__entry header-links__entry_search">
						<form class="header-links__search" method="GET" action="/title-search">
							<input type="text" name="q" placeholder="Title search" class="header-links__search-bar">
						</form>
					</li>
					`)
//line views/stuff.qtpl:31
	qw422016.N().S(UserMenuHTML(u))
//line views/stuff.qtpl:31
	qw422016.N().S(`
				</ul>
			</nav>
		</header>
		`)
//line views/stuff.qtpl:35
	qw422016.N().S(body)
//line views/stuff.qtpl:35
	qw422016.N().S(`
		<template id="dialog-template">
			<div class="dialog-backdrop"></div>
			<div class="dialog" tabindex="0">
				<div class="dialog__header">
					<h1 class="dialog__title"></h1>
					<button class="dialog__close-button" aria-label="Close this dialog"></button>
				</div>

				<div class="dialog__content"></div>
			</div>
		</template>
		`)
//line views/stuff.qtpl:47
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:47
	qw422016.N().S(`
	</body>
</html>
`)
//line views/stuff.qtpl:50
}

//line views/stuff.qtpl:50
func WriteBaseHTML(qq422016 qtio422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:50
	StreamBaseHTML(qw422016, title, body, u, headElements...)
//line views/stuff.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:50
}

//line views/stuff.qtpl:50
func BaseHTML(title, body string, u *user.User, headElements ...string) string {
//line views/stuff.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:50
	WriteBaseHTML(qb422016, title, body, u, headElements...)
//line views/stuff.qtpl:50
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:50
	return qs422016
//line views/stuff.qtpl:50
}

//line views/stuff.qtpl:52
func StreamTitleSearchHTML(qw422016 *qt422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:52
	qw422016.N().S(`
<div class="layout">
<main class="main-width title-search">
	<h1>Search results for ‘`)
//line views/stuff.qtpl:55
	qw422016.E().S(query)
//line views/stuff.qtpl:55
	qw422016.N().S(`’</h1>
	<p>Every hypha name has been compared with the query. Hyphae that have matched the query are listed below.</p>
	<ul class="title-search__results">
	`)
//line views/stuff.qtpl:58
	for hyphaName := range generator(query) {
//line views/stuff.qtpl:58
		qw422016.N().S(`
		<li class="title-search__entry">
			<a class="title-search__link wikilink" href="/hypha/`)
//line views/stuff.qtpl:60
		qw422016.E().S(hyphaName)
//line views/stuff.qtpl:60
		qw422016.N().S(`">`)
//line views/stuff.qtpl:60
		qw422016.E().S(util.BeautifulName(hyphaName))
//line views/stuff.qtpl:60
		qw422016.N().S(`</a>
		</li>
	`)
//line views/stuff.qtpl:62
	}
//line views/stuff.qtpl:62
	qw422016.N().S(`
</main>
</div>
`)
//line views/stuff.qtpl:65
}

//line views/stuff.qtpl:65
func WriteTitleSearchHTML(qq422016 qtio422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:65
	StreamTitleSearchHTML(qw422016, query, generator)
//line views/stuff.qtpl:65
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:65
}

//line views/stuff.qtpl:65
func TitleSearchHTML(query string, generator func(string) <-chan string) string {
//line views/stuff.qtpl:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:65
	WriteTitleSearchHTML(qb422016, query, generator)
//line views/stuff.qtpl:65
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:65
	return qs422016
//line views/stuff.qtpl:65
}

//line views/stuff.qtpl:67
func StreamHelpHTML(qw422016 *qt422016.Writer, content string) {
//line views/stuff.qtpl:67
	qw422016.N().S(`
<div class="layout">
<main class="main-width help">
	<article>
	`)
//line views/stuff.qtpl:71
	qw422016.N().S(content)
//line views/stuff.qtpl:71
	qw422016.N().S(`
	</article>
</main>
`)
//line views/stuff.qtpl:74
	qw422016.N().S(helpTopicsHTML())
//line views/stuff.qtpl:74
	qw422016.N().S(`
</div>
`)
//line views/stuff.qtpl:76
}

//line views/stuff.qtpl:76
func WriteHelpHTML(qq422016 qtio422016.Writer, content string) {
//line views/stuff.qtpl:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:76
	StreamHelpHTML(qw422016, content)
//line views/stuff.qtpl:76
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:76
}

//line views/stuff.qtpl:76
func HelpHTML(content string) string {
//line views/stuff.qtpl:76
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:76
	WriteHelpHTML(qb422016, content)
//line views/stuff.qtpl:76
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:76
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:76
	return qs422016
//line views/stuff.qtpl:76
}

//line views/stuff.qtpl:78
func StreamHelpEmptyErrorHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:78
	qw422016.N().S(`
<h1>This entry does not exist!</h1>
<p>Try finding a different entry that would help you.</p>
<p>If you want to write this entry by yourself, consider <a class="wikilink wikilink_external wikilink_https" href="https://github.com/bouncepaw/mycorrhiza">contributing</a> to Mycorrhiza Wiki directly.</p>
`)
//line views/stuff.qtpl:82
}

//line views/stuff.qtpl:82
func WriteHelpEmptyErrorHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:82
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:82
	StreamHelpEmptyErrorHTML(qw422016)
//line views/stuff.qtpl:82
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:82
}

//line views/stuff.qtpl:82
func HelpEmptyErrorHTML() string {
//line views/stuff.qtpl:82
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:82
	WriteHelpEmptyErrorHTML(qb422016)
//line views/stuff.qtpl:82
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:82
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:82
	return qs422016
//line views/stuff.qtpl:82
}

//line views/stuff.qtpl:84
func streamhelpTopicsHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:84
	qw422016.N().S(`
<aside class="help-topics layout-card">
	<h2 class="layout-card__title">Help topics</h2>
	<ul class="help-topics__list">
		<li>
			<a href="/help/en">Main</a>
		</li>
		<li>
			<a href="/help/en/hypha">Hypha</a>
			<ul>
				<li>
					<a href="/help/en/attachment">Attachment</a>
				</li>
			</ul
		</li>
	</ul>
</aside>
`)
//line views/stuff.qtpl:101
}

//line views/stuff.qtpl:101
func writehelpTopicsHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:101
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:101
	streamhelpTopicsHTML(qw422016)
//line views/stuff.qtpl:101
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:101
}

//line views/stuff.qtpl:101
func helpTopicsHTML() string {
//line views/stuff.qtpl:101
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:101
	writehelpTopicsHTML(qb422016)
//line views/stuff.qtpl:101
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:101
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:101
	return qs422016
//line views/stuff.qtpl:101
}

//line views/stuff.qtpl:103
func StreamUserListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:103
	qw422016.N().S(`
<div class="layout">
<main class="main-width user-list">
	<h1>List of users</h1>
`)
//line views/stuff.qtpl:108
	var (
		admins     = make([]string, 0)
		moderators = make([]string, 0)
		editors    = make([]string, 0)
	)
	for u := range user.YieldUsers() {
		switch u.Group {
		case "admin":
			admins = append(admins, u.Name)
		case "moderator":
			moderators = append(moderators, u.Name)
		case "editor", "trusted":
			editors = append(editors, u.Name)
		}
	}

//line views/stuff.qtpl:123
	qw422016.N().S(`
	<section>
		<h2>Admins</h2>
		<ol>`)
//line views/stuff.qtpl:126
	for _, name := range admins {
//line views/stuff.qtpl:126
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:127
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:127
		qw422016.N().S(`/`)
//line views/stuff.qtpl:127
		qw422016.E().S(name)
//line views/stuff.qtpl:127
		qw422016.N().S(`">`)
//line views/stuff.qtpl:127
		qw422016.E().S(name)
//line views/stuff.qtpl:127
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:128
	}
//line views/stuff.qtpl:128
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Moderators</h2>
		<ol>`)
//line views/stuff.qtpl:132
	for _, name := range moderators {
//line views/stuff.qtpl:132
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:133
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:133
		qw422016.N().S(`/`)
//line views/stuff.qtpl:133
		qw422016.E().S(name)
//line views/stuff.qtpl:133
		qw422016.N().S(`">`)
//line views/stuff.qtpl:133
		qw422016.E().S(name)
//line views/stuff.qtpl:133
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:134
	}
//line views/stuff.qtpl:134
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Editors</h2>
		<ol>`)
//line views/stuff.qtpl:138
	for _, name := range editors {
//line views/stuff.qtpl:138
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:139
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:139
		qw422016.N().S(`/`)
//line views/stuff.qtpl:139
		qw422016.E().S(name)
//line views/stuff.qtpl:139
		qw422016.N().S(`">`)
//line views/stuff.qtpl:139
		qw422016.E().S(name)
//line views/stuff.qtpl:139
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:140
	}
//line views/stuff.qtpl:140
	qw422016.N().S(`</ol>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:144
}

//line views/stuff.qtpl:144
func WriteUserListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:144
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:144
	StreamUserListHTML(qw422016)
//line views/stuff.qtpl:144
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:144
}

//line views/stuff.qtpl:144
func UserListHTML() string {
//line views/stuff.qtpl:144
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:144
	WriteUserListHTML(qb422016)
//line views/stuff.qtpl:144
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:144
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:144
	return qs422016
//line views/stuff.qtpl:144
}

//line views/stuff.qtpl:146
func StreamHyphaListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:146
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>List of hyphae</h1>
	<p>This wiki has `)
//line views/stuff.qtpl:150
	qw422016.N().D(hyphae.Count())
//line views/stuff.qtpl:150
	qw422016.N().S(` hyphae.</p>
	<ul class="hypha-list">
		`)
//line views/stuff.qtpl:152
	for h := range hyphae.YieldExistingHyphae() {
//line views/stuff.qtpl:152
		qw422016.N().S(`
		<li class="hypha-list__entry">
			<a class="hypha-list__link" href="/hypha/`)
//line views/stuff.qtpl:154
		qw422016.E().S(h.Name)
//line views/stuff.qtpl:154
		qw422016.N().S(`">`)
//line views/stuff.qtpl:154
		qw422016.E().S(util.BeautifulName(h.Name))
//line views/stuff.qtpl:154
		qw422016.N().S(`</a>
			`)
//line views/stuff.qtpl:155
		if h.BinaryPath != "" {
//line views/stuff.qtpl:155
			qw422016.N().S(`
			<span class="hypha-list__amnt-type">`)
//line views/stuff.qtpl:156
			qw422016.E().S(filepath.Ext(h.BinaryPath)[1:])
//line views/stuff.qtpl:156
			qw422016.N().S(`</span>
			`)
//line views/stuff.qtpl:157
		}
//line views/stuff.qtpl:157
		qw422016.N().S(`
		</li>
		`)
//line views/stuff.qtpl:159
	}
//line views/stuff.qtpl:159
	qw422016.N().S(`
	</ul>
</main>
</div>
`)
//line views/stuff.qtpl:163
}

//line views/stuff.qtpl:163
func WriteHyphaListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:163
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:163
	StreamHyphaListHTML(qw422016)
//line views/stuff.qtpl:163
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:163
}

//line views/stuff.qtpl:163
func HyphaListHTML() string {
//line views/stuff.qtpl:163
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:163
	WriteHyphaListHTML(qb422016)
//line views/stuff.qtpl:163
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:163
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:163
	return qs422016
//line views/stuff.qtpl:163
}

//line views/stuff.qtpl:165
func StreamAboutHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:165
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section>
		<h1>About `)
//line views/stuff.qtpl:169
	qw422016.E().S(cfg.WikiName)
//line views/stuff.qtpl:169
	qw422016.N().S(`</h1>
		<ul>
			<li><b><a href="https://mycorrhiza.wiki">Mycorrhiza Wiki</a> version:</b> 1.3.0</li>
`)
//line views/stuff.qtpl:172
	if cfg.UseAuth {
//line views/stuff.qtpl:172
		qw422016.N().S(`			<li><b>User count:</b> `)
//line views/stuff.qtpl:173
		qw422016.N().DUL(user.Count())
//line views/stuff.qtpl:173
		qw422016.N().S(`</li>
			<li><b>Home page:</b> <a href="/">`)
//line views/stuff.qtpl:174
		qw422016.E().S(cfg.HomeHypha)
//line views/stuff.qtpl:174
		qw422016.N().S(`</a></li>
			<li><b>Administrators:</b>`)
//line views/stuff.qtpl:175
		for i, username := range user.ListUsersWithGroup("admin") {
//line views/stuff.qtpl:176
			if i > 0 {
//line views/stuff.qtpl:176
				qw422016.N().S(`<span aria-hidden="true">, </span>
`)
//line views/stuff.qtpl:177
			}
//line views/stuff.qtpl:177
			qw422016.N().S(`				<a href="/hypha/`)
//line views/stuff.qtpl:178
			qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:178
			qw422016.N().S(`/`)
//line views/stuff.qtpl:178
			qw422016.E().S(username)
//line views/stuff.qtpl:178
			qw422016.N().S(`">`)
//line views/stuff.qtpl:178
			qw422016.E().S(username)
//line views/stuff.qtpl:178
			qw422016.N().S(`</a>`)
//line views/stuff.qtpl:178
		}
//line views/stuff.qtpl:178
		qw422016.N().S(`</li>
`)
//line views/stuff.qtpl:179
	} else {
//line views/stuff.qtpl:179
		qw422016.N().S(`			<li>This wiki does not use authorization</li>
`)
//line views/stuff.qtpl:181
	}
//line views/stuff.qtpl:181
	qw422016.N().S(`		</ul>
		<p>See <a href="/list">/list</a> for information about hyphae on this wiki.</p>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:187
}

//line views/stuff.qtpl:187
func WriteAboutHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:187
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:187
	StreamAboutHTML(qw422016)
//line views/stuff.qtpl:187
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:187
}

//line views/stuff.qtpl:187
func AboutHTML() string {
//line views/stuff.qtpl:187
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:187
	WriteAboutHTML(qb422016)
//line views/stuff.qtpl:187
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:187
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:187
	return qs422016
//line views/stuff.qtpl:187
}

//line views/stuff.qtpl:189
func StreamCommonScripts(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:189
	qw422016.N().S(`
`)
//line views/stuff.qtpl:190
	for _, scriptPath := range cfg.CommonScripts {
//line views/stuff.qtpl:190
		qw422016.N().S(`
<script src="`)
//line views/stuff.qtpl:191
		qw422016.E().S(scriptPath)
//line views/stuff.qtpl:191
		qw422016.N().S(`"></script>
`)
//line views/stuff.qtpl:192
	}
//line views/stuff.qtpl:192
	qw422016.N().S(`
`)
//line views/stuff.qtpl:193
}

//line views/stuff.qtpl:193
func WriteCommonScripts(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:193
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:193
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:193
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:193
}

//line views/stuff.qtpl:193
func CommonScripts() string {
//line views/stuff.qtpl:193
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:193
	WriteCommonScripts(qb422016)
//line views/stuff.qtpl:193
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:193
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:193
	return qs422016
//line views/stuff.qtpl:193
}
