package check

import (
	"os/exec"
)

type GoFmtCheck struct {
	Dir string
}

// Percentage returns the percentage of .go files that pass gofmt
func (g GoFmtCheck) Percentage() (float64, error) {
	files, err := GoFiles(g.Dir)
	if err != nil {
		return 0, nil
	}
	var failed []string
	for _, fi := range files {
		out, err := exec.Command("gofmt", "-l", fi).Output()
		if err != nil {
			return 0, nil
		}
		if string(out) != "" {
			failed = append(failed, fi)
		}
	}
	return float64(len(files)-len(failed)) / float64(len(files)), nil
}