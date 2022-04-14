package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["backend/controllers:AuditLogController"] = append(beego.GlobalControllerRouter["backend/controllers:AuditLogController"],
        beego.ControllerComments{
            Method: "GetAuditLogs",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "RefreshCookie",
            Router: "/refresh-cookie",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "RefreshToken",
            Router: "/refresh-token",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "SocialCallback",
            Router: "/social/callback",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "SocialRedirect",
            Router: "/social/redirect",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["backend/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Token",
            Router: "/token",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CommonController"] = append(beego.GlobalControllerRouter["backend/controllers:CommonController"],
        beego.ControllerComments{
            Method: "Contact",
            Router: "/contact",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "UpdateCustomer",
            Router: "/:id",
            AllowHTTPMethods: []string{"PATCH"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "CustomerInvoices",
            Router: "/:id/invoices",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "ListPaymentMethods",
            Router: "/:id/payment-methods",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "CustomerSubscription",
            Router: "/:id/subscription",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["backend/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "UpcomingInvoice",
            Router: "/upcoming",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:PaymentController"] = append(beego.GlobalControllerRouter["backend/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "Webhook",
            Router: "/webhook/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:PaymentMethodController"] = append(beego.GlobalControllerRouter["backend/controllers:PaymentMethodController"],
        beego.ControllerComments{
            Method: "CreatePaymentMethod",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:PaymentMethodController"] = append(beego.GlobalControllerRouter["backend/controllers:PaymentMethodController"],
        beego.ControllerComments{
            Method: "DeletePaymentMethod",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:ProductController"] = append(beego.GlobalControllerRouter["backend/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProducts",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:SubscriptionController"] = append(beego.GlobalControllerRouter["backend/controllers:SubscriptionController"],
        beego.ControllerComments{
            Method: "CreateSubscription",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:SubscriptionController"] = append(beego.GlobalControllerRouter["backend/controllers:SubscriptionController"],
        beego.ControllerComments{
            Method: "UpdateSubscription",
            Router: "/:id",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:SubscriptionController"] = append(beego.GlobalControllerRouter["backend/controllers:SubscriptionController"],
        beego.ControllerComments{
            Method: "CancelSubscription",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "SignUp",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/:id",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Authenticator",
            Router: "/:id/authenticator",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GenerateAuthenticator",
            Router: "/:id/generate-auth-code",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "ReadNotification",
            Router: "/:id/notifications",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePassword",
            Router: "/:id/password",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "RecoveryQuestions",
            Router: "/:id/recovery-questions",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendVerifySMS",
            Router: "/:id/send-verify-sms",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifyMobile",
            Router: "/:id/verify-mobile",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "DisableMFA",
            Router: "/disable-mfa",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "ForgotPassword",
            Router: "/forgot-password",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetRecoveryQuestions",
            Router: "/recovery-questions",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: "/reset-password",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:UserController"] = append(beego.GlobalControllerRouter["backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifyEmail",
            Router: "/verify-email",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
