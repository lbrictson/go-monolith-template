// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "go-monolith-template/pkg/models"

func ProfilePage(user models.User, isAdmin bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = Navbar("dashboard", user.Email, isAdmin).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"row\"><div class=\"col-md-4\"></div><div class=\"col-md-4\"><div class=\"card shadow-lg\"><div class=\"card-body\"><h4 class=\"text-center\" style=\"text-decoration: underline;\">Settings</h4><br><form class=\"row row-cols-lg-auto g-3 align-items-center\"><div class=\"col-12\"><label class=\"visually-hidden\" for=\"inlineFormInputGroupUsername\">Email</label><div class=\"input-group\"><div class=\"input-group-text\">Email</div><input type=\"text\" name=\"email\" class=\"form-control\" disabled value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(user.Email))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div></form><br><form class=\"row row-cols-lg-auto g-3 align-items-center\" action=\"/profile/password\" method=\"post\"><div class=\"col-12\"><label class=\"visually-hidden\" for=\"inlineFormInputGroupUsername\">Password</label><div class=\"input-group\"><div class=\"input-group-text\">New Password</div><input type=\"password\" name=\"password\" class=\"form-control\"></div></div><div class=\"col-12\"><button type=\"submit\" class=\"btn btn-secondary\">Change</button></div></form><br><h4 class=\"text-center\" style=\"text-decoration: underline;\">Advanced Security</h4>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if user.MFAEnabled {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"row row-cols-lg-auto g-3 d-flex align-items-center justify-content-center\"><div class=\"col-12\"><button type=\"button\" class=\"btn btn-secondary\" hx-delete=\"/profile/disable_mfa\" hx-confirm=\"Are you sure you want to disable MFA?\">Disable Multifactor Authentication</button></div></form>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"row row-cols-lg-auto g-3 d-flex align-items-center justify-content-center\"><div class=\"col-12\"><a href=\"/profile/enable_mfa\" class=\"btn btn-secondary\">Enable Multifactor Authentication</a></div></form>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div><div class=\"col-md-4\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PageEnableMFA(user models.User, isAdmin bool, qrCode string, mfaSecret string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = Navbar("profile", user.Email, isAdmin).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"row\"><div class=\"col-md-4\"></div><div class=\"col-md-4 d-flex align-items-center justify-content-center\"><div class=\"card shadow-lg\"><div class=\"card-body\"><h4 class=\"text-center\" style=\"text-decoration: underline;\">Enable Multifactor Authentication</h4><br><div class=\"text-center\"><img class=\"img-thumbnail\" alt=\"QR Code\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(qrCode))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><br><p>Scan the QR code with your authenticator app</p><p>Or enter the code manually: ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(mfaSecret)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/profile.templ`, Line: 73, Col: 57}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><form class=\"row row-cols-lg-auto g-3 align-items-center\" action=\"/profile/enable_mfa\" method=\"post\"><div class=\"col-12\"><label class=\"visually-hidden\" for=\"inlineFormInputGroupUsername\">Enter Code</label><div class=\"input-group\"><div class=\"input-group-text\">Enter Code</div><input type=\"text\" name=\"token\" class=\"form-control\" required></div></div><div class=\"col-12\"><button type=\"submit\" class=\"btn btn-secondary\">Enable</button></div></form></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}