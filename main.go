package main

import (
	"github.com/guoyk93/gg"
	"github.com/guoyk93/gg/ggos"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var (
		optDir  = "/data"
		optLoop = "0"
	)

	ggos.MustEnv("TRUNCLOG_DIR", &optDir)
	ggos.MustEnv("TRUNCLOG_LOOP", &optLoop)

redo:

	err := fs.WalkDir(os.DirFS(optDir), "/", func(path string, d fs.DirEntry, err error) (_ error) {
		if err != nil {
			gg.Log("failed walking: " + path + ": " + err.Error())
			return
		}
		if d.IsDir() {
			return
		}
		if !strings.HasSuffix(strings.ToLower(path), ".log") {
			return
		}
		if err = os.Truncate(filepath.Join(optDir, path), 0); err != nil {
			gg.Log("failed truncating: " + path + ": " + err.Error())
			return
		}
		gg.Log("truncated: " + path)
		return
	})

	if err != nil {
		gg.Log("failed walking: " + err.Error())
	}

	if loop, _ := time.ParseDuration(optLoop); loop > 0 {
		time.Sleep(loop)
		goto redo
	}
}
