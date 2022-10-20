package cmdb_meta

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anka-software/cmdb-sdk/pkg/models"
)

// CreateIdentifyReconcileReader is a Reader for the CreateIdentifyReconcile structure.
type GetCmdbMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCmdbMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCmdbMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCmdbMetaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCmdbMetaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCmdbMetaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match defined for this endpoint", response, response.Code())
	}
}

// NewCreateIdentifyReconcileCreated creates a CreateIdentifyReconcileCreated with default headers values
func NewGetCmdbMetaOK() *GetCmdbMetaOK {
	return &GetCmdbMetaOK{}
}

type GetCmdbMetaOK struct {
	Payload *models.GetCMDClassSchema
}

func (o *GetCmdbMetaOK) Error() string {
	return fmt.Sprintf("[GET /CmdbMeta/api/CmdbMetas/{className}][%d] getCmdbMetaUnauthorized %+v", 200, o.Payload)

}
func (o *GetCmdbMetaOK) GetPayload() *models.GetCMDClassSchema {
	return o.Payload
}

func (o *GetCmdbMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCMDClassSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
func NewGetCmdbMetaUnauthorized() *GetCmdbMetaUnauthorized {
	return &GetCmdbMetaUnauthorized{}
}

/* GetCmdbMetaUnauthorized describes a response with status code 401, with default header values.
Unauthorized
*/
type GetCmdbMetaUnauthorized struct {
}

func (o *GetCmdbMetaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /CmdbMeta/api/CmdbMetas/{className}][%d] getCmdbMetaUnauthorized ", 401)
}

func (o *GetCmdbMetaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCmdbMetaForbidden creates a GetCmdbMetaForbidden with default headers values
func NewGetCmdbMetaForbidden() *GetCmdbMetaForbidden {
	return &GetCmdbMetaForbidden{}
}

/* GetCmdbMetaForbidden describes a response with status code 403, with default header values.
Forbidden
*/
type GetCmdbMetaForbidden struct {
}

func (o *GetCmdbMetaForbidden) Error() string {
	return fmt.Sprintf("[GET /CmdbMeta/api/{className}][%d] getCmdbMetaForbidden ", 403)
}

func (o *GetCmdbMetaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCmdbMetaNotFound creates a GetCmdbMetaNotFound with default headers values
func NewGetCmdbMetaNotFound() *GetCmdbMetaNotFound {
	return &GetCmdbMetaNotFound{}
}

/* GetCmdbMetaNotFound describes a response with status code 404, with default header values.
Not Found
*/
type GetCmdbMetaNotFound struct {
	Payload *models.Error
}

func (o *GetCmdbMetaNotFound) Error() string {
	return fmt.Sprintf("[GET /CmdbMeta/api/CmdbMetas/{className}][%d] getCmdbMetaNotFound  %+v", 404, o.Payload)
}
func (o *GetCmdbMetaNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCmdbMetaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
