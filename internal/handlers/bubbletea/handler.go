package bubbletea

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jonsampson/bubbletea-playground/internal/domain"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type projectCreator interface {
	CreateProject(validBubbleTeaPlayground domain.ValidBubbleTeaPlayground) error
}

type keymap = struct {
	accept key.Binding
	next   key.Binding
	prev   key.Binding
	toggle key.Binding
	quit   key.Binding
}

var DefaultKeymap = keymap{
	accept: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("accept", "accept input"),
	),
	next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "move to next input"),
	),
	prev: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "move to previous input"),
	),
	quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "exit the program"),
	),
	toggle: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("toggle", "toggle component"),
	),
}

type teaModel struct {
	btp              domain.BubbleTeaPlayground
	componentList    list.Model
	focusedIndex     int
	err              error
	keymap           keymap
	projectNameInput textinput.Model
	teamNameInput    textinput.Model
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func itemsForComponents(components []domain.Component) []list.Item {
	items := make([]list.Item, len(components))
	for i, component := range components {
		items[i] = item{title: (string(component)), desc: ""}
	}
	return items
}

func initialModel() teaModel {
	model := teaModel{
		btp: domain.BubbleTeaPlayground{
			ProjectName:      "",
			TeamName:         "",
			ChosenComponents: make(map[domain.Component]struct{}),
			ComponentOptions: []domain.Component{
				domain.CLI,
				domain.OracleDB,
				domain.MongoDB,
				domain.NATSConsumer,
				domain.NATSProducer,
				domain.SqliteDB,
				domain.Web,
			},
		},
		keymap:       DefaultKeymap,
		focusedIndex: 0,
	}

	model.projectNameInput = textinput.New()
	model.projectNameInput.Placeholder = "enter your project name"
	model.projectNameInput.Focus()
	model.projectNameInput.PromptStyle = focusedStyle
	model.projectNameInput.TextStyle = focusedStyle
	model.projectNameInput.CharLimit = 64
	model.projectNameInput.Width = 32

	model.teamNameInput = textinput.New()
	model.teamNameInput.PromptStyle = blurredStyle
	model.teamNameInput.TextStyle = blurredStyle
	model.teamNameInput.Placeholder = "enter your team name"
	model.teamNameInput.CharLimit = 64
	model.teamNameInput.Width = 32

	model.componentList = list.New(itemsForComponents(model.btp.ComponentOptions), list.NewDefaultDelegate(), 0, 0)
	model.componentList.Title = "Choose components"
	model.componentList.Styles.Title = blurredStyle
	model.componentList.KeyMap.ForceQuit = DefaultKeymap.quit
	model.componentList.KeyMap.Quit = DefaultKeymap.quit
	// model.componentList.SetShowHelp(false)

	return model
}

func (m teaModel) Init() tea.Cmd {
	return nil
}

func (m teaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.next):
			m.focusedIndex++
			if m.focusedIndex > 2 {
				m.focusedIndex = 0
			}
		case key.Matches(msg, m.keymap.prev):
			m.focusedIndex--
			if m.focusedIndex < 0 {
				m.focusedIndex = 2
			}
		case key.Matches(msg, m.keymap.toggle):
			if m.projectNameInput.Focused() || m.teamNameInput.Focused() {
				return m, nil
			} else {
				selectedItem := m.componentList.SelectedItem().(item)
				component := domain.ComponentFromString(selectedItem.Title())
				_, ok := m.btp.ChosenComponents[component]
				if ok {
					delete(m.btp.ChosenComponents, component)
					selectedItem.desc = ""
					cmds = append(cmds, m.componentList.NewStatusMessage(fmt.Sprintf("%s removed", selectedItem.Title())))
				} else {
					m.btp.ChosenComponents[component] = struct{}{}
					selectedItem.desc = "[X]"
					cmds = append(cmds, m.componentList.NewStatusMessage(fmt.Sprintf("%s added", selectedItem.Title())))
				}
				cmds = append(cmds, m.componentList.SetItem(m.componentList.Index(), selectedItem))
				return m, tea.Batch(cmds...)
			}
		case key.Matches(msg, m.keymap.accept):
			m.btp.ProjectName = m.projectNameInput.Value()
		}
	case error:
		m.err = msg
		return m, nil
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.componentList.SetSize(msg.Width-v, msg.Height-h-8)

	}

	if m.focusedIndex == 0 {
		m.projectNameInput.Focus()
		m.projectNameInput.PromptStyle = focusedStyle
		m.projectNameInput.TextStyle = focusedStyle
		m.teamNameInput.Blur()
		m.teamNameInput.PromptStyle = blurredStyle
		m.teamNameInput.TextStyle = blurredStyle
		m.componentList.Styles.Title = blurredStyle
		cmds = append(cmds, m.projectNameInput.Cursor.BlinkCmd())
		m.projectNameInput, cmd = m.projectNameInput.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.focusedIndex == 1 {
		m.teamNameInput.Focus()
		m.teamNameInput.PromptStyle = focusedStyle
		m.teamNameInput.TextStyle = focusedStyle
		m.projectNameInput.Blur()
		m.projectNameInput.PromptStyle = blurredStyle
		m.projectNameInput.TextStyle = blurredStyle
		m.componentList.Styles.Title = blurredStyle
		cmds = append(cmds, m.teamNameInput.Cursor.BlinkCmd())
		m.teamNameInput, cmd = m.teamNameInput.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.focusedIndex == 2 {
		m.projectNameInput.Blur()
		m.teamNameInput.Blur()
		m.componentList.Styles.Title = focusedStyle
		m.projectNameInput.PromptStyle = blurredStyle
		m.projectNameInput.TextStyle = blurredStyle
		m.teamNameInput.PromptStyle = blurredStyle
		m.teamNameInput.TextStyle = blurredStyle

		m.componentList, cmd = m.componentList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m teaModel) View() string {
	return appStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top,
			lipgloss.JoinHorizontal(lipgloss.Center,
				fmt.Sprintf("\n\n%s\n\n", m.projectNameInput.View()),
				fmt.Sprintf("\n\n%s\n\n", m.teamNameInput.View()),
			),
			fmt.Sprintf("%s\n\n", m.componentList.View()),
		),
	)
}

func Run(projectCreator projectCreator) {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
