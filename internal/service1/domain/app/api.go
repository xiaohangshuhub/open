package app

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
	"github.com/xiaohangshuhub/open/internal/service1/domain/err"
)

type API struct {
	ddd.AggregateRoot[uuid.UUID]           // ID
	Name                         string    // 名称
	Version                      string    // 版本
	Description                  string    // 描述
	Endpoint                     string    // 访问地址
	Method                       string    // 请求方法
	AuthType                     string    // 鉴权类型
	Created                      int64     // 创建时间
	CreateBy                     uuid.UUID // 创建人
	Updated                      *int64    // 修改时间
	UpdateBy                     *string   // 修改人
	AppID                        uuid.UUID // 所属应用ID
	Docs                         string    // 文档地址
}

func NewAPI(name, version, description, endpoint, method, authType, docs string, appID, createBy uuid.UUID) (*API, *err.Error) {

	api := &API{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
	}

	if err := api.SetName(name); err != nil {
		return nil, err
	}
	if err := api.SetVersion(version); err != nil {
		return nil, err
	}
	if err := api.SetDescription(description); err != nil {
		return nil, err
	}
	if err := api.SetEndpoint(endpoint); err != nil {
		return nil, err
	}
	if err := api.SetMethod(method); err != nil {
		return nil, err
	}
	if err := api.SetAuthType(authType); err != nil {
		return nil, err
	}
	if err := api.SetAppID(appID); err != nil {
		return nil, err
	}
	if err := api.SetDocs(docs); err != nil {
		return nil, err
	}
	if createBy == uuid.Nil {
		return nil, err.ErrCreateByEmpty
	}

	return api, nil
}

// SetName 设置API名称
func (a *API) SetName(name string) *err.Error {
	if name == "" {
		return err.ErrNameEmpty
	}
	a.Name = name
	return nil
}

// SetAppID 设置所属应用ID
func (a *API) SetAppID(appID uuid.UUID) *err.Error {
	if appID == uuid.Nil {
		return err.ErrAppIDEmpty
	}
	a.AppID = appID
	return nil
}

// SetDescription 设置描述
func (a *API) SetDescription(description string) *err.Error {
	if description == "" {
		return err.ErrDescEmpty
	}
	a.Description = description
	return nil
}

// SetEndpoint 设置访问地址
func (a *API) SetEndpoint(endpoint string) *err.Error {
	if endpoint == "" {
		return err.ErrDescEmpty
	}
	a.Endpoint = endpoint
	return nil
}

// SetMethod 设置请求方法
func (a *API) SetMethod(method string) *err.Error {
	if method == "" {
		return err.ErrDescEmpty
	}
	a.Method = method
	return nil
}

// SetAuthType 设置鉴权类型
func (a *API) SetAuthType(authType string) *err.Error {
	if authType == "" {
		return err.ErrDescEmpty
	}
	a.AuthType = authType
	return nil
}

// SetVersion 设置版本
func (a *API) SetVersion(version string) *err.Error {
	if version == "" {
		return err.ErrDescEmpty
	}
	a.Version = version
	return nil
}

// SetDocs 设置文档地址
func (a *API) SetDocs(docs string) *err.Error {
	if docs == "" {
		return err.ErrDescEmpty
	}
	a.Docs = docs
	return nil
}
