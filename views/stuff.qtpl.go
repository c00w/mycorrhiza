// Code generated by qtc from "stuff.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/stuff.qtpl:1
package views

//line views/stuff.qtpl:1
import "path/filepath"

//line views/stuff.qtpl:2
import "sort"

//line views/stuff.qtpl:3
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/stuff.qtpl:4
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/stuff.qtpl:5
import "github.com/bouncepaw/mycorrhiza/user"

//line views/stuff.qtpl:6
import "github.com/bouncepaw/mycorrhiza/util"

//line views/stuff.qtpl:7
import "github.com/bouncepaw/mycorrhiza/l18n"

//line views/stuff.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/stuff.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/stuff.qtpl:9
func StreamBaseHTML(qw422016 *qt422016.Writer, title, body string, lc *l18n.Localizer, u *user.User, headElements ...string) {
//line views/stuff.qtpl:9
	qw422016.N().S(`
<!doctype html>
<html lang="`)
//line views/stuff.qtpl:11
	qw422016.E().S(lc.Locale)
//line views/stuff.qtpl:11
	qw422016.N().S(`">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>`)
//line views/stuff.qtpl:15
	qw422016.E().S(title)
//line views/stuff.qtpl:15
	qw422016.N().S(`</title>
		<link rel="shortcut icon" href="/static/favicon.ico">
		<link rel="stylesheet" href="/static/style.css">
		<script src="/static/shortcuts.js"></script>
		`)
//line views/stuff.qtpl:19
	for _, el := range headElements {
//line views/stuff.qtpl:19
		qw422016.N().S(el)
//line views/stuff.qtpl:19
	}
//line views/stuff.qtpl:19
	qw422016.N().S(`
	</head>
	<body>
		<header>
			<nav class="main-width top-bar">
				<ul class="top-bar__wrapper">
					<li class="top-bar__section top-bar__section_home">
						<div class="top-bar__home-link-wrapper">
							<a class="top-bar__home-link" href="/">`)
//line views/stuff.qtpl:27
	qw422016.E().S(cfg.WikiName)
//line views/stuff.qtpl:27
	qw422016.N().S(`</a>
						</div>
					</li>
					<li class="top-bar__section top-bar__section_search">
						<form class="top-bar__search" method="GET" action="/title-search">
							<input type="text" name="q" placeholder="`)
//line views/stuff.qtpl:32
	qw422016.E().S(lc.Get("ui.title_search"))
//line views/stuff.qtpl:32
	qw422016.N().S(`" class="top-bar__search-bar">
						</form>
					</li>
					<li class="top-bar__section top-bar__section_auth">
					`)
//line views/stuff.qtpl:36
	if cfg.UseAuth {
//line views/stuff.qtpl:36
		qw422016.N().S(`
						<ul class="top-bar__auth auth-links">
							<li class="auth-links__box auth-links__user-box">
								`)
//line views/stuff.qtpl:39
		if u.Group == "anon" {
//line views/stuff.qtpl:39
			qw422016.N().S(`
								<a href="/login" class="auth-links__link auth-links__login-link">`)
//line views/stuff.qtpl:40
			qw422016.E().S(lc.Get("ui.login"))
//line views/stuff.qtpl:40
			qw422016.N().S(`</a>
								`)
//line views/stuff.qtpl:41
		} else {
//line views/stuff.qtpl:41
			qw422016.N().S(`
								<a href="/hypha/`)
//line views/stuff.qtpl:42
			qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:42
			qw422016.N().S(`/`)
//line views/stuff.qtpl:42
			qw422016.E().S(u.Name)
//line views/stuff.qtpl:42
			qw422016.N().S(`" class="auth-links__link auth-links__user-link">`)
//line views/stuff.qtpl:42
			qw422016.E().S(util.BeautifulName(u.Name))
//line views/stuff.qtpl:42
			qw422016.N().S(`</a>
								`)
//line views/stuff.qtpl:43
		}
//line views/stuff.qtpl:43
		qw422016.N().S(`
							</li>
							`)
//line views/stuff.qtpl:45
		if cfg.AllowRegistration && u.Group == "anon" {
//line views/stuff.qtpl:45
			qw422016.N().S(`
							<li class="auth-links__box auth-links__register-box">
								<a href="/register" class="auth-links__link auth-links__register-link">`)
//line views/stuff.qtpl:47
			qw422016.E().S(lc.Get("ui.register"))
//line views/stuff.qtpl:47
			qw422016.N().S(`</a>
							</li>
							`)
//line views/stuff.qtpl:49
		}
//line views/stuff.qtpl:49
		qw422016.N().S(`
							`)
//line views/stuff.qtpl:50
		if u.Group == "admin" {
//line views/stuff.qtpl:50
			qw422016.N().S(`
							<li class="auth-links__box auth-links__admin-box">
								<a href="/admin" class="auth-links__link auth-links__admin-link">`)
//line views/stuff.qtpl:52
			qw422016.E().S(lc.Get("ui.admin_panel"))
//line views/stuff.qtpl:52
			qw422016.N().S(`</a>
							</li>
							`)
//line views/stuff.qtpl:54
		}
//line views/stuff.qtpl:54
		qw422016.N().S(`
						</ul>
					`)
//line views/stuff.qtpl:56
	}
//line views/stuff.qtpl:56
	qw422016.N().S(`
					</li>
					<li class="top-bar__section top-bar__section_highlights">
						<ul class="top-bar__highlights">
`)
//line views/stuff.qtpl:60
	for _, link := range cfg.HeaderLinks {
//line views/stuff.qtpl:60
		qw422016.N().S(`						`)
//line views/stuff.qtpl:61
		if link.Href != "/" {
//line views/stuff.qtpl:61
			qw422016.N().S(`
							<li class="top-bar__highlight">
								<a class="top-bar__highlight-link" href="`)
//line views/stuff.qtpl:63
			qw422016.E().S(link.Href)
//line views/stuff.qtpl:63
			qw422016.N().S(`">`)
//line views/stuff.qtpl:63
			qw422016.E().S(link.Display)
//line views/stuff.qtpl:63
			qw422016.N().S(`</a>
							</li>
						`)
//line views/stuff.qtpl:65
		}
//line views/stuff.qtpl:65
		qw422016.N().S(`
`)
//line views/stuff.qtpl:66
	}
//line views/stuff.qtpl:66
	qw422016.N().S(`						</ul>
					</li>
				</ul>
			</nav>
		</header>
		`)
//line views/stuff.qtpl:72
	qw422016.N().S(body)
//line views/stuff.qtpl:72
	qw422016.N().S(`
		<template id="dialog-template">
			<div class="dialog-backdrop"></div>
			<div class="dialog" tabindex="0">
				<div class="dialog__header">
					<h1 class="dialog__title"></h1>
					<button class="dialog__close-button" aria-label="`)
//line views/stuff.qtpl:78
	qw422016.E().S(lc.Get("ui.close_dialog"))
//line views/stuff.qtpl:78
	qw422016.N().S(`"></button>
				</div>

				<div class="dialog__content"></div>
			</div>
		</template>
		`)
//line views/stuff.qtpl:84
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:84
	qw422016.N().S(`
		<script src="/static/view.js"></script>
	</body>
</html>
`)
//line views/stuff.qtpl:88
}

