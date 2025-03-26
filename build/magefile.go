//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/autonomouskoi/mageutil"
	"github.com/magefile/mage/sh"
)

var (
	baseDir string
	Default = Protos
)

func init() {
	var err error
	baseDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	baseDir = filepath.Join(baseDir, "..")
}

func Clean() error {
	pbFiles, err := mageutil.DirGlob(baseDir, "*.pb.go")
	if err != nil {
		return fmt.Errorf("globbing %s: %w", baseDir, err)
	}
	for _, pbFile := range pbFiles {
		if err := sh.Rm(filepath.Join(baseDir, pbFile)); err != nil {
			return fmt.Errorf("removing %s: %w", pbFile, err)
		}
	}
	return nil
}

func Protos() error {
	akDir := filepath.Join(baseDir, "..")
	twitchDir := filepath.Join(akDir, "twitch")

	protos, err := mageutil.DirGlob(twitchDir, "*.proto")
	if err != nil {
		return fmt.Errorf("globbing %s: %w", twitchDir)
	}

	for _, protoFile := range protos {
		srcPath := filepath.Join(twitchDir, protoFile)
		dstPath := filepath.Join(baseDir, strings.TrimSuffix(protoFile, ".proto")+".pb.go")
		err := mageutil.TinyGoProto(dstPath, srcPath, twitchDir)
		if err != nil {
			return fmt.Errorf("generating from %s: %w", srcPath, err)
		}
	}

	return nil
}
