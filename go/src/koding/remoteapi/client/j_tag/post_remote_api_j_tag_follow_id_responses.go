package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTagFollowIDReader is a Reader for the PostRemoteAPIJTagFollowID structure.
type PostRemoteAPIJTagFollowIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagFollowIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagFollowIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagFollowIDOK creates a PostRemoteAPIJTagFollowIDOK with default headers values
func NewPostRemoteAPIJTagFollowIDOK() *PostRemoteAPIJTagFollowIDOK {
	return &PostRemoteAPIJTagFollowIDOK{}
}

/*PostRemoteAPIJTagFollowIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJTagFollowIDOK struct {
	Payload *models.JTag
}

func (o *PostRemoteAPIJTagFollowIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.follow/{id}][%d] postRemoteApiJTagFollowIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagFollowIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
