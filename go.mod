module netinfo

go 1.18

require (
	gek_net v0.0.0
	geolite2 v0.0.0
)

replace (
	gek_net => ../gek/gek_net
	geolite2 => ../geolite2
)
