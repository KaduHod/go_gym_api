package core

import (
	"net/http"
	"time"
)
type HttpDefaultHandler func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request)
type CsrfTokenInfo struct {
	Token      string
	Expiration time.Time
}
type UserSessionData struct {
    Login string `json:"login"`
    AccessToken string `json:"access_token"`
    CsrfToken CsrfTokenInfo `json:"csrf_token"`
}
type UserAPIToken struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	TokenName string     `json:"token_name"`
	TokenHash string     `json:"token_hash"` // Nunca deve ser exposto
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
type ApiUser struct {
	Login     string  `json:"login"`
	Id        int64   `json:"id"`
	AvatarURL string  `json:"avatar_url"`
	URL       string  `json:"url"`
	Tipo      string  `json:"type"`
	Nome      string  `json:"name"`
	Empresa   string  `json:"company"`
	Blog      string  `json:"blog"`
	Localizacao string `json:"location"`
	Email     *string `json:"email"` // Pode ser null
}
type MuscleMovementInfo struct {
	MuscleGroupName   string `json:"muscle_group_name,omitempty"`
	MusclePortionName string `json:"muscle_portion_name,omitempty"`
	MovementName      string `json:"movement_name,omitempty"`
	JointName         string `json:"joint_name,omitempty"`
}
type Movement struct {
    Id int `json:"id,omitempty"`
    Name string `json:"name"`
}
type MuscleGroup struct {
    Id *int `json:"id,omitempty"`
    Name string `json:"name"`
    Portions []MusclePortion `json:"portions,omitempty"`
}
type MusclePortion struct {
    Id *int `json:"id,omitempty"`
    Name string `json:"name"`
    MuscleGroupId *int `json:"muscle_group_id,omitempty"`
}
type Joint struct {
    Id int `json:"id,omitempty"`
    Name string `json:"name"`
}
