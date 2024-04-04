package memory_test

import (
	"github.com/JohnsonSmile/jet/pkg/cache"
	"github.com/JohnsonSmile/jet/pkg/cache/memory"
	"github.com/stretchr/testify/require"

	"testing"
)

func ExampleMemoryCache() {

}

func TestMemoryCache(t *testing.T) {
	mc := memory.NewMemoryCache()
	err := mc.SetMaxMemory(100 * cache.MB)
	require.NoError(t, err)
	err = mc.Set("int", 1)
	require.NoError(t, err, "mc.Set(\"int\", 1)")
	err = mc.Set("bool", false)
	require.NoError(t, err, "mc.Set(\"bool\", false)")
	err = mc.Set("data", map[string]any{"a": 1})
	require.NoError(t, err, "mc.Set(\"data\", map[string]any{\"a\": 1})")
	intVal, err := mc.Get("int")
	require.NoError(t, err, "mc.Get(\"int\")")
	require.Equal(t, 1, intVal, "mc.Get(\"int\")")
	err = mc.Del("int")
	count := mc.Count()
	require.Equal(t, 3, count, "mc.Count()")
	err = mc.Flush()
	require.NoError(t, err, "mc.Flush()")
	count = mc.Count()
}
