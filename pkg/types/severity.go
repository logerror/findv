package types

import "fmt"

const (
	SeverityUnknown Severity = iota
	SeverityLow
	SeverityMedium
	SeverityHigh
	SeverityCritical
)

type Severity int

var (
	SeverityNames = []string{
		"UNKNOWN",
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
)

func NewSeverity(severity string) (Severity, error) {
	for i, name := range SeverityNames {
		if severity == name {
			return Severity(i), nil
		}
	}
	return SeverityUnknown, fmt.Errorf("unknown severity: %s", severity)
}

func CompareSeverityString(sev1, sev2 string) int {
	s1, _ := NewSeverity(sev1)
	s2, _ := NewSeverity(sev2)
	return int(s2) - int(s1)
}
