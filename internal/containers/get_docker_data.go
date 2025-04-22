package containers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func fetchAllDockerContainers() (Containers, error) {
	// Build a custom client that works on the unix docker sock
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", defaultDockerSockPath)
			},
		},
	}

	request, err := http.NewRequest("GET", "http://docker/containers/json?all=true", nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("sending request to socket: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response status: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("decoding response body: %w", err)
	}

	var cnts []Container

	if err = json.Unmarshal(body, &cnts); err != nil {
		return nil, fmt.Errorf("Unmarshal: %w ", err)
	}

	return cnts, nil
}
