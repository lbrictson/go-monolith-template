// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func LoginPage(showSuccessReset bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"row\"><div class=\"col-md-4\"></div><div class=\"col-md-4\" style=\"padding-top: 8rem;\"><div class=\"card shadow-lg\"><div class=\"card-body\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if showSuccessReset {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"alert alert-secondary alert-dismissible fade show\" role=\"alert\"><strong>Success</strong> Password reset successful. You can now login.\r <button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"alert\" aria-label=\"Close\"></button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h3 class=\"text-center\">Application Name</h3><hr><form class=\"form\" action=\"/login\" method=\"post\"><div class=\"mb-3\"><input class=\"form-control\" type=\"email\" name=\"email\" required placeholder=\"Email\"></div><div class=\"mb-3\"><input class=\"form-control\" type=\"password\" name=\"password\" required placeholder=\"Password\"></div><div class=\"d-grid gap-2\"><button class=\"btn btn-secondary\" type=\"submit\">Login</button></div></form><hr><p>Forgot your password? <a href=\"/reset_password\">Reset it</a></p></div></div></div><div class=\"col-md-4\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MFAPage(email string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"row\"><div class=\"col-md-4\"></div><div class=\"col-md-4\" style=\"padding-top: 8rem;\"><div class=\"card shadow-lg\"><div class=\"card-body\"><h3 class=\"text-center\">Multifactor Authentication</h3><hr><form class=\"form\" action=\"/mfa\" method=\"post\" autocomplete=\"off\"><div class=\"mb-3\"><input class=\"form-control\" type=\"email\" name=\"email\" hidden value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(email))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div class=\"mb-3\"><input class=\"form-control\" type=\"text\" name=\"token\" required placeholder=\"Code\"></div><div class=\"d-grid gap-2\"><button class=\"btn btn-secondary\" type=\"submit\">Submit</button></div></form></div></div></div><div class=\"col-md-4\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ResetPasswordPage(success bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"row\"><div class=\"col-md-4\"></div><div class=\"col-md-4\" style=\"padding-top: 8rem;\"><div class=\"card shadow-lg\"><div class=\"card-body\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if success {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"alert alert-secondary alert-dismissible fade show\" role=\"alert\"><strong>Success</strong> If an account exists with the requested email a password reset email will be sent.\r <button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"alert\" aria-label=\"Close\"></button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h3 class=\"text-center\">Request a Password Reset Email</h3><hr><form class=\"form\" action=\"/reset_password\" method=\"post\"><div class=\"mb-3\"><input class=\"form-control\" type=\"email\" name=\"email\" required placeholder=\"Email\"></div><div class=\"d-grid gap-2\"><button class=\"btn btn-secondary\" type=\"submit\">Request Password Reset</button></div></form></div></div></div><div class=\"col-md-4\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SetPasswordPage(token string, email string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"row\"><div class=\"col-md-4\"></div><div class=\"col-md-4\" style=\"padding-top: 8rem;\"><div class=\"card shadow-lg\"><div class=\"card-body\"><h3 class=\"text-center\">Set a New Password</h3><hr><form class=\"form\" action=\"/set_password\" method=\"post\"><div class=\"mb-3\"><input class=\"form-control\" type=\"email\" name=\"email\" hidden value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(email))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div class=\"mb-3\"><input class=\"form-control\" type=\"text\" name=\"token\" hidden value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(token))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div class=\"mb-3\"><input class=\"form-control\" type=\"password\" name=\"password\" required placeholder=\"Password\"></div><div class=\"d-grid gap-2\"><button class=\"btn btn-secondary\" type=\"submit\">Set Password</button></div></form></div></div></div><div class=\"col-md-4\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
