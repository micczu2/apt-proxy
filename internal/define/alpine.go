package define

import "regexp"

var ALPINE_HOST_PATTERN = regexp.MustCompile(`/alpine/(.+)$`)

const ALPINE_BENCHMAKR_URL = "MIRRORS.txt"

// https://mirrors.alpinelinux.org/ 2022.11.19
// Sites that contain protocol headers, restrict access to resources using that protocol
var ALPINE_OFFICIAL_MIRRORS = []string{
	"alpine.sakamoto.pl/alpine/",
	"ftp.icm.edu.pl/pub/Linux/distributions/alpine/",
	"mirror.fel.cvut.cz/alpine/",
//	"dl-cdn.alpinelinux.org/alpine/", //some cert issues?
	"mirror.leaseweb.com/alpine/",
	"ftp.halifax.rwth-aachen.de/alpine/",
	"ftp.acc.umu.se/mirror/alpinelinux.org/",
//	"eu.edge.kernel.org/alpine/", //some signature issues; "BAD signature"
}

var ALPINE_CUSTOM_MIRRORS = []string{}

var BUILDIN_ALPINE_MIRRORS = GenerateBuildInList(ALPINE_OFFICIAL_MIRRORS, ALPINE_CUSTOM_MIRRORS)

var ALPINE_DEFAULT_CACHE_RULES = []Rule{
	{Pattern: regexp.MustCompile(`APKINDEX.tar.gz$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TYPE_LINUX_DISTROS_ALPINE},
	{Pattern: regexp.MustCompile(`tar.gz$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TYPE_LINUX_DISTROS_ALPINE},
	{Pattern: regexp.MustCompile(`apk$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TYPE_LINUX_DISTROS_ALPINE},
	{Pattern: regexp.MustCompile(`.*`), CacheControl: `max-age=100000`, Rewrite: true, OS: TYPE_LINUX_DISTROS_ALPINE},
}
