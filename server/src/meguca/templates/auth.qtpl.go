// Code generated by qtc from "auth.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line server/src/meguca/templates/auth.qtpl:1
package templates

//line server/src/meguca/templates/auth.qtpl:1
import "fmt"

//line server/src/meguca/templates/auth.qtpl:2
import "time"

//line server/src/meguca/templates/auth.qtpl:3
import "strconv"

//line server/src/meguca/templates/auth.qtpl:4
import "meguca/auth"

//line server/src/meguca/templates/auth.qtpl:5
import "meguca/config"

//line server/src/meguca/templates/auth.qtpl:6
import "meguca/lang"

//line server/src/meguca/templates/auth.qtpl:7
import "meguca/common"

//line server/src/meguca/templates/auth.qtpl:8
import "github.com/bakape/mnemonics"

// Header of a standalone HTML page

//line server/src/meguca/templates/auth.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line server/src/meguca/templates/auth.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line server/src/meguca/templates/auth.qtpl:11
func streamhtmlHeader(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/auth.qtpl:11
	qw422016.N().S(`<!DOCTYPE html><html><head><meta charset="utf-8"/></head><body>`)
//line server/src/meguca/templates/auth.qtpl:18
}

//line server/src/meguca/templates/auth.qtpl:18
func writehtmlHeader(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/auth.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/auth.qtpl:18
	streamhtmlHeader(qw422016)
	//line server/src/meguca/templates/auth.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/auth.qtpl:18
}

//line server/src/meguca/templates/auth.qtpl:18
func htmlHeader() string {
	//line server/src/meguca/templates/auth.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/auth.qtpl:18
	writehtmlHeader(qb422016)
	//line server/src/meguca/templates/auth.qtpl:18
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/auth.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/auth.qtpl:18
	return qs422016
//line server/src/meguca/templates/auth.qtpl:18
}

// End of a standalone HTML page

//line server/src/meguca/templates/auth.qtpl:21
func streamhtmlEnd(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/auth.qtpl:21
	qw422016.N().S(`</body></html>`)
//line server/src/meguca/templates/auth.qtpl:24
}

//line server/src/meguca/templates/auth.qtpl:24
func writehtmlEnd(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/auth.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/auth.qtpl:24
	streamhtmlEnd(qw422016)
	//line server/src/meguca/templates/auth.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/auth.qtpl:24
}

//line server/src/meguca/templates/auth.qtpl:24
func htmlEnd() string {
	//line server/src/meguca/templates/auth.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/auth.qtpl:24
	writehtmlEnd(qb422016)
	//line server/src/meguca/templates/auth.qtpl:24
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/auth.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/auth.qtpl:24
	return qs422016
//line server/src/meguca/templates/auth.qtpl:24
}

// BanPage renders a ban page for a banned user

//line server/src/meguca/templates/auth.qtpl:27
func StreamBanPage(qw422016 *qt422016.Writer, rec auth.BanRecord) {
	//line server/src/meguca/templates/auth.qtpl:28
	streamhtmlHeader(qw422016)
	//line server/src/meguca/templates/auth.qtpl:29
	ln := lang.Get().Templates["banPage"]

	//line server/src/meguca/templates/auth.qtpl:30
	if len(ln) < 3 {
		//line server/src/meguca/templates/auth.qtpl:31
		panic(fmt.Errorf("invalid ban format strings: %v", ln))

		//line server/src/meguca/templates/auth.qtpl:32
	}
	//line server/src/meguca/templates/auth.qtpl:32
	qw422016.N().S(`<div class="ban-page glass">`)
	//line server/src/meguca/templates/auth.qtpl:34
	qw422016.N().S(fmt.Sprintf(ln[0], bold(rec.Board), bold(rec.By)))
	//line server/src/meguca/templates/auth.qtpl:34
	qw422016.N().S(`<br><br><b>`)
	//line server/src/meguca/templates/auth.qtpl:38
	qw422016.E().S(rec.Reason)
	//line server/src/meguca/templates/auth.qtpl:38
	qw422016.N().S(`</b><br><br>`)
	//line server/src/meguca/templates/auth.qtpl:42
	exp := rec.Expires.Round(time.Second)

	//line server/src/meguca/templates/auth.qtpl:43
	date := exp.Format(time.UnixDate)

	//line server/src/meguca/templates/auth.qtpl:44
	till := exp.Sub(time.Now().Round(time.Second)).String()

	//line server/src/meguca/templates/auth.qtpl:45
	qw422016.N().S(fmt.Sprintf(ln[1], bold(date), bold(till)))
	//line server/src/meguca/templates/auth.qtpl:45
	qw422016.N().S(`<br><br>`)
	//line server/src/meguca/templates/auth.qtpl:48
	qw422016.N().S(fmt.Sprintf(ln[2], bold(rec.IP)))
	//line server/src/meguca/templates/auth.qtpl:48
	qw422016.N().S(`<br></div>`)
	//line server/src/meguca/templates/auth.qtpl:51
	streamhtmlEnd(qw422016)
//line server/src/meguca/templates/auth.qtpl:52
}

