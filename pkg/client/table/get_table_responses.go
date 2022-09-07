package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anka-software/cmdb-sdk/pkg/models"
)

// CreateIdentifyReconcileReader is a Reader for the CreateIdentifyReconcile structure.
type GetTableItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTableItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTableItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTableItemsUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTableItemsUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTableItemsUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match defined for this endpoint", response, response.Code())
	}
}

// NewCreateIdentifyReconcileCreated creates a CreateIdentifyReconcileCreated with default headers values
func NewGetTableItemsOK() *GetTableItemsOK {
	return &GetTableItemsOK{}
}

type GetTableItemsOK struct {
	Payload *models.GetTableItem
}

func (o *GetTableItemsOK) Error() string {
	return fmt.Sprintf("ERROR")
}
func (o *GetTableItemsOK) GetPayload() *models.GetTableItem {
	return o.Payload
}

func (o *GetTableItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTableItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
func NewGetTableItemsUsingGET1Unauthorized() *GetTableItemsUsingGET1Unauthorized {
	return &GetTableItemsUsingGET1Unauthorized{}
}

/* GetTableItemsUsingGET1Unauthorized describes a response with status code 401, with default header values.
Unauthorized
*/
type GetTableItemsUsingGET1Unauthorized struct {
}

func (o *GetTableItemsUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /TableItems/api/TableItemss/{TableItemsId}][%d] getTableItemsUsingGET1Unauthorized ", 401)
}

func (o *GetTableItemsUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTableItemsUsingGET1Forbidden creates a GetTableItemsUsingGET1Forbidden with default headers values
func NewGetTableItemsUsingGET1Forbidden() *GetTableItemsUsingGET1Forbidden {
	return &GetTableItemsUsingGET1Forbidden{}
}

/* GetTableItemsUsingGET1Forbidden describes a response with status code 403, with default header values.
Forbidden
*/
type GetTableItemsUsingGET1Forbidden struct {
}

func (o *GetTableItemsUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /TableItems/api/TableItemss/{TableItemsId}][%d] getTableItemsUsingGET1Forbidden ", 403)
}

func (o *GetTableItemsUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTableItemsUsingGET1NotFound creates a GetTableItemsUsingGET1NotFound with default headers values
func NewGetTableItemsUsingGET1NotFound() *GetTableItemsUsingGET1NotFound {
	return &GetTableItemsUsingGET1NotFound{}
}

/* GetTableItemsUsingGET1NotFound describes a response with status code 404, with default header values.
Not Found
*/
type GetTableItemsUsingGET1NotFound struct {
	Payload *models.Error
}

func (o *GetTableItemsUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /TableItems/api/TableItemss/{TableItemsId}][%d] getTableItemsUsingGET1NotFound  %+v", 404, o.Payload)
}
func (o *GetTableItemsUsingGET1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTableItemsUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
