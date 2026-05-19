package override

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/oops"
	"gopkg.in/yaml.v3"
)

// Config represents the override configuration file
type Config struct {
	Patches []PatchEntry `yaml:"patches"`
}

// PatchEntry represents a single patch entry
type PatchEntry struct {
	Target string `yaml:"target"` // Path suffix to match (e.g., "ghsa/2025/11/GHSA-xxxx.json")
	Diff   string `yaml:"diff"`   // Path to jd diff file (relative to overrides dir)
}

// Patches holds loaded override configuration
type Patches struct {
	entries      []PatchEntry
	overridesDir string
}

// Patch is a no-op stub. The upstream version wraps a jd.Diff and uses
// github.com/josephburnett/jd/v2 to apply JSON patches. The agent only
// consumes the prebuilt database, so this codepath is never reached
// at runtime.
type Patch struct{}

// Load reads config.yaml from the given directory
func Load(overridesDir string) (*Patches, error) {
	eb := oops.With("overrides_dir", overridesDir)

	configPath := filepath.Join(overridesDir, "config.yaml")
	f, err := os.Open(configPath)
	if err != nil {
		return nil, eb.Wrapf(err, "failed to open config file")
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, eb.Wrapf(err, "failed to parse config.yaml")
	}

	patches := &Patches{
		entries:      make([]PatchEntry, 0, len(cfg.Patches)),
		overridesDir: overridesDir,
	}

	for _, p := range cfg.Patches {
		eb := eb.With("target", p.Target).With("diff", p.Diff)
		if p.Diff == "" {
			return nil, eb.Errorf("patch entry missing 'diff' field")
		} else if !filepath.IsLocal(p.Diff) {
			return nil, eb.Errorf("diff path must be local")
		}

		target := filepath.ToSlash(p.Target) // Normalize to forward slashes
		if !strings.HasPrefix(target, "/") {
			return nil, eb.Errorf("target path must start with '/'")
		}

		patches.entries = append(patches.entries, PatchEntry{
			Target: target,
			Diff:   p.Diff,
		})
	}

	return patches, nil
}

// Match never returns a matching patch in this stripped build, so the
// jd-based diff/apply codepath is unreachable and the dependency drops
// out of the link.
func (p *Patches) Match(_ string) (*Patch, bool, error) {
	return nil, false, nil
}

// Apply returns the input unchanged. Unreachable in practice because
// Match never returns true.
func (p *Patch) Apply(original []byte) ([]byte, error) {
	return original, nil
}

// Count returns the number of patch entries
func (p *Patches) Count() int {
	if p == nil {
		return 0
	}
	return len(p.entries)
}
