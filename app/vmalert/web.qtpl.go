// Code generated by qtc from "web.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line web.qtpl:1
package main

//line web.qtpl:3
import (
	"path"
	"sort"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/notifier"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/tpl"
)

//line web.qtpl:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line web.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line web.qtpl:13
func StreamWelcome(qw422016 *qt422016.Writer) {
//line web.qtpl:13
	qw422016.N().S(`
    `)
//line web.qtpl:14
	tpl.StreamHeader(qw422016, "vmalert", navItems)
//line web.qtpl:14
	qw422016.N().S(`
    <p>
        API:<br>
        `)
//line web.qtpl:17
	for _, p := range apiLinks {
//line web.qtpl:17
		qw422016.N().S(`
            `)
//line web.qtpl:19
		p, doc := p[0], p[1]

//line web.qtpl:20
		qw422016.N().S(`
        	<a href="`)
//line web.qtpl:21
		qw422016.E().S(p)
//line web.qtpl:21
		qw422016.N().S(`">`)
//line web.qtpl:21
		qw422016.E().S(p)
//line web.qtpl:21
		qw422016.N().S(`</a> - `)
//line web.qtpl:21
		qw422016.E().S(doc)
//line web.qtpl:21
		qw422016.N().S(`<br/>
        `)
//line web.qtpl:22
	}
//line web.qtpl:22
	qw422016.N().S(`
    </p>
    `)
//line web.qtpl:24
	tpl.StreamFooter(qw422016)
//line web.qtpl:24
	qw422016.N().S(`
`)
//line web.qtpl:25
}

//line web.qtpl:25
func WriteWelcome(qq422016 qtio422016.Writer) {
//line web.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web.qtpl:25
	StreamWelcome(qw422016)
//line web.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line web.qtpl:25
}

//line web.qtpl:25
func Welcome() string {
//line web.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line web.qtpl:25
	WriteWelcome(qb422016)
//line web.qtpl:25
	qs422016 := string(qb422016.B)
//line web.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line web.qtpl:25
	return qs422016
//line web.qtpl:25
}

