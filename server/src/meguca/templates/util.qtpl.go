// Code generated by qtc from "util.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line server/src/meguca/templates/util.qtpl:1
package templates

//line server/src/meguca/templates/util.qtpl:1
import "strconv"

//line server/src/meguca/templates/util.qtpl:2
import "fmt"

//line server/src/meguca/templates/util.qtpl:3
import "meguca/common"

//line server/src/meguca/templates/util.qtpl:4
import "meguca/assets"

//line server/src/meguca/templates/util.qtpl:5
import "meguca/lang"

// Renders the tab selection butts in tabbed windows

//line server/src/meguca/templates/util.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line server/src/meguca/templates/util.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line server/src/meguca/templates/util.qtpl:8
func streamtabButts(qw422016 *qt422016.Writer, names []string) {
	//line server/src/meguca/templates/util.qtpl:8
	qw422016.N().S(`<div class="tab-butts">`)
	//line server/src/meguca/templates/util.qtpl:10
	for i, n := range names {
		//line server/src/meguca/templates/util.qtpl:10
		qw422016.N().S(`<a class="tab-link`)
		//line server/src/meguca/templates/util.qtpl:11
		if i == 0 {
			//line server/src/meguca/templates/util.qtpl:11
			qw422016.N().S(` `)
			//line server/src/meguca/templates/util.qtpl:11
			qw422016.N().S(`tab-sel`)
			//line server/src/meguca/templates/util.qtpl:11
		}
		//line server/src/meguca/templates/util.qtpl:11
		qw422016.N().S(`" data-id="`)
		//line server/src/meguca/templates/util.qtpl:11
		qw422016.N().D(i)
		//line server/src/meguca/templates/util.qtpl:11
		qw422016.N().S(`">`)
		//line server/src/meguca/templates/util.qtpl:12
		qw422016.N().S(n)
		//line server/src/meguca/templates/util.qtpl:12
		qw422016.N().S(`</a>`)
		//line server/src/meguca/templates/util.qtpl:14
	}
	//line server/src/meguca/templates/util.qtpl:14
	qw422016.N().S(`</div><hr>`)
//line server/src/meguca/templates/util.qtpl:17
}

//line server/src/meguca/templates/util.qtpl:17
func writetabButts(qq422016 qtio422016.Writer, names []string) {
	//line server/src/meguca/templates/util.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:17
	streamtabButts(qw422016, names)
	//line server/src/meguca/templates/util.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:17
}

//line server/src/meguca/templates/util.qtpl:17
func tabButts(names []string) string {
	//line server/src/meguca/templates/util.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:17
	writetabButts(qb422016, names)
	//line server/src/meguca/templates/util.qtpl:17
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:17
	return qs422016
//line server/src/meguca/templates/util.qtpl:17
}

// Render a link to another post. Can optionally be cross-thread.

//line server/src/meguca/templates/util.qtpl:20
func streampostLink(qw422016 *qt422016.Writer, link common.Link, cross, boardPage bool) {
	//line server/src/meguca/templates/util.qtpl:21
	idBuf := strconv.AppendUint(make([]byte, 0, 16), link.ID, 10)

	//line server/src/meguca/templates/util.qtpl:22
	url := make([]byte, 0, 64)

	//line server/src/meguca/templates/util.qtpl:23
	if cross {
		//line server/src/meguca/templates/util.qtpl:24
		url = append(url, '/')

		//line server/src/meguca/templates/util.qtpl:25
		url = append(url, link.Board...)

		//line server/src/meguca/templates/util.qtpl:26
		url = append(url, '/')

		//line server/src/meguca/templates/util.qtpl:27
		url = strconv.AppendUint(url, link.OP, 10)

		//line server/src/meguca/templates/util.qtpl:28
	}
	//line server/src/meguca/templates/util.qtpl:29
	url = append(url, "#p"...)

	//line server/src/meguca/templates/util.qtpl:30
	url = append(url, idBuf...)

	//line server/src/meguca/templates/util.qtpl:30
	qw422016.N().S(`<a class="post-link" data-id="`)
	//line server/src/meguca/templates/util.qtpl:31
	qw422016.N().Z(idBuf)
	//line server/src/meguca/templates/util.qtpl:31
	qw422016.N().S(`" href="`)
	//line server/src/meguca/templates/util.qtpl:31
	qw422016.N().Z(url)
	//line server/src/meguca/templates/util.qtpl:31
	qw422016.N().S(`">>>`)
	//line server/src/meguca/templates/util.qtpl:33
	qw422016.N().Z(idBuf)
	//line server/src/meguca/templates/util.qtpl:34
	if cross && !boardPage {
		//line server/src/meguca/templates/util.qtpl:35
		qw422016.N().S(` `)
		//line server/src/meguca/templates/util.qtpl:35
		qw422016.N().S(`➡`)
		//line server/src/meguca/templates/util.qtpl:36
	}
	//line server/src/meguca/templates/util.qtpl:36
	qw422016.N().S(`</a><a class="hash-link" href="`)
	//line server/src/meguca/templates/util.qtpl:38
	qw422016.N().Z(url)
	//line server/src/meguca/templates/util.qtpl:38
	qw422016.N().S(`"> #</a>`)
//line server/src/meguca/templates/util.qtpl:39
}

