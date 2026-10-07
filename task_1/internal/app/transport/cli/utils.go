package cli

import (
	"flag"
	"fmt"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

func parseFlags(flags *flag.FlagSet, args []string) error {
	if err := flags.Parse(args); err != nil {
		return fmt.Errorf("failed to parse %s flag: %w", flags.Name(), err)
	}

	if rest := flags.Args(); len(rest) != 0 {
		return fmt.Errorf(
			"unexpected arguments %q: %w",
			rest,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
