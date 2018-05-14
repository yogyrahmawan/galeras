package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/yogyrahmawan/galeras/app"
)

// NewRunTestCommand run test 3 nodes command
func NewRunTestCommand() *cobra.Command {
	runTestCmd := &cobra.Command{
		Use: "runtest",
		Run: runTest3Nodes,
	}

	return runTestCmd
}

func runTest3Nodes(cmd *cobra.Command, args []string) {
	// pull image
	fmt.Println("pulling image, please wait ....")
	app.DockerPull("yogyrahmawan/galera-mariadb:10.1")

	// remove all nodes
	fmt.Println("Removing node if exist")
	app.RemoveNode([]string{"galera-node-1", "galera-node-2", "galera-node-3"})

	// run node one by one
	// node 1
	fmt.Println("starting node 1,  Please wait for a moment")
	app.RunNode("galera-node-1", "172.25.0.2", "var/env_1.env", "galeranet", "172.25.0.2", buildRunDockerCommand("galera-node-2:172.25.0.3,galera-node-3:172.25.0.4", "3306,4444,4567,4568", "yogyrahmawan/galera-mariadb:10.1", "mysqld,init"))
	fmt.Println("==========================")
	time.Sleep(20 * time.Second)
	app.MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_cluster_size';")
	app.MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_incoming_addresses';")
	fmt.Println("==========================")

	// node 2
	fmt.Println("starting node 2, Please wait for a moment")
	app.RunNode("galera-node-2", "172.25.0.3", "var/env_2.env", "galeranet", "172.25.0.3", buildRunDockerCommand("galera-node-1:172.25.0.2,galera-node-3:172.25.0.4", "3306,4444,4567,4568", "yogyrahmawan/galera-mariadb:10.1", "mysqld,join"))
	fmt.Println("==========================")
	time.Sleep(20 * time.Second)
	app.MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_cluster_size';")
	app.MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_incoming_addresses';")
	fmt.Println("==========================")

	// node 3
	fmt.Println("starting node 3, Please wait for a moment")
	app.RunNode("galera-node-3", "172.25.0.4", "var/env_3.env", "galeranet", "172.25.0.4", buildRunDockerCommand("galera-node-1:172.25.0.2,galera-node-2:172.25.0.3", "3306,4444,4567,4568", "yogyrahmawan/galera-mariadb:10.1", "mysqld,join"))
	fmt.Println("==========================")
	time.Sleep(20 * time.Second)
	app.MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_cluster_size';")
	app.MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_incoming_addresses';")
	fmt.Println("==========================")
}