//line views/stuff.qtpl:88
func WriteBaseHTML(qq422016 qtio422016.Writer, title, body string, lc *l18n.Localizer, u *user.User, headElements ...string) {
//line views/stuff.qtpl:88
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:88
	StreamBaseHTML(qw422016, title, body, lc, u, headElements...)
//line views/stuff.qtpl:88
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:88
}

//line views/stuff.qtpl:88
func BaseHTML(title, body string, lc *l18n.Localizer, u *user.User, headElements ...string) string {
//line views/stuff.qtpl:88
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:88
	WriteBaseHTML(qb422016, title, body, lc, u, headElements...)
//line views/stuff.qtpl:88
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:88
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:88
	return qs422016
//line views/stuff.qtpl:88
}

//line views/stuff.qtpl:90
func StreamTitleSearchHTML(qw422016 *qt422016.Writer, query string, generator func(string) <-chan string, lc *l18n.Localizer) {
//line views/stuff.qtpl:90
	qw422016.N().S(`
<div class="layout">
<main class="main-width title-search">
	<h1>`)
//line views/stuff.qtpl:93
	qw422016.E().S(lc.Get("ui.search_results_query", &l18n.Replacements{"query": query}))
//line views/stuff.qtpl:93
	qw422016.N().S(`</h1>
	<p>`)
//line views/stuff.qtpl:94
	qw422016.E().S(lc.Get("ui.search_results_desc"))
//line views/stuff.qtpl:94
	qw422016.N().S(`</p>
	<ul class="title-search__results">
	`)
//line views/stuff.qtpl:96
	for hyphaName := range generator(query) {
//line views/stuff.qtpl:96
		qw422016.N().S(`
		<li class="title-search__entry">
			<a class="title-search__link wikilink" href="/hypha/`)
//line views/stuff.qtpl:98
		qw422016.E().S(hyphaName)
//line views/stuff.qtpl:98
		qw422016.N().S(`">`)
//line views/stuff.qtpl:98
		qw422016.E().S(util.BeautifulName(hyphaName))
//line views/stuff.qtpl:98
		qw422016.N().S(`</a>
		</li>
	`)
//line views/stuff.qtpl:100
	}
//line views/stuff.qtpl:100
	qw422016.N().S(`
</main>
</div>
`)
//line views/stuff.qtpl:103
}

