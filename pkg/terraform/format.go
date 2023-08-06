package terraform

import (
	"fmt"
)

func (c *Client) Format() error {
	if err := c.tf.FormatWrite(c.ctx); err != nil {
		return fmt.Errorf("error formatting terraform: %w", err)
	}

	return nil
}
