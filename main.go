package main

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/vmware/vra-sdk-go/pkg/client"
	"github.com/vmware/vra-sdk-go/pkg/client/cmdb"
	"github.com/vmware/vra-sdk-go/pkg/models"

	httptransport "github.com/go-openapi/runtime/client"
)

var apiHost = "dev125776.service-now.com"

func main() {

	transport := httptransport.New(apiHost, "/api/now", nil)
	transport.SetDebug(true)
	transport.DefaultAuthentication = httptransport.BasicAuth("admin", "F!hoPlCr7/B6")
	apiclient := client.New(transport, strfmt.Default)
	fmt.Println("apiclient:\n", apiclient)
	result, err := apiclient.Cmdb.CreateIdentifyReconcile(cmdb.NewCreateIdentifyReconcileParams().
		WithBody(&models.IdentifyReconcileItemList{
			Items: []*models.IdentifyReconcileItem{{ClassName: "cmdb_ci_windows_service", Values: &models.ItemValue{
				Name:      "Win Server 100",
				IpAddress: "10.20.30.40",
			}}},
		}))
	fmt.Println("RES:\n", result)
	fmt.Println("ERR:\n", err)

}
