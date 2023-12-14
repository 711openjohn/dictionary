package client

import "testing"

func TestCambridge(t *testing.T) {
	cambridge := NewCambridge()
	cambridge.Lookup("test", "chinese-traditional")
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	if exitCode != 0 {
	}
}
