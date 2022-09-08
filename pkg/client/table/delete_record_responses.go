package table

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/anka-software/cmdb-sdk/pkg/models"
)

// DeleteRecordeader is a Reader for the DeleteIdentifyReconcile structure.
type DeleteRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	/*case 200:
	result := NewDeleteRecordOK()
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	return result, nil*/
	case 204:
		result := NewDeleteRecordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRecordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRecordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match defined for this endpoint", response, response.Code())
	}
}

// NewDeleteIdentifyReconcileDeleted Deletes a DeleteIdentifyReconcileDeleted with default headers values
func NewDeleteRecordOK() *DeleteRecordOK {
	return &DeleteRecordOK{}
}

type DeleteRecordOK struct {
	Payload string
}

func (o *DeleteRecordOK) Error() string {
	return fmt.Sprintf("[DELETE /api/now/table/{tableName}/{sys_id}][%d] DeleteRecordForbidden %+v", 204, o.Payload)

}
func (o *DeleteRecordOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteRecordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
func NewDeleteRecordUnauthorized() *DeleteRecordUnauthorized {
	return &DeleteRecordUnauthorized{}
}

/* DeleteRecordUnauthorized describes a response with status code 401, with default header values.
Unauthorized
*/
type DeleteRecordUnauthorized struct {
}

func (o *DeleteRecordUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/now/table/{tableName}/{sys_id}][%d] DeleteRecordUnauthorized ", 401)
}

func (o *DeleteRecordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecordForbidden Deletes a DeleteRecordForbidden with default headers values
func NewDeleteRecordForbidden() *DeleteRecordForbidden {
	return &DeleteRecordForbidden{}
}

/* DeleteRecordForbidden describes a response with status code 403, with default header values.
Forbidden
*/
type DeleteRecordForbidden struct {
}

func (o *DeleteRecordForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/now/table/{tableName}/{sys_id}][%d] DeleteRecordForbidden ", 403)
}

func (o *DeleteRecordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecordNotFound Deletes a DeleteRecordNotFound with default headers values
func NewDeleteRecordNotFound() *DeleteRecordNotFound {
	return &DeleteRecordNotFound{}
}

/* DeleteRecordNotFound describes a response with status code 404, with default header values.
Not Found
*/
type DeleteRecordNotFound struct {
	Payload *models.Error
}

func (o *DeleteRecordNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/now/table/{tableName}/{sys_id}][%d] DeleteRecordNotFound  %+v", 404, o.Payload)
}
func (o *DeleteRecordNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRecordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
