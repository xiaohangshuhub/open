package app

import (
	"time"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
	"github.com/xiaohangshuhub/open/internal/service1/domain/dic/status"
	"github.com/xiaohangshuhub/open/internal/service1/domain/err"
)

// AppType 应用类型
type AppType int8

const (
	WebApp     AppType = iota + 1 // web应用
	DesptopApp                    // 桌面应用
	MobileApp                     // 移动应用
	MiniApp                       // 小程序
	ClientApp                     // 客户端应用
	IoTApp                        // 物联网应用
)

const (
	clientNameLengthMin = 2
	clientNameLengthMax = 10
)

// Client 客户端实体
type App struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Name                         string        // 名称
	Key                          string        // app key
	Security                     string        // app security
	appType                      AppType       // 客户端类型
	Desc                         string        // 描述
	Status                       status.Status // 状态
	Created                      time.Time     // 创建时间
	CreateBy                     uuid.UUID     // 创建人
	Updated                      *time.Time    // 修改时间
	UpdateBy                     *string       // 修改人
}

func NewApp(name string, createBy uuid.UUID, appType AppType) (*App, *err.Error) {

	app := App{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Key:           uuid.NewString(),
		Security:      uuid.NewString(),
		Status:        status.Review,
		Created:       time.Now(),
		appType:       appType,
	}

	if err := app.SetName(name); err != nil {
		return nil, err
	}

	if createBy == uuid.Nil {
		return nil, err.ErrCreateByEmpty
	}

	app.CreateBy = createBy

	return &app, nil
}

// SetName 设置客户端的名称
func (c *App) SetName(name string) *err.Error {
	if name == "" {
		return err.ErrNameEmpty
	}
	if len := len(name); len > clientNameLengthMax || len < clientNameLengthMin {
		return err.ErrLengthOutOfRange
	}
	c.Name = name
	return nil
}
