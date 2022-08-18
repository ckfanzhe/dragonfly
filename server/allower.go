package server

import (
	"log"
	"net"

	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

// Allower may be implemented to specifically allow or disallow players from joining a Server, by setting the specific
// Allower implementation through a call to Server.Allow.
type Allower interface {
	// Allow filters what connections are allowed to connect to the Server. The address, identity data, and client data
	// of the connection are passed. If Admit returns false, the connection is closed with the string returned
	// as the disconnect message. WARNING: Use the client data at your own risk, it cannot be trusted because
	// it can be freely changed by the player connecting.
	Allow(addr net.Addr, d login.IdentityData, c login.ClientData) (string, bool)
}

// allower is the standard Allower implementation. It accepts all connections.
type allower struct{}

// Allow always returns true.
func (allower) Allow(a net.Addr, d login.IdentityData, c login.ClientData) (string, bool) {
	log.Printf("User:%s form:%s Xuid:%s Identity:%s", d.DisplayName, a.String(), d.XUID, d.Identity)
	info := "没有权限登录，请联系管理员！"
	permit := false
	if checkallow(d.XUID) {
		info = "登录成功！"
		permit = true
	}
	return info, permit
}

func checkallow(xuid string) bool {
	xuid_list := []string{
		"2535441176471440", // xuekuilei2022
		"2535468093794442", // KuChaZi5467
		"2535423732992577", // mmyy13
		"2535419280628438", // liubai0003
		"2535466795063605", // Yonezu S
		"2535417232763992", // BladedAlarm7206
		"2535415141730477", // RivalStone74010
		"2535414938543844", // ifanzhe
	}
	for _, user := range xuid_list {
		if user == xuid {
			return true
		}
	}
	return false
}
