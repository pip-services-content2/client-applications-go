package version1

import (
	"encoding/json"

	"github.com/pip-services-content2/client-applications-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]any {
	var r map[string]any

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromApplication(application *ApplicationV1) *protos.Application {
	if application == nil {
		return nil
	}

	obj := &protos.Application{
		Id:           application.Id,
		Name:         application.Name,
		Description:  application.Description,
		Product:      application.Product,
		Group:        application.Group,
		Copyrights:   application.Copyrights,
		Url:          application.Url,
		Icon:         application.Icon,
		MinVer:       application.MinVer,
		MaxVer:       application.MaxVer,
		AccessRights: application.AccessRights,
	}

	return obj
}

func toApplication(obj *protos.Application) *ApplicationV1 {
	if obj == nil {
		return nil
	}

	application := &ApplicationV1{
		Id:           obj.Id,
		Name:         obj.Name,
		Description:  obj.Description,
		Product:      obj.Product,
		Group:        obj.Group,
		Copyrights:   obj.Copyrights,
		Url:          obj.Url,
		Icon:         obj.Icon,
		MinVer:       obj.MinVer,
		MaxVer:       obj.MaxVer,
		AccessRights: obj.AccessRights,
	}

	return application
}

func fromApplicationPage(page data.DataPage[*ApplicationV1]) *protos.ApplicationPage {
	obj := &protos.ApplicationPage{
		Total: int64(page.Total),
		Data:  make([]*protos.Application, len(page.Data)),
	}

	for i, v := range page.Data {
		application := v
		obj.Data[i] = fromApplication(application)
	}

	return obj
}

func toApplicationPage(obj *protos.ApplicationPage) data.DataPage[*ApplicationV1] {
	if obj == nil {
		return *data.NewEmptyDataPage[*ApplicationV1]()
	}

	applications := make([]*ApplicationV1, len(obj.Data))

	for i, v := range obj.Data {
		applications[i] = toApplication(v)
	}

	page := *data.NewDataPage(applications, int(obj.Total))

	return page
}
