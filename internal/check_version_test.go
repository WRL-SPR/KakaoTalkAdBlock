package internal

import "testing"

func TestHasNewRelease(t *testing.T) {
	tests := []struct {
		name    string
		current string
		latest  string
		want    bool
	}{
		{name: "new patch", current: "1.0.0", latest: "1.0.1", want: true},
		{name: "same version", current: "1.0.0", latest: "1.0.0", want: false},
		{name: "older release", current: "1.1.0", latest: "1.0.9", want: false},
		{name: "v prefix", current: "1.0.0", latest: "v1.1.0", want: true},
		{name: "build metadata", current: "1.0.0+private", latest: "1.0.0", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasNewRelease(tt.current, tt.latest); got != tt.want {
				t.Fatalf(
					"hasNewRelease(%q, %q) = %v, want %v",
					tt.current,
					tt.latest,
					got,
					tt.want,
				)
			}
		})
	}
}
