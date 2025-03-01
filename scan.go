package viamwifislam

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func DoScan(ctx context.Context) (*int, error) {

	cmd := exec.Command("nmcli", "-t", "-f", "SSID,SIGNAL,BSSID", "dev", "wifi", "list")
	raw, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("cannot run nmcli %v", err)
	}

	for _, l := range strings.Split(string(raw), "\n") {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		x := strings.SplitN(l, ":", 3)
		x[2] = strings.ReplaceAll(x[2], "\\", "")
		fmt.Printf("\t %d - signal: %v %v\n", len(x), x[1], x[2])

	}

	panic(1)
}