//line views/stuff.qtpl:103
func WriteTitleSearchHTML(qq422016 qtio422016.Writer, query string, generator func(string) <-chan string, lc *l18n.Localizer) {
//line views/stuff.qtpl:103
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:103
	StreamTitleSearchHTML(qw422016, query, generator, lc)
//line views/stuff.qtpl:103
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:103
}

//line views/stuff.qtpl:103
func TitleSearchHTML(query string, generator func(string) <-chan string, lc *l18n.Localizer) string {
//line views/stuff.qtpl:103
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:103
	WriteTitleSearchHTML(qb422016, query, generator, lc)
//line views/stuff.qtpl:103
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:103
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:103
	return qs422016
//line views/stuff.qtpl:103
}

// It outputs a poorly formatted JSON, but it works and is valid.

//line views/stuff.qtpl:106
func StreamTitleSearchJSON(qw422016 *qt422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:106
	qw422016.N().S(`
`)
//line views/stuff.qtpl:108
	// Lol
	counter := 0

//line views/stuff.qtpl:110
	qw422016.N().S(`
{
	"source_query": "`)
//line views/stuff.qtpl:112
	qw422016.E().S(query)
//line views/stuff.qtpl:112
	qw422016.N().S(`",
	"results": [
	`)
//line views/stuff.qtpl:114
	for hyphaName := range generator(query) {
//line views/stuff.qtpl:114
		qw422016.N().S(`
		`)
//line views/stuff.qtpl:115
		if counter > 0 {
//line views/stuff.qtpl:115
			qw422016.N().S(`, `)
//line views/stuff.qtpl:115
		}
//line views/stuff.qtpl:115
		qw422016.N().S(`{
			"canonical_name": "`)
//line views/stuff.qtpl:116
		qw422016.E().S(hyphaName)
//line views/stuff.qtpl:116
		qw422016.N().S(`",
			"beautiful_name": "`)
//line views/stuff.qtpl:117
		qw422016.E().S(util.BeautifulName(hyphaName))
//line views/stuff.qtpl:117
		qw422016.N().S(`",
			"url": "`)
//line views/stuff.qtpl:118
		qw422016.E().S(cfg.URL + "/hypha/" + hyphaName)
//line views/stuff.qtpl:118
		qw422016.N().S(`"
		}`)
//line views/stuff.qtpl:119
		counter++

//line views/stuff.qtpl:119
		qw422016.N().S(`
	`)
//line views/stuff.qtpl:120
	}
//line views/stuff.qtpl:120
	qw422016.N().S(`
	]
}
`)
//line views/stuff.qtpl:123
}

//line views/stuff.qtpl:123
func WriteTitleSearchJSON(qq422016 qtio422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:123
	StreamTitleSearchJSON(qw422016, query, generator)
//line views/stuff.qtpl:123
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:123
}

//line views/stuff.qtpl:123
func TitleSearchJSON(query string, generator func(string) <-chan string) string {
//line views/stuff.qtpl:123
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:123
	WriteTitleSearchJSON(qb422016, query, generator)
//line views/stuff.qtpl:123
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:123
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:123
	return qs422016
//line views/stuff.qtpl:123
}

//line views/stuff.qtpl:125
func StreamBacklinksHTML(qw422016 *qt422016.Writer, query string, generator func(string) <-chan string, lc *l18n.Localizer) {
//line views/stuff.qtpl:125
	qw422016.N().S(`
<div class="layout">
<main class="main-width backlinks">
	<h1>`)
//line views/stuff.qtpl:128
	qw422016.E().S(lc.Get("ui.backlinks_query", &l18n.Replacements{"query": query}))
//line views/stuff.qtpl:128
	qw422016.N().S(`</h1>
	<p>`)
//line views/stuff.qtpl:129
	qw422016.E().S(lc.Get("ui.backlinks_desc"))
//line views/stuff.qtpl:129
	qw422016.N().S(`</p>
	<ul class="backlinks__list">
	`)
//line views/stuff.qtpl:131
	for hyphaName := range generator(query) {
//line views/stuff.qtpl:131
		qw422016.N().S(`
		<li class="backlinks__entry">
			<a class="backlinks__link wikilink" href="/hypha/`)
//line views/stuff.qtpl:133
		qw422016.E().S(hyphaName)
//line views/stuff.qtpl:133
		qw422016.N().S(`">`)
//line views/stuff.qtpl:133
		qw422016.E().S(util.BeautifulName(hyphaName))
//line views/stuff.qtpl:133
		qw422016.N().S(`</a>
		</li>
	`)
//line views/stuff.qtpl:135
	}
//line views/stuff.qtpl:135
	qw422016.N().S(`
</main>
</div>
`)
//line views/stuff.qtpl:138
}

