package j_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJSessionRemoveIDReader is a Reader for the PostRemoteAPIJSessionRemoveID structure.
type PostRemoteAPIJSessionRemoveIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJSessionRemoveIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJSessionRemoveIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJSessionRemoveIDOK creates a PostRemoteAPIJSessionRemoveIDOK with default headers values
func NewPostRemoteAPIJSessionRemoveIDOK() *PostRemoteAPIJSessionRemoveIDOK {
	return &PostRemoteAPIJSessionRemoveIDOK{}
}

/*PostRemoteAPIJSessionRemoveIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJSessionRemoveIDOK struct {
	Payload *models.JSession
}

func (o *PostRemoteAPIJSessionRemoveIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JSession.remove/{id}][%d] postRemoteApiJSessionRemoveIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJSessionRemoveIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
