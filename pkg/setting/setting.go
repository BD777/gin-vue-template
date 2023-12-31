package setting

import (
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize    int
	IdentityKey string
)

func getAppRoot() (string, error) {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(b))))
}

func init() {
	var err error
	var appRoot string
	appRoot, err = getAppRoot()
	if err != nil {
		log.Fatalf("Fail to get app root: %v", err)
	}
	// log app.ini with absolute path
	log.Printf("app.ini path: %s", filepath.Join(appRoot, "conf/app.ini"))
	Cfg, err = ini.Load(filepath.Join(appRoot, "conf/app.ini"))
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	IdentityKey = sec.Key("IDENTITY_KEY").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