//line views/stuff.qtpl:138
func WriteBacklinksHTML(qq422016 qtio422016.Writer, query string, generator func(string) <-chan string, lc *l18n.Localizer) {
//line views/stuff.qtpl:138
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:138
	StreamBacklinksHTML(qw422016, query, generator, lc)
//line views/stuff.qtpl:138
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:138
}

//line views/stuff.qtpl:138
func BacklinksHTML(query string, generator func(string) <-chan string, lc *l18n.Localizer) string {
//line views/stuff.qtpl:138
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:138
	WriteBacklinksHTML(qb422016, query, generator, lc)
//line views/stuff.qtpl:138
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:138
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:138
	return qs422016
//line views/stuff.qtpl:138
}

//line views/stuff.qtpl:140
func StreamHelpHTML(qw422016 *qt422016.Writer, content, lang string, lc *l18n.Localizer) {
//line views/stuff.qtpl:140
	qw422016.N().S(`
<div class="layout">
<main class="main-width help">
	<article>
	`)
//line views/stuff.qtpl:144
	qw422016.N().S(content)
//line views/stuff.qtpl:144
	qw422016.N().S(`
	</article>
</main>
`)
//line views/stuff.qtpl:147
	qw422016.N().S(helpTopicsHTML(lang, lc))
//line views/stuff.qtpl:147
	qw422016.N().S(`
</div>
`)
//line views/stuff.qtpl:149
}

//line views/stuff.qtpl:149
func WriteHelpHTML(qq422016 qtio422016.Writer, content, lang string, lc *l18n.Localizer) {
//line views/stuff.qtpl:149
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:149
	StreamHelpHTML(qw422016, content, lang, lc)
//line views/stuff.qtpl:149
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:149
}

//line views/stuff.qtpl:149
func HelpHTML(content, lang string, lc *l18n.Localizer) string {
//line views/stuff.qtpl:149
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:149
	WriteHelpHTML(qb422016, content, lang, lc)
//line views/stuff.qtpl:149
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:149
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:149
	return qs422016
//line views/stuff.qtpl:149
}

//line views/stuff.qtpl:151
func StreamHelpEmptyErrorHTML(qw422016 *qt422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:151
	qw422016.N().S(`
<h1>`)
//line views/stuff.qtpl:152
	qw422016.E().S(lc.Get("help.empty_error_title"))
//line views/stuff.qtpl:152
	qw422016.N().S(`</h1>
<p>`)
//line views/stuff.qtpl:153
	qw422016.E().S(lc.Get("help.empty_error_line_1"))
//line views/stuff.qtpl:153
	qw422016.N().S(`</p>
<p>`)
//line views/stuff.qtpl:154
	qw422016.E().S(lc.Get("help.empty_error_line_2a"))
//line views/stuff.qtpl:154
	qw422016.N().S(` <a class="wikilink wikilink_external wikilink_https" href="https://github.com/bouncepaw/mycorrhiza">`)
//line views/stuff.qtpl:154
	qw422016.E().S(lc.Get("help.empty_error_link"))
//line views/stuff.qtpl:154
	qw422016.N().S(`</a> `)
//line views/stuff.qtpl:154
	qw422016.E().S(lc.Get("help.empty_error_line_2b"))
//line views/stuff.qtpl:154
	qw422016.N().S(`</p>
`)
//line views/stuff.qtpl:155
}

//line views/stuff.qtpl:155
func WriteHelpEmptyErrorHTML(qq422016 qtio422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:155
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:155
	StreamHelpEmptyErrorHTML(qw422016, lc)
//line views/stuff.qtpl:155
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:155
}

//line views/stuff.qtpl:155
func HelpEmptyErrorHTML(lc *l18n.Localizer) string {
//line views/stuff.qtpl:155
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:155
	WriteHelpEmptyErrorHTML(qb422016, lc)
//line views/stuff.qtpl:155
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:155
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:155
	return qs422016
//line views/stuff.qtpl:155
}

