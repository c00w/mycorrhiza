// Code generated by qtc from "unattach.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/unattach.qtpl:1
package templates

//line templates/unattach.qtpl:1
import "net/http"

//line templates/unattach.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/unattach.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/unattach.qtpl:2
func StreamUnattachAskHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line templates/unattach.qtpl:2
	qw422016.N().S(`
`)
//line templates/unattach.qtpl:3
	streamnavHTML(qw422016, rq, hyphaName, "unattach-ask")
//line templates/unattach.qtpl:3
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
`)
//line templates/unattach.qtpl:6
	if isOld {
//line templates/unattach.qtpl:6
		qw422016.N().S(`	<section>
		<h1>Unattach `)
//line templates/unattach.qtpl:8
		qw422016.E().S(hyphaName)
//line templates/unattach.qtpl:8
		qw422016.N().S(`?</h1>
		<p>Do you really want to unattach hypha <em>`)
//line templates/unattach.qtpl:9
		qw422016.E().S(hyphaName)
//line templates/unattach.qtpl:9
		qw422016.N().S(`</em>?</p>
		<p><a href="/unattach-confirm/`)
//line templates/unattach.qtpl:10
		qw422016.E().S(hyphaName)
//line templates/unattach.qtpl:10
		qw422016.N().S(`"><strong>Confirm</strong></a></p>
		<p><a href="/page/`)
//line templates/unattach.qtpl:11
		qw422016.E().S(hyphaName)
//line templates/unattach.qtpl:11
		qw422016.N().S(`">Cancel</a></p>
	</section>
`)
//line templates/unattach.qtpl:13
	} else {
//line templates/unattach.qtpl:13
		qw422016.N().S(`	`)
//line templates/unattach.qtpl:14
		streamcannotUnattachDueToNonExistence(qw422016, hyphaName)
//line templates/unattach.qtpl:14
		qw422016.N().S(`
`)
//line templates/unattach.qtpl:15
	}
//line templates/unattach.qtpl:15
	qw422016.N().S(`</main>
</div>
`)
//line templates/unattach.qtpl:18
}

//line templates/unattach.qtpl:18
func WriteUnattachAskHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line templates/unattach.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/unattach.qtpl:18
	StreamUnattachAskHTML(qw422016, rq, hyphaName, isOld)
//line templates/unattach.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line templates/unattach.qtpl:18
}

//line templates/unattach.qtpl:18
func UnattachAskHTML(rq *http.Request, hyphaName string, isOld bool) string {
//line templates/unattach.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/unattach.qtpl:18
	WriteUnattachAskHTML(qb422016, rq, hyphaName, isOld)
//line templates/unattach.qtpl:18
	qs422016 := string(qb422016.B)
//line templates/unattach.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/unattach.qtpl:18
	return qs422016
//line templates/unattach.qtpl:18
}

//line templates/unattach.qtpl:20
func streamcannotUnattachDueToNonExistence(qw422016 *qt422016.Writer, hyphaName string) {
//line templates/unattach.qtpl:20
	qw422016.N().S(`
	<section>
		<h1>Cannot unattach `)
//line templates/unattach.qtpl:22
	qw422016.E().S(hyphaName)
//line templates/unattach.qtpl:22
	qw422016.N().S(`</h1>
		<p>This hypha does not exist.</p>
		<p><a href="/page/`)
//line templates/unattach.qtpl:24
	qw422016.E().S(hyphaName)
//line templates/unattach.qtpl:24
	qw422016.N().S(`">Go back</a></p>
	</section>
`)
//line templates/unattach.qtpl:26
}

//line templates/unattach.qtpl:26
func writecannotUnattachDueToNonExistence(qq422016 qtio422016.Writer, hyphaName string) {
//line templates/unattach.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/unattach.qtpl:26
	streamcannotUnattachDueToNonExistence(qw422016, hyphaName)
//line templates/unattach.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line templates/unattach.qtpl:26
}

//line templates/unattach.qtpl:26
func cannotUnattachDueToNonExistence(hyphaName string) string {
//line templates/unattach.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/unattach.qtpl:26
	writecannotUnattachDueToNonExistence(qb422016, hyphaName)
//line templates/unattach.qtpl:26
	qs422016 := string(qb422016.B)
//line templates/unattach.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/unattach.qtpl:26
	return qs422016
//line templates/unattach.qtpl:26
}
