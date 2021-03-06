// Code generated by qtc from "forms.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line server/src/meguca/templates/forms.qtpl:1
package templates

//line server/src/meguca/templates/forms.qtpl:1
import "meguca/config"

//line server/src/meguca/templates/forms.qtpl:2
import "meguca/lang"

// OwnedBoard renders a form for selecting one of several boards owned by the user

//line server/src/meguca/templates/forms.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line server/src/meguca/templates/forms.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line server/src/meguca/templates/forms.qtpl:5
func StreamOwnedBoard(qw422016 *qt422016.Writer, boards config.BoardTitles) {
	//line server/src/meguca/templates/forms.qtpl:6
	if len(boards) != 0 {
		//line server/src/meguca/templates/forms.qtpl:6
		qw422016.N().S(`<select name="boards" required>`)
		//line server/src/meguca/templates/forms.qtpl:8
		for _, b := range boards {
			//line server/src/meguca/templates/forms.qtpl:8
			qw422016.N().S(`<option value="`)
			//line server/src/meguca/templates/forms.qtpl:9
			qw422016.N().S(b.ID)
			//line server/src/meguca/templates/forms.qtpl:9
			qw422016.N().S(`">`)
			//line server/src/meguca/templates/forms.qtpl:10
			streamformatTitle(qw422016, b.ID, b.Title)
			//line server/src/meguca/templates/forms.qtpl:10
			qw422016.N().S(`</option>`)
			//line server/src/meguca/templates/forms.qtpl:12
		}
		//line server/src/meguca/templates/forms.qtpl:12
		qw422016.N().S(`</select><br>`)
		//line server/src/meguca/templates/forms.qtpl:15
		streamsubmit(qw422016, true)
		//line server/src/meguca/templates/forms.qtpl:16
	} else {
		//line server/src/meguca/templates/forms.qtpl:17
		qw422016.N().S(lang.Get().UI["ownNoBoards"])
		//line server/src/meguca/templates/forms.qtpl:17
		qw422016.N().S(`<br><br>`)
		//line server/src/meguca/templates/forms.qtpl:20
		streamcancel(qw422016)
		//line server/src/meguca/templates/forms.qtpl:20
		qw422016.N().S(`<div class="form-response admin"></div>`)
		//line server/src/meguca/templates/forms.qtpl:22
	}
//line server/src/meguca/templates/forms.qtpl:23
}

//line server/src/meguca/templates/forms.qtpl:23
func WriteOwnedBoard(qq422016 qtio422016.Writer, boards config.BoardTitles) {
	//line server/src/meguca/templates/forms.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:23
	StreamOwnedBoard(qw422016, boards)
	//line server/src/meguca/templates/forms.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:23
}

//line server/src/meguca/templates/forms.qtpl:23
func OwnedBoard(boards config.BoardTitles) string {
	//line server/src/meguca/templates/forms.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:23
	WriteOwnedBoard(qb422016, boards)
	//line server/src/meguca/templates/forms.qtpl:23
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:23
	return qs422016
//line server/src/meguca/templates/forms.qtpl:23
}

//line server/src/meguca/templates/forms.qtpl:25
func streamformatTitle(qw422016 *qt422016.Writer, id, title string) {
	//line server/src/meguca/templates/forms.qtpl:25
	qw422016.N().S(`/`)
	//line server/src/meguca/templates/forms.qtpl:26
	qw422016.N().S(id)
	//line server/src/meguca/templates/forms.qtpl:26
	qw422016.N().S(`/`)
	//line server/src/meguca/templates/forms.qtpl:26
	qw422016.N().S(` `)
	//line server/src/meguca/templates/forms.qtpl:26
	qw422016.N().S(`-`)
	//line server/src/meguca/templates/forms.qtpl:26
	qw422016.N().S(` `)
	//line server/src/meguca/templates/forms.qtpl:26
	qw422016.E().S(title)
//line server/src/meguca/templates/forms.qtpl:27
}

//line server/src/meguca/templates/forms.qtpl:27
func writeformatTitle(qq422016 qtio422016.Writer, id, title string) {
	//line server/src/meguca/templates/forms.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:27
	streamformatTitle(qw422016, id, title)
	//line server/src/meguca/templates/forms.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:27
}

//line server/src/meguca/templates/forms.qtpl:27
func formatTitle(id, title string) string {
	//line server/src/meguca/templates/forms.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:27
	writeformatTitle(qb422016, id, title)
	//line server/src/meguca/templates/forms.qtpl:27
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:27
	return qs422016
//line server/src/meguca/templates/forms.qtpl:27
}

