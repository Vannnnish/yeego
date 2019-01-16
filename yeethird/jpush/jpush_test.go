/**
 * Created by angelina on 2017/5/4.
 */

package jpush

import (
	"testing"
)

func TestClient_PushToOne(t *testing.T) {
	client := NewClient(NewClientRequest{
		Name:         "xx",
		AppKey:       "xx",
		Secret:       "xx",
		IsIosProduct: false,
	})
	client.PushToOne("xx", "xx", "xx", nil)
}

func TestClient_PushToAll(t *testing.T) {
	client := NewClient(NewClientRequest{
		Name:         "xx",
		AppKey:       "xx",
		Secret:       "xx",
		IsIosProduct: false,
	})
	client.PushToAll("content")
}
