/**
 * Created by angelina on 2017/4/17.
 */

package yeeStrconv

import (
	"github.com/yeeyuntech/yeego"
	"testing"
)

func TestAtoIDefault0(t *testing.T) {
	yeego.Equal(AtoIDefault0("a"), 0)
	yeego.Equal(AtoIDefault0("3"), 3)
}

func TestFormatInt(t *testing.T) {
	yeego.Equal(FormatInt(1), "1")
}

func TestFormatFloat(t *testing.T) {
	yeego.Equal(FormatFloat(3.33), "3.33")
}

func TestFormatFloatPrec0(t *testing.T) {
	yeego.Equal(FormatFloatPrec0(3.333), "3")
}

func TestFormatFloatPrec2(t *testing.T) {
	yeego.Equal(FormatFloatPrec2(3.333), "3.33")
}

func TestFormatFloatPrec4(t *testing.T) {
	yeego.Equal(FormatFloatPrec4(3.3331111), "3.3331")
}
