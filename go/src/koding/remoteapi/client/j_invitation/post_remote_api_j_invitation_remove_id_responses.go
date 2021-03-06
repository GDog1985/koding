package j_invitation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJInvitationRemoveIDReader is a Reader for the PostRemoteAPIJInvitationRemoveID structure.
type PostRemoteAPIJInvitationRemoveIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJInvitationRemoveIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJInvitationRemoveIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJInvitationRemoveIDOK creates a PostRemoteAPIJInvitationRemoveIDOK with default headers values
func NewPostRemoteAPIJInvitationRemoveIDOK() *PostRemoteAPIJInvitationRemoveIDOK {
	return &PostRemoteAPIJInvitationRemoveIDOK{}
}

/*PostRemoteAPIJInvitationRemoveIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJInvitationRemoveIDOK struct {
	Payload *models.JInvitation
}

func (o *PostRemoteAPIJInvitationRemoveIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JInvitation.remove/{id}][%d] postRemoteApiJInvitationRemoveIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJInvitationRemoveIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JInvitation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
