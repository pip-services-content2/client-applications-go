package version1

type ApplicationV1 struct {
	Id           string            `json:"id"`
	Name         map[string]string `json:"name"`
	Description  map[string]string `json:"description"`
	Product      string            `json:"product"`
	Group        string            `json:"group"`
	Copyrights   string            `json:"copyrights"`
	Url          string            `json:"url"`
	Icon         string            `json:"icon"`
	MinVer       int32             `json:"min_ver"`
	MaxVer       int32             `json:"max_ver"`
	AccessRights []string          `json:"access_rights"`
}

func EmptyApplicationV1() *ApplicationV1 {
	return &ApplicationV1{}
}

func NewApplicationV1(id string, name string, product string, group string) *ApplicationV1 {
	return &ApplicationV1{
		Id:      id,
		Name:    map[string]string{"en": name},
		Product: product,
		Group:   group,
	}
}