//line server/src/meguca/templates/util.qtpl:39
func writepostLink(qq422016 qtio422016.Writer, link common.Link, cross, boardPage bool) {
	//line server/src/meguca/templates/util.qtpl:39
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:39
	streampostLink(qw422016, link, cross, boardPage)
	//line server/src/meguca/templates/util.qtpl:39
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:39
}

//line server/src/meguca/templates/util.qtpl:39
func postLink(link common.Link, cross, boardPage bool) string {
	//line server/src/meguca/templates/util.qtpl:39
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:39
	writepostLink(qb422016, link, cross, boardPage)
	//line server/src/meguca/templates/util.qtpl:39
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:39
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:39
	return qs422016
//line server/src/meguca/templates/util.qtpl:39
}

//line server/src/meguca/templates/util.qtpl:41
func streamexpandLink(qw422016 *qt422016.Writer, board, id string) {
	//line server/src/meguca/templates/util.qtpl:41
	qw422016.N().S(`<span class="act"><a href="/`)
	//line server/src/meguca/templates/util.qtpl:43
	qw422016.N().S(board)
	//line server/src/meguca/templates/util.qtpl:43
	qw422016.N().S(`/`)
	//line server/src/meguca/templates/util.qtpl:43
	qw422016.N().S(id)
	//line server/src/meguca/templates/util.qtpl:43
	qw422016.N().S(`">`)
	//line server/src/meguca/templates/util.qtpl:44
	qw422016.N().S(lang.Get().Common.Posts["expand"])
	//line server/src/meguca/templates/util.qtpl:44
	qw422016.N().S(`</a></span>`)
//line server/src/meguca/templates/util.qtpl:47
}

//line server/src/meguca/templates/util.qtpl:47
func writeexpandLink(qq422016 qtio422016.Writer, board, id string) {
	//line server/src/meguca/templates/util.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:47
	streamexpandLink(qw422016, board, id)
	//line server/src/meguca/templates/util.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:47
}

//line server/src/meguca/templates/util.qtpl:47
func expandLink(board, id string) string {
	//line server/src/meguca/templates/util.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:47
	writeexpandLink(qb422016, board, id)
	//line server/src/meguca/templates/util.qtpl:47
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:47
	return qs422016
//line server/src/meguca/templates/util.qtpl:47
}

//line server/src/meguca/templates/util.qtpl:49
func streamlast100Link(qw422016 *qt422016.Writer, board, id string) {
	//line server/src/meguca/templates/util.qtpl:49
	qw422016.N().S(`<span class="act"><a href="/`)
	//line server/src/meguca/templates/util.qtpl:51
	qw422016.N().S(board)
	//line server/src/meguca/templates/util.qtpl:51
	qw422016.N().S(`/`)
	//line server/src/meguca/templates/util.qtpl:51
	qw422016.N().S(id)
	//line server/src/meguca/templates/util.qtpl:51
	qw422016.N().S(`?last=100#bottom">`)
	//line server/src/meguca/templates/util.qtpl:52
	qw422016.N().S(lang.Get().Common.UI["last"])
	//line server/src/meguca/templates/util.qtpl:52
	qw422016.N().S(` `)
	//line server/src/meguca/templates/util.qtpl:52
	qw422016.N().S(`100</a></span>`)
//line server/src/meguca/templates/util.qtpl:55
}