//line web.qtpl:27
func StreamListGroups(qw422016 *qt422016.Writer, groups []APIGroup) {
//line web.qtpl:27
	qw422016.N().S(`
    `)
//line web.qtpl:28
	tpl.StreamHeader(qw422016, "Groups", navItems)
//line web.qtpl:28
	qw422016.N().S(`
    `)
//line web.qtpl:29
	if len(groups) > 0 {
//line web.qtpl:29
		qw422016.N().S(`
        `)
//line web.qtpl:31
		rOk := make(map[string]int)
		rNotOk := make(map[string]int)
		for _, g := range groups {
			for _, r := range g.Rules {
				if r.LastError != "" {
					rNotOk[g.Name]++
				} else {
					rOk[g.Name]++
				}
			}
		}

//line web.qtpl:42
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>
        `)
//line web.qtpl:45
		for _, g := range groups {
//line web.qtpl:45
			qw422016.N().S(`
              <div class="group-heading`)
//line web.qtpl:46
			if rNotOk[g.Name] > 0 {
//line web.qtpl:46
				qw422016.N().S(` alert-danger`)
//line web.qtpl:46
			}
//line web.qtpl:46
			qw422016.N().S(`"  data-bs-target="rules-`)
//line web.qtpl:46
			qw422016.E().S(g.ID)
//line web.qtpl:46
			qw422016.N().S(`">
                <span class="anchor" id="group-`)
//line web.qtpl:47
			qw422016.E().S(g.ID)
//line web.qtpl:47
			qw422016.N().S(`"></span>
                <a href="#group-`)
//line web.qtpl:48
			qw422016.E().S(g.ID)
//line web.qtpl:48
			qw422016.N().S(`">`)
//line web.qtpl:48
			qw422016.E().S(g.Name)
//line web.qtpl:48
			if g.Type != "prometheus" {
//line web.qtpl:48
				qw422016.N().S(` (`)
//line web.qtpl:48
				qw422016.E().S(g.Type)
//line web.qtpl:48
				qw422016.N().S(`)`)
//line web.qtpl:48
			}
//line web.qtpl:48
			qw422016.N().S(` (every `)
//line web.qtpl:48
			qw422016.N().FPrec(g.Interval, 0)
//line web.qtpl:48
			qw422016.N().S(`s)</a>
                 `)
//line web.qtpl:49
			if rNotOk[g.Name] > 0 {
//line web.qtpl:49
				qw422016.N().S(`<span class="badge bg-danger" title="Number of rules with status Error">`)
//line web.qtpl:49
				qw422016.N().D(rNotOk[g.Name])
//line web.qtpl:49
				qw422016.N().S(`</span> `)
//line web.qtpl:49
			}
//line web.qtpl:49
			qw422016.N().S(`
                <span class="badge bg-success" title="Number of rules withs status Ok">`)
//line web.qtpl:50
			qw422016.N().D(rOk[g.Name])
//line web.qtpl:50
			qw422016.N().S(`</span>
                <p class="fs-6 fw-lighter">`)
//line web.qtpl:51
			qw422016.E().S(g.File)
//line web.qtpl:51
			qw422016.N().S(`</p>
                `)
//line web.qtpl:52
			if len(g.Params) > 0 {
//line web.qtpl:52
				qw422016.N().S(`
                    <div class="fs-6 fw-lighter">Extra params
                    `)
//line web.qtpl:54
				for _, param := range g.Params {
//line web.qtpl:54
					qw422016.N().S(`
                            <span class="float-left badge bg-primary">`)
//line web.qtpl:55
					qw422016.E().S(param)
//line web.qtpl:55
					qw422016.N().S(`</span>
                    `)
//line web.qtpl:56
				}
//line web.qtpl:56
				qw422016.N().S(`
                    </div>
                `)
//line web.qtpl:58
			}
//line web.qtpl:58
			qw422016.N().S(`
            </div>
            <div class="collapse" id="rules-`)
//line web.qtpl:60
			qw422016.E().S(g.ID)
//line web.qtpl:60
			qw422016.N().S(`">
                <table class="table table-striped table-hover table-sm">
                    <thead>
                        <tr>
                            <th scope="col" style="width: 60%">Rule</th>
                            <th scope="col" style="width: 20%" class="text-center" title="How many samples were produced by the rule">Samples</th>
                            <th scope="col" style="width: 20%" class="text-center" title="How many seconds ago rule was executed">Updated</th>
                        </tr>
                    </thead>
                    <tbody>
                    `)
//line web.qtpl:70
			for _, r := range g.Rules {
//line web.qtpl:70
				qw422016.N().S(`
                        <tr`)
//line web.qtpl:71
				if r.LastError != "" {
//line web.qtpl:71
					qw422016.N().S(` class="alert-danger"`)
//line web.qtpl:71
				}
//line web.qtpl:71
				qw422016.N().S(`>
                            <td>
                                <div class="row">
                                    <div class="col-12 mb-2">
                                        `)
//line web.qtpl:75
				if r.Type == "alerting" {
//line web.qtpl:75
					qw422016.N().S(`
                                        <b>alert:</b> `)
//line web.qtpl:76
					qw422016.E().S(r.Name)
//line web.qtpl:76
					qw422016.N().S(` (for: `)
//line web.qtpl:76
					qw422016.E().V(r.Duration)
//line web.qtpl:76
					qw422016.N().S(` seconds)
                                        `)
//line web.qtpl:77
				} else {
//line web.qtpl:77
					qw422016.N().S(`
                                        <b>record:</b> `)
//line web.qtpl:78
					qw422016.E().S(r.Name)
//line web.qtpl:78
					qw422016.N().S(`
                                        `)
//line web.qtpl:79
				}
//line web.qtpl:79
				qw422016.N().S(`
                                    </div>
                                    <div class="col-12">
                                        <code><pre>`)
//line web.qtpl:82
				qw422016.E().S(r.Query)
//line web.qtpl:82
				qw422016.N().S(`</pre></code>
                                    </div>
                                    <div class="col-12 mb-2">
                                        `)
//line web.qtpl:85
				if len(r.Labels) > 0 {
//line web.qtpl:85
					qw422016.N().S(` <b>Labels:</b>`)
//line web.qtpl:85
				}
//line web.qtpl:85
				qw422016.N().S(`
                                        `)
//line web.qtpl:86
				for k, v := range r.Labels {
//line web.qtpl:86
					qw422016.N().S(`
                                                <span class="ms-1 badge bg-primary">`)
//line web.qtpl:87
					qw422016.E().S(k)
//line web.qtpl:87
					qw422016.N().S(`=`)
//line web.qtpl:87
					qw422016.E().S(v)
//line web.qtpl:87
					qw422016.N().S(`</span>
                                        `)
//line web.qtpl:88
				}
//line web.qtpl:88
				qw422016.N().S(`
                                    </div>
                                    `)
//line web.qtpl:90
				if r.LastError != "" {
//line web.qtpl:90
					qw422016.N().S(`
                                    <div class="col-12">
                                        <b>Error:</b>
                                        <div class="error-cell">
                                        `)
//line web.qtpl:94
					qw422016.E().S(r.LastError)
//line web.qtpl:94
					qw422016.N().S(`
                                        </div>
                                    </div>
                                    `)
//line web.qtpl:97
				}
//line web.qtpl:97
				qw422016.N().S(`
                                </div>
                            </td>
                            <td class="text-center">`)
//line web.qtpl:100
				qw422016.N().D(r.LastSamples)
//line web.qtpl:100
				qw422016.N().S(`</td>
                            <td class="text-center">`)
//line web.qtpl:101
				qw422016.N().FPrec(time.Since(r.LastEvaluation).Seconds(), 3)
//line web.qtpl:101
				qw422016.N().S(`s ago</td>
                        </tr>
                    `)
//line web.qtpl:103
			}
//line web.qtpl:103
			qw422016.N().S(`
                 </tbody>
                </table>
            </div>
        `)
//line web.qtpl:107
		}
//line web.qtpl:107
		qw422016.N().S(`

    `)
//line web.qtpl:109
	} else {
//line web.qtpl:109
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line web.qtpl:113
	}