//line views/stuff.qtpl:157
func streamhelpTopicsHTML(qw422016 *qt422016.Writer, lang string, lc *l18n.Localizer) {
//line views/stuff.qtpl:157
	qw422016.N().S(`
<aside class="help-topics layout-card">
	<h2 class="layout-card__title">`)
//line views/stuff.qtpl:159
	qw422016.E().S(lc.GetWithLocale(lang, "help.topics"))
//line views/stuff.qtpl:159
	qw422016.N().S(`</h2>
	<ul class="help-topics__list">
		<li><a href="/help/`)
//line views/stuff.qtpl:161
	qw422016.E().S(lang)
//line views/stuff.qtpl:161
	qw422016.N().S(`">`)
//line views/stuff.qtpl:161
	qw422016.E().S(lc.GetWithLocale(lang, "help.main"))
//line views/stuff.qtpl:161
	qw422016.N().S(`</a></li>
		<li><a href="/help/`)
//line views/stuff.qtpl:162
	qw422016.E().S(lang)
//line views/stuff.qtpl:162
	qw422016.N().S(`/hypha">`)
//line views/stuff.qtpl:162
	qw422016.E().S(lc.GetWithLocale(lang, "help.hypha"))
//line views/stuff.qtpl:162
	qw422016.N().S(`</a>
			<ul>
				<li><a href="/help/`)
//line views/stuff.qtpl:164
	qw422016.E().S(lang)
//line views/stuff.qtpl:164
	qw422016.N().S(`/attachment">`)
//line views/stuff.qtpl:164
	qw422016.E().S(lc.GetWithLocale(lang, "help.attachment"))
//line views/stuff.qtpl:164
	qw422016.N().S(`</a></li>
			</ul>
		</li>
		<li><a href="/help/`)
//line views/stuff.qtpl:167
	qw422016.E().S(lang)
//line views/stuff.qtpl:167
	qw422016.N().S(`/mycomarkup">`)
//line views/stuff.qtpl:167
	qw422016.E().S(lc.GetWithLocale(lang, "help.mycomarkup"))
//line views/stuff.qtpl:167
	qw422016.N().S(`</a></li>
		<li>`)
//line views/stuff.qtpl:168
	qw422016.E().S(lc.GetWithLocale(lang, "help.interface"))
//line views/stuff.qtpl:168
	qw422016.N().S(`
			<ul>
				<li><a href="/help/`)
//line views/stuff.qtpl:170
	qw422016.E().S(lang)
//line views/stuff.qtpl:170
	qw422016.N().S(`/prevnext">`)
//line views/stuff.qtpl:170
	qw422016.E().S(lc.GetWithLocale(lang, "help.prevnext"))
//line views/stuff.qtpl:170
	qw422016.N().S(`</a></li>
				<li><a href="/help/`)
//line views/stuff.qtpl:171
	qw422016.E().S(lang)
//line views/stuff.qtpl:171
	qw422016.N().S(`/top_bar">`)
//line views/stuff.qtpl:171
	qw422016.E().S(lc.GetWithLocale(lang, "help.top_bar"))
//line views/stuff.qtpl:171
	qw422016.N().S(`</a></li>
				<li><a href="/help/`)
//line views/stuff.qtpl:172
	qw422016.E().S(lang)
//line views/stuff.qtpl:172
	qw422016.N().S(`/sibling_hyphae_section">`)
//line views/stuff.qtpl:172
	qw422016.E().S(lc.GetWithLocale(lang, "help.sibling_hyphae"))
//line views/stuff.qtpl:172
	qw422016.N().S(`</a></li>
				<li>...</li>
			</ul>
		</li>
		<li>`)
//line views/stuff.qtpl:176
	qw422016.E().S(lc.GetWithLocale(lang, "help.special_pages"))
//line views/stuff.qtpl:176
	qw422016.N().S(`
			<ul>
				<li><a href="/help/`)
//line views/stuff.qtpl:178
	qw422016.E().S(lang)
//line views/stuff.qtpl:178
	qw422016.N().S(`/recent_changes">`)
//line views/stuff.qtpl:178
	qw422016.E().S(lc.GetWithLocale(lang, "help.recent_changes"))
//line views/stuff.qtpl:178
	qw422016.N().S(`</a></li>
				<li><a href="/help/`)
//line views/stuff.qtpl:179
	qw422016.E().S(lang)
//line views/stuff.qtpl:179
	qw422016.N().S(`/feeds">`)
//line views/stuff.qtpl:179
	qw422016.E().S(lc.GetWithLocale(lang, "help.feeds"))
//line views/stuff.qtpl:179
	qw422016.N().S(`</a></li>
			</ul>
		</li>
		<li>`)
//line views/stuff.qtpl:182
	qw422016.E().S(lc.GetWithLocale(lang, "help.configuration"))
//line views/stuff.qtpl:182
	qw422016.N().S(`
			<ul>
				<li><a href="/help/`)
//line views/stuff.qtpl:184
	qw422016.E().S(lang)
//line views/stuff.qtpl:184
	qw422016.N().S(`/lock">`)
//line views/stuff.qtpl:184
	qw422016.E().S(lc.GetWithLocale(lang, "help.lock"))
//line views/stuff.qtpl:184
	qw422016.N().S(`</a></li>
				<li><a href="/help/`)
//line views/stuff.qtpl:185
	qw422016.E().S(lang)
//line views/stuff.qtpl:185
	qw422016.N().S(`/whitelist">`)
//line views/stuff.qtpl:185
	qw422016.E().S(lc.GetWithLocale(lang, "help.whitelist"))
//line views/stuff.qtpl:185
	qw422016.N().S(`</a></li>
				<li><a href="/help/`)
//line views/stuff.qtpl:186
	qw422016.E().S(lang)
//line views/stuff.qtpl:186
	qw422016.N().S(`/telegram">`)
//line views/stuff.qtpl:186
	qw422016.E().S(lc.GetWithLocale(lang, "help.telegram"))
//line views/stuff.qtpl:186
	qw422016.N().S(`</a></li>
				<li>...</li>
			</ul>
		</li>
	</ul>
</aside>
`)
//line views/stuff.qtpl:192
}

