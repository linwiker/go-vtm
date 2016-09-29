package stingray

import (
	"fmt"
	"log"
	"strings"
)

// fileResource represents a file resource.
type fileResource struct {
	resource
	Content []byte
	Note    string
}

func (f *fileResource) String() string {
	if f.GetNote() == "" {
		return fmt.Sprintf("#=-%v\n", f.Note) + string(f.Content)
	} else {
		return string(f.Content)
	}
}

func (f *fileResource) GetNote() string {
	firstln := strings.Split(string(f.Content), "\n")[0]
	if strings.Contains(firstln, "#=-") {
		log.Printf("[DEBUG] Found note in resource: %s", firstln)
		return string([]rune(firstln)[3:])
	}
	log.Printf("[DEBUG] Did not find note in resource: %s", firstln)
	return f.Note
}

func (f *fileResource) decode(data []byte) error {
	f.Content = data
	return nil
}

func (f *fileResource) contentType() string {
	return "application/octet-stream"
}