//line server/src/meguca/templates/util.qtpl:55
func writelast100Link(qq422016 qtio422016.Writer, board, id string) {
	//line server/src/meguca/templates/util.qtpl:55
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:55
	streamlast100Link(qw422016, board, id)
	//line server/src/meguca/templates/util.qtpl:55
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:55
}

//line server/src/meguca/templates/util.qtpl:55
func last100Link(board, id string) string {
	//line server/src/meguca/templates/util.qtpl:55
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:55
	writelast100Link(qb422016, board, id)
	//line server/src/meguca/templates/util.qtpl:55
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:55
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:55
	return qs422016
//line server/src/meguca/templates/util.qtpl:55
}

// Render the class attribute of a post

//line server/src/meguca/templates/util.qtpl:58
func streampostClass(qw422016 *qt422016.Writer, p common.Post, op uint64) {
	//line server/src/meguca/templates/util.qtpl:58
	qw422016.N().S(`class="glass`)
	//line server/src/meguca/templates/util.qtpl:60
	if p.Editing {
		//line server/src/meguca/templates/util.qtpl:61
		qw422016.N().S(` `)
		//line server/src/meguca/templates/util.qtpl:61
		qw422016.N().S(`editing`)
		//line server/src/meguca/templates/util.qtpl:62
	}
	//line server/src/meguca/templates/util.qtpl:63
	if p.IsDeleted() {
		//line server/src/meguca/templates/util.qtpl:64
		qw422016.N().S(` `)
		//line server/src/meguca/templates/util.qtpl:64
		qw422016.N().S(`deleted`)
		//line server/src/meguca/templates/util.qtpl:65
	}
	//line server/src/meguca/templates/util.qtpl:66
	if p.Image != nil {
		//line server/src/meguca/templates/util.qtpl:67
		qw422016.N().S(` `)
		//line server/src/meguca/templates/util.qtpl:67
		qw422016.N().S(`media`)
		//line server/src/meguca/templates/util.qtpl:68
	}
	//line server/src/meguca/templates/util.qtpl:69
	if p.ID == op {
		//line server/src/meguca/templates/util.qtpl:70
		qw422016.N().S(` `)
		//line server/src/meguca/templates/util.qtpl:70
		qw422016.N().S(`op`)
		//line server/src/meguca/templates/util.qtpl:71
	}
	//line server/src/meguca/templates/util.qtpl:71
	qw422016.N().S(`"`)
//line server/src/meguca/templates/util.qtpl:73
}

//line server/src/meguca/templates/util.qtpl:73
func writepostClass(qq422016 qtio422016.Writer, p common.Post, op uint64) {
	//line server/src/meguca/templates/util.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:73
	streampostClass(qw422016, p, op)
	//line server/src/meguca/templates/util.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:73
}

//line server/src/meguca/templates/util.qtpl:73
func postClass(p common.Post, op uint64) string {
	//line server/src/meguca/templates/util.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:73
	writepostClass(qb422016, p, op)
	//line server/src/meguca/templates/util.qtpl:73
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:73
	return qs422016
//line server/src/meguca/templates/util.qtpl:73
}

// Renders a stylized deleted post display toggle

//line server/src/meguca/templates/util.qtpl:76
func streamdeletedToggle(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/util.qtpl:76
	qw422016.N().S(`<input type="checkbox" class="deleted-toggle">`)
//line server/src/meguca/templates/util.qtpl:78
}

//line server/src/meguca/templates/util.qtpl:78
func writedeletedToggle(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/util.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:78
	streamdeletedToggle(qw422016)
	//line server/src/meguca/templates/util.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:78
}

//line server/src/meguca/templates/util.qtpl:78
func deletedToggle() string {
	//line server/src/meguca/templates/util.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:78
	writedeletedToggle(qb422016)
	//line server/src/meguca/templates/util.qtpl:78
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:78
	return qs422016
//line server/src/meguca/templates/util.qtpl:78
}

// Notice widget, that reveals text on hover