//line views/stuff.qtpl:192
func writehelpTopicsHTML(qq422016 qtio422016.Writer, lang string, lc *l18n.Localizer) {
//line views/stuff.qtpl:192
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:192
	streamhelpTopicsHTML(qw422016, lang, lc)
//line views/stuff.qtpl:192
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:192
}

//line views/stuff.qtpl:192
func helpTopicsHTML(lang string, lc *l18n.Localizer) string {
//line views/stuff.qtpl:192
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:192
	writehelpTopicsHTML(qb422016, lang, lc)
//line views/stuff.qtpl:192
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:192
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:192
	return qs422016
//line views/stuff.qtpl:192
}

//line views/stuff.qtpl:194
func streamhelpTopicBadgeHTML(qw422016 *qt422016.Writer, lang, topic string) {
//line views/stuff.qtpl:194
	qw422016.N().S(`
<a class="help-topic-badge" href="/help/`)
//line views/stuff.qtpl:195
	qw422016.E().S(lang)
//line views/stuff.qtpl:195
	qw422016.N().S(`/`)
//line views/stuff.qtpl:195
	qw422016.E().S(topic)
//line views/stuff.qtpl:195
	qw422016.N().S(`">?</a>
`)
//line views/stuff.qtpl:196
}

//line views/stuff.qtpl:196
func writehelpTopicBadgeHTML(qq422016 qtio422016.Writer, lang, topic string) {
//line views/stuff.qtpl:196
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:196
	streamhelpTopicBadgeHTML(qw422016, lang, topic)
//line views/stuff.qtpl:196
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:196
}

//line views/stuff.qtpl:196
func helpTopicBadgeHTML(lang, topic string) string {
//line views/stuff.qtpl:196
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:196
	writehelpTopicBadgeHTML(qb422016, lang, topic)
//line views/stuff.qtpl:196
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:196
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:196
	return qs422016
//line views/stuff.qtpl:196
}

//line views/stuff.qtpl:198
func StreamUserListHTML(qw422016 *qt422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:198
	qw422016.N().S(`
<div class="layout">
<main class="main-width user-list">
	<h1>`)
//line views/stuff.qtpl:201
	qw422016.E().S(lc.Get("ui.users_heading"))
//line views/stuff.qtpl:201
	qw422016.N().S(`</h1>
`)
//line views/stuff.qtpl:203
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
	sort.Strings(admins)
	sort.Strings(moderators)
	sort.Strings(editors)

//line views/stuff.qtpl:221
	qw422016.N().S(`
	<section>
		<h2>`)
//line views/stuff.qtpl:223
	qw422016.E().S(lc.Get("ui.users_admins"))
//line views/stuff.qtpl:223
	qw422016.N().S(`</h2>
		<ol>`)
//line views/stuff.qtpl:224
	for _, name := range admins {
//line views/stuff.qtpl:224
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:225
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:225
		qw422016.N().S(`/`)
//line views/stuff.qtpl:225
		qw422016.E().S(name)
//line views/stuff.qtpl:225
		qw422016.N().S(`">`)
//line views/stuff.qtpl:225
		qw422016.E().S(name)
//line views/stuff.qtpl:225
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:226
	}
//line views/stuff.qtpl:226
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>`)
//line views/stuff.qtpl:229
	qw422016.E().S(lc.Get("ui.users_moderators"))
//line views/stuff.qtpl:229
	qw422016.N().S(`</h2>
		<ol>`)
//line views/stuff.qtpl:230
	for _, name := range moderators {
//line views/stuff.qtpl:230
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:231
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:231
		qw422016.N().S(`/`)
//line views/stuff.qtpl:231
		qw422016.E().S(name)
//line views/stuff.qtpl:231
		qw422016.N().S(`">`)
//line views/stuff.qtpl:231
		qw422016.E().S(name)
//line views/stuff.qtpl:231
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:232
	}
