package app

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// NewImage create docker image
func NewImage(imageName, dockerDirPath string) error {
	dockerArgs := append([]string{"build",
		"-t",
		imageName,
		dockerDirPath})
	cmd := exec.Command("docker", dockerArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("docker build out | %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return err
	}

	return nil
}

// InspectNetwork inspect network
func InspectNetwork(name, subnet string) error {
	dockerArgs := append([]string{
		"network",
		"inspect",
		name,
	})

	out, err := exec.Command("docker", dockerArgs...).Output()
	if err != nil {
		if strings.Contains(string(out), "[]") {
			// create new network
			createNetworkArgs := append([]string{
				"network",
				"create",
				subnet,
				name,
			})
			_, err = exec.Command("docker", createNetworkArgs...).Output()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error create network", err)
				return err
			}

			return nil
		}
		fmt.Fprintln(os.Stderr, "Error run command", err)
		return err
	}

	return nil
}

// RemoveNetwork remove network by name
func RemoveNetwork(name string) error {
	dockerArgs := []string{
		"network",
		"rm",
		name,
	}
	out, err := exec.Command("docker", dockerArgs...).Output()
	fmt.Println(string(out))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error rm network", err)
		return err
	}

	return nil
}

// RunNode run galera node
func RunNode(name, host, envFile, net, ip string, args []string) error {
	var out bytes.Buffer
	var stderr bytes.Buffer

	InspectNetwork("galeranet", "--subnet=172.25.0.0/16")
	dockerArgs := append([]string{
		"run", "-d", "--name", name,
		"-h", host,
		"--env-file", envFile,
		"--net", net,
		"--ip", ip},
		args...)
	cmd := exec.Command("docker", dockerArgs...)

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("error run node")
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		exec.Command("docker", "rm", "-f", name).Run()
		return err
	}
	fmt.Println(out.String())

	fmt.Println("running node " + name)
	return nil
}

// RemoveNode stop and remove container of galera node
func RemoveNode(name []string) error {
	var out bytes.Buffer
	var stderr bytes.Buffer

	dockerArgs := append([]string{
		"rm",
		"-vf",
	},
		name...)
	cmd := exec.Command("docker", dockerArgs...)

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	fmt.Println(out.String())

	return nil
}

// MonitorNode monitor cluster node
func MonitorNode(username, password, node, query string) (*MonitorResp, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	dockerArgs := []string{
		"exec",
		node,
		"mysql",
		"-u" + username,
		"-p" + password,
		"-e",
		query,
	}

	cmd := exec.Command("docker", dockerArgs...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return nil, err
	}

	strValue := strings.TrimSpace(out.String())
	mRes, err := parseResponse(strValue)
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return nil, err
	}

	fmt.Println(mRes.String())

	return mRes, nil
}

func parseResponse(str string) (*MonitorResp, error) {
	vals := strings.Split(str, "\n")
	var name, value string

	if len(vals) > 1 {
		lastVals := vals[1]
		if strings.Contains(lastVals, "wsrep_cluster_size") {
			name = "wsrep_cluster_size"
			lastVals = strings.Replace(lastVals, name, "", -1)
			value = strings.TrimSpace(lastVals)
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			NumberOfStartedNode = num
		} else if strings.Contains(lastVals, "wsrep_incoming_addresses") {
			name = "wsrep_incoming_addresses"
			lastVals = strings.Replace(lastVals, name, "", -1)
			value = strings.TrimSpace(lastVals)

		}
	}

	return &MonitorResp{
		Name:  name,
		Value: value,
	}, nil
}

// DockerPull pull image
func DockerPull(imageName string) error {
	dockerArgs := append([]string{"pull",
		imageName,
	})
	cmd := exec.Command("docker", dockerArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("docker pull out | %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return err
	}

	return nil
}
