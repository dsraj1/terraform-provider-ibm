// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetImagesReader is a Reader for the GetImages structure.
type GetImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetImagesOK creates a GetImagesOK with default headers values
func NewGetImagesOK() *GetImagesOK {
	return &GetImagesOK{}
}

/*GetImagesOK handles this case with default header values.

dummy
*/
type GetImagesOK struct {
	Payload *models.GetImagesOKBody
}

func (o *GetImagesOK) Error() string {
	return fmt.Sprintf("[GET /images][%d] getImagesOK  %+v", 200, o.Payload)
}

func (o *GetImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetImagesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesDefault creates a GetImagesDefault with default headers values
func NewGetImagesDefault(code int) *GetImagesDefault {
	return &GetImagesDefault{
		_statusCode: code,
	}
}

/*GetImagesDefault handles this case with default header values.

unexpectederror
*/
type GetImagesDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get images default response
func (o *GetImagesDefault) Code() int {
	return o._statusCode
}

func (o *GetImagesDefault) Error() string {
	return fmt.Sprintf("[GET /images][%d] GetImages default  %+v", o._statusCode, o.Payload)
}

func (o *GetImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}