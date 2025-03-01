package viamwifislam

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type APPower struct {
	BSSID  string
	Signal int
}

func DoScan(ctx context.Context, ssid string) ([]APPower, error) {

	cmd := exec.Command("nmcli", "-t", "-f", "SSID,SIGNAL,BSSID", "dev", "wifi", "list")
	raw, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("cannot run nmcli %v", err)
	}

	res := []APPower{}

	for _, l := range strings.Split(string(raw), "\n") {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		x := strings.SplitN(l, ":", 3)
		x[2] = strings.ReplaceAll(x[2], "\\", "")

		if len(ssid) > 0 && x[0] != ssid {
			continue
		}

		p := APPower{BSSID: x[2]}
		p.Signal, err = strconv.Atoi(x[1])
		if err != nil {
			return nil, fmt.Errorf("bad signal [%s]", x[1])
		}

		res = append(res, p)
	}

	return res, nil
}
