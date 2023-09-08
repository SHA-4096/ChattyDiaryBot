package util

type RespBingImageAPISuccess struct {
	RespType     string `json:"_type"`        //"_type": "Images",
	ReadLink     string `json:"readLink"`     //"readLink": "images/search?q=black cocktail dresses",
	WebSearchUrl string `json:"webSearchUrl"` //"webSearchUrl": "https://www.bing.com/images/search?q=black cocktail dresses&FORM=OIIARP",
	//QueryContext          string      `json:"queryContext"`          //"queryContext": { ... },
	TotalEstimatedMatches int           `json:"totalEstimatedMatches"` //"totalEstimatedMatches": 835,
	NextOffset            int           `json:"nextOffset"`            //"nextOffset": 36,
	CurrentOffset         int           `json:"currentOffset"`         //"currentOffset": 0,
	Value                 []ValueStruct `json:"value"`
	//"queryExpansions": [ { ... } ],
	//"pivotSuggestions": [ { ... } ],
	//"relatedSearches": [ { ... } ]
}

type ValueStruct struct {
	WebSearchUrl string `json:"webSearchUrl"` //: "https://www.bing.com/images/search?view=detailv2&FORM=OIIRPO...",
	Name         string `json:"name"`         //: "The Perfect Black Cocktail Dress...",
	//"thumbnailUrl": "https://tse4.mm.bing.net/th?id=OIP.fGuCgUtRUl_f2c8...",
	//"datePublished": "2014-11-02T12:00:00.0000000Z",
	//"isFamilyFriendly": true,
	ContentUrl  string `json:"contentUrl"`  //: "http://contoso.com/wp-content/uploads/2014/11/black...",
	HostPageUrl string `json:"hostPageUrl"` //: "http://contoso.com/2014/11/02/the-perfect-black-cocktail-dress...",
	//"contentSize": "171202 B",
	EncodingFormat string `json:"encodingFormat"`     //: "jpeg",
	HostDiaplayUrl string `json:"hostPageDisplayUrl"` //: "contoso.com/2014/11/02/the-perfect-black...",
	//"width": 996,
	//"height": 1500,
	//"thumbnail": {
	//  "width": 474,
	//  "height": 713
	//},
	//"imageInsightsToken": "ccid_fGuCgUtR*cp_D72F7E52B27BF10...",
	//"insightsMetadata": {
	//  "shoppingSourcesCount": 0,
	//  "recipeSourcesCount": 0,
	//  "pagesIncludingCount": 31,
	// "availableSizesCount": 15
	//},
	//"imageId": "1A8B199F1D3ED20462C904A1F5353C57736A2D44",
	//"accentColor": "6F5C5F"
}
