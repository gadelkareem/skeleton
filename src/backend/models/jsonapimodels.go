package models

type (
    AuthToken struct {
        ID           int    `jsonapi:"primary,token-requests,omitempty"`
        Token        string `jsonapi:"attr,token,omitempty"`
        UserId       int64  `jsonapi:"attr,user_id"`
        RefreshToken string `jsonapi:"attr,refresh_token,omitempty"`
    }
    Login struct {
        ID         int    `jsonapi:"primary,login-requests,omitempty"`
        Username   string `jsonapi:"attr,username"`
        Password   string `jsonapi:"attr,password"`
        RememberMe bool   `jsonapi:"attr,remember_me,omitempty"`
        Code       string `jsonapi:"attr,code,omitempty"`
    }
    VerifyEmail struct {
        ID    int    `jsonapi:"primary,email-verify-requests,omitempty"`
        Email string `jsonapi:"attr,email"`
        Token string `jsonapi:"attr,token"`
    }
    VerifyMobile struct {
        ID   int    `jsonapi:"primary,mobile-verify-requests,omitempty"`
        Code string `jsonapi:"attr,code"`
    }
    ResetPassword struct {
        ID       int    `jsonapi:"primary,reset-password-requests,omitempty"`
        Email    string `jsonapi:"attr,email"`
        Username string `jsonapi:"attr,username"`
        Token    string `jsonapi:"attr,token,omitempty"`
        Password string `jsonapi:"attr,password,omitempty"`
    }
    UpdatePassword struct {
        ID          int    `jsonapi:"primary,update-password-requests,omitempty"`
        OldPassword string `jsonapi:"attr,old_pass"`
        Password    string `jsonapi:"attr,password"`
    }
    SocialAuth struct {
        ID          int    `jsonapi:"primary,social-login-requests,omitempty"`
        Provider    string `jsonapi:"attr,provider,omitempty"`
        ClientId    string `jsonapi:"attr,client_id,omitempty"`
        Code        string `jsonapi:"attr,code,omitempty"`
        State       string `jsonapi:"attr,state,omitempty"`
        RedirectUri string `jsonapi:"attr,redirect_uri,omitempty"`
    }
    Authenticator struct {
        ID      int    `jsonapi:"primary,authenticator-requests,omitempty"`
        Image   string `jsonapi:"attr,image,omitempty"`
        Seed    string `jsonapi:"attr,seed,omitempty"`
        URL     string `jsonapi:"attr,url,omitempty"`
        Code    string `jsonapi:"attr,code,omitempty"`
        Refresh bool   `jsonapi:"attr,refresh,omitempty"`
        Enable  bool   `jsonapi:"attr,enable,omitempty"`
    }
    DisableMFA struct {
        ID                int                 `jsonapi:"primary,mfa-disable-requests,omitempty"`
        Username          string              `jsonapi:"attr,username"`
        Password          string              `jsonapi:"attr,password"`
        RecoveryQuestions []*RecoveryQuestion `jsonapi:"relation,questions"`
    }
    RecoveryQuestions struct {
        ID        int                 `jsonapi:"primary,recovery-questions-requests,omitempty"`
        Questions []*RecoveryQuestion `jsonapi:"relation,questions"`
    }
    RecoveryQuestion struct {
        Question string `jsonapi:"primary,questions"`
        Answer   string `jsonapi:"attr,answer"`
    }
)