// BoardNavigation renders a board selection and search form

//line server/src/meguca/templates/forms.qtpl:30
func StreamBoardNavigation(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:31
	ln := lang.Get().Common.UI

	//line server/src/meguca/templates/forms.qtpl:31
	qw422016.N().S(`<input type="text" class="full-width" name="search" placeholder="`)
	//line server/src/meguca/templates/forms.qtpl:32
	qw422016.N().S(ln["search"])
	//line server/src/meguca/templates/forms.qtpl:32
	qw422016.N().S(`"><br><form><span class="flex">`)
	//line server/src/meguca/templates/forms.qtpl:36
	streamsubmit(qw422016, true)
	//line server/src/meguca/templates/forms.qtpl:36
	qw422016.N().S(`<label><input type="checkbox" name="pointToCatalog">`)
	//line server/src/meguca/templates/forms.qtpl:39
	qw422016.N().S(ln["pointToCatalog"])
	//line server/src/meguca/templates/forms.qtpl:39
	qw422016.N().S(`</label></span><div class="board-list">`)
	//line server/src/meguca/templates/forms.qtpl:43
	for _, b := range config.GetBoardTitles() {
		//line server/src/meguca/templates/forms.qtpl:43
		qw422016.N().S(`<label class="board"><input type="checkbox" name="`)
		//line server/src/meguca/templates/forms.qtpl:45
		qw422016.N().S(b.ID)
		//line server/src/meguca/templates/forms.qtpl:45
		qw422016.N().S(`"><a href="/`)
		//line server/src/meguca/templates/forms.qtpl:46
		qw422016.N().S(b.ID)
		//line server/src/meguca/templates/forms.qtpl:46
		qw422016.N().S(`/">`)
		//line server/src/meguca/templates/forms.qtpl:47
		streamformatTitle(qw422016, b.ID, b.Title)
		//line server/src/meguca/templates/forms.qtpl:47
		qw422016.N().S(`</a><br></label>`)
		//line server/src/meguca/templates/forms.qtpl:51
	}
	//line server/src/meguca/templates/forms.qtpl:51
	qw422016.N().S(`</div></form>`)
//line server/src/meguca/templates/forms.qtpl:54
}

//line server/src/meguca/templates/forms.qtpl:54
func WriteBoardNavigation(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:54
	StreamBoardNavigation(qw422016)
	//line server/src/meguca/templates/forms.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:54
}

//line server/src/meguca/templates/forms.qtpl:54
func BoardNavigation() string {
	//line server/src/meguca/templates/forms.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:54
	WriteBoardNavigation(qb422016)
	//line server/src/meguca/templates/forms.qtpl:54
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:54
	return qs422016
//line server/src/meguca/templates/forms.qtpl:54
}

// CreateBoard renders a the form for creating new boards

//line server/src/meguca/templates/forms.qtpl:57
func StreamCreateBoard(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:58
	streamtable(qw422016, specs["createBoard"])
	//line server/src/meguca/templates/forms.qtpl:59
	StreamCaptchaConfirmation(qw422016)
//line server/src/meguca/templates/forms.qtpl:60
}

//line server/src/meguca/templates/forms.qtpl:60
func WriteCreateBoard(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:60
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:60
	StreamCreateBoard(qw422016)
	//line server/src/meguca/templates/forms.qtpl:60
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:60
}

//line server/src/meguca/templates/forms.qtpl:60
func CreateBoard() string {
	//line server/src/meguca/templates/forms.qtpl:60
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:60
	WriteCreateBoard(qb422016)
	//line server/src/meguca/templates/forms.qtpl:60
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:60
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:60
	return qs422016
//line server/src/meguca/templates/forms.qtpl:60
}

// CaptchaConfirmation renders a confirmation form with an optional captcha

//line server/src/meguca/templates/forms.qtpl:63
func StreamCaptchaConfirmation(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:64
	streamcaptcha(qw422016, "all")
	//line server/src/meguca/templates/forms.qtpl:65
	streamsubmit(qw422016, true)
//line server/src/meguca/templates/forms.qtpl:66
}

//line server/src/meguca/templates/forms.qtpl:66
func WriteCaptchaConfirmation(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:66
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:66
	StreamCaptchaConfirmation(qw422016)
	//line server/src/meguca/templates/forms.qtpl:66
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:66
}

