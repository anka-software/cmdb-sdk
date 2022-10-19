package cmdb_meta

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
		result := NewGetTableItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTableItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTableItemsNotFound()
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
	Payload *models.GetCMDClassSchema
}

func (o *GetTableItemsOK) Error() string {
	return fmt.Sprintf("[GET /TableItems/api/TableItemss/{TableItemsId}][%d] getTableItemsUnauthorized %+v", 200, o.Payload)

}
func (o *GetTableItemsOK) GetPayload() *models.GetCMDClassSchema {
	return o.Payload
}

func (o *GetTableItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCMDClassSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
func NewGetTableItemsUnauthorized() *GetTableItemsUnauthorized {
	return &GetTableItemsUnauthorized{}
}

/* GetTableItemsUnauthorized describes a response with status code 401, with default header values.
Unauthorized
*/
type GetTableItemsUnauthorized struct {
}

func (o *GetTableItemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /TableItems/api/TableItemss/{TableItemsId}][%d] getTableItemsUnauthorized ", 401)
}

func (o *GetTableItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTableItemsForbidden creates a GetTableItemsForbidden with default headers values
func NewGetTableItemsForbidden() *GetTableItemsForbidden {
	return &GetTableItemsForbidden{}
}

/* GetTableItemsForbidden describes a response with status code 403, with default header values.
Forbidden
*/
type GetTableItemsForbidden struct {
}

func (o *GetTableItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /TableItems/api/TableItemss/{TableItemsId}][%d] getTableItemsForbidden ", 403)
}

func (o *GetTableItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTableItemsNotFound creates a GetTableItemsNotFound with default headers values
func NewGetTableItemsNotFound() *GetTableItemsNotFound {
	return &GetTableItemsNotFound{}
}

/* GetTableItemsNotFound describes a response with status code 404, with default header values.
Not Found
*/
type GetTableItemsNotFound struct {
	Payload *models.Error
}

func (o *GetTableItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /TableItems/api/TableItemss/{TableItemsId}][%d] getTableItemsNotFound  %+v", 404, o.Payload)
}
func (o *GetTableItemsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTableItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
