//go:build linux || darwin

package main

import (
	"runtime"

	"github.com/shirou/gopsutil/v3/process"
)

func getProcessStats() ([]ProcessStat, error) {
	numCores := float64(runtime.NumCPU())
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var stats []ProcessStat
	for _, p := range processes {
		name, _ := p.Name()
		cpu, err := p.CPUPercent()

		if err == nil && cpu > 0 {
			stats = append(stats, ProcessStat{
				Name: name,
				CPU:  cpu / numCores,
			})
		}
	}
	return stats, nil
}
