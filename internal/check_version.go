package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const latestReleaseURL = "https://api.github.com/repos/WRL-SPR/KakaoGuard/releases/latest"

var versionHTTPClient = &http.Client{Timeout: 5 * time.Second}

func CheckLatestVersion(currentVersion string) (string, bool) {
	response, err := versionHTTPClient.Get(latestReleaseURL)
	if err != nil {
		return currentVersion, false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return currentVersion, false
	}

	var data struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil || data.TagName == "" {
		return currentVersion, false
	}

	if hasNewRelease(currentVersion, data.TagName) {
		return data.TagName, true
	}

	return currentVersion, false
}

func hasNewRelease(current, latest string) bool {
	v1Parts := versionParts(current)
	v2Parts := versionParts(latest)

	for i := 0; i < len(v1Parts) && i < len(v2Parts); i++ {
		v1Part, err := strconv.Atoi(v1Parts[i])
		if err != nil {
			return false
		}
		v2Part, err := strconv.Atoi(v2Parts[i])
		if err != nil {
			return false
		}

		if v1Part > v2Part {
			return false
		}
		if v1Part < v2Part {
			return true
		}
	}
	return len(v2Parts) > len(v1Parts)
}

func versionParts(version string) []string {
	version = strings.TrimPrefix(strings.TrimSpace(version), "v")
	if suffix := strings.IndexAny(version, "-+"); suffix >= 0 {
		version = version[:suffix]
	}
	return strings.Split(version, ".")
}