//line server/src/meguca/templates/auth.qtpl:52
func WriteBanPage(qq422016 qtio422016.Writer, rec auth.BanRecord) {
	//line server/src/meguca/templates/auth.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/auth.qtpl:52
	StreamBanPage(qw422016, rec)
	//line server/src/meguca/templates/auth.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/auth.qtpl:52
}

//line server/src/meguca/templates/auth.qtpl:52
func BanPage(rec auth.BanRecord) string {
	//line server/src/meguca/templates/auth.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/auth.qtpl:52
	WriteBanPage(qb422016, rec)
	//line server/src/meguca/templates/auth.qtpl:52
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/auth.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/auth.qtpl:52
	return qs422016
//line server/src/meguca/templates/auth.qtpl:52
}

// Renders a list of bans for a specific page with optional unbanning API links

//line server/src/meguca/templates/auth.qtpl:55
func StreamBanList(qw422016 *qt422016.Writer, bans []auth.BanRecord, board string, canUnban bool) {
	//line server/src/meguca/templates/auth.qtpl:56
	streamhtmlHeader(qw422016)
	//line server/src/meguca/templates/auth.qtpl:57
	streamtableStyle(qw422016)
	//line server/src/meguca/templates/auth.qtpl:57
	qw422016.N().S(`<form method="post" action="/api/unban/`)
	//line server/src/meguca/templates/auth.qtpl:58
	qw422016.N().S(board)
	//line server/src/meguca/templates/auth.qtpl:58
	qw422016.N().S(`"><table>`)
	//line server/src/meguca/templates/auth.qtpl:60
	headers := []string{
		"reason", "by", "post", "posterID", "expires",
	}

	//line server/src/meguca/templates/auth.qtpl:63
	if canUnban {
		//line server/src/meguca/templates/auth.qtpl:64
		headers = append(headers, "unban")

		//line server/src/meguca/templates/auth.qtpl:65
	}
	//line server/src/meguca/templates/auth.qtpl:66
	streamtableHeaders(qw422016, headers...)
	//line server/src/meguca/templates/auth.qtpl:67
	salt := config.Get().Salt

	//line server/src/meguca/templates/auth.qtpl:68
	for _, b := range bans {
		//line server/src/meguca/templates/auth.qtpl:68
		qw422016.N().S(`<tr><td>`)
		//line server/src/meguca/templates/auth.qtpl:70
		qw422016.E().S(b.Reason)
		//line server/src/meguca/templates/auth.qtpl:70
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:71
		qw422016.E().S(b.By)
		//line server/src/meguca/templates/auth.qtpl:71
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:72
		streamstaticPostLink(qw422016, b.ForPost)
		//line server/src/meguca/templates/auth.qtpl:72
		qw422016.N().S(`</td>`)
		//line server/src/meguca/templates/auth.qtpl:73
		buf := make([]byte, 0, len(salt)+len(b.IP))

		//line server/src/meguca/templates/auth.qtpl:74
		buf = append(buf, salt...)

		//line server/src/meguca/templates/auth.qtpl:75
		buf = append(buf, b.IP...)

		//line server/src/meguca/templates/auth.qtpl:75
		qw422016.N().S(`<td>`)
		//line server/src/meguca/templates/auth.qtpl:76
		qw422016.E().S(mnemonic.FantasyName(buf))
		//line server/src/meguca/templates/auth.qtpl:76
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:77
		qw422016.E().S(b.Expires.Format(time.UnixDate))
		//line server/src/meguca/templates/auth.qtpl:77
		qw422016.N().S(`</td>`)
		//line server/src/meguca/templates/auth.qtpl:78
		if canUnban {
			//line server/src/meguca/templates/auth.qtpl:78
			qw422016.N().S(`<td><input type="checkbox" name="`)
			//line server/src/meguca/templates/auth.qtpl:80
			qw422016.E().S(strconv.FormatUint(b.ForPost, 10))
			//line server/src/meguca/templates/auth.qtpl:80
			qw422016.N().S(`"></td>`)
			//line server/src/meguca/templates/auth.qtpl:82
		}
		//line server/src/meguca/templates/auth.qtpl:82
		qw422016.N().S(`</tr>`)
		//line server/src/meguca/templates/auth.qtpl:84
	}
	//line server/src/meguca/templates/auth.qtpl:84
	qw422016.N().S(`</table>`)
	//line server/src/meguca/templates/auth.qtpl:86
	if canUnban {
		//line server/src/meguca/templates/auth.qtpl:87
		streamsubmit(qw422016, false)
		//line server/src/meguca/templates/auth.qtpl:88
	}
	//line server/src/meguca/templates/auth.qtpl:88
	qw422016.N().S(`</form>`)
	//line server/src/meguca/templates/auth.qtpl:90
	streamhtmlEnd(qw422016)
//line server/src/meguca/templates/auth.qtpl:91
}

