package wavefront

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Dashboard represents a single Wavefront Dashboard
type Dashboard struct {
	// Name is the name given to an Dashboard
	Name string `json:"name"`

	// ID is the Wavefront-assigned ID of an existing Dashboard
	ID string `json:"id"`

	// Tags are the tags applied to the Dashboard
	Tags []string `json:"tags"`

	// Description is a description given to the Dashboard
	Description string `json:"description"`

	// Url is the relative url to access the dashboard by on a cluster
	Url string `json:"url"`

	// Sections is an array of Section that split up the dashboard
	Sections []Section `json:"sections"`

	// More dashboard settings (TODO: decriptions)
	ChartTitleBgColor             string `json:"chartTitleBgColor,omitempty"`
	ChartTitleColor               string `json:"chartTitleColor,omitempty"`
	ChartTitleScalar              int    `json:"chartTitleScalar,omitempty"`
	Customer                      string `json:"customer"`
	DefaultEndTime                int    `json:"defaultEndTime,omitempty"`
	DefaultStartTime              int    `json:"defaultStartTime,omitempty"`
	DefaultTimeWindow             string `json:"defaultTimeWindow"`
	Deleted                       bool   `json:"deleted"`
	DisplayDescription            bool   `json:"displayDescription"`
	DisplayQueryParameters        bool   `json:"displayQueryParameters"`
	DisplaySectionTableOfContents bool   `json:"displaySectionTableOfContents"`
	EventFilterType               string `json:"eventFilterType"`
	EventQuery                    string `json:"eventQuery"`
	Favorite                      bool   `json:"favorite"`
	Hidden                        bool   `json:"hidden"`
	NumCharts                     int    `json:"numCharts"`
	NumFavorites                  int    `json:"numFavorites"`
	SystemOwned                   bool   `json:"systemOwned"`

	// Parameters sets variables (TODO: for what?)
	Parameters Parameter `json:"parameters"`
	
	// ParameterDetails sets variables that can be used within queries
	ParameterDetails map[string]ParameterDetail `json:"parameterDetails"`
}

// Parameter represents a parameter in dashboard (TODO: why?)
type Parameter struct {}

// ParameterDetail represents a parameter to dashboard that can be consumed in queries
type ParameterDetail struct {
	// Label represents the name of the variable
	Label string `json:"label"`

	// DefaultValue maps to keys in the map ValuesToReadableStrings
	DefaultValue string `json:"defaultValue"`

	// HideFromView Whether to hide from the view of the user viewing the Dashboard
	HideFromView bool `json:"hideFromView"`

	// ParameterType (SIMPLE, LIST, DYNAMIC)
	ParameterType string `json:"parameterType"`

	// ValuesToReadableStrings
	ValuesToReadableStrings map[string]string `json:"valuesToReadableStrings"`

	// QueryValue
	QueryValue string `json:"queryValue,omitempty"`

	// TagKey Only required for a DynamicFieldType of TAG_KEY
	TagKey string `json:"tagKey,omitempty"`

	// DynamicFieldType (TAG_KEY, MATCHING_SOURCE_TAG, SOURCE_TAG, SOURCE, METRIC_NAME) Only required for a Parameter type of Dynamic.
	DynamicFieldType string `json:"dynamicFieldType,omitempty"`
}

// Section Represents a Single section within a Dashboard
type Section struct {
	// Name is the name given to this section
	Name string `json:"name"`

	// Rows is an array of Rows in this section
	Rows []Row `json:"rows"`
}

// Row represents a single Row withing a Section of a Wavefront Dashboard
type Row struct {
	// Name represents the display name of the Row
	Name string `json:"name"`

	// HeightFactor  represents the height of the Row
	HeightFactor int `json:"heightFactor"`

	// Charts is an array of Chart that this row contains
	Charts []Chart `json:"charts"`
}

// Chart represents a single Chart, on a single Row with in Section of a Wavefront Dashboard
type Chart struct {
	// Name is the name of a chart
	Name string `json:"name"`

	// Description is the description of the chart
	Description string `json:"description"`

	// More settings (TODO: decriptions)
	Base                   int    `json:"base"`
	IncludeObsoleteMetrics bool   `json:"includeObsoleteMetrics"`
	InterpolatePoints      bool   `json:"interpolatePoints"`
	NoDefaultEvents        bool   `json:"noDefaultEvents"`
	Summarization          string `json:"summarization"`

	// Sources is an Array of Source
	Sources []Source `json:"sources"`

	// Units are the units to use for the y axis
	Units string `json:"units,omitempty"`

	// ChartSettings are custom settings for the chart
	ChartSettings ChartSetting `json:"chartSettings"`
}

