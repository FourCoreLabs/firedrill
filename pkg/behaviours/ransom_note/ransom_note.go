package ransom_note

import (
	"context"
	"encoding/base64"
	"os"
	"os/user"
	"path/filepath"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

const (
	ID   = "ransom_note"
	Name = "Ransomware note"

	ransomMessage      = `KkFsbCB5b3VyIGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQhKiAKCkFsbCB5b3VyIGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQgZHVlIHRvIGEgc2VjdXJpdHkgcHJvYmxlbSB3aXRoIHlvdXIgUEMgSUQuIApJZiB5b3Ugd2FudCB0byByZXN0b3JlIHRoZW0sIHdyaXRlIHVzIHRvIHRoZSBlLW1haWw6IGZpcmVkcmlsbEBoaWRlLmJpei5zdCAKCllvdSBoYXZlIHRvIHBheSBmb3IgZGVjcnlwdGlvbiBpbiBCaXRjb2lucy4gVGhlIHByaWNlIGRlcGVuZHMgb24gaG93IGZhc3QgeW91IHdyaXRlIHRvIHVzLiAKQWZ0ZXIgcGF5bWVudCwgd2Ugd2lsbCBzZW5kIHlvdSB0aGUgdG9vbCB0aGF0IHdpbGwgZGVjcnlwdCBhbGwgeW91ciBmaWxlcy4gRnJlZSBkZWNyeXB0aW9uIGFzIGd1YXJhbnRlZS4gQmVmb3JlIHBheWluZyB5b3UgY2FuIHNlbmQgdXMgdXAgdG8gNSBmaWxlcyBmb3IgZnJlZSBkZWNyeXB0aW9uLiAKVGhlIHRvdGFsIHNpemUgb2YgZmlsZXMgbXVzdCBiZSBsZXNzIHRoYW4gNE1iIChub24tYXJjaGl2ZWQpLCBhbmQgZmlsZXMgc2hvdWxkIG5vdCBjb250YWluIHZhbHVhYmxlIGluZm9ybWF0aW9uLiAoZGF0YWJhc2VzLCBiYWNrdXBzLCBsYXJnZSBleGNlbCBzaGVldHMsIGV0Yy4pIEF0dGVudGlvbiEgRG8gbm90IHJlbmFtZSBlbmNyeXB0ZWQgZmlsZXMuIApEbyBub3QgdHJ5IHRvIGRlY3J5cHQgeW91ciBkYXRhIHVzaW5nIHRoaXJkLXBhcnR5IHNvZnR3YXJlLCBpdCBtYXkgY2F1c2UgcGVybWFuZW50IGRhdGEgbG9zcy4gClRoZSBkZWNyeXB0aW9uIG9mIHlvdXIgZmlsZXMgd2l0aCB0aGUgaGVscCBvZiB0aGlyZCBwYXJ0aWVzIG1heSBjYXVzZSBpbmNyZWFzZWQgcHJpY2UgKHRoZXkgYWRkIHRoZWlyIGZlZSB0byBvdXIpIG9yIHlvdSBjYW4gYmVjb21lIGEgdmljdGltIG9mIGEgc2NhbS4gQ2hlZXJzIQ==`
	ransomNoteFileName = "ransomnote.txt"
)

type RansomNoteOptions struct {
	NoteFileName string
	Note         string
}

type RansomNote struct {
	note         string
	noteFileName string
	base64       bool
}

func NewRansomNote(opts ...RansomNoteOptions) sergeant.Runnable {
	var options RansomNoteOptions = RansomNoteOptions{
		Note:         ransomMessage,
		NoteFileName: ransomNoteFileName,
	}

	if len(opts) > 0 {
		options = opts[0]
	}

	return &RansomNote{note: options.Note, noteFileName: options.NoteFileName, base64: true}
}

func (e *RansomNote) ID() string {
	return ID
}

func (e *RansomNote) Name() string {
	return Name
}

func (e *RansomNote) Run(ctx context.Context, logger *zap.Logger) error {
	desktopPath := UserDesktop()
	logger.Sugar().Infof("User desktop path for ransom note: %s", desktopPath)

	ransomNoteFileName := e.noteFileName

	ransomNotePath := filepath.Join(desktopPath, ransomNoteFileName)

	if e.base64 {
		uDec, _ := base64.URLEncoding.DecodeString(e.note)

		e.note = string(uDec)
	}

	if err := os.WriteFile(ransomNotePath, []byte(e.note), 0644); err != nil {
		logger.Sugar().Warnf("Failed to drop ransom note at %s: %s", ransomNotePath, err.Error())
		return err
	}

	logger.Sugar().Infof("Dropped ransom note at %s", ransomNotePath)

	return nil
}

func UserDesktop() string {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	homeDir := curUser.HomeDir
	return filepath.Join(homeDir, "Desktop")
}
