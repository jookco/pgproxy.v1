/*
   Copyright 2018 Jook.co

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package cli

import (
	"github.com/spf13/pflag"
)

var host string
var port string

var format string

type flagInfoString struct {
	Name        string
	Shorthand   string
	Description string
	Default     string
}

type flagInfoBool struct {
	Name        string
	Shorthand   string
	Description string
	Default     bool
}

var (
	FlagApiHost = flagInfoString{
		Name:        "host",
		Description: "proxy api server address",
		Default:     "127.0.0.1",
	}

	FlagApiPort = flagInfoString{
		Name:        "port",
		Description: "proxy api server port",
		Default:     "8000",
	}

	FlagConfigPath = flagInfoString{
		Name:        "config",
		Shorthand:   "c",
		Description: "path to configuration file",
	}

	FlagLogLevel = flagInfoString{
		Name:        "log-level",
		Shorthand:   "l",
		Description: "logging level",
		Default:     "info",
	}

	FlagBackground = flagInfoBool{
		Name:        "background",
		Shorthand:   "d",
		Description: "run process in background",
		Default:     false,
	}
)

func stringFlag(f *pflag.FlagSet, valPtr *string, flagInfo flagInfoString) {
	f.StringVarP(valPtr,
		flagInfo.Name,
		flagInfo.Shorthand,
		flagInfo.Default,
		flagInfo.Description)
}

func boolFlag(f *pflag.FlagSet, valPtr *bool, flagInfo flagInfoBool) {
	f.BoolVarP(valPtr,
		flagInfo.Name,
		flagInfo.Shorthand,
		flagInfo.Default,
		flagInfo.Description)
}
