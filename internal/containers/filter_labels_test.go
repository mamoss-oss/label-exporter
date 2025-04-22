package containers

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_filter_container_labels(t *testing.T) {
	input := Container{
		Names: []string{"Container1"},
		Image: "busybox",
		State: "running",
		Labels: map[string]string{
			"a": "1",
			"b": "2",
		},
	}
	var prefixes = []string{"a", "b"}

	expected := labels{
		"a": "1",
		"b": "2",
	}

	testname := fmt.Sprint("Filter for container labels")

	t.Run(testname, func(t *testing.T) {
		ans := filter_container_labels(input, prefixes)
		if !reflect.DeepEqual(ans, expected) {
			t.Errorf("got %v, want %v", ans, expected)
		}
	})

}
