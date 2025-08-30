package network

import (
	"fmt"
	"net"
	"os/exec"
	"time"

	"github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/api"
)

// networkOperator implements the api.NetworkOperator interface.
type networkOperator struct {
	// Potentially a logger or other dependencies
}

// NewNetworkOperator creates a new network operator.
func NewNetworkOperator() api.NetworkOperator {
	return &networkOperator{}
}

// CheckConnectivity attempts to establish a TCP connection to a given host and port.
func (o *networkOperator) CheckConnectivity(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)
	timeout := 5 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return errors.Wrapf(err, errors.NetworkError, "connectivity check failed for %s", address)
	}
	defer conn.Close()

	return nil
}

// ManageVIP adds or removes the virtual IP from a network interface.
// This is a simplified placeholder. A real implementation would be more robust.
// It might need to know the interface name.
func (o *networkOperator) ManageVIP(action string, vip string) error {
	var cmd *exec.Cmd

	// This is highly OS-specific. This example is for Linux.
	// A better implementation would use a library like `vishvananda/netlink`.
	switch action {
	case "add":
		// A common way to do this is to add the IP to the loopback interface,
		// and use something like `arping` to announce it.
		// e.g., ip addr add ${VIP}/32 dev lo
		// e.g., arping -c 1 -A -I eth0 ${VIP}
		cmd = exec.Command("ip", "addr", "add", vip+"/32", "dev", "lo")
	case "del":
		cmd = exec.Command("ip", "addr", "del", vip+"/32", "dev", "lo")
	default:
		return errors.Newf(errors.ValidationError, "invalid action for ManageVIP: %s", action)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, errors.NetworkError, "failed to manage VIP: %s", string(output))
	}

	return nil
}

//Personal.AI order the ending