//line server/src/meguca/templates/forms.qtpl:66
func CaptchaConfirmation() string {
	//line server/src/meguca/templates/forms.qtpl:66
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:66
	WriteCaptchaConfirmation(qb422016)
	//line server/src/meguca/templates/forms.qtpl:66
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:66
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:66
	return qs422016
//line server/src/meguca/templates/forms.qtpl:66
}

//line server/src/meguca/templates/forms.qtpl:68
func streamcaptcha(qw422016 *qt422016.Writer, board string) {
	//line server/src/meguca/templates/forms.qtpl:69
	if !config.Get().Captcha {
		//line server/src/meguca/templates/forms.qtpl:70
		return
		//line server/src/meguca/templates/forms.qtpl:71
	}
	//line server/src/meguca/templates/forms.qtpl:71
	qw422016.N().S(`<div class="captcha-container full-width"><noscript><iframe width="462" height="525" scrolling="no" marginwidth="0" marginheight="0" src="/api/captcha/`)
	//line server/src/meguca/templates/forms.qtpl:74
	qw422016.N().S(board)
	//line server/src/meguca/templates/forms.qtpl:74
	qw422016.N().S(`"></iframe></noscript></div>`)
//line server/src/meguca/templates/forms.qtpl:77
}

//line server/src/meguca/templates/forms.qtpl:77
func writecaptcha(qq422016 qtio422016.Writer, board string) {
	//line server/src/meguca/templates/forms.qtpl:77
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:77
	streamcaptcha(qw422016, board)
	//line server/src/meguca/templates/forms.qtpl:77
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:77
}

//line server/src/meguca/templates/forms.qtpl:77
func captcha(board string) string {
	//line server/src/meguca/templates/forms.qtpl:77
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:77
	writecaptcha(qb422016, board)
	//line server/src/meguca/templates/forms.qtpl:77
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:77
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:77
	return qs422016
//line server/src/meguca/templates/forms.qtpl:77
}

// Form for inputting key-value map-like data

//line server/src/meguca/templates/forms.qtpl:80
func streamkeyValueForm(qw422016 *qt422016.Writer, k, v string) {
	//line server/src/meguca/templates/forms.qtpl:80
	qw422016.N().S(`<span><input type="text" class="map-field" value="`)
	//line server/src/meguca/templates/forms.qtpl:82
	qw422016.E().S(k)
	//line server/src/meguca/templates/forms.qtpl:82
	qw422016.N().S(`"><input type="text" class="map-field" value="`)
	//line server/src/meguca/templates/forms.qtpl:83
	qw422016.E().S(v)
	//line server/src/meguca/templates/forms.qtpl:83
	qw422016.N().S(`"><a class="map-remove">[X]</a><br></span>`)
//line server/src/meguca/templates/forms.qtpl:89
}

//line server/src/meguca/templates/forms.qtpl:89
func writekeyValueForm(qq422016 qtio422016.Writer, k, v string) {
	//line server/src/meguca/templates/forms.qtpl:89
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:89
	streamkeyValueForm(qw422016, k, v)
	//line server/src/meguca/templates/forms.qtpl:89
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:89
}

//line server/src/meguca/templates/forms.qtpl:89
func keyValueForm(k, v string) string {
	//line server/src/meguca/templates/forms.qtpl:89
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:89
	writekeyValueForm(qb422016, k, v)
	//line server/src/meguca/templates/forms.qtpl:89
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:89
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:89
	return qs422016
//line server/src/meguca/templates/forms.qtpl:89
}

// Form for inputting one array-like form item

//line server/src/meguca/templates/forms.qtpl:92
func streamarrayItemForm(qw422016 *qt422016.Writer, v string) {
	//line server/src/meguca/templates/forms.qtpl:92
	qw422016.N().S(`<span><input type="text" class="array-field" value="`)
	//line server/src/meguca/templates/forms.qtpl:94
	qw422016.E().S(v)
	//line server/src/meguca/templates/forms.qtpl:94
	qw422016.N().S(`"><a class="array-remove">[X]</a><br></span>`)
//line server/src/meguca/templates/forms.qtpl:100
}

//line server/src/meguca/templates/forms.qtpl:100
func writearrayItemForm(qq422016 qtio422016.Writer, v string) {
	//line server/src/meguca/templates/forms.qtpl:100
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:100
	streamarrayItemForm(qw422016, v)
	//line server/src/meguca/templates/forms.qtpl:100
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:100
}