//line views/stuff.qtpl:232
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>`)
//line views/stuff.qtpl:235
	qw422016.E().S(lc.Get("ui.users_editors"))
//line views/stuff.qtpl:235
	qw422016.N().S(`</h2>
		<ol>`)
//line views/stuff.qtpl:236
	for _, name := range editors {
//line views/stuff.qtpl:236
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:237
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:237
		qw422016.N().S(`/`)
//line views/stuff.qtpl:237
		qw422016.E().S(name)
//line views/stuff.qtpl:237
		qw422016.N().S(`">`)
//line views/stuff.qtpl:237
		qw422016.E().S(name)
//line views/stuff.qtpl:237
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:238
	}
//line views/stuff.qtpl:238
	qw422016.N().S(`</ol>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:242
}

//line views/stuff.qtpl:242
func WriteUserListHTML(qq422016 qtio422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:242
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:242
	StreamUserListHTML(qw422016, lc)
//line views/stuff.qtpl:242
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:242
}

//line views/stuff.qtpl:242
func UserListHTML(lc *l18n.Localizer) string {
//line views/stuff.qtpl:242
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:242
	WriteUserListHTML(qb422016, lc)
//line views/stuff.qtpl:242
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:242
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:242
	return qs422016
//line views/stuff.qtpl:242
}

//line views/stuff.qtpl:244
func StreamHyphaListHTML(qw422016 *qt422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:244
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>`)
//line views/stuff.qtpl:247
	qw422016.E().S(lc.Get("ui.list_heading"))
//line views/stuff.qtpl:247
	qw422016.N().S(`</h1>
	<p>`)
//line views/stuff.qtpl:248
	qw422016.E().S(lc.GetPlural("ui.list_desc", hyphae.Count()))
//line views/stuff.qtpl:248
	qw422016.N().S(`</p>
	<ul class="hypha-list">
		`)
//line views/stuff.qtpl:251
	hyphaNames := make(chan string)
	sortedHypha := hyphae.PathographicSort(hyphaNames)
	for hypha := range hyphae.YieldExistingHyphae() {
		hyphaNames <- hypha.Name
	}
	close(hyphaNames)

//line views/stuff.qtpl:257
	qw422016.N().S(`
		`)
//line views/stuff.qtpl:258
	for hyphaName := range sortedHypha {
//line views/stuff.qtpl:258
		qw422016.N().S(`
		`)
//line views/stuff.qtpl:259
		hypha := hyphae.ByName(hyphaName)

//line views/stuff.qtpl:259
		qw422016.N().S(`
		<li class="hypha-list__entry">
			<a class="hypha-list__link" href="/hypha/`)
//line views/stuff.qtpl:261
		qw422016.E().S(hypha.Name)
//line views/stuff.qtpl:261
		qw422016.N().S(`">`)
//line views/stuff.qtpl:261
		qw422016.E().S(util.BeautifulName(hypha.Name))
//line views/stuff.qtpl:261
		qw422016.N().S(`</a>
			`)
//line views/stuff.qtpl:262
		if hypha.BinaryPath != "" {
//line views/stuff.qtpl:262
			qw422016.N().S(`
			<span class="hypha-list__amnt-type">`)
//line views/stuff.qtpl:263
			qw422016.E().S(filepath.Ext(hypha.BinaryPath)[1:])
//line views/stuff.qtpl:263
			qw422016.N().S(`</span>
			`)
//line views/stuff.qtpl:264
		}
//line views/stuff.qtpl:264
		qw422016.N().S(`
		</li>
		`)
//line views/stuff.qtpl:266
	}
//line views/stuff.qtpl:266
	qw422016.N().S(`
	</ul>
</main>
</div>
`)
//line views/stuff.qtpl:270
}

//line views/stuff.qtpl:270
func WriteHyphaListHTML(qq422016 qtio422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:270
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:270
	StreamHyphaListHTML(qw422016, lc)
//line views/stuff.qtpl:270
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:270
}

//line views/stuff.qtpl:270
func HyphaListHTML(lc *l18n.Localizer) string {
//line views/stuff.qtpl:270
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:270
	WriteHyphaListHTML(qb422016, lc)
//line views/stuff.qtpl:270
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:270
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:270
	return qs422016
//line views/stuff.qtpl:270
}

