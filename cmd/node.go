package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/yogyrahmawan/galeras/app"
)

// NewNodeCommand create new node command
func NewNodeCommand() *cobra.Command {
	nodeCmd := &cobra.Command{
		Use:   "node <subcommands>",
		Short: "node <subcommands>",
	}

	nodeCmd.AddCommand(NewRunNodeCommand())
	nodeCmd.AddCommand(NewRemoveNodeCommand())

	return nodeCmd
}

// NewRunNodeCommand create run node command
func NewRunNodeCommand() *cobra.Command {
	nr := &cobra.Command{
		Use: "run --name <name> --host <host> --env-file <env-file> --net <net> --ip <ip> --add-host <host1,host2> --port <port1,port2,port3> --image <image> --additional-command <command1,command2>",
		Run: runNode,
	}

	nr.Flags().StringP("name", "n", "name", "container name")
	nr.Flags().StringP("host", "s", "host", "host name")
	nr.Flags().StringP("env-file", "e", "env file", "path of environment file")
	nr.Flags().StringP("net", "t", "net", "docker network")
	nr.Flags().StringP("ip", "i", "ip", "this node ip")
	nr.Flags().StringP("add-host", "a", "add-host", "host of other nodes, separate by coma")
	nr.Flags().StringP("port", "p", "port", "port separated by coma")
	nr.Flags().StringP("image", "g", "image", "docker image")
	nr.Flags().StringP("additional-command", "d", "additional-command", "additional docker command")
	return nr

}

func runNode(cmd *cobra.Command, args []string) {
	var name, host, envFile, net, ip, addHost, port, image, additionalCommand string
	var err error

	name, err = cmd.Flags().GetString("name")
	if err != nil {
		ExitWithError(128, err)
	}

	host, err = cmd.Flags().GetString("host")
	if err != nil {
		ExitWithError(128, err)
	}

	envFile, err = cmd.Flags().GetString("env-file")
	if err != nil {
		ExitWithError(128, err)
	}

	net, err = cmd.Flags().GetString("net")
	if err != nil {
		ExitWithError(128, err)
	}

	addHost, err = cmd.Flags().GetString("add-host")
	if err != nil {
		ExitWithError(128, err)
	}

	port, err = cmd.Flags().GetString("port")
	if err != nil {
		ExitWithError(128, err)
	}

	image, err = cmd.Flags().GetString("image")
	if err != nil {
		ExitWithError(128, err)
	}

	additionalCommand, err = cmd.Flags().GetString("additional-command")
	if err != nil {
		ExitWithError(128, err)
	}
	ip, err = cmd.Flags().GetString("ip")
	if err != nil {
		ExitWithError(128, err)
	}

	//run
	app.RunNode(name, host, envFile, net, ip, buildRunDockerCommand(addHost, port, image, additionalCommand))
}

func buildRunDockerCommand(addHost, port, image, additionalCommand string) []string {
	var result []string

	// build add-host
	addHostArr := strings.Split(addHost, ",")
	for i := 0; i < len(addHostArr); i++ {
		result = append(result, "--add-host", addHostArr[i])
	}

	// build port
	portArr := strings.Split(port, ",")
	for i := 0; i < len(portArr); i++ {
		result = append(result, "-p", portArr[i])
	}

	result = append(result, image)

	addCmds := strings.Split(additionalCommand, ",")
	for i := 0; i < len(addCmds); i++ {
		result = append(result, addCmds[i])
	}

	return result
}

// NewRemoveNodeCommand remove node command
func NewRemoveNodeCommand() *cobra.Command {
	nr := &cobra.Command{
		Use: "rm --name <name> --name <name>",
		Run: removeNode,
	}

	nr.Flags().StringArrayP("name", "n", []string{"name"}, "container name")
	return nr
}

func removeNode(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetStringArray("name")
	if err != nil {
		ExitWithError(128, err)
	}

	app.RemoveNode(name)
}
