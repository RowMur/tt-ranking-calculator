// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "strconv"

func isItemDefaultSelected(item Ranking, me string) string {
	if strconv.Itoa(item.Id) == me {
		return "selected"
	}

	return "false"
}

func page(options []Ranking, tournamentOptions []TournamentOption) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC\" crossorigin=\"anonymous\"><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM\" crossorigin=\"anonymous\"></script><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"></head><body><div><p><b>Latest Ranking List:</b> September 2024</p><form id=\"form\"><div class=\"mb-3\"><label for=\"me-select\" class=\"form-label\">You</label> <select id=\"me-select\" form=\"form\" name=\"me\" list=\"me-options\" class=\"form-control\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range options {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(parseName(item.Name))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 29, Col: 44}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(parseName(item.Name))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 29, Col: 69}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div><div class=\"mb-3\"><label for=\"tournament-select\" class=\"form-label\">Tournament</label> <select id=\"tournament-select\" form=\"form\" name=\"tournament\" class=\"form-control\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, tournament := range tournamentOptions {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(tournament.Id)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 37, Col: 37}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(tournament.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 37, Col: 57}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div><fieldset id=\"formFields\"><legend>Opponents</legend>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = opponent(options).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</fieldset><button id=\"add-opponent-button\" type=\"button\" class=\"btn btn-secondary\">Add opponent</button> <button type=\"submit\" class=\"btn btn-primary\">Submit</button></form><div class=\"hidden\" id=\"points-wrapper\"><p>Points: <span id=\"points\">0</span></p></div></div></body><style>\n\t\tbody {\n\t\t\tdisplay: grid;\n\t\t\tplace-items: center;\n\t\t\theight: 100vh;\n\t\t}\n\n\t\tform {\n\t\t\twidth: 100%;\n\t\t\tmin-width: 0px;\n\t\t\tmargin: 0 auto;\n\t\t}\n\n\t\t@media (max-width: 768px) {\n\t\t\tform {\n\t\t\t\twidth: 80%;\n\t\t\t}\n\t\t}\n\n\t\tp {\n\t\t\ttext-align: center;\n\t\t}\n\n\t\t.grid {\n\t\t\tdisplay: grid;\n\t\t\tplace-items: center;\n\t\t}\n\n\t\t.opponent {\n\t\t\tdisplay: flex;\n\t\t}\n\n\t\t.win {\n\t\t\tmargin-left: 1rem;\n\t\t\tmin-width: 60px;\n\t\t}\n\n\t\t.hidden {\n\t\t\tdisplay: none;\n\t\t}\n\n\t\t#points-wrapper {\n\t\t\tmargin-top: 1rem;\n\t\t}\n\t\t\n\t</style><script type=\"text/javascript\">\n\t\tdocument.getElementById(\"add-opponent-button\").addEventListener(\"click\",\n\t\t\tfunction(event) {\n\t\t\t\tevent.preventDefault();\n\t\t\t\taddOpponent();\n\t\t\t}\n\t\t);\n\n\t\tfunction addOpponent() {\n\t\t\tconst lastOpponent = document.getElementById(\"formFields\").lastChild;\n\t\t\tconst clonedOpponent = lastOpponent.cloneNode(true);\n\t\t\tconst formFields = document.getElementById(\"formFields\")\n\t\t\tformFields.appendChild(clonedOpponent);\n\t\t\t\n\t\t\tconst key = document.getElementById(\"formFields\").children.length - 2;\n\t\t\tconst clonedElementInDom = document.getElementById(\"formFields\").lastChild;\n\t\t\t\n\t\t\tconst checkbox = clonedElementInDom.querySelector(\"input[type=\\\"checkbox\\\"]\");\n\t\t\tcheckbox.name = \"win\" + key;\n\t\t\tcheckbox.id = \"win\" + key;\n\t\t\tcheckbox.checked = true;\n\n\t\t\tclonedElementInDom.querySelector(\"label\").for = \"win\" + key;\n\n\t\t\tconst input = clonedElementInDom.querySelector(\"select\");\n\t\t\tinput.name = \"opponent\" + key;\n\t\t\tinput.id = \"opponent-select\" + key;\n\t\t}\n\n\t\tconst form = document.getElementById(\"form\");\n\t\tform.addEventListener(\"submit\", async (e) => {\n\t\t\tevent.preventDefault();\n\t\t\tconst data = new FormData(e.target);\n\t\t\tlet resp = await fetch(\"/calculate\", {\n\t\t\t\tmethod: \"POST\",\n\t\t\t\theaders: {\n\t\t\t\t\t\"Content-Type\": \"application/x-www-form-urlencoded\",\n\t\t\t\t},\n\t\t\t\tbody: new URLSearchParams(data),\n\t\t\t});\n\n\t\t\tlet res = await resp.json();\n\t\t\tdocument.getElementById(\"points\").innerText = res.points;\n\t\t\tdocument.getElementById(\"points-wrapper\").classList.remove(\"hidden\");\n\t\t})\n\t</script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func dropdown(items []Ranking) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<select id=\"opponent-select0\" name=\"opponent0\" form=\"form\" type=\"text\" class=\"form-control\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range items {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(parseName(item.Name))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 159, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(parseName(item.Name))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 159, Col: 64}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func opponent(options []Ranking) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"opponent\" class=\"mb-3 opponent\"><div class=\"\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dropdown(options).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"grid win\"><div><input class=\"form-check-input\" type=\"checkbox\" name=\"win0\" id=\"win0\" checked> <label for=\"win0\" class=\"form-check-label\">Win?</label></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
