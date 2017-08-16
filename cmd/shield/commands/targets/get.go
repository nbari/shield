package targets

import (
	"strings"

	"github.com/starkandwayne/shield/cmd/shield/log"
	"github.com/starkandwayne/shield/cmd/shield/commands"
	"github.com/starkandwayne/shield/cmd/shield/commands/internal"
)

func init() {
	target := commands.Register("target", cliGetTarget)
	target.Summarize("Print detailed information about a specific backup target")
	target.Aliases("show target", "view target", "display target", "list target", "ls target")
	target.Help(commands.HelpInfo{
		Flags: []commands.FlagInfo{
			commands.TargetNameFlag,
			commands.FlagInfo{
				Name: "uuid", Desc: "Return UUID of target",
			},
		},
		JSONOutput: `{
			"uuid":"8add3e57-95cd-4ec0-9144-4cd5c50cd392",
			"name":"SampleTarget",
			"summary":"A Sample Target",
			"plugin":"postgres",
			"endpoint":"{\"endpoint\":\"127.0.0.1:5432\"}",
			"agent":"127.0.0.1:1234"
		}`,
	})
}

//Print detailed information about a specific backup target
func cliGetTarget(opts *commands.Options, args ...string) error {
	log.DEBUG("running 'show target' command")

	target, _, err := internal.FindTarget(strings.Join(args, " "), *commands.Opts.Raw)
	if err != nil {
		return err
	}

	if *opts.Raw {
		internal.RawJSON(target)
		return nil
	}

	if *opts.ShowUUID {
		internal.RawUUID(target.UUID)
		return nil
	}
	internal.ShowTarget(target)
	return nil
}
