package objects

type Object interface{}

type SpaceObjectResponse struct {
	Count string        `json:"count"`
	Data  []SpaceObject `json:"data"`
}

type SpaceObject struct {
	DesignationObject string `json:"des"`
	Diameter          string `json:"diameter"`
	Name              string `json:"fullname"`
	H                 string `json:"h"`
	Id                string `json:"id"`
	ImpactProbability string `json:"ip"`
	LastObserved      string `json:"last_obs"`
	PotentialImpacts  string `json:"n_imp"`
	PalermoScaleCum   string `json:"ps_cum"`
	PalermoScaleMax   string `json:"ps_max"`
	Range             string `json:"range"`
	TorinoScale       string `json:"ts_max"`
	VelocityInfinity  string `json:"v_inf"`
}

type SummaryData struct {
	DateImpact string `json:"date"`
	Distance   string `json:"dist"`
	Width      string `json:"width"`
}

type SpaceObjectSummaryResponse struct {
	Data    []SummaryData      `json:"data"`
	Summary SpaceObjectSummary `json:"summary"`
}

type SpaceObjectSummary struct {
	SpaceObject
	DesignationObject        string `json:"des"`
	Diameter                 string `json:"diameter"`
	ImpactComputedDate       string `json:"cdate"`
	SpanningDays             string `json:"darc"`
	Energy                   string `json:"energy"`
	FirstObservation         string `json:"first_obs"`
	ImpactEnergy             string `json:"ip"`
	Mass                     string `json:"mass"`
	Method                   string `json:"method"`
	Observations             string `json:"nobs"`
	VelocityImpact           string `json:"v_imp"`
	ImpactEstimatedDate      string `json:"impactEstimatedDate"`
	EstimatedDistanceToEarth string `json:"estimatedDistanceToEarth"`
}

type SpaceObjectDetail struct {
	SummaryData
	SpaceObjectSummary
}