//line web.qtpl:113
	qw422016.N().S(`

    `)
//line web.qtpl:115
	tpl.StreamFooter(qw422016)
//line web.qtpl:115
	qw422016.N().S(`

`)
//line web.qtpl:117
}

//line web.qtpl:117
func WriteListGroups(qq422016 qtio422016.Writer, groups []APIGroup) {
//line web.qtpl:117
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web.qtpl:117
	StreamListGroups(qw422016, groups)
//line web.qtpl:117
	qt422016.ReleaseWriter(qw422016)
//line web.qtpl:117
}

//line web.qtpl:117
func ListGroups(groups []APIGroup) string {
//line web.qtpl:117
	qb422016 := qt422016.AcquireByteBuffer()
//line web.qtpl:117
	WriteListGroups(qb422016, groups)
//line web.qtpl:117
	qs422016 := string(qb422016.B)
//line web.qtpl:117
	qt422016.ReleaseByteBuffer(qb422016)
//line web.qtpl:117
	return qs422016
//line web.qtpl:117
}

//line web.qtpl:120
func StreamListAlerts(qw422016 *qt422016.Writer, pathPrefix string, groupAlerts []GroupAlerts) {
//line web.qtpl:120
	qw422016.N().S(`
    `)
//line web.qtpl:121
	tpl.StreamHeader(qw422016, "Alerts", navItems)
//line web.qtpl:121
	qw422016.N().S(`
    `)
//line web.qtpl:122
	if len(groupAlerts) > 0 {
//line web.qtpl:122
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>
         `)
//line web.qtpl:125
		for _, ga := range groupAlerts {
//line web.qtpl:125
			qw422016.N().S(`
            `)
//line web.qtpl:126
			g := ga.Group

//line web.qtpl:126
			qw422016.N().S(`
            <div class="group-heading alert-danger" data-bs-target="rules-`)
//line web.qtpl:127
			qw422016.E().S(g.ID)
//line web.qtpl:127
			qw422016.N().S(`">
                <span class="anchor" id="group-`)
//line web.qtpl:128
			qw422016.E().S(g.ID)
//line web.qtpl:128
			qw422016.N().S(`"></span>
                <a href="#group-`)
//line web.qtpl:129
			qw422016.E().S(g.ID)
//line web.qtpl:129
			qw422016.N().S(`">`)
//line web.qtpl:129
			qw422016.E().S(g.Name)
//line web.qtpl:129
			if g.Type != "prometheus" {
//line web.qtpl:129
				qw422016.N().S(` (`)
//line web.qtpl:129
				qw422016.E().S(g.Type)
//line web.qtpl:129
				qw422016.N().S(`)`)
//line web.qtpl:129
			}
//line web.qtpl:129
			qw422016.N().S(`</a>
                <span class="badge bg-danger" title="Number of active alerts">`)
//line web.qtpl:130
			qw422016.N().D(len(ga.Alerts))
//line web.qtpl:130
			qw422016.N().S(`</span>
                <br>
                <p class="fs-6 fw-lighter">`)
//line web.qtpl:132
			qw422016.E().S(g.File)
//line web.qtpl:132
			qw422016.N().S(`</p>
            </div>
            `)
//line web.qtpl:135
			var keys []string
			alertsByRule := make(map[string][]*APIAlert)
			for _, alert := range ga.Alerts {
				if len(alertsByRule[alert.RuleID]) < 1 {
					keys = append(keys, alert.RuleID)
				}
				alertsByRule[alert.RuleID] = append(alertsByRule[alert.RuleID], alert)
			}
			sort.Strings(keys)

//line web.qtpl:144
			qw422016.N().S(`
            <div class="collapse" id="rules-`)
//line web.qtpl:145
			qw422016.E().S(g.ID)
//line web.qtpl:145
			qw422016.N().S(`">
                `)
//line web.qtpl:146
			for _, ruleID := range keys {
//line web.qtpl:146
				qw422016.N().S(`
                    `)
//line web.qtpl:148
				defaultAR := alertsByRule[ruleID][0]
				var labelKeys []string
				for k := range defaultAR.Labels {
					labelKeys = append(labelKeys, k)
				}
				sort.Strings(labelKeys)

//line web.qtpl:154
				qw422016.N().S(`
                    <br>
                    <b>alert:</b> `)
//line web.qtpl:156
				qw422016.E().S(defaultAR.Name)
//line web.qtpl:156
				qw422016.N().S(` (`)
//line web.qtpl:156
				qw422016.N().D(len(alertsByRule[ruleID]))
//line web.qtpl:156
				qw422016.N().S(`)
                     | <span><a target="_blank" href="`)
//line web.qtpl:157
				qw422016.E().S(defaultAR.SourceLink)
//line web.qtpl:157
				qw422016.N().S(`">Source</a></span>
                    <br>
                    <b>expr:</b><code><pre>`)
//line web.qtpl:159
				qw422016.E().S(defaultAR.Expression)
//line web.qtpl:159
				qw422016.N().S(`</pre></code>
                    <table class="table table-striped table-hover table-sm">
                        <thead>
                            <tr>
                                <th scope="col">Labels</th>
                                <th scope="col">State</th>
                                <th scope="col">Active at</th>
                                <th scope="col">Value</th>
                                <th scope="col">Link</th>
                            </tr>
                        </thead>
                        <tbody>
                        `)
//line web.qtpl:171
				for _, ar := range alertsByRule[ruleID] {
//line web.qtpl:171
					qw422016.N().S(`
                            <tr>
                                <td>
                                    `)
//line web.qtpl:174
					for _, k := range labelKeys {
//line web.qtpl:174
						qw422016.N().S(`
                                        <span class="ms-1 badge bg-primary">`)
//line web.qtpl:175
						qw422016.E().S(k)
//line web.qtpl:175
						qw422016.N().S(`=`)
//line web.qtpl:175
						qw422016.E().S(ar.Labels[k])
//line web.qtpl:175
						qw422016.N().S(`</span>
                                    `)
//line web.qtpl:176
					}
//line web.qtpl:176
					qw422016.N().S(`
                                </td>
                                <td>`)
//line web.qtpl:178
					streambadgeState(qw422016, ar.State)
//line web.qtpl:178
					qw422016.N().S(`</td>
                                <td>
                                    `)
//line web.qtpl:180
					qw422016.E().S(ar.ActiveAt.Format("2006-01-02T15:04:05Z07:00"))
//line web.qtpl:180
					qw422016.N().S(`
                                    `)
//line web.qtpl:181
					if ar.Restored {
//line web.qtpl:181
						streambadgeRestored(qw422016)
//line web.qtpl:181
					}
//line web.qtpl:181
					qw422016.N().S(`
                                </td>
                                <td>`)
//line web.qtpl:183
					qw422016.E().S(ar.Value)
//line web.qtpl:183
					qw422016.N().S(`</td>
                                <td>
                                    <a href="`)
//line web.qtpl:185
					qw422016.E().S(path.Join(pathPrefix, g.ID, ar.ID, "status"))
//line web.qtpl:185
					qw422016.N().S(`">Details</a>
                                </td>
                            </tr>
                        `)
//line web.qtpl:188
				}
//line web.qtpl:188
				qw422016.N().S(`
                     </tbody>
                    </table>
                `)
//line web.qtpl:191
			}
//line web.qtpl:191
			qw422016.N().S(`
            </div>
            <br>
        `)
//line web.qtpl:194
		}
//line web.qtpl:194
		qw422016.N().S(`

    `)
//line web.qtpl:196
	} else {
//line web.qtpl:196
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line web.qtpl:200
	}
