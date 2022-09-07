package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anka-software/cmdb-sdk/pkg/client"
	"github.com/anka-software/cmdb-sdk/pkg/client/cmdb"
	"github.com/anka-software/cmdb-sdk/pkg/models"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"

	httptransport "github.com/go-openapi/runtime/client"
)

var apiHost = "dev125776.service-now.com"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	transport := httptransport.New(apiHost, "/api/now", nil)
	transport.SetDebug(true)
	transport.DefaultAuthentication = httptransport.BasicAuth(os.Getenv("SN_USERNAME"), os.Getenv("SN_PASSWORD"))
	apiclient := client.New(transport, strfmt.Default)
	fmt.Println("apiclient:\n", apiclient)
	sysparam := "ServiceNow"

	result, err := apiclient.Cmdb.CreateIdentifyReconcile(cmdb.NewCreateIdentifyReconcileParams().WithSysParamDataSource(&sysparam).
		WithBody(&models.IdentifyReconcileItemList{
			Items: []*models.IdentifyReconcileItem{{ClassName: "cmdb_ci_win_server", Values: &models.ItemValue{
				Name: "Win Server 100",
				/*IpAddress:    "10.20.30.40",
				SysClassName: "cmdb_ci_win_server",*/
			}}},
		}))
	fmt.Println("RES:\n", result)
	fmt.Println("ERR:\n", err)
	//?sysparm_data_source=ServiceNow

}
