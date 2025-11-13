package domain

import "fmt"

type Status string

const (
	StatusDraft   Status = "draft"
	StatusPublish Status = "publish"
)

func ValidateStatus(status Status) error {
	if status != StatusDraft && status != StatusPublish {
		return fmt.Errorf("must be either 'draft' or 'publish'")
	}
	return nil
}
