// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/remediation/models"
)

// RemediateResourceAsyncReader is a Reader for the RemediateResourceAsync structure.
type RemediateResourceAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemediateResourceAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemediateResourceAsyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemediateResourceAsyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemediateResourceAsyncInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemediateResourceAsyncOK creates a RemediateResourceAsyncOK with default headers values
func NewRemediateResourceAsyncOK() *RemediateResourceAsyncOK {
	return &RemediateResourceAsyncOK{}
}

/*RemediateResourceAsyncOK handles this case with default header values.

OK
*/
type RemediateResourceAsyncOK struct {
}

func (o *RemediateResourceAsyncOK) Error() string {
	return fmt.Sprintf("[POST /remediateasync][%d] remediateResourceAsyncOK ", 200)
}

func (o *RemediateResourceAsyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemediateResourceAsyncBadRequest creates a RemediateResourceAsyncBadRequest with default headers values
func NewRemediateResourceAsyncBadRequest() *RemediateResourceAsyncBadRequest {
	return &RemediateResourceAsyncBadRequest{}
}

/*RemediateResourceAsyncBadRequest handles this case with default header values.

Bad request
*/
type RemediateResourceAsyncBadRequest struct {
	Payload *models.Error
}

func (o *RemediateResourceAsyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /remediateasync][%d] remediateResourceAsyncBadRequest  %+v", 400, o.Payload)
}

func (o *RemediateResourceAsyncBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemediateResourceAsyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemediateResourceAsyncInternalServerError creates a RemediateResourceAsyncInternalServerError with default headers values
func NewRemediateResourceAsyncInternalServerError() *RemediateResourceAsyncInternalServerError {
	return &RemediateResourceAsyncInternalServerError{}
}

/*RemediateResourceAsyncInternalServerError handles this case with default header values.

Internal server error
*/
type RemediateResourceAsyncInternalServerError struct {
}

func (o *RemediateResourceAsyncInternalServerError) Error() string {
	return fmt.Sprintf("[POST /remediateasync][%d] remediateResourceAsyncInternalServerError ", 500)
}

func (o *RemediateResourceAsyncInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}