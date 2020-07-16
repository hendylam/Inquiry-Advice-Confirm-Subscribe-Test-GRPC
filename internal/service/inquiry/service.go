package inquiry

import (
	"context"

	mic "tapera.integrasi/grpc/client/mitraintegrasi/v1"
)

type (
	Service interface {
		InquiryConfirmSubs(ctx context.Context, parm *InquiryConfirmSubData) (*InquiryConfirmSubResponse, error)
		ConvertStringtoJson(parm *InquiryConfirmSubResponse) *InquiryConfirmSubResponseSuccessJson
	}

	TokenResult struct {
		OrganizationName string `json:"organization_name"`
		IssuedAt         string `json:"issued_at"`
		ExpiresIn        string `json:"expires_in"`
		ClientID         string `json:"client_id"`
		AccessToken      string `json:"access_token"`
		ApplicationName  string `json:"application_name"`
	}

	service struct {
		miClientMgr mic.GrpcClientManager
	}

	InquiryConfirmSubData struct {
		PaymentPoolId string `json:"payment_pool_id"`
	}

	InquiryConfirmSubResponse struct {
		ResponseCode        string `json:"response_code"`
		ResponseDescription string `json:"response_description"`
		ResponseStatus      string `json:"response_status"`
		Data                Data
	}

	InquiryConfirmSubResponseSuccess struct {
		ResponseCode        string `json:"response_code"`
		ResponseDescription string `json:"response_description"`
		Status              string `json:"response_status"`
		Data                Data
	}

	InquiryConfirmSubResponseSuccessJson struct {
		ResponseCode        string   `json:"response_code"`
		ResponseDescription string   `json:"response_description"`
		Status              string   `json:"response_status"`
		Data                FailData `json:"data"`
	}

	FailData struct {
		PaymentPoolId string `json:"payment_pool_id"`
	}

	Data struct {
		PaymentPoolId string `json:"payment_pool_id"`
		NavDate       string `json:"nav_date"`
		Status        string `json:"status"`
	}

	InquiryConfirmSubResponseError struct {
		ResponseCode        string `json:"response_code"`
		ResponseDescription string `json:"response_description"`
		Status              Status
	}

	Status struct {
		Code string `json:"code"`
		Desc string `json:"desc"`
	}
)

func NewService(miClientMgr mic.GrpcClientManager) Service {
	return &service{miClientMgr: miClientMgr}
}
