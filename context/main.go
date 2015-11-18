package context

var Domain string = "localhost"
var Port string = "8080"
var IsSecure bool = false

func SetContext(domain string, port string, isSecure bool) {
	Domain = domain
	Port = port
	IsSecure = isSecure
}

func GetSiteRootPath() string {
	prefix := "http://"
	if IsSecure {
		prefix = "https://"
	}

	return prefix + Domain + ":" + Port
}
