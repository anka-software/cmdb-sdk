package cmdb

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anka-software/cmdb-sdk/pkg/models"
)

// CreateIdentifyReconcileReader is a Reader for the CreateIdentifyReconcile structure.
type CreateIdentifyReconcileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIdentifyReconcileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIdentifyReconcileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateIdentifyReconcileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateIdentifyReconcileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateIdentifyReconcileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateIdentifyReconcileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint", response, response.Code())
	}
}

// NewCreateIdentifyReconcileCreated creates a CreateIdentifyReconcileCreated with default headers values
func NewCreateIdentifyReconcileCreated() *CreateIdentifyReconcileCreated {
	return &CreateIdentifyReconcileCreated{}
}

/* CreateIdentifyReconcileCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateIdentifyReconcileCreated struct {
	Payload *models.IdentifyReconcileItem
}

func (o *CreateIdentifyReconcileCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/IdentifyReconciles][%d] createIdentifyReconcileCreated  %+v", 201, o.Payload)
}
func (o *CreateIdentifyReconcileCreated) GetPayload() *models.IdentifyReconcileItem {
	return o.Payload
}

func (o *CreateIdentifyReconcileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentifyReconcileItem)
	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIdentifyReconcileBadRequest creates a CreateIdentifyReconcileBadRequest with default headers values
func NewCreateIdentifyReconcileBadRequest() *CreateIdentifyReconcileBadRequest {
	return &CreateIdentifyReconcileBadRequest{}
}

/* CreateIdentifyReconcileBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateIdentifyReconcileBadRequest struct {
	Payload *models.Error
}

func (o *CreateIdentifyReconcileBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/IdentifyReconciles][%d] createIdentifyReconcileBadRequest  %+v", 400, o.Payload)
}
func (o *CreateIdentifyReconcileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIdentifyReconcileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIdentifyReconcileForbidden creates a CreateIdentifyReconcileForbidden with default headers values
func NewCreateIdentifyReconcileForbidden() *CreateIdentifyReconcileForbidden {
	return &CreateIdentifyReconcileForbidden{}
}

/* CreateIdentifyReconcileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateIdentifyReconcileForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *CreateIdentifyReconcileForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/IdentifyReconciles][%d] createIdentifyReconcileForbidden  %+v", 403, o.Payload)
}
func (o *CreateIdentifyReconcileForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateIdentifyReconcileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payloadx
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