//line server/src/meguca/templates/util.qtpl:82
func streamhoverReveal(qw422016 *qt422016.Writer, tag, text, label string) {
	//line server/src/meguca/templates/util.qtpl:83
	if text == "" {
		//line server/src/meguca/templates/util.qtpl:84
		return
		//line server/src/meguca/templates/util.qtpl:85
	}
	//line server/src/meguca/templates/util.qtpl:85
	qw422016.N().S(`<`)
	//line server/src/meguca/templates/util.qtpl:86
	qw422016.N().S(tag)
	//line server/src/meguca/templates/util.qtpl:86
	qw422016.N().S(` `)
	//line server/src/meguca/templates/util.qtpl:86
	qw422016.N().S(`class="hover-reveal`)
	//line server/src/meguca/templates/util.qtpl:86
	if tag == "aside" {
		//line server/src/meguca/templates/util.qtpl:86
		qw422016.N().S(` `)
		//line server/src/meguca/templates/util.qtpl:86
		qw422016.N().S(`glass`)
		//line server/src/meguca/templates/util.qtpl:86
	}
	//line server/src/meguca/templates/util.qtpl:86
	qw422016.N().S(`"><span class="act">`)
	//line server/src/meguca/templates/util.qtpl:88
	qw422016.N().S(label)
	//line server/src/meguca/templates/util.qtpl:88
	qw422016.N().S(`</span><span class="popup-menu glass">`)
	//line server/src/meguca/templates/util.qtpl:91
	qw422016.E().S(text)
	//line server/src/meguca/templates/util.qtpl:91
	qw422016.N().S(`</span></`)
	//line server/src/meguca/templates/util.qtpl:93
	qw422016.N().S(tag)
	//line server/src/meguca/templates/util.qtpl:93
	qw422016.N().S(`>`)
//line server/src/meguca/templates/util.qtpl:94
}

//line server/src/meguca/templates/util.qtpl:94
func writehoverReveal(qq422016 qtio422016.Writer, tag, text, label string) {
	//line server/src/meguca/templates/util.qtpl:94
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:94
	streamhoverReveal(qw422016, tag, text, label)
	//line server/src/meguca/templates/util.qtpl:94
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:94
}

//line server/src/meguca/templates/util.qtpl:94
func hoverReveal(tag, text, label string) string {
	//line server/src/meguca/templates/util.qtpl:94
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:94
	writehoverReveal(qb422016, tag, text, label)
	//line server/src/meguca/templates/util.qtpl:94
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:94
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:94
	return qs422016
//line server/src/meguca/templates/util.qtpl:94
}

// Render pin signifying a thread is sticky

//line server/src/meguca/templates/util.qtpl:97
func streamrenderSticky(qw422016 *qt422016.Writer, sticky bool) {
	//line server/src/meguca/templates/util.qtpl:98
	if !sticky {
		//line server/src/meguca/templates/util.qtpl:99
		return
		//line server/src/meguca/templates/util.qtpl:100
	}
	//line server/src/meguca/templates/util.qtpl:100
	qw422016.N().S(`<svg class="sticky" xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M1.34 0a.5.5 0 0 0 .16 1h.5v2h-1c-.55 0-1 .45-1 1h3v3l.44 1 .56-1v-3h3c0-.55-.45-1-1-1h-1v-2h.5a.5.5 0 1 0 0-1h-4a.5.5 0 0 0-.09 0 .5.5 0 0 0-.06 0z" /></svg>`)
//line server/src/meguca/templates/util.qtpl:104
}

//line server/src/meguca/templates/util.qtpl:104
func writerenderSticky(qq422016 qtio422016.Writer, sticky bool) {
	//line server/src/meguca/templates/util.qtpl:104
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:104
	streamrenderSticky(qw422016, sticky)
	//line server/src/meguca/templates/util.qtpl:104
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:104
}

//line server/src/meguca/templates/util.qtpl:104
func renderSticky(sticky bool) string {
	//line server/src/meguca/templates/util.qtpl:104
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:104
	writerenderSticky(qb422016, sticky)
	//line server/src/meguca/templates/util.qtpl:104
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:104
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:104
	return qs422016
//line server/src/meguca/templates/util.qtpl:104
}

// Render lock signifying a thread has posting disabled

//line server/src/meguca/templates/util.qtpl:107
func streamrenderLocked(qw422016 *qt422016.Writer, locked bool) {
	//line server/src/meguca/templates/util.qtpl:108
	if !locked {
		//line server/src/meguca/templates/util.qtpl:109
		return
		//line server/src/meguca/templates/util.qtpl:110
	}
	//line server/src/meguca/templates/util.qtpl:110
	qw422016.N().S(`<svg class="locked" xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M3 0c-1.1 0-2 .9-2 2v1h-1v4h6v-4h-1v-1c0-1.1-.9-2-2-2zm0 1c.56 0 1 .44 1 1v1h-2v-1c0-.56.44-1 1-1z" transform="translate(1)" /></svg>`)
//line server/src/meguca/templates/util.qtpl:114
}