//line server/src/meguca/templates/forms.qtpl:100
func arrayItemForm(v string) string {
	//line server/src/meguca/templates/forms.qtpl:100
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:100
	writearrayItemForm(qb422016, v)
	//line server/src/meguca/templates/forms.qtpl:100
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:100
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:100
	return qs422016
//line server/src/meguca/templates/forms.qtpl:100
}

// Form formatted as a table, with cancel and submit buttons

//line server/src/meguca/templates/forms.qtpl:103
func streamtableForm(qw422016 *qt422016.Writer, specs []inputSpec, needCaptcha bool) {
	//line server/src/meguca/templates/forms.qtpl:104
	streamtable(qw422016, specs)
	//line server/src/meguca/templates/forms.qtpl:105
	if needCaptcha {
		//line server/src/meguca/templates/forms.qtpl:106
		streamcaptcha(qw422016, "all")
		//line server/src/meguca/templates/forms.qtpl:107
	}
	//line server/src/meguca/templates/forms.qtpl:108
	streamsubmit(qw422016, true)
//line server/src/meguca/templates/forms.qtpl:109
}

//line server/src/meguca/templates/forms.qtpl:109
func writetableForm(qq422016 qtio422016.Writer, specs []inputSpec, needCaptcha bool) {
	//line server/src/meguca/templates/forms.qtpl:109
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:109
	streamtableForm(qw422016, specs, needCaptcha)
	//line server/src/meguca/templates/forms.qtpl:109
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:109
}

//line server/src/meguca/templates/forms.qtpl:109
func tableForm(specs []inputSpec, needCaptcha bool) string {
	//line server/src/meguca/templates/forms.qtpl:109
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:109
	writetableForm(qb422016, specs, needCaptcha)
	//line server/src/meguca/templates/forms.qtpl:109
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:109
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:109
	return qs422016
//line server/src/meguca/templates/forms.qtpl:109
}

// Render a map form for inputting map-like data

//line server/src/meguca/templates/forms.qtpl:112
func streamrenderMap(qw422016 *qt422016.Writer, spec inputSpec) {
	//line server/src/meguca/templates/forms.qtpl:113
	ln := lang.Get()

	//line server/src/meguca/templates/forms.qtpl:113
	qw422016.N().S(`<div class="map-form" name="`)
	//line server/src/meguca/templates/forms.qtpl:114
	qw422016.N().S(spec.ID)
	//line server/src/meguca/templates/forms.qtpl:114
	qw422016.N().S(`" title="`)
	//line server/src/meguca/templates/forms.qtpl:114
	qw422016.N().S(ln.Forms[spec.ID][1])
	//line server/src/meguca/templates/forms.qtpl:114
	qw422016.N().S(`">`)
	//line server/src/meguca/templates/forms.qtpl:115
	for k, v := range spec.Val.(map[string]string) {
		//line server/src/meguca/templates/forms.qtpl:116
		streamkeyValueForm(qw422016, k, v)
		//line server/src/meguca/templates/forms.qtpl:117
	}
	//line server/src/meguca/templates/forms.qtpl:117
	qw422016.N().S(`<a class="map-add">`)
	//line server/src/meguca/templates/forms.qtpl:119
	qw422016.N().S(ln.UI["add"])
	//line server/src/meguca/templates/forms.qtpl:119
	qw422016.N().S(`</a><br></div>`)
//line server/src/meguca/templates/forms.qtpl:123
}

//line server/src/meguca/templates/forms.qtpl:123
func writerenderMap(qq422016 qtio422016.Writer, spec inputSpec) {
	//line server/src/meguca/templates/forms.qtpl:123
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:123
	streamrenderMap(qw422016, spec)
	//line server/src/meguca/templates/forms.qtpl:123
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:123
}

//line server/src/meguca/templates/forms.qtpl:123
func renderMap(spec inputSpec) string {
	//line server/src/meguca/templates/forms.qtpl:123
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:123
	writerenderMap(qb422016, spec)
	//line server/src/meguca/templates/forms.qtpl:123
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:123
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:123
	return qs422016
//line server/src/meguca/templates/forms.qtpl:123
}

// Render form for inputting array-like data

