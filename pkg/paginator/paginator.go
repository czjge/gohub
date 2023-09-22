package paginator

import (
	"fmt"
	"math"
	"strings"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 分页数据
type Paging struct {
	CurrentPage int
	PerPage     int
	TotalPage   int
	TotalCount  int64
	NextPageURL string
	PrevPageURL string
}

// 分页操作类
type Paginator struct {
	BaseURL    string // 用以拼接 URL
	PerPage    int
	Page       int
	Offset     int // 数据库读取数据时 Offset 的值
	TotalCount int64
	TotalPage  int
	Sort       string
	Order      string

	query *gorm.DB // db query 句柄
	ctx   *gin.Context
}

// Paginate 分页
// c —— gin.context 用来获取分页的 URL 参数
// db —— GORM 查询句柄，用以查询数据集和获取数据总数
// baseURL —— 用以分页链接
// data —— 模型数组，传址获取数据
// PerPage —— 每页条数，优先从 url 参数里取，否则使用 perPage 的值
// 用法:
//
//	query := database.DB.Model(Topic{}).Where("category_id = ?", cid)
//	var topics []Topic
//	paging := paginator.Paginate(
//	    c,
//	    query,
//	    &topics,
//	    app.APIURL(database.TableName(&Topic{})),
//	    perPage,
//	)
func Paginate(c *gin.Context, db *gorm.DB, data any, baseURL string, perPage int) Paging {

	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseURL)

	// clause.Associations can work with Preload similar like Select when creating/updating, you can use it to Preload all associations
	// clause.Associations won’t preload nested associations, but you can use it with Nested Preloading together
	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

func (p Paginator) getPageLink(page int) string {
	config := config.GetConfig().Paging
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.UrlQuerySort,
		p.Sort,
		config.UrlQueryOrder,
		p.Order,
		config.UrlQueryPerPage,
		p.PerPage,
	)
}

// 初始化分页必须用到的属性，基于这些属性查询数据库
func (p *Paginator) initProperties(perPage int, baseURL string) {

	p.BaseURL = p.formatBaseURL(baseURL)
	p.PerPage = p.getPerPage(perPage)

	p.Order = p.ctx.DefaultQuery(config.GetConfig().Paging.UrlQueryOrder, "desc")
	p.Sort = p.ctx.DefaultQuery(config.GetConfig().Paging.UrlQuerySort, "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p Paginator) getPerPage(perPage int) int {

	// 优先使用请求 per_page 参数
	queryPerPage := p.ctx.Query(config.GetConfig().Paging.UrlQueryPerPage)
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	// 没有传参，使用默认
	if perPage <= 0 {
		perPage = config.GetConfig().Paging.Perpage
	}

	return perPage
}

func (p Paginator) getCurrentPage() int {

	// 优先取用户请求的 page
	page := cast.ToInt(p.ctx.Query(config.GetConfig().Paging.UrlQueryPage))
	if page <= 0 {
		page = 1
	}

	// TotalPage 等于 0 ，意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}

	// 请求页数大于总页数，返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}

	return page
}

func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

// 兼容 URL 带与不带 `?` 的情况
func (p *Paginator) formatBaseURL(baseURL string) string {
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.GetConfig().Paging.UrlQueryPage + "="
	} else {
		baseURL = baseURL + "?" + config.GetConfig().Paging.UrlQueryPage + "="
	}
	return baseURL
}

func (p Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

func (p Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.PerPage - 1)
}
