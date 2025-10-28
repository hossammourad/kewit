package helpers

import (
	"fmt"
)

func InputToInt(input string) (int, error) {
	var id int
	_, err := fmt.Sscanf(input, "%d", &id)
	if err == nil && input != fmt.Sprintf("%d", id) {
		return 0, fmt.Errorf("invalid input: %q", input)
	}
	if err != nil {
		return 0, err
	}
	return id, nil
}
