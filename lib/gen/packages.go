package gen

import "strings"

const (
	typesPkg   = "github.com/weimil/krpc-go/types"
	krpcPkg    = "github.com/weimil/krpc-go"
	servicePkg = "github.com/weimil/krpc-go/lib/service"
	encodePkg  = "github.com/weimil/krpc-go/lib/encode"
	tracerrPkg = "github.com/ztrue/tracerr"
)

func getServicePackage(serviceName string) string {
	return "github.com/weimil/krpc-go/" + strings.ToLower(serviceName)
}
