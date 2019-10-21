// +build !darwin,!linux,!freebsd,!openbsd,!windows

package net

import (
	"context"

	"github.com/shirou/gopsutil/internal/common"
)

func IOCounters(pernic bool) ([]IOCountersStat, error) {
	return IOCountersWithContext(context.Background(), pernic)
}

func IOCountersWithContext(ctx context.Context, pernic bool) ([]IOCountersStat, error) {
	return []IOCountersStat{}, common.ErrNotImplementedError
}

func FilterCounters() ([]FilterStat, error) {
	return FilterCountersWithContext(context.Background())
}

func FilterCountersWithContext(ctx context.Context) ([]FilterStat, error) {
	return []FilterStat{}, common.ErrNotImplementedError
}

func ConntrackStats(percpu bool) ([]ConntrackStat, error) {
	return ConntrackStatsWithContext(context.Background(), percpu)
}

func ConntrackStatsWithContext(ctx context.Context, percpu bool) ([]ConntrackStat, error) {
	return nil, common.ErrNotImplementedError
}

func ProtoCounters(protocols []string) ([]ProtoCountersStat, error) {
	return ProtoCountersWithContext(context.Background(), protocols)
}

func ProtoCountersWithContext(ctx context.Context, protocols []string) ([]ProtoCountersStat, error) {
	return []ProtoCountersStat{}, common.ErrNotImplementedError
}

func Connections(kind string, config ...ConnectionStatConfigurer) ([]ConnectionStat, error) {
	return ConnectionsWithContext(context.Background(), kind, config...)
}

func ConnectionsWithContext(ctx context.Context, kind string, config ...ConnectionStatConfigurer) ([]ConnectionStat, error) {
	return []ConnectionStat{}, common.ErrNotImplementedError
}

func ConnectionsMax(kind string, max int, config ...ConnectionStatConfigurer) ([]ConnectionStat, error) {
	return ConnectionsMaxWithContext(context.Background(), kind, max, config...)
}

func ConnectionsMaxWithContext(ctx context.Context, kind string, max int, config ...ConnectionStatConfigurer) ([]ConnectionStat, error) {
	return []ConnectionStat{}, common.ErrNotImplementedError
}
