package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

var (
	Quiet = false
)

type Spinner struct {
	client *spinner.Spinner
}

func NewSpinner(suffix string) *Spinner {
	if Quiet {
		return &Spinner{}
	}
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Suffix = suffix
	return &Spinner{client: s}
}

func (s *Spinner) Start() {
	if s.client == nil {
		return
	}
	s.client.Start()
}
func (s *Spinner) Stop() {
	if s.client == nil {
		return
	}
	s.client.Stop()
}

// ProgressBar is a no-op stub. The upstream version wraps
// github.com/cheggaaa/pb/v3.ProgressBar, whose package-level
//
//	var elements = map[string]Element{ "percent": ElementPercent, ... }
//
// converts ElementFunc -> Element at init, creating an itab that keeps
// (*ProgressBar).render -> text/template -> reflect.Value.MethodByName
// reachable. Once any REFLECTMETHOD function is reachable, the Go
// linker keeps every exported method of every reachable <UsedInIface>
// type — see cmd/link/internal/ld/deadcode.go. The agent only consumes
// the prebuilt database, so the progress UI is never used at runtime.
type ProgressBar struct{}

// NewProgressBar returns an empty ProgressBar.
func NewProgressBar(_ int) *ProgressBar { return &ProgressBar{} }

// Increment is a no-op.
func (p *ProgressBar) Increment() {}

// Finish is a no-op.
func (p *ProgressBar) Finish() {}
