package api_version

type ApiVersion struct {
	AppName    *string `json:"app_name"`
	AppVersion *string `json:"app_version"`
}

type ApiVersionResolver struct {
	ApiVersion
}

func (vr *ApiVersionResolver) AppName() *string {
	return vr.ApiVersion.AppName
}

func (vr *ApiVersionResolver) AppVersion() *string {
	return vr.ApiVersion.AppVersion
}
