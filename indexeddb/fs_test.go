// +build wasm

package indexeddb

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/hack-pad/go-indexeddb/idb"
	"github.com/hack-pad/hackpadfs/fstest"
	"github.com/hack-pad/hackpadfs/internal/assert"
)

const (
	testDBPrefix = "hackpadfs-test-"
)

func TestFS(t *testing.T) {
	t.Parallel()
	options := fstest.FSOptions{
		Name: "indexeddb",
		TestFS: func(tb testing.TB) fstest.SetupFS {
			n, err := rand.Int(rand.Reader, big.NewInt(1000))
			assert.NoError(tb, err)
			name := fmt.Sprintf("%s%s/%d", testDBPrefix, tb.Name(), n.Int64())

			factory := idb.Global()

			fs, err := NewFS(name, factory)
			if err != nil {
				tb.Fatal(err)
			}
			tb.Cleanup(func() {
				req, err := factory.DeleteDatabase(name)
				assert.NoError(tb, err)
				assert.NoError(tb, req.Await(context.Background()))
			})
			return fs
		},
	}
	fstest.FS(t, options)
	fstest.File(t, options)
}