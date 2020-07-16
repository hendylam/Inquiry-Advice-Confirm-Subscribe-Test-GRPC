package inquiry

func (s *service) ConvertStringtoJson(parm *InquiryConfirmSubResponse) *InquiryConfirmSubResponseSuccessJson {

	FailData := FailData{
		PaymentPoolId: parm.Data.PaymentPoolId,
	}

	InquiryResponse := &InquiryConfirmSubResponseSuccessJson{
		Status:              parm.ResponseStatus,
		ResponseCode:        parm.ResponseCode,
		ResponseDescription: parm.ResponseDescription,
		Data:                FailData,
	}

	return InquiryResponse

}
