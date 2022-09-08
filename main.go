package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anka-software/cmdb-sdk/pkg/client"
	"github.com/anka-software/cmdb-sdk/pkg/client/cmdb"
	"github.com/anka-software/cmdb-sdk/pkg/client/table"
	"github.com/anka-software/cmdb-sdk/pkg/models"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"

	httptransport "github.com/go-openapi/runtime/client"
)

var apiHost = "dev125776.service-now.com"

func main() {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatalf("Some error occured. Err: %s", err1)
	}

	transport := httptransport.New(apiHost, "/api/now", nil)
	transport.SetDebug(true)
	transport.DefaultAuthentication = httptransport.BasicAuth(os.Getenv("SN_USERNAME"), os.Getenv("SN_PASSWORD"))
	apiclient := client.New(transport, strfmt.Default)
	/*apiclient := client.New(transport, strfmt.Default)
	 */

	sysparam := "ServiceNow"

	result, err := apiclient.Cmdb.CreateIdentifyReconcile(cmdb.NewCreateIdentifyReconcileParams().WithSysParamDataSource(&sysparam).
		WithBody(&models.IdentifyReconcileItemList{
			Items: []*models.IdentifyReconcileItem{{ClassName: "cmdb_ci_win_server", Values: &models.ItemValue{
				Name: "testUpdate",
				Ram:  "1",
			}}},
		}))
	fmt.Println("RES:\n", result)
	fmt.Println("ERR:\n", err)
	tableName := "cmdb_ci_win_server"
	query := "name=testUpdate"
	result1, err1 := apiclient.Table.GetTableItems(table.NewGetTableItemParams().WithTableName(tableName).WithQuery(&query))
	fmt.Println("RES:\n", result1.Payload.Result)
	fmt.Println("ERR:\n", err1)

	//
}
