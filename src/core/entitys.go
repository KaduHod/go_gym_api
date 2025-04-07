package core
type GitHubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
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
