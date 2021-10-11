package api

//type BriefDecoder interface {
//	DecodeGetBriefById(r *http.Request) (*GetBriefByIdRequest, error)
//	DecodeInsertBrief(r *http.Request) (*InsertBriefRequest, error)
//	DecodeUpdateBrief(r *http.Request) (*UpdateBriefRequest, error)
//}
//
//type briefDecoder struct {
//	validator *validator.Validate
//}
//
//func (decoder *briefDecoder) DecodeGetBriefById(r *http.Request) (*GetBriefByIdRequest, error) {
//	dto := &GetBriefByIdRequest{
//		Id: chi.URLParam(r, "id"),
//	}
//	err := decoder.validator.Struct(dto)
//	if err != nil {
//		err = errors.Wrap(utils.ErrInvalidInputParams, err.Error())
//		return nil, err
//	} else {
//		return dto, nil
//	}
//}
//
//func (decoder *briefDecoder) DecodeInsertBrief(r *http.Request) (*InsertBriefRequest, error) {
//	request := &InsertBriefRequest{}
//	err := json.NewDecoder(r.Body).Decode(request)
//	if err != nil {
//		return nil, errors.Wrap(utils.ErrDecodingRequest, err.Error())
//	}
//	err = decoder.validator.Struct(request)
//	if err != nil {
//		return nil, errors.Wrap(utils.ErrInvalidInputParams, err.Error())
//	}
//	return request, err
//}
//
//func (decoder *briefDecoder) DecodeUpdateBrief(r *http.Request) (*UpdateBriefRequest, error) {
//	request := &UpdateBriefRequest{}
//	err := json.NewDecoder(r.Body).Decode(&request)
//	if err != nil {
//		return nil, errors.Wrap(utils.ErrDecodingRequest, err.Error())
//	}
//
//	err = decoder.validator.Struct(request)
//	if err != nil {
//		return nil, errors.Wrap(utils.ErrInvalidInputParams, err.Error())
//	}
//
//	return request, err
//}
//
//func NewBriefDecoder() BriefDecoder {
//	decoder := &briefDecoder{
//		validator: validator.New(),
//	}
//	return decoder
//}
//