//line server/src/meguca/templates/forms.qtpl:126
func streamrenderArray(qw422016 *qt422016.Writer, spec inputSpec) {
	//line server/src/meguca/templates/forms.qtpl:127
	ln := lang.Get()

	//line server/src/meguca/templates/forms.qtpl:127
	qw422016.N().S(`<div class="array-form" name="`)
	//line server/src/meguca/templates/forms.qtpl:128
	qw422016.N().S(spec.ID)
	//line server/src/meguca/templates/forms.qtpl:128
	qw422016.N().S(`" title="`)
	//line server/src/meguca/templates/forms.qtpl:128
	qw422016.N().S(ln.Forms[spec.ID][1])
	//line server/src/meguca/templates/forms.qtpl:128
	qw422016.N().S(`">`)
	//line server/src/meguca/templates/forms.qtpl:129
	for _, v := range spec.Val.([]string) {
		//line server/src/meguca/templates/forms.qtpl:130
		streamarrayItemForm(qw422016, v)
		//line server/src/meguca/templates/forms.qtpl:131
	}
	//line server/src/meguca/templates/forms.qtpl:131
	qw422016.N().S(`<a class="array-add">`)
	//line server/src/meguca/templates/forms.qtpl:133
	qw422016.N().S(ln.UI["add"])
	//line server/src/meguca/templates/forms.qtpl:133
	qw422016.N().S(`</a><br></div>`)
//line server/src/meguca/templates/forms.qtpl:137
}

//line server/src/meguca/templates/forms.qtpl:137
func writerenderArray(qq422016 qtio422016.Writer, spec inputSpec) {
	//line server/src/meguca/templates/forms.qtpl:137
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:137
	streamrenderArray(qw422016, spec)
	//line server/src/meguca/templates/forms.qtpl:137
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:137
}

//line server/src/meguca/templates/forms.qtpl:137
func renderArray(spec inputSpec) string {
	//line server/src/meguca/templates/forms.qtpl:137
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:137
	writerenderArray(qb422016, spec)
	//line server/src/meguca/templates/forms.qtpl:137
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:137
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:137
	return qs422016
//line server/src/meguca/templates/forms.qtpl:137
}

// Render submit and cancel buttons

//line server/src/meguca/templates/forms.qtpl:140
func streamsubmit(qw422016 *qt422016.Writer, cancel bool) {
	//line server/src/meguca/templates/forms.qtpl:140
	qw422016.N().S(`<input type="submit" value="`)
	//line server/src/meguca/templates/forms.qtpl:141
	qw422016.N().S(lang.Get().Common.UI["submit"])
	//line server/src/meguca/templates/forms.qtpl:141
	qw422016.N().S(`">`)
	//line server/src/meguca/templates/forms.qtpl:142
	if cancel {
		//line server/src/meguca/templates/forms.qtpl:143
		streamcancel(qw422016)
		//line server/src/meguca/templates/forms.qtpl:144
	}
	//line server/src/meguca/templates/forms.qtpl:144
	qw422016.N().S(`<div class="form-response admin"></div>`)
//line server/src/meguca/templates/forms.qtpl:146
}

//line server/src/meguca/templates/forms.qtpl:146
func writesubmit(qq422016 qtio422016.Writer, cancel bool) {
	//line server/src/meguca/templates/forms.qtpl:146
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:146
	streamsubmit(qw422016, cancel)
	//line server/src/meguca/templates/forms.qtpl:146
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:146
}

//line server/src/meguca/templates/forms.qtpl:146
func submit(cancel bool) string {
	//line server/src/meguca/templates/forms.qtpl:146
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:146
	writesubmit(qb422016, cancel)
	//line server/src/meguca/templates/forms.qtpl:146
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:146
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:146
	return qs422016
//line server/src/meguca/templates/forms.qtpl:146
}

// Renders a cancel button

//line server/src/meguca/templates/forms.qtpl:149
func streamcancel(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:149
	qw422016.N().S(`<input type="button" name="cancel" value="`)
	//line server/src/meguca/templates/forms.qtpl:150
	qw422016.N().S(lang.Get().Common.UI["cancel"])
	//line server/src/meguca/templates/forms.qtpl:150
	qw422016.N().S(`">`)
//line server/src/meguca/templates/forms.qtpl:151
}

//line server/src/meguca/templates/forms.qtpl:151
func writecancel(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:151
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:151
	streamcancel(qw422016)
	//line server/src/meguca/templates/forms.qtpl:151
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:151
}

//line server/src/meguca/templates/forms.qtpl:151
func cancel() string {
	//line server/src/meguca/templates/forms.qtpl:151
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:151
	writecancel(qb422016)
	//line server/src/meguca/templates/forms.qtpl:151
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:151
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:151
	return qs422016
//line server/src/meguca/templates/forms.qtpl:151
}

// Render link to request new noscript captcha

