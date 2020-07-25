package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["backend/controllers:AuditLogController"] = append(beego.GlobalControllerRouter["backend/controllers:AuditLogController"],
        beego.ControllerComments{
            Method: "GetAuditLogs",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "RefreshCookie",
            Router: `/refresh-cookie`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "RefreshToken",
            Router: `/refresh-token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "SocialCallback",
            Router: `/social/callback`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "SocialRedirect",
            Router: `/social/redirect`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Token",
            Router: `/token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "SignUp",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Authenticator",
            Router: `/:id/authenticator`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GenerateAuthenticator",
            Router: `/:id/generate-auth-code`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePassword",
            Router: `/:id/password`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "RecoveryQuestions",
            Router: `/:id/recovery-questions`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendVerifySMS",
            Router: `/:id/send-verify-sms`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifyMobile",
            Router: `/:id/verify-mobile`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "DisableMFA",
            Router: `/disable-mfa`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "ForgotPassword",
            Router: `/forgot-password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetRecoveryQuestions",
            Router: `/recovery-questions`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/reset-password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifyEmail",
            Router: `/verify-email`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
