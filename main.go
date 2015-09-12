package main

import (
	"flag"
	"log"
)

var commands = []*Command{
	cmdInit,
	cmdAdd,
}

func main() {
	debug := flag.Bool("debug", false, "debug flag")
	flag.Usage = usage
	flag.Parse()

	if *debug {
		log.SetFlags(log.Lshortfile)
	} else {
		log.SetFlags(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			return
		}
	}
	log.Printf("dotf: unknown subcommand %q\nRun goimp help' for usage.\n", args[0])

}
