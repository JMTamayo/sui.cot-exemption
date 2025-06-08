package server

import (
	"fmt"
	"os/exec"

	"sui.cot-exemption/app/config"
	"sui.cot-exemption/app/models"
)

// VPNClient handles VPN operations
type VPNClient struct{}

// NewVPNClient creates a new VPN client instance
//
// Arguments:
//   - None.
//
// Returns:
//   - A new VPN client instance.
func NewVPNClient() *VPNClient {
	return &VPNClient{}
}

// Login authenticates with VPN using the provided token
//
// Arguments:
//   - None.
//
// Returns:
//   - An error if the login fails.
func (c *VPNClient) Login() *models.Error {
	cmd := exec.Command("nordvpn", "login", "--token", config.Conf.GetVpnAccessToken())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.Error{
			Details: fmt.Sprintf("failed to login to NordVPN: %v, output: %s", err, string(output)),
		}
	}
	return nil
}

// Connect connects to a specific VPN server
//
// Arguments:
//   - None.
//
// Returns:
//   - An error if the connection fails.
func (c *VPNClient) Connect() *models.Error {
	cmd := exec.Command("nordvpn", "connect", config.Conf.GetVpnServer())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &models.Error{
			Details: fmt.Sprintf("failed to connect to VPN: %v, output: %s", err, string(output)),
		}
	}
	return nil
}
