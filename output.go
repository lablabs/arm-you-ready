package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func renderTable(targetArchitecture string, outputs []Output) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"name", "namespace", "resource", "image", targetArchitecture + " supported", "architectures"})

	for _, o := range outputs {
		t.AppendRows([]table.Row{
			{o.Name, o.Namespace, o.Resource, o.Image, architectureIsSupported(targetArchitecture, o.Architectures), o.Architectures },
		})
	}

	t.Render()
}