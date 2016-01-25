package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/read

// Imports
import (
	"fmt"
)

// UserEmailAccountsFolderMessageReadResponse ...
type UserEmailAccountsFolderMessageReadResponse struct {
	Success int `json:"success,omitempty"`
}

// MarkUserEmailAccountsFolderMessageRead marks the message as read.
// formValues may optionally contain CioParams.Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#post
func (cioLite *CioLite) MarkUserEmailAccountsFolderMessageRead(userID string, label string, folder string, messageID string, formValues CioParams) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, folder, messageID),
		formValues: formValues,
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// MarkUserEmailAccountsFolderMessageUnRead marks the message as unread.
// formValues may optionally contain CioParams.Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#delete
func (cioLite *CioLite) MarkUserEmailAccountsFolderMessageUnRead(userID string, label string, folder string, messageID string, formValues CioParams) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := clientRequest{
		method:     "DELETE",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, folder, messageID),
		formValues: formValues,
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
