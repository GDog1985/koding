package j_machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJMachineShareIDReader is a Reader for the PostRemoteAPIJMachineShareID structure.
type PostRemoteAPIJMachineShareIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJMachineShareIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJMachineShareIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJMachineShareIDOK creates a PostRemoteAPIJMachineShareIDOK with default headers values
func NewPostRemoteAPIJMachineShareIDOK() *PostRemoteAPIJMachineShareIDOK {
	return &PostRemoteAPIJMachineShareIDOK{}
}

/*PostRemoteAPIJMachineShareIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJMachineShareIDOK struct {
	Payload *models.JMachine
}

func (o *PostRemoteAPIJMachineShareIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JMachine.share/{id}][%d] postRemoteApiJMachineShareIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJMachineShareIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
