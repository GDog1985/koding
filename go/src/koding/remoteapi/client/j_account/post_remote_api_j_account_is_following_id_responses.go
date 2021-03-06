package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAccountIsFollowingIDReader is a Reader for the PostRemoteAPIJAccountIsFollowingID structure.
type PostRemoteAPIJAccountIsFollowingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountIsFollowingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountIsFollowingIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountIsFollowingIDOK creates a PostRemoteAPIJAccountIsFollowingIDOK with default headers values
func NewPostRemoteAPIJAccountIsFollowingIDOK() *PostRemoteAPIJAccountIsFollowingIDOK {
	return &PostRemoteAPIJAccountIsFollowingIDOK{}
}

/*PostRemoteAPIJAccountIsFollowingIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountIsFollowingIDOK struct {
	Payload *models.JAccount
}

func (o *PostRemoteAPIJAccountIsFollowingIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.isFollowing/{id}][%d] postRemoteApiJAccountIsFollowingIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountIsFollowingIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
