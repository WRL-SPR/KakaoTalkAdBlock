//go:build windows

package win

import "testing"

func TestStartupCommandQuotesExecutablePath(t *testing.T) {
	originalExecutablePath := executablePath
	executablePath = func() (string, error) {
		return `C:\Users\Test User\Apps\KakaoGuard.exe`, nil
	}
	t.Cleanup(func() {
		executablePath = originalExecutablePath
	})

	got, err := startupCommand()
	if err != nil {
		t.Fatal(err)
	}

	want := `"C:\Users\Test User\Apps\KakaoGuard.exe" --startup`
	if got != want {
		t.Fatalf("startupCommand() = %q, want %q", got, want)
	}
}

func TestIsStartupApprovedState(t *testing.T) {
	tests := []struct {
		name  string
		state []byte
		want  bool
	}{
		{name: "enabled", state: []byte{0x02, 0x00, 0x00, 0x00}, want: true},
		{name: "disabled", state: []byte{0x03, 0x00, 0x00, 0x00}, want: false},
		{name: "missing data", state: nil, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStartupApprovedState(tt.state); got != tt.want {
				t.Fatalf("isStartupApprovedState(%v) = %v, want %v", tt.state, got, tt.want)
			}
		})
	}
}
