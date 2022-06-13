// main.go

import (
  "github.com/hashicorp/packer-plugin-sdk/plugin"
)

// Assume this implements the packer.Builder interface
type ExampleBuilder struct{}

// Assume this implements the packer.PostProcessor interface
type FooPostProcessor struct{}

// Assume this implements the packer.Provisioner interface
type BarProvisioner struct{}

func main() {
    pps := plugin.NewSet()
    pps.RegisterBuilder("example", new(ExampleBuilder))
    pps.RegisterBuilder(plugin.DEFAULT_NAME, new(AnotherBuilder))
    pps.RegisterPostProcessor("foo", new(FooPostProcessor))
    pps.RegisterProvisioner("bar", new(BarProvisioner))
    err := pps.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
