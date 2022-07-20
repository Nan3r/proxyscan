/*

Copyright (c) 2017 xsec.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package cmd

import (
	"proxy_scanner/models"
	"proxy_scanner/proxy"

	"github.com/urfave/cli"
)

var Scan = cli.Command{
	Name:        "scan",
	Usage:       "start to scan proxy",
	Description: "start to scan proxy",
	Action:      proxy.Scan,
	Flags: []cli.Flag{
		boolFlag("debug, d", "debug mode"),
		intFlag("scan_num, n", 1000, "scan num"),
		intFlag("timeout, t", 5, "timeout"),
		stringFlag("filename, f", "iplist.txt", "filename"),
		stringFlag("target, u", "http://email.163.com/", "target url"),
		stringFlag("targetr, r", "网易免费邮箱", "target response"),
	},
}

var Dump = cli.Command{
	Name:        "dump",
	Usage:       "dump proxies to a text file",
	Description: "dump proxies to a text file",
	Action:      models.Dump,
	Flags: []cli.Flag{
		stringFlag("dumpfile, o", "dump_proxies.txt", "filename"),
	},
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
