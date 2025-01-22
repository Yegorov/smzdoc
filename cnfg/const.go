package cnfg

const (
	ProjectName string = "smzdoc"
	Version     string = "0.0.1"
)

func WelcomeMsg() string {
	return ProjectName + ", version: " + Version
}
