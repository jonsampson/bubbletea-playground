package domain

import "errors"

var (
	ErrProjectNameIsRequired      = errors.New("project name is required")
	ErrTeamNameIsRequired         = errors.New("team name is required")
	ErrChosenComponentsIsRequired = errors.New("at least one component must be chosen")
)

type Component string

const (
	CLI          Component = "CLI"
	OracleDB     Component = "OracleDB"
	MongoDB      Component = "MongoDB"
	NATSConsumer Component = "NATSConsumer"
	NATSProducer Component = "NATSProducer"
	SqliteDB     Component = "SqliteDB"
	Web          Component = "Web"
)

func ComponentFromString(s string) Component {
	switch s {
	case "CLI":
		return CLI
	case "OracleDB":
		return OracleDB
	case "MongoDB":
		return MongoDB
	case "NATSConsumer":
		return NATSConsumer
	case "NATSProducer":
		return NATSProducer
	case "SqliteDB":
		return SqliteDB
	case "Web":
		return Web
	}

	return ""
}

// BubbleTeaPlayground is a struct that represents the model of bubble tea playground
type BubbleTeaPlayground struct {
	ProjectName      string
	TeamName         string
	ChosenComponents map[Component]struct{}
	ComponentOptions []Component
}

func NewBubbleTeaPlayground(projectName, teamName string, chosenComponents map[Component]struct{}) *BubbleTeaPlayground {
	return &BubbleTeaPlayground{
		ProjectName:      projectName,
		TeamName:         teamName,
		ChosenComponents: chosenComponents,
	}
}

func (b BubbleTeaPlayground) Validate() error {
	if b.ProjectName == "" {
		return ErrProjectNameIsRequired
	}

	if b.TeamName == "" {
		return ErrTeamNameIsRequired
	}

	if len(b.ChosenComponents) == 0 {
		return ErrChosenComponentsIsRequired
	}

	return nil
}

type ValidBubbleTeaPlayground struct {
	btp     BubbleTeaPlayground
	isValid bool
}

func NewValidBubbleTeaPlayground(btp BubbleTeaPlayground) (*ValidBubbleTeaPlayground, error) {
	if err := btp.Validate(); err != nil {
		return nil, err
	}

	return &ValidBubbleTeaPlayground{
		btp:     btp,
		isValid: true,
	}, nil
}
