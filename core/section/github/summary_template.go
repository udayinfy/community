// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package github

const summaryTemplate = `
<div class="section-github-render">

	<p>
	Activity since {{.Config.Since}}{{.Config.DateMessage}} for {{.Config.Owner}} repositories:
		{{range $data := .Config.Lists}}
			{{if $data.Included}}
				<a class="link" href="{{$data.URL}}">
					{{$data.Repo}}{{if $data.Comma}},{{end}}
				</a>
			{{end}}
		{{end}}
	</p>

	{{if .HasSharedLabels}}
		<h3>Common Labels</h3>
		<p>There
		{{if eq 1 (len .SharedLabels)}} is {{else}} are {{end}}
		{{len .SharedLabels}}
		shared
		{{if eq 1 (len .SharedLabels)}} label {{else}} labels {{end}}
		across the repositories.</p>
		<table class="width-100">
		    <tbody class="github">
			{{range $slabel := .SharedLabels}}
		        <tr>
		            <td class="width-100">
						<span class="github-issue-label" style="background-color:#{{$slabel.Color}}">{{$slabel.Name}} ({{$slabel.Count}})</span> in {{$slabel.Repos}}
					</td>
		        </tr>
			{{end}}
		    </tbody>
		</table>
	{{end}}
</div>
`
