package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJGroupFetchAdminsWithEmailIDReader is a Reader for the PostRemoteAPIJGroupFetchAdminsWithEmailID structure.
type PostRemoteAPIJGroupFetchAdminsWithEmailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupFetchAdminsWithEmailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupFetchAdminsWithEmailIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupFetchAdminsWithEmailIDOK creates a PostRemoteAPIJGroupFetchAdminsWithEmailIDOK with default headers values
func NewPostRemoteAPIJGroupFetchAdminsWithEmailIDOK() *PostRemoteAPIJGroupFetchAdminsWithEmailIDOK {
	return &PostRemoteAPIJGroupFetchAdminsWithEmailIDOK{}
}

/*PostRemoteAPIJGroupFetchAdminsWithEmailIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupFetchAdminsWithEmailIDOK struct {
	Payload *models.JGroup
}

func (o *PostRemoteAPIJGroupFetchAdminsWithEmailIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.fetchAdminsWithEmail/{id}][%d] postRemoteApiJGroupFetchAdminsWithEmailIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupFetchAdminsWithEmailIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
