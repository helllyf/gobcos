module github.com/KasperLiu/gobcos

go 1.12

require (
	github.com/aristanetworks/goarista v0.0.0-20190712234253-ed1100a1c015
	github.com/btcsuite/btcd v0.0.0-20190807005414-4063feeff79a
	github.com/deckarep/golang-set v1.7.1
	github.com/pborman/uuid v1.2.0
	github.com/rjeczalik/notify v0.9.2
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.2.2
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/sys v0.0.0-20190412213103-97732733099d
	gopkg.in/urfave/cli.v1 v1.20.0
)

replace (
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 => golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/sys v0.0.0-20180926160741-c2ed4eda69e7 => golang.org/x/sys v0.0.0-20190524122548-abf6ff778158
)
