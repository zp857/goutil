package urlutil

import (
	"net/url"
	"path/filepath"
	"strings"
)

func GetFileName(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	file := filepath.Base(u.Path)
	return file
}

func GetFileWithoutExt(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	u.RawQuery = ""
	file := filepath.Base(u.Path)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	return file
}

func GetFileExt(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	u.RawQuery = ""
	ext := filepath.Ext(u.Path)
	ext = strings.ToLower(ext)
	return ext
}
