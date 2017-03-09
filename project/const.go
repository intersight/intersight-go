package project

import "github.com/intersight/intersight-go"

const (
	ModeDevelopment      intersight.ProjectMode = "dev"
	ModeTesting          intersight.ProjectMode = "test"
	ModeQualityAssurance intersight.ProjectMode = "qa"
	ModeStaging          intersight.ProjectMode = "stage"
	ModeProduction       intersight.ProjectMode = "prod"

	RegionUS   intersight.ProjectRegion = "us"
	RegionEU   intersight.ProjectRegion = "eu"
	RegionAsia intersight.ProjectRegion = "ap"
)