//line server/src/meguca/templates/util.qtpl:114
func writerenderLocked(qq422016 qtio422016.Writer, locked bool) {
	//line server/src/meguca/templates/util.qtpl:114
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:114
	streamrenderLocked(qw422016, locked)
	//line server/src/meguca/templates/util.qtpl:114
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:114
}

//line server/src/meguca/templates/util.qtpl:114
func renderLocked(locked bool) string {
	//line server/src/meguca/templates/util.qtpl:114
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:114
	writerenderLocked(qb422016, locked)
	//line server/src/meguca/templates/util.qtpl:114
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:114
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:114
	return qs422016
//line server/src/meguca/templates/util.qtpl:114
}

// Render an image or video asset

//line server/src/meguca/templates/util.qtpl:117
func streamasset(qw422016 *qt422016.Writer, url, mime string) {
	//line server/src/meguca/templates/util.qtpl:118
	if mime == "video/webm" {
		//line server/src/meguca/templates/util.qtpl:118
		qw422016.N().S(`<video src="`)
		//line server/src/meguca/templates/util.qtpl:119
		qw422016.N().S(url)
		//line server/src/meguca/templates/util.qtpl:119
		qw422016.N().S(`" autoplay loop>`)
		//line server/src/meguca/templates/util.qtpl:120
	} else {
		//line server/src/meguca/templates/util.qtpl:120
		qw422016.N().S(`<img src="`)
		//line server/src/meguca/templates/util.qtpl:121
		qw422016.N().S(url)
		//line server/src/meguca/templates/util.qtpl:121
		qw422016.N().S(`">`)
		//line server/src/meguca/templates/util.qtpl:122
	}
//line server/src/meguca/templates/util.qtpl:123
}

//line server/src/meguca/templates/util.qtpl:123
func writeasset(qq422016 qtio422016.Writer, url, mime string) {
	//line server/src/meguca/templates/util.qtpl:123
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:123
	streamasset(qw422016, url, mime)
	//line server/src/meguca/templates/util.qtpl:123
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:123
}

//line server/src/meguca/templates/util.qtpl:123
func asset(url, mime string) string {
	//line server/src/meguca/templates/util.qtpl:123
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:123
	writeasset(qb422016, url, mime)
	//line server/src/meguca/templates/util.qtpl:123
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:123
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:123
	return qs422016
//line server/src/meguca/templates/util.qtpl:123
}

//line server/src/meguca/templates/util.qtpl:125
func streamloadingImage(qw422016 *qt422016.Writer, board string) {
	//line server/src/meguca/templates/util.qtpl:125
	qw422016.N().S(`<div id="loading-image" class="noscript-hide">`)
	//line server/src/meguca/templates/util.qtpl:127
	streamasset(qw422016, fmt.Sprintf("/assets/loading/%s", board), assets.Loading.Get(board).Mime)
	//line server/src/meguca/templates/util.qtpl:127
	qw422016.N().S(`</div>`)
//line server/src/meguca/templates/util.qtpl:129
}

//line server/src/meguca/templates/util.qtpl:129
func writeloadingImage(qq422016 qtio422016.Writer, board string) {
	//line server/src/meguca/templates/util.qtpl:129
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:129
	streamloadingImage(qw422016, board)
	//line server/src/meguca/templates/util.qtpl:129
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:129
}

//line server/src/meguca/templates/util.qtpl:129
func loadingImage(board string) string {
	//line server/src/meguca/templates/util.qtpl:129
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:129
	writeloadingImage(qb422016, board)
	//line server/src/meguca/templates/util.qtpl:129
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:129
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:129
	return qs422016
//line server/src/meguca/templates/util.qtpl:129
}

// Render localized table headers by UI translation ID

//line server/src/meguca/templates/util.qtpl:132
func streamtableHeaders(qw422016 *qt422016.Writer, ids ...string) {
	//line server/src/meguca/templates/util.qtpl:133
	ln := lang.Get().UI

	//line server/src/meguca/templates/util.qtpl:133
	qw422016.N().S(`<tr>`)
	//line server/src/meguca/templates/util.qtpl:135
	for _, id := range ids {
		//line server/src/meguca/templates/util.qtpl:135
		qw422016.N().S(`<th>`)
		//line server/src/meguca/templates/util.qtpl:136
		qw422016.N().S(ln[id])
		//line server/src/meguca/templates/util.qtpl:136
		qw422016.N().S(`</th>`)
		//line server/src/meguca/templates/util.qtpl:137
	}
	//line server/src/meguca/templates/util.qtpl:137
	qw422016.N().S(`</tr>`)
//line server/src/meguca/templates/util.qtpl:139
}