//line web.qtpl:200
	qw422016.N().S(`

    `)
//line web.qtpl:202
	tpl.StreamFooter(qw422016)
//line web.qtpl:202
	qw422016.N().S(`

`)
//line web.qtpl:204
}

//line web.qtpl:204
func WriteListAlerts(qq422016 qtio422016.Writer, pathPrefix string, groupAlerts []GroupAlerts) {
//line web.qtpl:204
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web.qtpl:204
	StreamListAlerts(qw422016, pathPrefix, groupAlerts)
//line web.qtpl:204
	qt422016.ReleaseWriter(qw422016)
//line web.qtpl:204
}

//line web.qtpl:204
func ListAlerts(pathPrefix string, groupAlerts []GroupAlerts) string {
//line web.qtpl:204
	qb422016 := qt422016.AcquireByteBuffer()
//line web.qtpl:204
	WriteListAlerts(qb422016, pathPrefix, groupAlerts)
//line web.qtpl:204
	qs422016 := string(qb422016.B)
//line web.qtpl:204
	qt422016.ReleaseByteBuffer(qb422016)
//line web.qtpl:204
	return qs422016
//line web.qtpl:204
}

//line web.qtpl:206
func StreamListTargets(qw422016 *qt422016.Writer, targets map[notifier.TargetType][]notifier.Target) {
//line web.qtpl:206
	qw422016.N().S(`
    `)
//line web.qtpl:207
	tpl.StreamHeader(qw422016, "Notifiers", navItems)
//line web.qtpl:207
	qw422016.N().S(`
    `)
//line web.qtpl:208
	if len(targets) > 0 {
//line web.qtpl:208
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>

         `)
//line web.qtpl:213
		var keys []string
		for key := range targets {
			keys = append(keys, string(key))
		}
		sort.Strings(keys)

//line web.qtpl:218
		qw422016.N().S(`

         `)
//line web.qtpl:220
		for i := range keys {
//line web.qtpl:220
			qw422016.N().S(`
           `)
//line web.qtpl:221
			typeK, ns := keys[i], targets[notifier.TargetType(keys[i])]
			count := len(ns)

//line web.qtpl:223
			qw422016.N().S(`
           <div class="group-heading data-bs-target="rules-`)
//line web.qtpl:224
			qw422016.E().S(typeK)
//line web.qtpl:224
			qw422016.N().S(`">
             <span class="anchor" id="notifiers-`)
//line web.qtpl:225
			qw422016.E().S(typeK)
//line web.qtpl:225
			qw422016.N().S(`"></span>
             <a href="#notifiers-`)
//line web.qtpl:226
			qw422016.E().S(typeK)
//line web.qtpl:226
			qw422016.N().S(`">`)
//line web.qtpl:226
			qw422016.E().S(typeK)
//line web.qtpl:226
			qw422016.N().S(` (`)
//line web.qtpl:226
			qw422016.N().D(count)
//line web.qtpl:226
			qw422016.N().S(`)</a>
         </div>
         <div class="collapse show" id="notifiers-`)
//line web.qtpl:228
			qw422016.E().S(typeK)
//line web.qtpl:228
			qw422016.N().S(`">
             <table class="table table-striped table-hover table-sm">
                 <thead>
                     <tr>
                         <th scope="col">Labels</th>
                         <th scope="col">Address</th>
                     </tr>
                 </thead>
                 <tbody>
                 `)
//line web.qtpl:237
			for _, n := range ns {
//line web.qtpl:237
				qw422016.N().S(`
                     <tr>
                         <td>
                              `)
//line web.qtpl:240
				for _, l := range n.Labels {
//line web.qtpl:240
					qw422016.N().S(`
                                      <span class="ms-1 badge bg-primary">`)
//line web.qtpl:241
					qw422016.E().S(l.Name)
//line web.qtpl:241
					qw422016.N().S(`=`)
//line web.qtpl:241
					qw422016.E().S(l.Value)
//line web.qtpl:241
					qw422016.N().S(`</span>
                              `)
//line web.qtpl:242
				}
//line web.qtpl:242
				qw422016.N().S(`
                          </td>
                         <td>`)
//line web.qtpl:244
				qw422016.E().S(n.Notifier.Addr())
//line web.qtpl:244
				qw422016.N().S(`</td>
                     </tr>
                 `)
//line web.qtpl:246
			}
//line web.qtpl:246
			qw422016.N().S(`
              </tbody>
             </table>
         </div>
     `)
//line web.qtpl:250
		}
//line web.qtpl:250
		qw422016.N().S(`

    `)
//line web.qtpl:252
	} else {
//line web.qtpl:252
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line web.qtpl:256
	}