// Source represents a single Source for a Chart
type Source struct {
	// Name is the name given to the source
	Name string `json:"name"`

	// Query is a wavefront Query
	Query string `json:"query"`

	// Disabled indicated whether the source is disabled from being rendered on the chart
	Disabled bool `json:"disabled,omitempty"`

	// ScatterPlotSource
	ScatterPlotSource string `json:"scatterPlotSource"`

	// QuerybuilderEnabled
	QuerybuilderEnabled bool `json:"querybuilderEnabled"`

	// SourceDescription
	SourceDescription string `json:"sourceDescription"`
}

// ChartSetting represents settings for a Chart
type ChartSetting struct {
	AutoColumnTags                     bool     `json:"autoColumnTags,omitempty"`
	ColumnTags                         string   `json:"columnTags,omitempty"`
	CustomTags                         []string `json:"customTags,omitempty"`
	ExpectedDataSpacing                int      `json:"expectedDataSpacing,omitempty"`
	FixedLegendDisplayStats            []string `json:"fixedLegendDisplayStats,omitempty"`
	FixedLegendEnabled                 bool     `json:"fixedLegendEnabled,omitempty"`
	FixedLegendFilterField             string   `json:"fixedLegendFilterField,omitempty"`
	FixedLegendFilterLimit             int      `json:"fixedLegendFilterLimit,omitempty"`
	FixedLegendFilterSort              string   `json:"fixedLegendFilterSort,omitempty"`
	FixedLegendHideLabel               bool     `json:"fixedLegendHideLabel,omitempty"`
	FixedLegendPosition                string   `json:"fixedLegendPosition,omitempty"`
	FixedLegendUseRawStats             bool     `json:"fixedLegendUseRawStats,omitempty"`
	GroupBySource                      bool     `json:"groupBySource,omitempty"`
	InvertDynamicLegendHoverControl    bool     `json:"invertDynamicLegendHoverControl,omitempty"`
	LineType                           string   `json:"lineType,omitempty"`
	Max                                float32  `json:"max,omitempty"`
	Min                                float32  `json:"min,omitempty"`
	NumTags                            int      `json:"numTags,omitempty"`
	PlainMarkdownContent               string   `json:"plainMarkdownContent,omitempty"`
	ShowHosts                          bool     `json:"showHosts,omitempty"`
	ShowLabels                         bool     `json:"showLabels,omitempty"`
	ShowRawValues                      bool     `json:"showRawValues,omitempty"`
	SortValuesDescending               bool     `json:"sortValuesDescending,omitempty"`
	SparklineDecimalPrecision          int      `json:"sparklineDecimalPrecision,omitempty"`
	SparklineDisplayColor              string   `json:"sparklineDisplayColor,omitempty"`
	SparklineDisplayFontSize           string   `json:"sparklineDisplayFontSize,omitempty"`
	SparklineDisplayHorizontalPosition string   `json:"sparklineDisplayHorizontalPosition,omitempty"`
	SparklineDisplayPostfix            string   `json:"sparklineDisplayPostfix,omitempty"`
	SparklineDisplayPrefix             string   `json:"sparklineDisplayPrefix,omitempty"`
	SparklineDisplayValueType          string   `json:"sparklineDisplayValueType,omitempty"`
	SparklineDisplayVerticalPosition   string   `json:"sparklineDisplayVerticalPosition,omitempty"`
	SparklineFillColor                 string   `json:"sparklineFillColor,omitempty"`
	SparklineLineColor                 string   `json:"sparklineLineColor,omitempty"`
	SparklineSize                      string   `json:"sparklineSize,omitempty"`
	SparklineValueColorMapApplyTo      string   `json:"sparklineValueColorMapApplyTo,omitempty"`
	SparklineValueColorMapColors       []string `json:"sparklineValueColorMapColors,omitempty,omitempty"`
	SparklineValueColorMapValues       []int    `json:"sparklineValueColorMapValues,omitempty,omitempty"`
	SparklineValueColorMapValuesV2     []int    `json:"sparklineValueColorMapValuesV2,omitempty,omitempty"`
	SparklineValueTextMapText          []string `json:"sparklineValueTextMapText,omitempty,omitempty"`
	SparklineValueTextMapThresholds    []int    `json:"sparklineValueTextMapThresholds,omitempty,omitempty"`
	StackType                          string   `json:"stackType,omitempty"`
	TagMode                            string   `json:"tagMode,omitempty"`
	TimeBasedColoring                  bool     `json:"timeBasedColoring,omitempty"`
	Type                               string   `json:"type,omitempty"`
	Windowing                          string   `json:"windowing,omitempty"`
	WindowSize                         int      `json:"windowSize,omitempty"`
	Xmax                               float32  `json:"xmax,omitempty"`
	Xmin                               float32  `json:"xmin,omitempty"`
	Y0ScaleSIBy1024                    bool     `json:"y0ScaleSIBy1024,omitempty"`
	Y0UnitAutoscaling                  bool     `json:"y0UnitAutoscaling,omitempty"`
	Y1Max                              float32  `json:"y1Max,omitempty"`
	Y1Min                              float32  `json:"y1Min,omitempty"`
	Y1ScaleSIBy1024                    bool     `json:"y1ScaleSIBy1024,omitempty"`
	Y1UnitAutoscaling                  bool     `json:"y1UnitAutoscaling,omitempty"`
	Y1Units                            string   `json:"y1Units,omitempty"`
	Ymax                               float32  `json:"ymax,omitempty"`
	Ymin                               float32  `json:"ymin,omitempty"`
}

