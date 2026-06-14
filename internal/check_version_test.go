package internal

import "testing"

func TestHasNewRelease(t *testing.T) {
	tests := []struct {
		name    string
		current string
		latest  string
		want    bool
	}{
		{name: "new patch", current: "2.2.3", latest: "2.2.4", want: true},
		{name: "same version", current: "2.2.3", latest: "2.2.3", want: false},
		{name: "older release", current: "2.3.0", latest: "2.2.9", want: false},
		{name: "v prefix", current: "2.2.3", latest: "v2.3.0", want: true},
		{name: "build metadata", current: "2.2.3+ux1", latest: "2.2.3", want: false},
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
