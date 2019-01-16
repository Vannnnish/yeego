/**
 * Created by angelina-zf on 17/2/28.
 */
package yeecache

import (
	"github.com/vannnnish/yeego"
	"testing"
	"time"
)

func TestNewMemoryCache(t *testing.T) {
	mc := NewMemoryCache(DefaultExpiration, 0)
	a, found := mc.Get("a")
	yeego.Equal(found, false)
	yeego.Equal(a, nil)
	b, found := mc.Get("b")
	yeego.Equal(found, false)
	yeego.Equal(b, nil)
	c, found := mc.Get("c")
	yeego.Equal(found, false)
	yeego.Equal(c, nil)
	mc.Set("a", 1, DefaultExpiration)
	mc.Set("b", "b", DefaultExpiration)
	mc.Set("c", 3.5, DefaultExpiration)
	x, found := mc.Get("a")
	yeego.Equal(found, true)
	yeego.Equal(x, 1)
	y, found := mc.Get("b")
	yeego.Equal(found, true)
	yeego.Equal(y, "b")
	z, found := mc.Get("c")
	yeego.Equal(found, true)
	yeego.Equal(z, 3.5)
}

func TestMemoryCacheTime(t *testing.T) {
	var found bool
	mc := NewMemoryCache(40*time.Millisecond, time.Millisecond)
	mc.Set("a", 1, DefaultExpiration)
	mc.Set("b", 2, NoExpiration)
	mc.Set("c", 3, 30*time.Millisecond)
	mc.Set("d", 4, 50*time.Millisecond)
	<-time.After(30 * time.Millisecond)
	_, found = mc.Get("c")
	yeego.Equal(found, false)
	<-time.After(15 * time.Millisecond)
	_, found = mc.Get("a")
	yeego.Equal(found, false)
	<-time.After(15 * time.Millisecond)
	_, found = mc.Get("d")
	yeego.Equal(found, false)
	_, found = mc.Get("b")
	yeego.Equal(found, true)
}
