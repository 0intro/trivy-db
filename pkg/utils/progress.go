package utils

// The upstream version of this file imports github.com/cheggaaa/pb/v3 and
// github.com/briandowns/spinner to render a CLI progress UI while
// trivy-db is downloading or building the vulnerability database. Both
// are never reached by the DataDog Agent (which consumes the prebuilt
// database), and cheggaaa/pb/v3 in particular reaches text/template ->
// reflect.Value.MethodByName at init, defeating the Go linker's DCE.
//
// Stubbing here keeps trivy-db's other packages compiling against the
// same API while breaking the dependency chain.

// Quiet exists for source compatibility; it has no effect.
var Quiet = false

// Spinner is a no-op stub.
type Spinner struct{}

// NewSpinner returns an empty Spinner.
func NewSpinner(_ string) *Spinner { return &Spinner{} }

// Start is a no-op.
func (s *Spinner) Start() {}

// Stop is a no-op.
func (s *Spinner) Stop() {}

// ProgressBar is a no-op stub.
type ProgressBar struct{}

// NewProgressBar returns an empty ProgressBar.
func NewProgressBar(_ int) *ProgressBar { return &ProgressBar{} }

// Increment is a no-op.
func (p *ProgressBar) Increment() {}

// Finish is a no-op.
func (p *ProgressBar) Finish() {}
