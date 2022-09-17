package html_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/gopl/ch07/html"
)

func TestFindLinksFromString(t *testing.T) {
	t.Parallel()

	contents := `<html>
	<a href="http://yshibata.blog.so-net.ne.jp"/>
	<a href="http://www001.upp.so-net.ne.jp/yshibata" />
	</html>`

	want := []string{
		"http://yshibata.blog.so-net.ne.jp",
		"http://www001.upp.so-net.ne.jp/yshibata",
	}
	got := html.LinksFromString(contents)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
