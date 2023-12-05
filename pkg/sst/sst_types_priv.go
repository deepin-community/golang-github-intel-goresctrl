// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -I/usr/src/linux -I/usr/src/linux/include -I/usr/src/linux/arch/x86/include/generated/ _sst_types_priv.go

package sst

const (
	CONFIG_TDP                        = 0x7f
	CONFIG_TDP_GET_LEVELS_INFO        = 0x0
	CONFIG_TDP_GET_TDP_CONTROL        = 0x1
	CONFIG_TDP_SET_TDP_CONTROL        = 0x2
	CONFIG_TDP_GET_TDP_INFO           = 0x3
	CONFIG_TDP_GET_PWR_INFO           = 0x4
	CONFIG_TDP_GET_TJMAX_INFO         = 0x5
	CONFIG_TDP_GET_CORE_MASK          = 0x6
	CONFIG_TDP_GET_TURBO_LIMIT_RATIOS = 0x7
	CONFIG_TDP_SET_LEVEL              = 0x8
	CONFIG_TDP_GET_UNCORE_P0_P1_INFO  = 0x9
	CONFIG_TDP_GET_P1_INFO            = 0xa
	CONFIG_TDP_GET_MEM_FREQ           = 0xb

	CONFIG_TDP_GET_FACT_HP_TURBO_LIMIT_NUMCORES = 0x10
	CONFIG_TDP_GET_FACT_HP_TURBO_LIMIT_RATIOS   = 0x11
	CONFIG_TDP_GET_FACT_LP_CLIPPING_RATIO       = 0x12

	CONFIG_TDP_PBF_GET_CORE_MASK_INFO = 0x20
	CONFIG_TDP_PBF_GET_P1HI_P1LO_INFO = 0x21
	CONFIG_TDP_PBF_GET_TJ_MAX_INFO    = 0x22
	CONFIG_TDP_PBF_GET_TDP_INFO       = 0x23

	CONFIG_CLOS        = 0xd0
	CLOS_PM_QOS_CONFIG = 0x2
	CLOS_PQR_ASSOC     = 0x0
	CLOS_PM_CLOS       = 0x1
	CLOS_STATUS        = 0x3

	MBOX_CMD_WRITE_BIT = 0x8

	READ_PM_CONFIG  = 0x94
	WRITE_PM_CONFIG = 0x95
	PM_FEATURE      = 0x3

	PM_QOS_INFO_OFFSET   = 0x0
	PM_QOS_CONFIG_OFFSET = 0x4
	PM_CLOS_OFFSET       = 0x8
	PQR_ASSOC_OFFSET     = 0x20

	MSR_PM_ENABLE = 0x770
)