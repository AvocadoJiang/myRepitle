package jsonEntity

//https://mholt.github.io/json-to-go/
//岗位列表实体类

type EngineJds struct {
	Type              string        `json:"type"`
	Jt                string        `json:"jt"`
	Tags              []interface{} `json:"tags"`
	AdTrack           string        `json:"ad_track"`
	Jobid             string        `json:"jobid"`
	Coid              string        `json:"coid"`
	Effect            string        `json:"effect"`
	IsSpecialJob      string        `json:"is_special_job"`
	JobHref           string        `json:"job_href"`
	JobName           string        `json:"job_name"`
	JobTitle          string        `json:"job_title"`
	CompanyHref       string        `json:"company_href"`
	CompanyName       string        `json:"company_name"`
	ProvidesalaryText string        `json:"providesalary_text"`
	Workarea          string        `json:"workarea"`
	WorkareaText      string        `json:"workarea_text"`
	Updatedate        string        `json:"updatedate"`
	Iscommunicate     string        `json:"iscommunicate"`
	CompanytypeText   string        `json:"companytype_text"`
	Degreefrom        string        `json:"degreefrom"`
	Workyear          string        `json:"workyear"`
	Issuedate         string        `json:"issuedate"`
	IsFromXyz         string        `json:"isFromXyz"`
	IsIntern          string        `json:"isIntern"`
	Jobwelf           string        `json:"jobwelf"`
	JobwelfList       []string      `json:"jobwelf_list"`
	Isdiffcity        string        `json:"isdiffcity"`
	AttributeText     []string      `json:"attribute_text"`
	CompanysizeText   string        `json:"companysize_text"`
	CompanyindText    string        `json:"companyind_text"`
	Adid              string        `json:"adid"`
}
type JobListResult struct {
	TopAds                []interface{} `json:"top_ads"`
	AuctionAds            []interface{} `json:"auction_ads"`
	MarketAds             []interface{} `json:"market_ads"`
	EngineJds             []EngineJds   `json:"engine_jds"`
	JobidCount            string        `json:"jobid_count"`
	Group                 []interface{} `json:"group"`
	BannerAds             string        `json:"banner_ads"`
	IsCollapseexpansion   string        `json:"is_collapseexpansion"`
	CoAds                 []interface{} `json:"co_ads"`
	KeywordRecommendation struct {
		Title    string `json:"title"`
		DataType string `json:"data_type"`
		Keyword  string `json:"keyword"`
		Data     []struct {
			Href  string `json:"href"`
			Text  string `json:"text"`
			Click string `json:"click"`
		} `json:"data"`
	} `json:"keyword_recommendation"`
	SearchCondition struct {
		Lang           string `json:"lang"`
		Keywordtype    string `json:"keywordtype"`
		OrdField       string `json:"ord_field"`
		Jobarea        string `json:"jobarea"`
		CurrPage       string `json:"curr_page"`
		Workarea       string `json:"workarea"`
		District       string `json:"district"`
		Dibiaoid       string `json:"dibiaoid"`
		Postchannel    string `json:"postchannel"`
		Reservechannel string `json:"reservechannel"`
		Issuedate      string `json:"issuedate"`
		Providesalary  string `json:"providesalary"`
		Degreefrom     string `json:"degreefrom"`
		Companysize    string `json:"companysize"`
		Cotype         string `json:"cotype"`
		Workyear       string `json:"workyear"`
		Industrytype   string `json:"industrytype"`
		Funtype        string `json:"funtype"`
		Jobterm        string `json:"jobterm"`
		Keyword        string `json:"keyword"`
		Welfare        string `json:"welfare"`
		Address        string `json:"address"`
		Line           string `json:"line"`
		Confirmdate    string `json:"confirmdate"`
		Radius         string `json:"radius"`
		Lonlat         string `json:"lonlat"`
		Exparea        string `json:"exparea"`
		Groupid        string `json:"groupid"`
		Coid           string `json:"coid"`
	} `json:"search_condition"`
	SearchedCondition   string        `json:"searched_condition"`
	CurrPage            string        `json:"curr_page"`
	TotalPage           string        `json:"total_page"`
	KeywordAds          []interface{} `json:"keyword_ads"`
	JobSearchAssistance string        `json:"job_search_assistance"`
	SeoTitle            string        `json:"seo_title"`
	SeoDescription      string        `json:"seo_description"`
	SeoKeywords         string        `json:"seo_keywords"`
}