//line views/stuff.qtpl:272
func StreamAboutHTML(qw422016 *qt422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:272
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section>
		<h1>`)
//line views/stuff.qtpl:276
	qw422016.E().S(lc.Get("ui.about_title", &l18n.Replacements{"name": cfg.WikiName}))
//line views/stuff.qtpl:276
	qw422016.N().S(`</h1>
		<ul>
			<li><b>`)
//line views/stuff.qtpl:278
	qw422016.N().S(lc.Get("ui.about_version", &l18n.Replacements{"pre": "<a href=\"https://mycorrhiza.wiki\">", "post": "</a>"}))
//line views/stuff.qtpl:278
	qw422016.N().S(`</b> 1.6.0</li>
`)
//line views/stuff.qtpl:279
	if cfg.UseAuth {
//line views/stuff.qtpl:279
		qw422016.N().S(`			<li><b>`)
//line views/stuff.qtpl:280
		qw422016.E().S(lc.Get("ui.about_usercount"))
//line views/stuff.qtpl:280
		qw422016.N().S(`</b> `)
//line views/stuff.qtpl:280
		qw422016.N().DUL(user.Count())
//line views/stuff.qtpl:280
		qw422016.N().S(`</li>
			<li><b>`)
//line views/stuff.qtpl:281
		qw422016.E().S(lc.Get("ui.about_homepage"))
//line views/stuff.qtpl:281
		qw422016.N().S(`</b> <a href="/">`)
//line views/stuff.qtpl:281
		qw422016.E().S(cfg.HomeHypha)
//line views/stuff.qtpl:281
		qw422016.N().S(`</a></li>
			<li><b>`)
//line views/stuff.qtpl:282
		qw422016.E().S(lc.Get("ui.about_admins"))
//line views/stuff.qtpl:282
		qw422016.N().S(`</b>`)
//line views/stuff.qtpl:282
		for i, username := range user.ListUsersWithGroup("admin") {
//line views/stuff.qtpl:283
			if i > 0 {
//line views/stuff.qtpl:283
				qw422016.N().S(`<span aria-hidden="true">, </span>
`)
//line views/stuff.qtpl:284
			}
//line views/stuff.qtpl:284
			qw422016.N().S(`				<a href="/hypha/`)
//line views/stuff.qtpl:285
			qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:285
			qw422016.N().S(`/`)
//line views/stuff.qtpl:285
			qw422016.E().S(username)
//line views/stuff.qtpl:285
			qw422016.N().S(`">`)
//line views/stuff.qtpl:285
			qw422016.E().S(username)
//line views/stuff.qtpl:285
			qw422016.N().S(`</a>`)
//line views/stuff.qtpl:285
		}
//line views/stuff.qtpl:285
		qw422016.N().S(`</li>
`)
//line views/stuff.qtpl:286
	} else {
//line views/stuff.qtpl:286
		qw422016.N().S(`			<li>`)
//line views/stuff.qtpl:287
		qw422016.E().S(lc.Get("ui.about_noauth"))
//line views/stuff.qtpl:287
		qw422016.N().S(`</li>
`)
//line views/stuff.qtpl:288
	}
//line views/stuff.qtpl:288
	qw422016.N().S(`		</ul>
		<p>`)
//line views/stuff.qtpl:290
	qw422016.N().S(lc.Get("ui.about_hyphae", &l18n.Replacements{"link": "<a href=\"/list\">/list</a>"}))
//line views/stuff.qtpl:290
	qw422016.N().S(`</p>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:294
}

//line views/stuff.qtpl:294
func WriteAboutHTML(qq422016 qtio422016.Writer, lc *l18n.Localizer) {
//line views/stuff.qtpl:294
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:294
	StreamAboutHTML(qw422016, lc)
//line views/stuff.qtpl:294
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:294
}

//line views/stuff.qtpl:294
func AboutHTML(lc *l18n.Localizer) string {
//line views/stuff.qtpl:294
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:294
	WriteAboutHTML(qb422016, lc)
//line views/stuff.qtpl:294
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:294
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:294
	return qs422016
//line views/stuff.qtpl:294
}

//line views/stuff.qtpl:296
func StreamCommonScripts(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:296
	qw422016.N().S(`
`)
//line views/stuff.qtpl:297
	for _, scriptPath := range cfg.CommonScripts {
//line views/stuff.qtpl:297
		qw422016.N().S(`
<script src="`)
//line views/stuff.qtpl:298
		qw422016.E().S(scriptPath)
//line views/stuff.qtpl:298
		qw422016.N().S(`"></script>
`)
//line views/stuff.qtpl:299
	}
//line views/stuff.qtpl:299
	qw422016.N().S(`
`)
//line views/stuff.qtpl:300
}

//line views/stuff.qtpl:300
func WriteCommonScripts(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:300
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:300
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:300
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:300
}

//line views/stuff.qtpl:300
func CommonScripts() string {
//line views/stuff.qtpl:300
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:300
	WriteCommonScripts(qb422016)
//line views/stuff.qtpl:300
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:300
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:300
	return qs422016
//line views/stuff.qtpl:300
}