// Dashboards is used to perform Dashboard-related operations against the Wavefront API
type Dashboards struct {
	// client is the Wavefront client used to perform Dashboard-related operations
	client Wavefronter
}

const baseDashboardPath = "/api/v2/dashboard"

// UnmarshalJSON is a custom JSON unmarshaller for an Dashboard, used in order to
// populate the Tags field in a more intuitive fashion
func (a *Dashboard) UnmarshalJSON(b []byte) error {
	type dashboard Dashboard
	temp := struct {
		Tags map[string][]string `json:"tags,omitempty"`
		*dashboard
	}{
		dashboard: (*dashboard)(a),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	a.Tags = temp.Tags["customerTags"]
	return nil
}

func (a *Dashboard) MarshalJSON() ([]byte, error) {
	type dashboard Dashboard
	return json.Marshal(&struct {
		Tags map[string][]string `json:"tags,omitempty"`
		*dashboard
	}{
		Tags: map[string][]string{
			"customerTags": a.Tags,
		},
		dashboard: (*dashboard)(a),
	})
}

// Dashboards is used to return a client for Dashboard-related operations
func (c *Client) Dashboards() *Dashboards {
	return &Dashboards{client: c}
}

// Find returns all Dashboards filtered by the given search conditions.
// If filter is nil, all Dashboards are returned.
func (a Dashboards) Find(filter []*SearchCondition) ([]*Dashboard, error) {
	search := &Search{
		client: a.client,
		Type:   "dashboard",
		Params: &SearchParams{
			Conditions: filter,
		},
	}

	var results []*Dashboard
	moreItems := true
	for moreItems == true {
		resp, err := search.Execute()
		if err != nil {
			return nil, err
		}
		var tmpres []*Dashboard
		err = json.Unmarshal(resp.Response.Items, &tmpres)
		if err != nil {
			return nil, err
		}
		results = append(results, tmpres...)
		moreItems = resp.Response.MoreItems
		search.Params.Offset = resp.NextOffset
	}

	return results, nil
}

// Create is used to create an Dashboard in Wavefront.
// If successful, the ID field of the Dashboard will be populated.
func (a Dashboards) Create(dashboard *Dashboard) error {
	return a.crudDashboard("POST", baseDashboardPath, dashboard)
}

// Update is used to update an existing Dashboard.
// The ID field of the Dashboard must be populated
func (a Dashboards) Update(dashboard *Dashboard) error {
	if dashboard.ID == "" {
		return fmt.Errorf("Dashboard id field not set")
	}

	return a.crudDashboard("PUT", fmt.Sprintf("%s/%s", baseDashboardPath, dashboard.ID), dashboard)

}

// Get is used to retrieve an existing Dashboard by ID.
// The ID field must be provided
func (a Dashboards) Get(dashboard *Dashboard) error {
	if dashboard.ID == "" {
		return fmt.Errorf("Dashboard id field is not set")
	}

	return a.crudDashboard("GET", fmt.Sprintf("%s/%s", baseDashboardPath, dashboard.ID), dashboard)
}

// Delete is used to delete an existing Dashboard.
// The ID field of the Dashboard must be populated
func (a Dashboards) Delete(dashboard *Dashboard) error {
	if dashboard.ID == "" {
		return fmt.Errorf("Dashboard id field not set")
	}

	err := a.crudDashboard("DELETE", fmt.Sprintf("%s/%s", baseDashboardPath, dashboard.ID), dashboard)
	if err != nil {
		return err
	}

	//reset the ID field so deletion is not attempted again
	dashboard.ID = ""
	return nil

}

func (a Dashboards) crudDashboard(method, path string, dashboard *Dashboard) error {
	payload, err := json.Marshal(dashboard)
	if err != nil {
		return err
	}
	req, err := a.client.NewRequest(method, path, nil, payload)
	if err != nil {
		return err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Close()

	body, err := ioutil.ReadAll(resp)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &struct {
		Response *Dashboard `json:"response"`
	}{
		Response: dashboard,
	})

}
