package barreldb

import "strings"

// Register a label (type).  This is used by things that wrap the
// barrel.
func (b *Barrel) Register(label string) error {
	b.Lock()
	defer b.Unlock()

	label = strings.ToLower(label)
	_, exists := b.labels[label]
	if exists {
		return ErrDuplicateLabel
	}
	b.labels[label] = true
	return nil
}

// Unregister a lable (type)
func (b *Barrel) Unregister(label string) {
	b.Lock()
	defer b.Unlock()

	label = strings.ToLower(label)
	delete(b.labels, label)
}
