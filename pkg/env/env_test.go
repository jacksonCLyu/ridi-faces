package env

import "testing"

func TestGetInfo(t *testing.T) {
	info := GetInfo()
	if info.envType != Local {
		t.Errorf("env.GetInfo().EnvType != Local")
	}
	if info.appRootPath == "" {
		t.Errorf("env.GetInfo().AppRootPath == \"\"")
	} else {
		t.Logf("env.GetInfo().AppRootPath = %s", info.appRootPath)
	}
}
