package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	log.SetFlags(0)

	// Everything should work without them.
	os.Unsetenv("GOPATH")
	os.Unsetenv("GO111MODULE")

	steps := []struct {
		name string
		fn   func() error
	}{
		{"check go version", checkGoVersion},
		{"install ruleguard to this module", installRuleguard},
		{"run ruleguard", runRuleguard},
		{"install rules bundle", installRulesBundle},
		{"run ruleguard with bundle rules", runRuleguardWithBundleRules},
		{"run gorule tests", runTests},
	}
	for _, step := range steps {
		log.Printf("* running %s step", step.name)
		if err := step.fn(); err != nil {
			log.Fatalf("* * error: %s: %v", step.name, err)
		}
	}
	log.Printf("Everything looks good!")
}

func checkGoVersion() error {
	// Could use runtime.Version(), but it's also important to check
	// that "go" binary is available through the systam PATH.
	// ruleguard may invoke `go list` in some cases.
	out, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	ok := bytes.Contains(out, []byte("go1.17")) ||
		bytes.Contains(out, []byte("go1.16")) ||
		bytes.Contains(out, []byte("go1.15")) ||
		bytes.Contains(out, []byte("devel"))
	if !ok {
		return fmt.Errorf("unrecognized Go version: %s", out)
	}
	return nil
}

func installRuleguard() error {
	out, err := exec.Command("go", "get", "github.com/quasilyte/go-ruleguard").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	out, err = exec.Command("go", "get", "github.com/quasilyte/go-ruleguard/dsl").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	out, err = exec.Command("go", "build", "-o", "ruleguard.exe", "github.com/quasilyte/go-ruleguard/cmd/ruleguard").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	return nil
}

func runRuleguard() error {
	absPath, err := filepath.Abs("ruleguard.exe")
	if err != nil {
		return err
	}
	out, err := exec.Command(absPath, "-rules", "rules.go", filepath.Join("test_target", "example.go")).CombinedOutput()
	// Linting errors give a non-zero exit, but we're expecting a warning here.
	// So we check for the output *before* checking for an error.
	if bytes.Contains(out, []byte("nolintGocritic: that is heartbreaking")) {
		return nil // OK
	}
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	return fmt.Errorf("unexpected ruleguard output: %s", out)
}

func installRulesBundle() error {
	out, err := exec.Command("go", "get", "github.com/quasilyte/go-ruleguard/rules").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	return nil
}

func runRuleguardWithBundleRules() error {
	absPath, err := filepath.Abs("ruleguard.exe")
	if err != nil {
		return err
	}
	out, err := exec.Command(absPath, "-rules", "rules2.go", filepath.Join("test_target", "example.go")).CombinedOutput()
	if bytes.Contains(out, []byte("emptyDecl: empty var() block")) {
		return nil // OK
	}
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	return fmt.Errorf("unexpected ruleguard output: %s", out)
}

func runTests() error {
	out, err := exec.Command("go", "test", "-v", "./analysistest").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}
	return nil
}
