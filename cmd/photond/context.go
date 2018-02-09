package main

import (
	"errors"

	"github.com/spf13/cobra"
	tmconfig "github.com/tendermint/tendermint/config"
)

const (
	flagRootDir = "data"
)

type cmdRunner func(cmd *cobra.Command, args []string) error
type ctxRunner func(ctx Context, cmd *cobra.Command, args []string) error

func withContext(fn ctxRunner) cmdRunner {
	return func(cmd *cobra.Command, args []string) error {
		ctx := newContext(cmd)
		return fn(ctx, cmd, args)
	}
}

func requireRootDir(fn ctxRunner) ctxRunner {
	return func(ctx Context, cmd *cobra.Command, args []string) error {
		if root := ctx.RootDir(); root == "" {
			return errors.New("root directory unset")
		}
		return fn(ctx, cmd, args)
	}
}

type Context interface {
	RootDir() string
	TMConfig() (*tmconfig.Config, error)
}

type context struct {
	cmd   *cobra.Command
	tmcfg *tmconfig.Config
}

func newContext(cmd *cobra.Command) Context {
	return &context{cmd: cmd}
}

func (ctx *context) RootDir() string {
	root, _ := ctx.cmd.Flags().GetString(flagRootDir)
	return root
}

func (ctx *context) TMConfig() (*tmconfig.Config, error) {
	if ctx.tmcfg != nil {
		return ctx.tmcfg, nil
	}
	root := ctx.RootDir()
	if root == "" {
		return nil, errors.New("root dir required")
	}
	cfg := tmconfig.DefaultConfig()
	cfg.SetRoot(root)
	tmconfig.EnsureRoot(root)
	return cfg, nil
}
