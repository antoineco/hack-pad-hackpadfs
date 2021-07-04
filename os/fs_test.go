package os

import (
	"strings"
	"testing"

	"github.com/hack-pad/hackpadfs/fstest"
	"github.com/hack-pad/hackpadfs/internal/assert"
)

func TestFSTest(t *testing.T) {
	oldmask := setUmask(0)
	t.Cleanup(func() {
		setUmask(oldmask)
	})

	fstest.FS(t, fstest.FSOptions{
		Name: "osfs.FS",
		TestFS: func(tb testing.TB) fstest.SetupFS {
			dir := tb.TempDir()
			dir = strings.TrimPrefix(dir, "/") // TODO support Windows root path
			fs, err := NewFS().Sub(dir)
			if !assert.NoError(tb, err) {
				tb.FailNow()
			}
			return fs.(*FS)
		},
	})
}