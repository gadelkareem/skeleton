package models

type (
    AuthToken struct {
        ID           int    `jsonapi:"primary,token-requests,omitempty"`
        Token        string `jsonapi:"attr,token,omitempty" valid:"Match(/^[A-Za-z0-9\\-_\\.]{100,}$/)"`
        UserId       int64  `jsonapi:"attr,user_id" valid:"Numeric"`
        RefreshToken string `jsonapi:"attr,refresh_token,omitempty"  valid:"Required;Match(/^[A-Za-z0-9\\-_\\.]{100,}$/)"`
    }
    Login struct {
        ID         int    `jsonapi:"primary,login-requests,omitempty"`
        Username   string `jsonapi:"attr,username" valid:"Required;Match(/^[A-Za-z0-9]+(?:[-_.][A-Za-z0-9]+)*$/)"`
        Password   string `jsonapi:"attr,password" valid:"MinSize(6);MaxSize(100)"`
        RememberMe bool   `jsonapi:"attr,remember_me,omitempty"`
        Code       string `jsonapi:"attr,code,omitempty" valid:"Length(6)"`
    }
    VerifyEmail struct {
        ID    int    `jsonapi:"primary,email-verify-requests,omitempty"`
        Email string `jsonapi:"attr,email" valid:"Required;Email;MaxSize(100)"`
        Token string `jsonapi:"attr,token" valid:"Required;Length(32)"`
    }
    VerifyMobile struct {
        ID   int    `jsonapi:"primary,mobile-verify-requests,omitempty"`
        Code string `jsonapi:"attr,code" valid:"Required;Length(6)"`
    }
    ResetPassword struct {
        ID       int    `jsonapi:"primary,reset-password-requests,omitempty" `
        Email    string `jsonapi:"attr,email" valid:"Email;MaxSize(100)"`
        Username string `jsonapi:"attr,username" valid:"Match(/^[A-Za-z0-9]+(?:[-_.][A-Za-z0-9]+)*$/)"`
        Token    string `jsonapi:"attr,token,omitempty" valid:"Length(32)"`
        Password string `jsonapi:"attr,password,omitempty" valid:"MinSize(6);MaxSize(100)"`
    }
    UpdatePassword struct {
        ID          int    `jsonapi:"primary,update-password-requests,omitempty"`
        OldPassword string `jsonapi:"attr,old_pass" valid:"Required;MinSize(6);MaxSize(100)"`
        Password    string `jsonapi:"attr,password" valid:"Required;MinSize(6);MaxSize(100)"`
    }
    SocialAuth struct {
        ID          int    `jsonapi:"primary,social-login-requests,omitempty"`
        Provider    string `jsonapi:"attr,provider,omitempty" valid:"MinSize(2)"`
        Code        string `jsonapi:"attr,code,omitempty" valid:"MinSize(2)"`
        State       string `jsonapi:"attr,state,omitempty" valid:"MinSize(2)"`
        RedirectUri string `jsonapi:"attr,redirect_uri,omitempty" valid:"MinSize(2)"`
    }
    Authenticator struct {
        ID      int    `jsonapi:"primary,authenticator-requests,omitempty"`
        Image   string `jsonapi:"attr,image,omitempty"`
        Seed    string `jsonapi:"attr,seed,omitempty"`
        URL     string `jsonapi:"attr,url,omitempty"`
        Code    string `jsonapi:"attr,code,omitempty" valid:"Required;Length(6)"`
        Refresh bool   `jsonapi:"attr,refresh,omitempty"`
        Enable  bool   `jsonapi:"attr,enable,omitempty"`
    }
    DisableMFA struct {
        ID                int                 `jsonapi:"primary,mfa-disable-requests,omitempty"`
        Username          string              `jsonapi:"attr,username" valid:"Required;Match(/^[A-Za-z0-9]+(?:[-_.][A-Za-z0-9]+)*$/)"`
        Password          string              `jsonapi:"attr,password" valid:"Required;MinSize(6);MaxSize(100)"`
        RecoveryQuestions []*RecoveryQuestion `jsonapi:"relation,questions" valid:"Required"`
    }
    RecoveryQuestions struct {
        ID        int                 `jsonapi:"primary,recovery-questions-requests,omitempty"`
        Questions []*RecoveryQuestion `jsonapi:"relation,questions" valid:"Required"`
    }
    RecoveryQuestion struct {
        Question string `jsonapi:"primary,questions" valid:"Required"`
        Answer   string `jsonapi:"attr,answer" valid:"Required"`
    }
    Contact struct {
        ID      int    `jsonapi:"primary,contact-requests,omitempty"`
        Name    string `jsonapi:"attr,name" valid:"Required"`
        Email   string `jsonapi:"attr,email" valid:"Required;Email;MaxSize(100)"`
        Message string `jsonapi:"attr,message" valid:"MinSize(10)"`
    }
)