//line server/src/meguca/templates/auth.qtpl:91
func WriteBanList(qq422016 qtio422016.Writer, bans []auth.BanRecord, board string, canUnban bool) {
	//line server/src/meguca/templates/auth.qtpl:91
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/auth.qtpl:91
	StreamBanList(qw422016, bans, board, canUnban)
	//line server/src/meguca/templates/auth.qtpl:91
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/auth.qtpl:91
}

//line server/src/meguca/templates/auth.qtpl:91
func BanList(bans []auth.BanRecord, board string, canUnban bool) string {
	//line server/src/meguca/templates/auth.qtpl:91
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/auth.qtpl:91
	WriteBanList(qb422016, bans, board, canUnban)
	//line server/src/meguca/templates/auth.qtpl:91
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/auth.qtpl:91
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/auth.qtpl:91
	return qs422016
//line server/src/meguca/templates/auth.qtpl:91
}

// Common style for plain html tables

//line server/src/meguca/templates/auth.qtpl:94
func streamtableStyle(qw422016 *qt422016.Writer) {
	//line server/src/meguca/templates/auth.qtpl:94
	qw422016.N().S(`<style>table, th, td {border: 1px solid black;}.hash-link {display: none;}</style>`)
//line server/src/meguca/templates/auth.qtpl:103
}

//line server/src/meguca/templates/auth.qtpl:103
func writetableStyle(qq422016 qtio422016.Writer) {
	//line server/src/meguca/templates/auth.qtpl:103
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/auth.qtpl:103
	streamtableStyle(qw422016)
	//line server/src/meguca/templates/auth.qtpl:103
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/auth.qtpl:103
}

//line server/src/meguca/templates/auth.qtpl:103
func tableStyle() string {
	//line server/src/meguca/templates/auth.qtpl:103
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/auth.qtpl:103
	writetableStyle(qb422016)
	//line server/src/meguca/templates/auth.qtpl:103
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/auth.qtpl:103
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/auth.qtpl:103
	return qs422016
//line server/src/meguca/templates/auth.qtpl:103
}

// Post link, that will redirect to the post from any page

//line server/src/meguca/templates/auth.qtpl:106
func streamstaticPostLink(qw422016 *qt422016.Writer, id uint64) {
	//line server/src/meguca/templates/auth.qtpl:107
	streampostLink(qw422016, common.Link{id, id, "all"}, true, true)
//line server/src/meguca/templates/auth.qtpl:108
}

//line server/src/meguca/templates/auth.qtpl:108
func writestaticPostLink(qq422016 qtio422016.Writer, id uint64) {
	//line server/src/meguca/templates/auth.qtpl:108
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/auth.qtpl:108
	streamstaticPostLink(qw422016, id)
	//line server/src/meguca/templates/auth.qtpl:108
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/auth.qtpl:108
}

//line server/src/meguca/templates/auth.qtpl:108
func staticPostLink(id uint64) string {
	//line server/src/meguca/templates/auth.qtpl:108
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/auth.qtpl:108
	writestaticPostLink(qb422016, id)
	//line server/src/meguca/templates/auth.qtpl:108
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/auth.qtpl:108
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/auth.qtpl:108
	return qs422016
//line server/src/meguca/templates/auth.qtpl:108
}

// Renders a moderation log page

