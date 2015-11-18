package context

var Port string = "8080"
var Domain string = "localhost"
var IsSecure bool = false

func SetContext(port string, domain string, isSecure bool) {
	Port = port
	Domain = domain
	IsSecure = isSecure
}

func GetSiteRootPath() string {
	prefix := "http://"
	if IsSecure {
		prefix = "https://"
	}

	return prefix + Domain + ":" + Port
}
