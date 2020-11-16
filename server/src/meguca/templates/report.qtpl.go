// Code generated by qtc from "report.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line server/src/meguca/templates/report.qtpl:1
package templates

//line server/src/meguca/templates/report.qtpl:1
import "strconv"

//line server/src/meguca/templates/report.qtpl:2
import "time"

//line server/src/meguca/templates/report.qtpl:3
import "meguca/lang"

//line server/src/meguca/templates/report.qtpl:4
import "meguca/auth"

//line server/src/meguca/templates/report.qtpl:5
import "meguca/common"

// Report submission form

//line server/src/meguca/templates/report.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line server/src/meguca/templates/report.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line server/src/meguca/templates/report.qtpl:8
func StreamReportForm(qw422016 *qt422016.Writer, id uint64) {
	//line server/src/meguca/templates/report.qtpl:9
	ln := lang.Get().UI

	//line server/src/meguca/templates/report.qtpl:9
	qw422016.N().S(`<input type=text name=target value="`)
	//line server/src/meguca/templates/report.qtpl:10
	qw422016.N().S(strconv.FormatUint(id, 10))
	//line server/src/meguca/templates/report.qtpl:10
	qw422016.N().S(`" hidden><input type=text name=reason placeholder="`)
	//line server/src/meguca/templates/report.qtpl:11
	qw422016.N().S(ln["reason"])
	//line server/src/meguca/templates/report.qtpl:11
	qw422016.N().S(`" maxlength="`)
	//line server/src/meguca/templates/report.qtpl:11
	qw422016.N().D(common.MaxLenReason)
	//line server/src/meguca/templates/report.qtpl:11
	qw422016.N().S(`"><br><label><input type=checkbox name=illegal>`)
	//line server/src/meguca/templates/report.qtpl:15
	qw422016.N().S(ln["illegal"])
	//line server/src/meguca/templates/report.qtpl:15
	qw422016.N().S(`<br></label>`)
	//line server/src/meguca/templates/report.qtpl:18
	streamcaptcha(qw422016, "all")
	//line server/src/meguca/templates/report.qtpl:19
	streamsubmit(qw422016, true)
//line server/src/meguca/templates/report.qtpl:20
}

//line server/src/meguca/templates/report.qtpl:20
func WriteReportForm(qq422016 qtio422016.Writer, id uint64) {
	//line server/src/meguca/templates/report.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/report.qtpl:20
	StreamReportForm(qw422016, id)
	//line server/src/meguca/templates/report.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/report.qtpl:20
}

//line server/src/meguca/templates/report.qtpl:20
func ReportForm(id uint64) string {
	//line server/src/meguca/templates/report.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/report.qtpl:20
	WriteReportForm(qb422016, id)
	//line server/src/meguca/templates/report.qtpl:20
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/report.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/report.qtpl:20
	return qs422016
//line server/src/meguca/templates/report.qtpl:20
}

// Render list of all reports on board

//line server/src/meguca/templates/report.qtpl:23
func StreamReportList(qw422016 *qt422016.Writer, reports []auth.Report) {
	//line server/src/meguca/templates/report.qtpl:24
	streamtableStyle(qw422016)
	//line server/src/meguca/templates/report.qtpl:24
	qw422016.N().S(`<table>`)
	//line server/src/meguca/templates/report.qtpl:26
	streamtableHeaders(qw422016, "id", "post", "reason", "time")
	//line server/src/meguca/templates/report.qtpl:27
	for _, r := range reports {
		//line server/src/meguca/templates/report.qtpl:27
		qw422016.N().S(`<tr><td>`)
		//line server/src/meguca/templates/report.qtpl:29
		qw422016.N().S(strconv.FormatUint(r.ID, 10))
		//line server/src/meguca/templates/report.qtpl:29
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/report.qtpl:30
		streamstaticPostLink(qw422016, r.Target)
		//line server/src/meguca/templates/report.qtpl:30
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/report.qtpl:31
		qw422016.E().S(r.Reason)
		//line server/src/meguca/templates/report.qtpl:31
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/report.qtpl:32
		qw422016.E().S(r.Created.Format(time.UnixDate))
		//line server/src/meguca/templates/report.qtpl:32
		qw422016.N().S(`</td></tr>`)
		//line server/src/meguca/templates/report.qtpl:34
	}
	//line server/src/meguca/templates/report.qtpl:34
	qw422016.N().S(`</table>`)
//line server/src/meguca/templates/report.qtpl:36
}

//line server/src/meguca/templates/report.qtpl:36
func WriteReportList(qq422016 qtio422016.Writer, reports []auth.Report) {
	//line server/src/meguca/templates/report.qtpl:36
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/report.qtpl:36
	StreamReportList(qw422016, reports)
	//line server/src/meguca/templates/report.qtpl:36
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/report.qtpl:36
}

//line server/src/meguca/templates/report.qtpl:36
func ReportList(reports []auth.Report) string {
	//line server/src/meguca/templates/report.qtpl:36
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/report.qtpl:36
	WriteReportList(qb422016, reports)
	//line server/src/meguca/templates/report.qtpl:36
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/report.qtpl:36
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/report.qtpl:36
	return qs422016
//line server/src/meguca/templates/report.qtpl:36
}
