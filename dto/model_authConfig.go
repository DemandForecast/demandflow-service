package dto

type AuthConfig struct {
	AUTH0_DOMAIN   string
	AUTH0_AUDIENCE string
}

type UserCacheData struct {
	Sub      string
	Role     string
	OrganizationId string
}