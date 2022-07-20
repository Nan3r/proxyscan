package main

import (
	"os"
	"runtime"

	"github.com/urfave/cli"

	"proxy_scanner/cmd"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "proxy scanner"
	app.Author = "nil"
	app.Email = "nil"
	app.Version = "20220717"
	app.Usage = "A SOCKS4/SOCKS4a/SOCKS5/HTTP/HTTPS proxy scanner"
	app.Commands = []cli.Command{cmd.Scan, cmd.Dump}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)
	app.Flags = append(app.Flags, cmd.Dump.Flags...)
	app.Run(os.Args)
}
