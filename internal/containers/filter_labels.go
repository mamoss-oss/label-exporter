package containers

import (
	"slices"
)

// filter_container_labels matches labels for desired prefixes. Returns empty map when no matches are found.
func filter_container_labels(container Container, prefixes []string) labels {
	if len(container.Labels) == 0 {
		return labels{}
	}
	var matched_labels = make(labels)

	for k, v := range container.Labels {
		if slices.Contains(prefixes, k) {
			matched_labels[k] = v
		}
	}
	return matched_labels
}
