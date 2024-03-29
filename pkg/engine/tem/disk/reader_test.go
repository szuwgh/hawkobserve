package disk

import (
	"testing"

	"github.com/szuwgh/hawkobserve/pkg/engine/tem/cache"
)

func Test_IndexReaderPrint(t *testing.T) {
	bcache := cache.NewCache(cache.NewLRU(DefaultBlockCacheCapacity))
	tcache := cache.NewCache(cache.NewLRU(defaultSegmentSize * 10))
	b := &cache.NamespaceGetter{Cache: bcache, NS: 1}
	m := &cache.NamespaceGetter{Cache: tcache, NS: 1}
	reader1 := NewIndexReader("/opt/goproject/hawkobserve/src/github.com/szuwgh/hawkobserve/data/01GJ4ZCT4ZVC69F95QHRYPYADV/index", 6, b, m)
	reader1.print()
}
