package buildinfo

import (
	"bytes"
	"log"
	"runtime"
	"text/template"
)

// Following variables are set via -ldflags
// nolint:gochecknoglobals
var (
	// AppVersion git SHA at build time
	AppVersion string
	// BuildTime time of build
	BuildTime string
	// VCSRef name of branch at build time
	VCSRef string
)

var (
	textTemplate = `
version:           {{.AppVersion}}
Go version:        {{.GoVersion}}
Git commit:        {{.VCSRef}}
Built:             {{.BuildTime}}
OS/Arch:           {{.GoOs}}/{{.GoArch}}`
)

func Print() {
	templateBytes := new(bytes.Buffer)
	templateValues := struct {
		AppVersion string
		GoVersion  string
		VCSRef     string
		BuildTime  string
		GoOs       string
		GoArch     string
	}{
		AppVersion: AppVersion,
		GoVersion:  runtime.Version(),
		VCSRef:     VCSRef,
		BuildTime:  BuildTime,
		GoOs:       runtime.GOOS,
		GoArch:     runtime.GOARCH,
	}

	err := template.Must(template.New("buildInfo").Parse(textTemplate)).Execute(templateBytes, templateValues)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(templateBytes.String())
}