//line web.qtpl:256
	qw422016.N().S(`

    `)
//line web.qtpl:258
	tpl.StreamFooter(qw422016)
//line web.qtpl:258
	qw422016.N().S(`

`)
//line web.qtpl:260
}

//line web.qtpl:260
func WriteListTargets(qq422016 qtio422016.Writer, targets map[notifier.TargetType][]notifier.Target) {
//line web.qtpl:260
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web.qtpl:260
	StreamListTargets(qw422016, targets)
//line web.qtpl:260
	qt422016.ReleaseWriter(qw422016)
//line web.qtpl:260
}

//line web.qtpl:260
func ListTargets(targets map[notifier.TargetType][]notifier.Target) string {
//line web.qtpl:260
	qb422016 := qt422016.AcquireByteBuffer()
//line web.qtpl:260
	WriteListTargets(qb422016, targets)
//line web.qtpl:260
	qs422016 := string(qb422016.B)
//line web.qtpl:260
	qt422016.ReleaseByteBuffer(qb422016)
//line web.qtpl:260
	return qs422016
//line web.qtpl:260
}

//line web.qtpl:262
func StreamAlert(qw422016 *qt422016.Writer, pathPrefix string, alert *APIAlert) {
//line web.qtpl:262
	qw422016.N().S(`
    `)
//line web.qtpl:263
	tpl.StreamHeader(qw422016, "", navItems)
//line web.qtpl:263
	qw422016.N().S(`
    `)
//line web.qtpl:265
	var labelKeys []string
	for k := range alert.Labels {
		labelKeys = append(labelKeys, k)
	}
	sort.Strings(labelKeys)

	var annotationKeys []string
	for k := range alert.Annotations {
		annotationKeys = append(annotationKeys, k)
	}
	sort.Strings(annotationKeys)

//line web.qtpl:276
	qw422016.N().S(`
    <div class="display-6 pb-3 mb-3">`)
//line web.qtpl:277
	qw422016.E().S(alert.Name)
//line web.qtpl:277
	qw422016.N().S(`<span class="ms-2 badge `)
//line web.qtpl:277
	if alert.State == "firing" {
//line web.qtpl:277
		qw422016.N().S(`bg-danger`)
//line web.qtpl:277
	} else {
//line web.qtpl:277
		qw422016.N().S(` bg-warning text-dark`)
//line web.qtpl:277
	}
//line web.qtpl:277
	qw422016.N().S(`">`)
//line web.qtpl:277
	qw422016.E().S(alert.State)
//line web.qtpl:277
	qw422016.N().S(`</span></div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Active at
        </div>
        <div class="col">
          `)
//line web.qtpl:284
	qw422016.E().S(alert.ActiveAt.Format("2006-01-02T15:04:05Z07:00"))
//line web.qtpl:284
	qw422016.N().S(`
        </div>
      </div>
      </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Expr
        </div>
        <div class="col">
          <code><pre>`)
//line web.qtpl:294
	qw422016.E().S(alert.Expression)
//line web.qtpl:294
	qw422016.N().S(`</pre></code>
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Labels
        </div>
        <div class="col">
           `)
//line web.qtpl:304
	for _, k := range labelKeys {
//line web.qtpl:304
		qw422016.N().S(`
                <span class="m-1 badge bg-primary">`)
//line web.qtpl:305
		qw422016.E().S(k)
//line web.qtpl:305
		qw422016.N().S(`=`)
//line web.qtpl:305
		qw422016.E().S(alert.Labels[k])
//line web.qtpl:305
		qw422016.N().S(`</span>
          `)
//line web.qtpl:306
	}
//line web.qtpl:306
	qw422016.N().S(`
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Annotations
        </div>
        <div class="col">
           `)
//line web.qtpl:316
	for _, k := range annotationKeys {
//line web.qtpl:316
		qw422016.N().S(`
                <b>`)
//line web.qtpl:317
		qw422016.E().S(k)
//line web.qtpl:317
		qw422016.N().S(`:</b><br>
                <p>`)
//line web.qtpl:318
		qw422016.E().S(alert.Annotations[k])
//line web.qtpl:318
		qw422016.N().S(`</p>
          `)
//line web.qtpl:319
	}
//line web.qtpl:319
	qw422016.N().S(`
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Group
        </div>
        <div class="col">
           <a target="_blank" href="`)
//line web.qtpl:329
	qw422016.E().S(path.Join(pathPrefix, "groups"))
//line web.qtpl:329
	qw422016.N().S(`#group-`)
//line web.qtpl:329
	qw422016.E().S(alert.GroupID)
//line web.qtpl:329
	qw422016.N().S(`">`)
//line web.qtpl:329
	qw422016.E().S(alert.GroupID)
//line web.qtpl:329
	qw422016.N().S(`</a>
        </div>
      </div>
    </div>
     <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Source link
        </div>
        <div class="col">
           <a target="_blank" href="`)
//line web.qtpl:339
	qw422016.E().S(alert.SourceLink)
//line web.qtpl:339
	qw422016.N().S(`">Link</a>
        </div>
      </div>
    </div>
    `)
//line web.qtpl:343
	tpl.StreamFooter(qw422016)
//line web.qtpl:343
	qw422016.N().S(`

`)
//line web.qtpl:345
}

