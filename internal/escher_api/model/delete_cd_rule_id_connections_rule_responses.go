// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCdRuleIDConnectionsRuleReader is a Reader for the DeleteCdRuleIDConnectionsRule structure.
type DeleteCdRuleIDConnectionsRuleReader struct {
	Formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCdRuleIDConnectionsRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCdRuleIDConnectionsRuleNoContent()
		if err := result.readResponse(response, consumer, o.Formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCdRuleIDConnectionsRuleNoContent creates a DeleteCdRuleIDConnectionsRuleNoContent with default headers values
func NewDeleteCdRuleIDConnectionsRuleNoContent() *DeleteCdRuleIDConnectionsRuleNoContent {
	return &DeleteCdRuleIDConnectionsRuleNoContent{}
}

/*DeleteCdRuleIDConnectionsRuleNoContent handles this case with default header values.

deleted
*/
type DeleteCdRuleIDConnectionsRuleNoContent struct {
}

func (o *DeleteCdRuleIDConnectionsRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cd/{ruleId}/connectionsRule][%d] deleteCdRuleIdConnectionsRuleNoContent ", 204)
}

func (o *DeleteCdRuleIDConnectionsRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}