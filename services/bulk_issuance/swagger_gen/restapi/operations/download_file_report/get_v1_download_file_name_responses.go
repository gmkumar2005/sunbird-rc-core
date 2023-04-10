// Code generated by go-swagger; DO NOT EDIT.

package download_file_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bulk_issuance/swagger_gen/models"
)

// GetV1DownloadFileNameOKCode is the HTTP code returned for type GetV1DownloadFileNameOK
const GetV1DownloadFileNameOKCode int = 200

/*GetV1DownloadFileNameOK OK

swagger:response getV1DownloadFileNameOK
*/
type GetV1DownloadFileNameOK struct {
	/*

	 */
	ContentDisposition string `json:"Content-Disposition"`

	/*
	  In: Body
	*/
	Payload models.FileDownload `json:"body,omitempty"`
}

// NewGetV1DownloadFileNameOK creates GetV1DownloadFileNameOK with default headers values
func NewGetV1DownloadFileNameOK() *GetV1DownloadFileNameOK {

	return &GetV1DownloadFileNameOK{}
}

// WithContentDisposition adds the contentDisposition to the get v1 download file name o k response
func (o *GetV1DownloadFileNameOK) WithContentDisposition(contentDisposition string) *GetV1DownloadFileNameOK {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the get v1 download file name o k response
func (o *GetV1DownloadFileNameOK) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithPayload adds the payload to the get v1 download file name o k response
func (o *GetV1DownloadFileNameOK) WithPayload(payload models.FileDownload) *GetV1DownloadFileNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 download file name o k response
func (o *GetV1DownloadFileNameOK) SetPayload(payload models.FileDownload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1DownloadFileNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