//line web.qtpl:345
func WriteAlert(qq422016 qtio422016.Writer, pathPrefix string, alert *APIAlert) {
//line web.qtpl:345
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web.qtpl:345
	StreamAlert(qw422016, pathPrefix, alert)
//line web.qtpl:345
	qt422016.ReleaseWriter(qw422016)
//line web.qtpl:345
}

//line web.qtpl:345
func Alert(pathPrefix string, alert *APIAlert) string {
//line web.qtpl:345
	qb422016 := qt422016.AcquireByteBuffer()
//line web.qtpl:345
	WriteAlert(qb422016, pathPrefix, alert)
//line web.qtpl:345
	qs422016 := string(qb422016.B)
//line web.qtpl:345
	qt422016.ReleaseByteBuffer(qb422016)
//line web.qtpl:345
	return qs422016
//line web.qtpl:345
}

//line web.qtpl:347
func streambadgeState(qw422016 *qt422016.Writer, state string) {
//line web.qtpl:347
	qw422016.N().S(`
`)
//line web.qtpl:349
	badgeClass := "bg-warning text-dark"
	if state == "firing" {
		badgeClass = "bg-danger"
	}

//line web.qtpl:353
	qw422016.N().S(`
<span class="badge `)
//line web.qtpl:354
	qw422016.E().S(badgeClass)
//line web.qtpl:354
	qw422016.N().S(`">`)
//line web.qtpl:354
	qw422016.E().S(state)
//line web.qtpl:354
	qw422016.N().S(`</span>
`)
//line web.qtpl:355
}

