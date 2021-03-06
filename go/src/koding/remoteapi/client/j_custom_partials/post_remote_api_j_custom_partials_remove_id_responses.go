package j_custom_partials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJCustomPartialsRemoveIDReader is a Reader for the PostRemoteAPIJCustomPartialsRemoveID structure.
type PostRemoteAPIJCustomPartialsRemoveIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJCustomPartialsRemoveIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJCustomPartialsRemoveIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJCustomPartialsRemoveIDOK creates a PostRemoteAPIJCustomPartialsRemoveIDOK with default headers values
func NewPostRemoteAPIJCustomPartialsRemoveIDOK() *PostRemoteAPIJCustomPartialsRemoveIDOK {
	return &PostRemoteAPIJCustomPartialsRemoveIDOK{}
}

/*PostRemoteAPIJCustomPartialsRemoveIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJCustomPartialsRemoveIDOK struct {
	Payload *models.JCustomPartials
}

func (o *PostRemoteAPIJCustomPartialsRemoveIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCustomPartials.remove/{id}][%d] postRemoteApiJCustomPartialsRemoveIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJCustomPartialsRemoveIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JCustomPartials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
