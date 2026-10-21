package core

import (
	"io"
	"os"
)

// newStdout 统一输出到标准输出（容器/本地都可直接采集）
func newStdout() io.Writer {
	return os.Stdout
}