//line web.qtpl:355
func writebadgeState(qq422016 qtio422016.Writer, state string) {
//line web.qtpl:355
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web.qtpl:355
	streambadgeState(qw422016, state)
//line web.qtpl:355
	qt422016.ReleaseWriter(qw422016)
//line web.qtpl:355
}

//line web.qtpl:355
func badgeState(state string) string {
//line web.qtpl:355
	qb422016 := qt422016.AcquireByteBuffer()
//line web.qtpl:355
	writebadgeState(qb422016, state)
//line web.qtpl:355
	qs422016 := string(qb422016.B)
//line web.qtpl:355
	qt422016.ReleaseByteBuffer(qb422016)
//line web.qtpl:355
	return qs422016
//line web.qtpl:355
}

//line web.qtpl:357
func streambadgeRestored(qw422016 *qt422016.Writer) {
//line web.qtpl:357
	qw422016.N().S(`
<span class="badge bg-warning text-dark" title="Alert state was restored after the service restart from remote storage">restored</span>
`)
//line web.qtpl:359
}

//line web.qtpl:359
func writebadgeRestored(qq422016 qtio422016.Writer) {
//line web.qtpl:359
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web.qtpl:359
	streambadgeRestored(qw422016)
//line web.qtpl:359
	qt422016.ReleaseWriter(qw422016)
//line web.qtpl:359
}

//line web.qtpl:359
func badgeRestored() string {
//line web.qtpl:359
	qb422016 := qt422016.AcquireByteBuffer()
//line web.qtpl:359
	writebadgeRestored(qb422016)
//line web.qtpl:359
	qs422016 := string(qb422016.B)
//line web.qtpl:359
	qt422016.ReleaseByteBuffer(qb422016)
//line web.qtpl:359
	return qs422016
//line web.qtpl:359
}
