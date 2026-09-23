package vpn

import (
	"fmt"
)

type VPNService struct {
	Users           []User
	Interfaces      []string
	UsageStatistics []string
	IsConnected     bool
}

type Peer struct {
	PublicKey           string
	Endpoint            string
	AllowedIPs          string
	PersistentKeepalive int
}

type UserConfig struct {
	UserId     int
	PrivateKey string
	Address    string
	DNS        string
	Peer       Peer
	QRcode     []byte
}

type User struct {
	UserName   string
	UserId     int
	UserConf   UserConfig
	UserStatus string
}

func (vs *VPNService) Connect() string {
	(*vs).IsConnected = true
	return "Status changed to connected"
}
func (vs *VPNService) Disconnect() string {
	(*vs).IsConnected = false
	return "Status changed to disconnected"
}
func (vs *VPNService) CreateProfile() UserConfig {
	fmt.Println("wow new request, let's make one")
	return UserConfig{
		UserId:     777,
		PrivateKey: "hey, that's personal",
		Address:    "would't you wanna know? maybe I should give you appartment keys also?",
		DNS:        "magical thing this DNS",
		Peer: Peer{
			PublicKey:           "hey look at my public key I'm a peer",
			Endpoint:            "hehe.com the best web page online!",
			AllowedIPs:          "everything is allowed! all the ips welcome",
			PersistentKeepalive: 25,
		},
		QRcode: []byte("HE-HE-HE!"),
	}
}
func (vs *VPNService) GetGeneralStats() string {
	return "Everything ok, and you?"
}
func (vs *VPNService) GetUserStats(userName string) string {
	return fmt.Sprintf("You know what, go ask %s yourself ok? I have to go and see ChatGPT we have a party tonight", userName)
}
