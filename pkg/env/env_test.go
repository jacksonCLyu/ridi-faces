package env

import "testing"

func TestGetInfo(t *testing.T) {
	info := GetInfo()
	if info.EnvType != Local {
		t.Errorf("env.GetInfo().EnvType != Local")
	}
	if info.AppRootPath == "" {
		t.Errorf("env.GetInfo().AppRootPath == \"\"")
	} else {
		t.Logf("env.GetInfo().AppRootPath = %s", info.AppRootPath)
	}
}
