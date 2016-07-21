package qemu

import (
	"github.com/emc-advanced-dev/unik/pkg/compilers"
	"github.com/emc-advanced-dev/unik/pkg/providers"
)

func (p *QemuProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
		SupportedCompilers: []string{
			compilers.RUMP_GO_QEMU,
			compilers.RUMP_NODEJS_QEMU,
			compilers.RUMP_PYTHON_QEMU,
		},
	}
}
