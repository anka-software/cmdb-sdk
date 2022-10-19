package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anka-software/cmdb-sdk/pkg/client"
	"github.com/anka-software/cmdb-sdk/pkg/client/cmdb_meta"
	"github.com/joho/godotenv"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
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

	// CREATE&UPDATE IDENTIFY RECONCILER
	/*apiclient := client.New(transport, strfmt.Default)
	sysparam := "ServiceNow"

	result, err := apiclient.Cmdb.CreateIdentifyReconcile(cmdb.NewCreateIdentifyReconcileParams().WithSysParamDataSource(&sysparam).
		WithBody(&models.IdentifyReconcileItemList{
			Items: []*models.IdentifyReconcileItem{{ClassName: "cmdb_ci_win_server", Values: &models.ItemValue{
				Name: "testUpdate",
				Ram:  "1",
			}}},
		}))
	fmt.Println("RES:\n", result)
	fmt.Println("ERR:\n", err)*/

	// GET TABLE ITEMS
	/*
		apiclient := client.New(transport, strfmt.Default)
		tableName := "cmdb_ci_win_server"
		query := "name=testUpdate"
		result1, err1 := apiclient.Table.GetTableItems(table.NewGetTableItemParams().WithTableName(tableName).WithQuery(&query))
		fmt.Println("RES:\n", result1.Payload.Result)
		fmt.Println("ERR:\n", err1)*/

	//DELETE A RECORD FROM TABLE

	apiclient := client.New(transport, strfmt.Default)
	tableName := "cmdb_ci_win_server"
	//SysId := "bb53f3f947f511504b6fcc88f36d4374"
	/*result, err := apiclient.CmdbMeta.GetCmdbMetaByClassName(cmdb_meta.NewGetTableItemParams().WithClassName(tableName))
	if err != nil {
		fmt.Println("ERR:\n", err)
	}

	// Fields of the specific class Name
	for _, v := range result.Payload.Result.Attributes {
		fmt.Println(v.Element)
	}*/
	mandatoryFields := RetrieveMandatoryFields(apiclient, tableName)
	fmt.Println("MandatoryFields:", mandatoryFields)
	//result1, err1 := apiclient.Table.DeleteRecord(table.NewDeleteRecordParams().WithTableName(tableName).WithSysID(SysId))
	//fmt.Println("RES:\n", result1.Payload)
	//fmt.Println("ERR:\n", err1)
}

func RetrieveMandatoryFields(apiclient *client.MainClient, tableName string) []string {

	result, err := apiclient.CmdbMeta.GetCmdbMetaByClassName(cmdb_meta.NewGetTableItemParams().WithClassName(tableName))
	if err != nil {
		fmt.Println("ERR:\n", err)
	}
	//fmt.Println("RES:\n", result.Payload.Result.Attributes)
	var mandatoryFields []string
	// Fields of the specific class Name
	for _, v := range result.Payload.Result.Attributes {
		if v.IsMandatory == "false" && v.IsDisplay == "true" {
			mandatoryFields = append(mandatoryFields, v.Element)
		}
		//fmt.Println(v.Element)
	}
	return mandatoryFields
}
