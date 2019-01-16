/**
 * Created by angelina on 2017/4/21.
 */

package cache

import (
	"testing"
	"time"
	"github.com/vannnnish/yeego"
)

func TestFileTtlCache(t *testing.T) {
	cachePath := "cache/test.cache"
	d, err := FileTtlCache(cachePath, func() (b []byte, ttl time.Duration, err error) {
		return []byte("1"), time.Second, nil
	})
	yeego.Equal(err, nil)
	yeego.Equal(d, []byte("1"))
	d, err = FileTtlCache(cachePath, func() (b []byte, ttl time.Duration, err error) {
		return []byte("2"), time.Second, nil
	})
	yeego.Equal(err, nil)
	yeego.Equal(d, []byte("1"))
	time.Sleep(time.Second)
	d, err = FileTtlCache(cachePath, func() (b []byte, ttl time.Duration, err error) {
		return []byte("2"), time.Second, nil
	})
	yeego.Equal(err, nil)
	yeego.Equal(d, []byte("2"))
}
