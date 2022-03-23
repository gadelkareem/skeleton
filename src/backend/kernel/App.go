package kernel

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/gadelkareem/go-helpers"
)

const (
	Dev        = "dev"
	Prod       = "prod"
	Test       = "test"
	Staging    = "staging"
	ListLimit  = 10
	MaxInt     = int(^uint(0) >> 1)
	BinaryName = "backend"
	SiteName   = "Skeleton"
	APIVersion = "v1"
)

type (
	app struct {
		Config                                      config.Configer
		RunMode, Host, FrontEndURL, APIURL, APIPath string
		DisableCache, EnableControls                bool
	}
	Controller struct {
		beego.Controller
	}
	ControllerInterface interface {
		beego.ControllerInterface
	}
	Command interface {
		Run(args []string)
		Help()
	}
)

var (
	App        *app
	Commands   map[string]Command
	TrustedIPs = []string{
		"127.0.0.1",
		"::1",
	}
	allowedHosts = []string{
		"skeleton-gadelkareem.herokuapp.com",
	}
)

func Bootstrap() {
	if App != nil {
		return
	}

	configMode := os.Getenv("BEEGO_RUNMODE")
	if configMode == "" {
		configMode = Prod
	}
	var (
		path string
		err  error
	)
	// go test
	if strings.HasSuffix(os.Args[0], ".test") {
		configMode = Test
		path = filepath.Join(os.Getenv("PWD"), "..")
		for {
			if _, err := os.Stat(filepath.Join(path, "conf")); err == nil {
				break
			}
			path = filepath.Join(path, "..")
		}
	} else {
		path, err = filepath.Abs(filepath.Dir(os.Args[0]))
		h.PanicOnError(err)
	}

	err = os.Chdir(path)
	h.PanicOnError(err)
	os.Args[0] = filepath.Join(path, BinaryName)

	handleProdConfig(configMode)

	err = beego.LoadAppConfig("ini", "conf/app."+configMode+".ini")
	h.PanicOnError(err)

	App = &app{
		Config: beego.AppConfig,
	}

	App.RunMode = App.Config.DefaultString("runmode", Prod)
	beego.BConfig.RunMode = App.RunMode
	beego.BConfig.Listen.HTTPAddr = App.ConfigOrEnvVar("httpaddr", "HTTP_ADDR")
	beego.BConfig.Listen.HTTPPort = App.ConfigOrEnvInt("httpport", "PORT")
	App.EnableControls = App.RunMode != Prod && App.Config.DefaultBool("enableControls", false)
	App.DisableCache = App.EnableControls && App.Config.DefaultBool("disableCache", false)
	// host
	App.Host = App.ConfigOrEnvVar("host", "SKELETON_HOST")
	allowedHosts = append(allowedHosts, App.Host)
	App.FrontEndURL = App.ConfigOrEnvVar("frontEndURL", "SKELETON_FRONTEND")
	App.APIURL = App.Config.String("apiURL")
	App.APIPath = App.Config.String("apiPath")

	// db
	initDBConfig()

}

func handleProdConfig(m string) {
	if m != Prod {
		return
	}
	p := "conf/app.prod.ini.secret"
	if h.FileExists(p) {
		return
	}
	s := os.Getenv("PROD_CONFIG_SECRET_FILE")
	var err error
	if s == "" {
		s, err = h.ReadFile("conf/app.dev.ini.secret.example")
		h.PanicOnError(err)
	} else {
		s, err = h.Base64Decode(s)
		h.PanicOnError(err)
	}
	err = h.WriteFile(p, s)
	h.PanicOnError(err)
}

func (a *app) ConfigOrEnvVar(k, e string) string {
	s := os.Getenv(e)
	if s == "" {
		return App.Config.String(k)
	}
	return s
}

func (a *app) ConfigOrEnvInt(k, e string) int {
	s := os.Getenv(e)
	if s == "" {
		return App.Config.DefaultInt(k, 0)
	}
	i, _ := strconv.Atoi(s)
	return i
}

func (a *app) SetupServer(errController ControllerInterface) {
	// leave it high until https://github.com/golang/go/issues/16100
	beego.BConfig.Listen.ServerTimeOut = 15
	beego.ErrorController(errController)

	a.logging()

	if !IsDev() {
		beego.DelStaticPath("/static")
		beego.DelStaticPath("/favicon.ico")
	}
	beego.SetStaticPath("/", "dist")
	beego.InsertFilter("*", beego.BeforeStatic, func(c *context.Context) {
		p := c.Request.URL.Path
		// CORS
		if c.Input.Method() == http.MethodOptions {
			SetCORS(c)
			c.Output.SetStatus(http.StatusOK)
			logs.AccessLog(&logs.AccessLogRecord{
				RequestMethod: http.MethodOptions,
				Request:       c.Request.URL.String(),
				Host:          c.Request.Host,
				HTTPReferrer:  c.Request.Referer(),
				HTTPUserAgent: c.Request.UserAgent(),
				Status:        http.StatusOK,
			}, "")
			panic(beego.ErrAbort)
		} else if c.Input.Method() == http.MethodGet &&
			!strings.HasSuffix(p, "/") && !strings.HasPrefix(p, "/api") && filepath.Ext(p) == "" {
			c.Request.URL.Path += "/"
			c.Redirect(http.StatusMovedPermanently, c.Request.URL.String())
			panic(beego.ErrAbort)
		}
	}, true)
}

func SetCORS(c *context.Context) {
	c.Output.Header("Access-Control-Allow-Origin", App.FrontEndURL)
	c.Output.Header("Access-Control-Allow-Headers", "content-type,authorization,x-requested-with")
	c.Output.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
	c.Output.Header("Access-Control-Allow-Credentials", "true")
}

func (a *app) logging() {
	if App.RunMode == Prod {
		err := logs.GetBeeLogger().DelLogger(logs.AdapterConsole)
		h.PanicOnError(err)
		err = logs.SetLogger(logs.AdapterConsole, `{"color":false}`)
		h.PanicOnError(err)
	} else {
		logs.EnableFuncCallDepth(true)
	}
	logs.SetLevel(App.Config.DefaultInt("logLevel", logs.LevelWarning))
}

func (a *app) Run(params ...string) {
	beego.Run(params...)
}

func (a *app) IsCLI() bool {
	return len(os.Args) > 1
}

func (a *app) RunCommand() {
	isHelp := func(s string) bool {
		return s == "help" || s == "-h" || s == "--help"
	}
	if !a.IsCLI() {
		return
	}
	args := os.Args
	if isHelp(args[1]) {
		fmt.Printf("Available Commands:\n")
		for k, cmd := range Commands {
			fmt.Printf("----- %s:", k)
			cmd.Help()
		}
		os.Exit(0)
	}
	if cmd, ok := Commands[args[1]]; ok {
		if len(args) > 2 && isHelp(args[2]) {
			cmd.Help()
			os.Exit(0)
		}
		cmd.Run(args[2:])
		os.Exit(0)
	}
}

func IsDev() bool {
	if App == nil {
		return true
	}
	return App.RunMode == Dev
}

func IsIPTrusted(ip string) bool {
	for _, v := range TrustedIPs {
		if ip == v {
			return true
		}
	}
	return false
}

func IsHostAllowed(host string) bool {
	if App.RunMode != Prod {
		return true
	}
	for _, a := range allowedHosts {
		if host == a {
			return true
		}
	}
	return false
}
