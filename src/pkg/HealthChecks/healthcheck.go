package HealthChecks

import (
	"strings"
	"time"
)

//healthCheckStatus Generic status enum for health checks
type healthCheckStatus int

const (
	Healthy healthCheckStatus = iota
	Degraded
	Unhealthy
)

//String Returns the string equivalent of the enum.
func (hcs healthCheckStatus) String() string {
	return [...]string{"Healthy", "Degraded", "Unhealthy"}[hcs]
}

//ClusterItem All clusters covered by health check should be added as HealthCheckResultItem
type ClusterItem struct {
	Name     string
	Status   healthCheckStatus
	Time     time.Time
	Duration time.Duration
}

//HealthCheckResult Expected return type for Health Checks.
type HealthCheckResult struct {
	Status   healthCheckStatus
	Clusters []ClusterItem
	Time     time.Time
	Duration time.Duration
}

//healthCheck Keeps all features of Health Checks
type healthCheck struct {
	Name        string
	Tags        []string
	Description string
	Critical    bool
	LastStatus  *HealthCheckResult
	ValidIn     time.Duration
	Checker     func() *HealthCheckResult
}

//ShouldRenewNow Determines whether the health check result should be updated.
func (hc *healthCheck) ShouldRenewNow() bool {
	return hc.LastStatus.Time.Add(hc.ValidIn).Before(time.Now())
}

//ServiceHealth The general structure that holds the system health checks.
type ServiceHealth struct {
	HealthChecks map[string]healthCheck
}

//NewServiceHealth Creates a new instance of service health.
func NewServiceHealth() *ServiceHealth {
	return &ServiceHealth{HealthChecks: map[string]healthCheck{}}
}

//AddHealthCheck Adds a health check. Tags must be seperated with comma. Critical holds the info that this health check's
//general effect to the system health. Setting the "Critical" value to true causes the system's general response to be unhealthy for unhealthy health check result.
//Setting the "Critical" value to false causes the overall response of the system to be degraded for unhealthy health check result.
//The value "ValidUntil" indicates how long the result from the health check will be considered valid without being refreshed.
//CheckerFunc can be anything with a function named "Check" and returns *HealthCheckResult
//Name must be unique.
func (s *ServiceHealth) AddHealthCheck(Name, Tags, Desc string, critical bool, ValidIn time.Duration, CheckerFunc func() *HealthCheckResult) {
	//hc := NewHealthCheck(Name, Tags, Desc, ValidIn)
	s.HealthChecks[Name] = healthCheck{
		Name:        Name,
		Tags:        strings.Split(Tags, ","),
		Description: Desc,
		LastStatus: &HealthCheckResult{
			Status:   Unhealthy,
			Time:     time.Time{},
			Duration: 0,
		},
		Critical: critical,
		ValidIn:  ValidIn,
		Checker:  CheckerFunc,
	}
}

/*
//Checker All health checks must be a Checker, which means having a Check method with *HealthCheckResult response.
type Checker interface {
	Check() *HealthCheckResult
}*/

//HardUpdateHealthChecks Ignores "ValidUntil" and re-executes all the health checks.
func (s *ServiceHealth) HardUpdateHealthChecks() {
	for key, value := range s.HealthChecks {
		value.LastStatus = value.Checker() // lock
		s.HealthChecks[key] = value
	}
}

//SoftUpdateHealthChecks Updates all health checks if Valid time is up.
func (s *ServiceHealth) SoftUpdateHealthChecks() {
	for key, value := range s.HealthChecks {
		if value.ShouldRenewNow() {
			value.LastStatus = value.Checker() // lock
			s.HealthChecks[key] = value
		}
	}
}

//GetHealthCheckSummary Returns the health check result for external usage. No need for the pre-using SoftUpdateHealthChecks method.
//GetHealthCheckSummary calls SoftUpdateHealthChecks internally. GetHealthCheckSummary returns only a string: "Healthy", "Unhealthy"
//, or "Degraded". If one or more critical service are "Unhealthy" the response will be unhealthy. Uncritical "unhealthy" health checks
//and "degraded" health checks cause the system response to be "degraded". Response will be "healthy" if all the health checks
//are "healthy".
func (s *ServiceHealth) GetHealthCheckSummary() string {
	s.SoftUpdateHealthChecks()
	degraded := false
	for _, value := range s.HealthChecks {
		if value.LastStatus.Status == Unhealthy {
			if value.Critical {
				return Unhealthy.String()
			}
			degraded = true
		}
		if value.LastStatus.Status == Degraded {
			degraded = true
		}
	}
	if degraded {
		return Degraded.String()
	}
	return Healthy.String()
}

type clusterHealthCheckResultItem struct {
	Name                string
	Status              string
	LastCheckTime       time.Time
	LastCheckDurationMS int64
}
type healthCheckResultItem struct {
	ServiceName         string
	Status              string
	Tags                []string
	Description         string
	LastCheckTime       time.Time
	LastCheckDurationMS int64
	IsCritical          bool
	Clusters            *[]clusterHealthCheckResultItem
}

type DetailedHealthCheckResult struct {
	Summary      string
	HealthChecks []healthCheckResultItem
}

//GetHealthCheckResult Returns the health check result for external usage. No need for the pre-using SoftUpdateHealthChecks method.
//GetHealthCheckResult calls SoftUpdateHealthChecks internally.
//If one or more critical service are "Unhealthy" the response will be unhealthy. Uncritical "unhealthy" health checks
//and "degraded" health checks cause the system response to be "degraded". Response will be "healthy" if all the health checks
//are "healthy"
func (s *ServiceHealth) GetHealthCheckResult() DetailedHealthCheckResult {
	result := DetailedHealthCheckResult{
		Summary:      s.GetHealthCheckSummary(), //FIXME: health check LastCheckTime goroutine RACE
		HealthChecks: []healthCheckResultItem{},
	}
	for _, value := range s.HealthChecks {
		var clusters []clusterHealthCheckResultItem
		for _, cluster := range value.LastStatus.Clusters {
			clusters = append(clusters, clusterHealthCheckResultItem{
				Name:                cluster.Name,
				Status:              cluster.Status.String(),
				LastCheckTime:       cluster.Time,
				LastCheckDurationMS: cluster.Duration.Milliseconds(),
			})
		}
		result.HealthChecks = append(result.HealthChecks, healthCheckResultItem{
			ServiceName:         value.Name,
			Status:              value.LastStatus.Status.String(),
			Tags:                value.Tags,
			Description:         value.Description,
			LastCheckTime:       value.LastStatus.Time,
			LastCheckDurationMS: value.LastStatus.Duration.Milliseconds(),
			IsCritical:          value.Critical,
			Clusters:            &clusters,
		})
	}
	return result
}
