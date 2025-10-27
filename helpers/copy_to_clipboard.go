package helpers

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func CopyToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	if err != nil {
		return fmt.Errorf("failed to copy to clipboard: %w", err)
	}
	return nil
}
