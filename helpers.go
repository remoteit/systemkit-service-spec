package spec

// NewEmptySERVICE - return an empty SERVICE, all fields are initialized
func NewEmptySERVICE() SERVICE {
	return SERVICE{
		Name:                    "",
		Description:             "",
		Documentation:           "",
		Executable:              "",
		Args:                    []string{},
		WorkingDirectory:        "",
		Environment:             map[string]string{},
		DependsOn:               []ServiceType{},
		DependsOnOverrideByOS:   map[OsType][]ServiceType{},
		DependsOnOverrideByInit: map[InitType][]ServiceType{},
		Start:                   StartConfig{},
		Logging: LoggingConfig{
			StdOut: LoggingConfigOut{
				Disabled: true,
			},
			StdErr: LoggingConfigOut{
				Disabled: true,
			},
		},
		Credentials: CredentialsConfig{},
	}
}

// KnownOsTypes -
var KnownOsTypes = []OsType{
	OsFreeBSD,
	OsLinux,
	OsMacOS,
	OsWindows,
	OsOpenBSD,
	OsNetBSD,
}

// KnownInitTypes -
var KnownInitTypes = []InitType{
	InitRC_D,
	InitSystemd,
	InitSystemV,
	InitUpstart,
	InitUknown,
}

// KnownServiceTypes -
var KnownServiceTypes = []ServiceType{
	ServiceNetwork,
	ServiceBluetooth,
}

// ServiceMappings -
var ServiceMappings = map[InitType]map[ServiceType]string{

	// based on /etc/rc.d/* and /usr/local/etc/rc.d/*
	InitRC_D: {
		ServiceNetwork:   "NETWORKING",
		ServiceBluetooth: "bluetooth",
	},

	// based on /etc/systemd/system/* and /usr/lib/systemd/system/*
	InitSystemd: {
		ServiceNetwork:   "network.target",
		ServiceBluetooth: "bluetooth.target",
	},
}
