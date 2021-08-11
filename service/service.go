package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/davidfragalaureano/sentry/errors"
	"github.com/davidfragalaureano/sentry/httputils"
	objects "github.com/davidfragalaureano/sentry/space-object"
)

var (
	httpClient *http.Client
)

type sentryService struct {
	ServerURL string
}

type SentryService interface {
	GetAllObjects() (object *objects.SpaceObjectResponse, err error)
	GetObjectByName(objectName string) (object *objects.SpaceObjectDetail, err error)
}

func NewSentryService(client *http.Client, serverURL string) SentryService {
	httpClient = client
	return &sentryService{serverURL}
}

func (service *sentryService) GetAllObjects() (object *objects.SpaceObjectResponse, err error) {
	var obj objects.SpaceObjectResponse
	req, err := httputils.Get(service.ServerURL + "?www=1")
	req.Header.Add("Content-Type", `application/json`)

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, errors.NewUnexpectedError(err, "Unable to get all objects")
	}

	err = jsonUnmarshall(resp, &obj)

	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (service *sentryService) GetObjectByName(objectName string) (object *objects.SpaceObjectDetail, err error) {
	var obj objects.SpaceObjectSummaryResponse

	req, err := httputils.Get(fmt.Sprintf("%s?des=%s&www=1", service.ServerURL, url.QueryEscape(objectName)))

	if err != nil {
		return nil, errors.NewUnexpectedError(err, "Unable to get object")
	}

	req.Header.Add("Content-Type", `application/json`)

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, errors.NewUnexpectedError(err, "Unable to get object")
	}

	err = jsonUnmarshall(resp, &obj)

	if err != nil {
		return nil, err
	}

	objectDetail := &objects.SpaceObjectDetail{
		objects.SummaryData{},
		obj.Summary,
	}

	if data := obj.Data; len(data) > 0 {
		objectDetail.DateImpact = data[0].DateImpact
		objectDetail.Distance = data[0].Distance
		objectDetail.Width = data[0].Width
	}

	distance, timeImpact := calculateTotalDistanceToEarth(objectDetail.DateImpact, objectDetail.VelocityImpact, objectDetail.VelocityInfinity)

	objectDetail.EstimatedDistanceToEarth = distance
	objectDetail.ImpactEstimatedDate = timeImpact

	return objectDetail, nil
}

func (service *sentryService) GetSpaceObjetsDetailed() (object *objects.SpaceObjectResponse, err error) {
	var objectsResp objects.SpaceObjectResponse
	req, err := httputils.Get(service.ServerURL + "?www=1")
	req.Header.Add("Content-Type", `application/json`)

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, errors.NewUnexpectedError(err, "Unable to get all objects")
	}

	err = jsonUnmarshall(resp, &objectsResp)

	if err != nil {
		return nil, err
	}

	return &objectsResp, nil
}

func calculateTotalDistanceToEarth(dateImpact string, velocityImpact string, velocityInfinity string) (distance string, timeImpact string) {
	log.Print(strings.Split(dateImpact, ".")[0])

	dateImpactParsed, err := time.Parse(DATE_LAYOUT, strings.Split(dateImpact, ".")[0])

	if err != nil {
		log.Fatal("Error while parsing date :", err)
	}

	velocityImpactParsed, err := strconv.ParseFloat(velocityImpact, 64)
	if err != nil {
		log.Fatal("Error converting velocityImpact to float", err)
	}

	velocityInfinityParsed, err := strconv.ParseFloat(velocityInfinity, 64)
	if err != nil {
		log.Fatal("Error converting velocityInfinity to float", err)
	}

	atmosphereImpactTime := DISTANCE_FROM_EARTH_TO_ATMOSPHERE / velocityImpactParsed
	distanceInfinity := velocityInfinityParsed * float64(dateImpactParsed.Unix())

	return fmt.Sprintf("%f", distanceInfinity+DISTANCE_FROM_EARTH_TO_ATMOSPHERE), fmt.Sprintf("%d", int64(float64(dateImpactParsed.Unix())+math.Round(atmosphereImpactTime)))
}

func jsonUnmarshall(resp *http.Response, objType interface{}) (reqErr error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return errors.NewIOReadError(err, "Error reading stream")
	}

	defer resp.Body.Close()

	log.Printf(string(body))

	jsonErr := json.Unmarshal(body, objType)
	if jsonErr != nil {
		log.Print(objType)
	}

	return nil
}
