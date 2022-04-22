package env

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// Env custom environment type
type Env int8

const (
	// Unknown environment
	Unknown Env = iota
	// Local environment
	Local
	// Dev development environment
	Dev
	// Test environment
	Test
	// Prod production environment
	Prod
	// Docker docker environment
	Docker
	// K8s k8s environment
	K8s
)

// Info is a struct that contains the environment information.
type Info struct {
	// once is used to ensure that the environment is only set once.
	once sync.Once
	// EnvType is the type of the environment.
	EnvType Env
	// AppRootPath is the home directory of the application context.
	AppRootPath string
	// Context for extra
	Context context.Context
}

var env = &Info{Context: context.Background()}

// GetInfo returns the environment information.
func GetInfo() *Info {
	ensure()
	return env
}

// IsLocal is local environment
func IsLocal() bool {
	return GetInfo().EnvType == Local
}

// IsDev is development environment
func IsDev() bool {
	return GetInfo().EnvType == Dev
}

// IsTest is test environment
func IsTest() bool {
	return GetInfo().EnvType == Test
}

// IsProd is production environment
func IsProd() bool {
	return GetInfo().EnvType == Prod
}

// IsDocker is docker environment
func IsDocker() bool {
	return GetInfo().EnvType == Docker
}

// IsK8s is k8s environment
func IsK8s() bool {
	return GetInfo().EnvType == K8s
}

// TypeString returns the string representation of the environment type.
func (e *Info) TypeString() string {
	return Type2String(e.EnvType)
}

// Type2String returns the string representation of the given environment type.
func Type2String(t Env) string {
	switch t {
	case Local:
		return "local"
	case Dev:
		return "dev"
	case Test:
		return "test"
	case Prod:
		return "prod"
	case Docker:
		return "docker"
	case K8s:
		return "k8s"
	case Unknown:
		return "unknown"
	default:
		return "unknown"
	}
}

// AppRootPath return the context path
func AppRootPath() string {
	return GetInfo().AppRootPath
}

// Init initializes the environment information.
func Init() error {
	var gErr error
	env.once.Do(func() {
		iMode := os.Getenv("DEV_UP_ENV_MODE")
		if iMode == "" {
			iMode = "0"
		}
		var mode int
		mode, gErr = strconv.Atoi(iMode)
		if gErr != nil {
			return
		}
		env.EnvType = Env(mode)
		var appAbsPath string
		appAbsPath, gErr = getAppAbsPath()
		if gErr != nil {
			return
		}
		env.AppRootPath = appAbsPath
	})
	return gErr
}

func ensure() {
	if err := Init(); err != nil {
		panic(err)
	}
}

// SetContextVal set env info's context value
func SetContextVal(key any, value any) {
	info := GetInfo()
	parent := info.Context
	info.Context = context.WithValue(parent, key, value)
}

// GetContextVal get env info's context value by key
func GetContextVal(key any) any {
	info := GetInfo()
	ctx := info.Context
	return ctx.Value(key)
}

func getAppAbsPath() (string, error) {
	var root string
	var gErr error
	if root, gErr = getAbsPathByExecutable(); gErr != nil {
		if root, gErr = getAbsPathByCaller(); gErr != nil {
			return root, gErr
		}
	}
	// judge if is test debugging
	tempDir, _ := getTempDir()
	if strings.Contains(root, tempDir) {
		env.EnvType = Local
		root, gErr = os.Getwd()
		if gErr != nil {
			return root, gErr
		}
	}
	return root, gErr
}

func getTempDir() (string, error) {
	return filepath.EvalSymlinks(os.TempDir())
}

func getAbsPathByCaller() (string, error) {
	_, fileName, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Dir(fileName), nil
	} else {
		return "", errors.New("get abs path by caller error")
	}
}

func getAbsPathByExecutable() (string, error) {
	var executablePath string
	exAbs, err := os.Executable()
	if err != nil {
		return executablePath, err
	}
	executablePath, err = filepath.EvalSymlinks(exAbs)
	if err != nil {
		return executablePath, err
	}
	return filepath.Dir(executablePath), err
}
