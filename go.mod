module jw.sys

go 1.17

require jw.lib v0.0.0-00010101000000-000000000000

require (
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace jw.lib => ../jw_lib
