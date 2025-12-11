package main

type ManifestConfig struct {
	ProjectName     string
	ImageRepository string
	DBImageName     string
	DNS             string
	VolumeHandler   string
	DBTagName       string
	DBPort          string
	APPport         string
	TagNameDev      string
	TagNameProd     string
}
type RutasConfig struct {
	PathDev          string
	PathProd         string
	PathSitesDev     string
	PathSitesProd    string
	PathBase         string
	PathBaseCert     string
	PathBaseConfig   string
	PathBasePvPvc    string
	PathBaseDatabase string
	PathBaseBackend  string
	PathBaseIngress  string
}
