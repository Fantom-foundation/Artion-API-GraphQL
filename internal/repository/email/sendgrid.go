// Package email provides emailing service interface.
package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// DynamicTemplateData represents a list of key->value pairs for dynamic template data.
type DynamicTemplateData map[string]interface{}

// SendGridDeliverDynamicTemplate sends the given email to the recipient using SendGrid Web API.
// The email body is composed on the SendGrid side using specified dynamic template.
func SendGridDeliverDynamicTemplate(
	sender *mail.Email,
	recipient *mail.Email,
	templateID string,
	subject string,
	data DynamicTemplateData,
) error {
	// make sure we have the API key
	if cfg.Notifications.SendGrid == nil {
		log.Warningf("missing SendGrid API configuration")
		return nil
	}

	// init the email
	m := mail.NewV3MailInit(sender, subject, recipient)

	// push dynamic data and the template ID
	m.TemplateID = templateID
	m.Personalizations[0].DynamicTemplateData = data

	// make the request
	req := sendgrid.GetRequest(cfg.Notifications.SendGrid.ApiKey, "/v3/mail/send", cfg.Notifications.SendGrid.ApiAddress)
	req.Method = "POST"
	req.Body = mail.GetRequestBody(m)

	// send API request
	resp, err := sendgrid.API(req)
	if err != nil {
		log.Errorf("failed to submit email; %s", err.Error())
		return err
	}

	log.Noticef("sent email notification to %s, response code #%d", recipient.Address, resp.StatusCode)
	return nil
}
