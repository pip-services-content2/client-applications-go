package version1

import (
	"context"
	"strings"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ApplicationsMockClientV1 struct {
	applications []*ApplicationV1
}

func NewApplicationsMockClientV1() *ApplicationsMockClientV1 {
	return &ApplicationsMockClientV1{
		applications: make([]*ApplicationV1, 0),
	}
}

func (c *ApplicationsMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *ApplicationsMockClientV1) matchSearch(item *ApplicationV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Id, search) {
		return true
	}
	if c.matchString(item.Product, search) {
		return true
	}
	if c.matchString(item.Copyrights, search) {
		return true
	}
	return false
}

func (c *ApplicationsMockClientV1) composeFilter(filter *data.FilterParams) func(item *ApplicationV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	product, productOk := filter.GetAsNullableString("product")
	group, groupOk := filter.GetAsNullableString("group")

	return func(item *ApplicationV1) bool {
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		if idOk && id != item.Id {
			return false
		}
		if productOk && product != item.Product {
			return false
		}
		if groupOk && group != item.Group {
			return false
		}

		return true
	}
}

func (c *ApplicationsMockClientV1) GetApplications(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*ApplicationV1], err error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*ApplicationV1, 0)
	for _, v := range c.applications {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.applications)), nil
}

func (c *ApplicationsMockClientV1) GetApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	for _, v := range c.applications {
		if v.Id == id {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *ApplicationsMockClientV1) CreateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	buf := *application
	c.applications = append(c.applications, &buf)
	return application, nil
}

func (c *ApplicationsMockClientV1) UpdateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	if application == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.applications {
		if v.Id == application.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	buf := *application
	c.applications[index] = &buf
	return application, nil
}

func (c *ApplicationsMockClientV1) DeleteApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	var index = -1
	for i, v := range c.applications {
		if v.Id == id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	var item = c.applications[index]
	c.applications = append(c.applications[:index], c.applications[index+1:]...)
	return item, nil
}