//line server/src/meguca/templates/util.qtpl:139
func writetableHeaders(qq422016 qtio422016.Writer, ids ...string) {
	//line server/src/meguca/templates/util.qtpl:139
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:139
	streamtableHeaders(qw422016, ids...)
	//line server/src/meguca/templates/util.qtpl:139
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:139
}

//line server/src/meguca/templates/util.qtpl:139
func tableHeaders(ids ...string) string {
	//line server/src/meguca/templates/util.qtpl:139
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:139
	writetableHeaders(qb422016, ids...)
	//line server/src/meguca/templates/util.qtpl:139
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:139
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:139
	return qs422016
//line server/src/meguca/templates/util.qtpl:139
}

//line server/src/meguca/templates/util.qtpl:141
func streamthreadWatcherToggle(qw422016 *qt422016.Writer, id uint64) {
	//line server/src/meguca/templates/util.qtpl:141
	qw422016.N().S(`<a class="watcher-toggle svg-link noscript-hide" title="`)
	//line server/src/meguca/templates/util.qtpl:142
	qw422016.N().S(lang.Get().Common.UI["watchThread"])
	//line server/src/meguca/templates/util.qtpl:142
	qw422016.N().S(`" data-id="`)
	//line server/src/meguca/templates/util.qtpl:142
	qw422016.N().S(strconv.FormatUint(id, 10))
	//line server/src/meguca/templates/util.qtpl:142
	qw422016.N().S(`"><svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M4.03 0c-2.53 0-4.03 3-4.03 3s1.5 3 4.03 3c2.47 0 3.97-3 3.97-3s-1.5-3-3.97-3zm-.03 1c1.11 0 2 .9 2 2 0 1.11-.89 2-2 2-1.1 0-2-.89-2-2 0-1.1.9-2 2-2zm0 1c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1c0-.1-.04-.19-.06-.28-.08.16-.24.28-.44.28-.28 0-.5-.22-.5-.5 0-.2.12-.36.28-.44-.09-.03-.18-.06-.28-.06z" transform="translate(0 1)" /></svg></a>`)
//line server/src/meguca/templates/util.qtpl:147
}

//line server/src/meguca/templates/util.qtpl:147
func writethreadWatcherToggle(qq422016 qtio422016.Writer, id uint64) {
	//line server/src/meguca/templates/util.qtpl:147
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:147
	streamthreadWatcherToggle(qw422016, id)
	//line server/src/meguca/templates/util.qtpl:147
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:147
}

//line server/src/meguca/templates/util.qtpl:147
func threadWatcherToggle(id uint64) string {
	//line server/src/meguca/templates/util.qtpl:147
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:147
	writethreadWatcherToggle(qb422016, id)
	//line server/src/meguca/templates/util.qtpl:147
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:147
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:147
	return qs422016
//line server/src/meguca/templates/util.qtpl:147
}

//line server/src/meguca/templates/util.qtpl:149
func streamcontrolLink(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/util.qtpl:149
	qw422016.N().S(`<a class="control svg-link noscript-hide"><svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M1.5 0l-1.5 1.5 4 4 4-4-1.5-1.5-2.5 2.5-2.5-2.5z" transform="translate(0 1)" /></svg></a>`)
//line server/src/meguca/templates/util.qtpl:155
}

//line server/src/meguca/templates/util.qtpl:155
func writecontrolLink(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/util.qtpl:155
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/util.qtpl:155
	streamcontrolLink(qw422016)
	//line server/src/meguca/templates/util.qtpl:155
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/util.qtpl:155
}

//line server/src/meguca/templates/util.qtpl:155
func controlLink() string {
	//line server/src/meguca/templates/util.qtpl:155
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/util.qtpl:155
	writecontrolLink(qb422016)
	//line server/src/meguca/templates/util.qtpl:155
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/util.qtpl:155
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/util.qtpl:155
	return qs422016
//line server/src/meguca/templates/util.qtpl:155
}
