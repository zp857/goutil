package urlutil

import (
	"fmt"
	"testing"
)

func TestGetFileExt(t *testing.T) {
	t.Run("GetFileExt", func(t *testing.T) {
		ext := GetFileExt("https://10.1.2.130/ui/bower_components/vui-bootstrap/css/vui-bootstrap.min.css")
		fmt.Println(ext)
	})
}