//line server/src/meguca/templates/auth.qtpl:111
func StreamModLog(qw422016 *qt422016.Writer, log []auth.ModLogEntry) {
	//line server/src/meguca/templates/auth.qtpl:112
	streamhtmlHeader(qw422016)
	//line server/src/meguca/templates/auth.qtpl:113
	ln := lang.Get()

	//line server/src/meguca/templates/auth.qtpl:114
	streamtableStyle(qw422016)
	//line server/src/meguca/templates/auth.qtpl:114
	qw422016.N().S(`<table>`)
	//line server/src/meguca/templates/auth.qtpl:116
	streamtableHeaders(qw422016, "type", "by", "post", "time", "data", "duration")
	//line server/src/meguca/templates/auth.qtpl:117
	for _, l := range log {
		//line server/src/meguca/templates/auth.qtpl:117
		qw422016.N().S(`<tr><td>`)
		//line server/src/meguca/templates/auth.qtpl:120
		switch l.Type {
		//line server/src/meguca/templates/auth.qtpl:121
		case common.BanPost:
			//line server/src/meguca/templates/auth.qtpl:122
			qw422016.E().S(ln.UI["ban"])
		//line server/src/meguca/templates/auth.qtpl:123
		case common.UnbanPost:
			//line server/src/meguca/templates/auth.qtpl:124
			qw422016.E().S(ln.UI["unban"])
		//line server/src/meguca/templates/auth.qtpl:125
		case common.DeletePost:
			//line server/src/meguca/templates/auth.qtpl:126
			qw422016.E().S(ln.UI["deletePost"])
		//line server/src/meguca/templates/auth.qtpl:127
		case common.DeleteImage:
			//line server/src/meguca/templates/auth.qtpl:128
			qw422016.E().S(ln.UI["deleteImage"])
		//line server/src/meguca/templates/auth.qtpl:129
		case common.SpoilerImage:
			//line server/src/meguca/templates/auth.qtpl:130
			qw422016.E().S(ln.UI["spoilerImage"])
		//line server/src/meguca/templates/auth.qtpl:131
		case common.LockThread:
			//line server/src/meguca/templates/auth.qtpl:132
			qw422016.E().S(ln.Common.UI["lockThread"])
		//line server/src/meguca/templates/auth.qtpl:133
		case common.DeleteBoard:
			//line server/src/meguca/templates/auth.qtpl:134
			qw422016.E().S(ln.Common.UI["deleteBoard"])
		//line server/src/meguca/templates/auth.qtpl:135
		case common.MeidoVision:
			//line server/src/meguca/templates/auth.qtpl:136
			qw422016.E().S(ln.Common.UI["meidoVisionPost"])
		//line server/src/meguca/templates/auth.qtpl:137
		case common.PurgePost:
			//line server/src/meguca/templates/auth.qtpl:138
			qw422016.E().S(ln.UI["purgePost"])
			//line server/src/meguca/templates/auth.qtpl:139
		}
		//line server/src/meguca/templates/auth.qtpl:139
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:141
		qw422016.E().S(l.By)
		//line server/src/meguca/templates/auth.qtpl:141
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:143
		if l.ID != 0 {
			//line server/src/meguca/templates/auth.qtpl:144
			streamstaticPostLink(qw422016, l.ID)
			//line server/src/meguca/templates/auth.qtpl:145
		}
		//line server/src/meguca/templates/auth.qtpl:145
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:147
		qw422016.E().S(l.Created.Format(time.UnixDate))
		//line server/src/meguca/templates/auth.qtpl:147
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:148
		qw422016.E().S(l.Data)
		//line server/src/meguca/templates/auth.qtpl:148
		qw422016.N().S(`</td><td>`)
		//line server/src/meguca/templates/auth.qtpl:150
		if l.Length != 0 {
			//line server/src/meguca/templates/auth.qtpl:151
			qw422016.E().S((time.Second * time.Duration(l.Length)).String())
			//line server/src/meguca/templates/auth.qtpl:152
		}
		//line server/src/meguca/templates/auth.qtpl:152
		qw422016.N().S(`</td></tr>`)
		//line server/src/meguca/templates/auth.qtpl:155
	}
	//line server/src/meguca/templates/auth.qtpl:155
	qw422016.N().S(`</table>`)
	//line server/src/meguca/templates/auth.qtpl:157
	streamhtmlEnd(qw422016)
//line server/src/meguca/templates/auth.qtpl:158
}

//line server/src/meguca/templates/auth.qtpl:158
func WriteModLog(qq422016 qtio422016.Writer, log []auth.ModLogEntry) {
	//line server/src/meguca/templates/auth.qtpl:158
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line server/src/meguca/templates/auth.qtpl:158
	StreamModLog(qw422016, log)
	//line server/src/meguca/templates/auth.qtpl:158
	qt422016.ReleaseWriter(qw422016)
//line server/src/meguca/templates/auth.qtpl:158
}

//line server/src/meguca/templates/auth.qtpl:158
func ModLog(log []auth.ModLogEntry) string {
	//line server/src/meguca/templates/auth.qtpl:158
	qb422016 := qt422016.AcquireByteBuffer()
	//line server/src/meguca/templates/auth.qtpl:158
	WriteModLog(qb422016, log)
	//line server/src/meguca/templates/auth.qtpl:158
	qs422016 := string(qb422016.B)
	//line server/src/meguca/templates/auth.qtpl:158
	qt422016.ReleaseByteBuffer(qb422016)
	//line server/src/meguca/templates/auth.qtpl:158
	return qs422016
//line server/src/meguca/templates/auth.qtpl:158
}