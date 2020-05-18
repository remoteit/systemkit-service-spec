# ![](https://fonts.gstatic.com/s/i/materialiconsoutlined/flare/v4/24px.svg) Service
[![](https://img.shields.io/github/v/release/codemodify/systemkit-service-spec?style=flat-square)](https://github.com/codemodify/systemkit-service-spec/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-service-spec?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-service-spec?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-service-spec/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-service-spec?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-service-spec?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-service-spec)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-service-spec)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-service-spec?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-service-spec?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-service-spec?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-service-spec?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-service-spec?style=flat-square)


# ![](https://fonts.gstatic.com/s/i/materialicons/extension/v5/24px.svg) What
>_After reading a ton of material on `init` frameworks and service management for different OSes;
Going through mental debate of smartness of one versus another and their historical backlog and tech debt,
as well as developer and user friendliness of the material provided which in itself can drive away companies and
engineering force to provide support for one or another platform - __THE TAKE AWAY IS__ - In the end what
matters is to have an OS efficient and developer friendly mechanics to manipulate a service._

>__This repo proposes to create a portable `SERVICE` definition that can be used by developers and businesses
to implement their products on a range of platforms without worrying about misunderstanding the genius on each of those.__


# ![](https://fonts.gstatic.com/s/i/materialicons/code/v5/24px.svg) Reference Implementations
&nbsp;							| &nbsp;
---:							| ---
Go implementation				| https://github.com/codemodify/systemkit-service
Command line tool 				| https://github.com/codemodify/systemkit-service-cli
Samples &amp; Tests				| https://github.com/codemodify/systemkit-service-tests
Init Frameworks					| https://github.com/codemodify/systemkit-service-encoders-rc_d
&nbsp;							| https://github.com/codemodify/systemkit-service-encoders-launchd
&nbsp;							| https://github.com/codemodify/systemkit-service-encoders-systemd
&nbsp;							| https://github.com/codemodify/systemkit-service-encoders-systemv
&nbsp;							| https://github.com/codemodify/systemkit-service-encoders-upstart


# ![](https://fonts.gstatic.com/s/i/materialicons/all_inclusive/v5/24px.svg) Where
>- Use Case (A) - __code__
>	- `NewServiceFromSERVICE()`
>		- creates and controls a service based on `SERVICE` portable description
>	- `NewServiceFromName()`
>		- finds and controls an existing service based on its name; __previously created by our app or something else__
>	- `NewServiceFromPlatformTemplate()`
>		- create a service from a platform dependent template
>		- template is already a service definition as it is defined by the one of `init` frameworks

>- Use Case (B) - __scripting__
>	- manipulate a service based on a SERVICE file
>		- `systemkit-service-spec-cli create/delete/start/stop/info -service SERVICE-FILE`
>	- manipulate a service based on commands and flags
>		- `systemkit-service-spec-cli create/delete/start/stop/info -name test-service -executable vim`
>	- compile your own service manipulator binary based on the `systemkit-service-spec-cli` source code

>- Use Case (C) - __convert__
>	- used as an ETL mechanism, used by scripting or cloud based services
>	- convert `SERVICE` file to a platform dependent format
>	- convert platform dependent format to `SERVICE` format
>	- convert platform dependent format to a different platform dependent format, by going through `SERVICE` format


# ![](https://fonts.gstatic.com/s/i/materialicons/layers/v5/24px.svg) SERVICE format
```json
{
	"name": "test-service",
	"description": "Test Service",
	"documentation": "http://test-service.com",
	"executable": "/bin/sleep",
	"args": [
		"40"
	],
	"workingDirectory": "/tmp",
	"environment": {
		"API_URL": "https://api.test-service.com"
	},
	"dependsOn": [
		"network"
	],
	"dependsOnOverrides": {
		"init_rc.d": {
			"remove": [
				"*"
			],
			"add": [
				"bumblebee",
				"mumbled"
			]
		},
		"init_systemd": {
			"add": [
				"printer"
			]
		},
		"os_freebsd": {
			"add": [
				"linux"
			]
		},
		"os_linux": {
			"add": [
				"bluetooth"
			]
		}
	},
	"start": {
		"atBoot": true,
		"restart": true,
		"restartTimeout": 10
	},
	"logging": {
		"stdout": {
			"disabled": false,
			"useDefault": false,
			"value": "/var/log/test-service-stdout.log"
		},
		"stderr": {
			"disabled": false,
			"useDefault": false,
			"value": "/var/log/test-service-stderr.log"
		}
	},
	"credentials": {
		"user": "user",
		"group": "group"
	}
}
```


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) References & Credits
**Any missing credits are the author's unconscious error**
- https://en.wikipedia.org/wiki/Operating_system_service_management
- https://nosystemd.org
- https://ungleich.ch/en-us/cms/blog/2019/05/20/linux-distros-without-systemd
- https://lwn.net/Articles/578209/
- https://lwn.net/Articles/578210/
- https://en.wikipedia.org/wiki/Init
- https://www.freebsd.org/doc/en_US.ISO8859-1/articles/linux-users/startup.html
- https://sosheskaz.github.io/tutorial/2017/03/28/FreeBSD-rcd-Setup.html
- https://www.freebsd.org/doc/en_US.ISO8859-1/articles/rc-scripting/index.html
- https://www.freedesktop.org/software/systemd/man/systemd.unit.html
- https://www.freedesktop.org/software/systemd/man/systemd.service.html
- https://www.freedesktop.org/software/systemd/man/systemd.directives.html
- https://www.manpagez.com/man/5/launchd.plist