//line server/src/meguca/templates/forms.qtpl:154
func StreamNoscriptCaptchaLink(qw422016 *qt422016.Writer, board string) {
	//line server/src/meguca/templates/forms.qtpl:154
	qw422016.N().S(`<a href="/api/captcha/`)
	//line server/src/meguca/templates/forms.qtpl:155
	qw422016.E().S(board)
	//line server/src/meguca/templates/forms.qtpl:155
	qw422016.N().S(`" style="display: flex; width: 100%; height: 100%;"><span style="align-self: center; margin: auto;">`)
	//line server/src/meguca/templates/forms.qtpl:157
	qw422016.N().S(lang.Get().UI["loadCaptcha"])
	//line server/src/meguca/templates/forms.qtpl:157
	qw422016.N().S(`</span></a>`)
//line server/src/meguca/templates/forms.qtpl:160
}

//line server/src/meguca/templates/forms.qtpl:160
func WriteNoscriptCaptchaLink(qq422016 qtio422016.Writer, board string) {
	//line server/src/meguca/templates/forms.qtpl:160
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:160
	StreamNoscriptCaptchaLink(qw422016, board)
	//line server/src/meguca/templates/forms.qtpl:160
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:160
}

//line server/src/meguca/templates/forms.qtpl:160
func NoscriptCaptchaLink(board string) string {
	//line server/src/meguca/templates/forms.qtpl:160
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:160
	WriteNoscriptCaptchaLink(qb422016, board)
	//line server/src/meguca/templates/forms.qtpl:160
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:160
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:160
	return qs422016
//line server/src/meguca/templates/forms.qtpl:160
}

//line server/src/meguca/templates/forms.qtpl:162
func StreamBannerForm(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:162
	qw422016.N().S(`<div style="white-space: normal;">`)
	//line server/src/meguca/templates/forms.qtpl:164
	qw422016.N().S(lang.Get().UI["bannerSpecs"])
	//line server/src/meguca/templates/forms.qtpl:164
	qw422016.N().S(`</div><br><input type="file" name="banners" multiple accept="image/png, image/gif, image/jpeg, video/webm"><br>`)
	//line server/src/meguca/templates/forms.qtpl:169
	streamcaptcha(qw422016, "all")
	//line server/src/meguca/templates/forms.qtpl:170
	streamsubmit(qw422016, true)
//line server/src/meguca/templates/forms.qtpl:171
}

//line server/src/meguca/templates/forms.qtpl:171
func WriteBannerForm(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:171
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:171
	StreamBannerForm(qw422016)
	//line server/src/meguca/templates/forms.qtpl:171
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:171
}

//line server/src/meguca/templates/forms.qtpl:171
func BannerForm() string {
	//line server/src/meguca/templates/forms.qtpl:171
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:171
	WriteBannerForm(qb422016)
	//line server/src/meguca/templates/forms.qtpl:171
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:171
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:171
	return qs422016
//line server/src/meguca/templates/forms.qtpl:171
}

//line server/src/meguca/templates/forms.qtpl:173
func StreamLoadingAnimationForm(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:173
	qw422016.N().S(`<div style="white-space: normal;">`)
	//line server/src/meguca/templates/forms.qtpl:175
	qw422016.N().S(lang.Get().UI["loadingSpecs"])
	//line server/src/meguca/templates/forms.qtpl:175
	qw422016.N().S(`</div><br><input type="file" name="image" accept="image/gif, video/webm"><br>`)
	//line server/src/meguca/templates/forms.qtpl:180
	streamcaptcha(qw422016, "all")
	//line server/src/meguca/templates/forms.qtpl:181
	streamsubmit(qw422016, true)
//line server/src/meguca/templates/forms.qtpl:182
}

//line server/src/meguca/templates/forms.qtpl:182
func WriteLoadingAnimationForm(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/forms.qtpl:182
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/forms.qtpl:182
	StreamLoadingAnimationForm(qw422016)
	//line server/src/meguca/templates/forms.qtpl:182
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/forms.qtpl:182
}

//line server/src/meguca/templates/forms.qtpl:182
func LoadingAnimationForm() string {
	//line server/src/meguca/templates/forms.qtpl:182
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/forms.qtpl:182
	WriteLoadingAnimationForm(qb422016)
	//line server/src/meguca/templates/forms.qtpl:182
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/forms.qtpl:182
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/forms.qtpl:182
	return qs422016
//line server/src/meguca/templates/forms.qtpl:182
}
