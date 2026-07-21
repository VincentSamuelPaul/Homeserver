package main

import (
	"fmt"
	"sync"

	"github.com/VincentSamuelPaul/greet/ui"
)

func main() {
	var (
		sysInfo    SystemInfo
		containers []ContainerInfo
		netInfo    NetworkInfo
		stoInfo    StorageInfo
		secInfo    SecurityInfo
	)

	var wg sync.WaitGroup
	wg.Add(5)

	go func() { defer wg.Done(); sysInfo = getSystemInfo() }()
	go func() { defer wg.Done(); containers = getContainers() }()
	go func() { defer wg.Done(); netInfo = getNetworkInfo() }()
	go func() { defer wg.Done(); stoInfo = getStorageInfo() }()
	go func() { defer wg.Done(); secInfo = getSecurityInfo() }()

	wg.Wait()

	uiContainers := make([]ui.ContainerData, len(containers))
	for i, c := range containers {
		uiContainers[i] = ui.ContainerData{Name: c.Name, State: c.State, Uptime: c.Uptime}
	}

	uiEntries := make([]ui.StorageEntry, len(stoInfo.Entries))
	for i, e := range stoInfo.Entries {
		uiEntries[i] = ui.StorageEntry{Name: e.Name, Size: e.Size}
	}

	fmt.Print(ui.Render(
		ui.SysData{
			Host:      sysInfo.Host,
			Uptime:    sysInfo.Uptime,
			Kernel:    sysInfo.Kernel,
			CPUPct:    sysInfo.CPUPct,
			MemUsed:   sysInfo.MemUsed,
			MemTotal:  sysInfo.MemTotal,
			RootUsed:  sysInfo.RootUsed,
			RootTotal: sysInfo.RootTotal,
			HomeUsed:  sysInfo.HomeUsed,
			HomeTotal: sysInfo.HomeTotal,
		},
		uiContainers,
		ui.NetData{
			Interface: netInfo.Interface,
			Address:   netInfo.Address,
			Tunnel:    netInfo.Tunnel,
			Internet:  netInfo.Internet,
		},
		ui.StoData{
			Entries:   uiEntries,
			TotalUsed: stoInfo.TotalUsed,
			TotalSize: stoInfo.TotalSize,
		},
		ui.SecData{
			FailedCount:  secInfo.FailedCount,
			LastAttacker: secInfo.LastAttacker,
			LastLogin:    secInfo.LastLogin,
		},
	))
}