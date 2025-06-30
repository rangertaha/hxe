/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 Rangertaha <rangertaha@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	duration "github.com/golang/protobuf/ptypes/duration"
	"github.com/rangertaha/hxe/internal/api/client"
	models "github.com/rangertaha/hxe/internal/api/models"
)

// Model represents the main application state
type Model struct {
	client       *client.Client
	table        table.Model
	help         help.Model
	timer        timer.Model
	services     []*models.Service
	selectedID   uint32
	loading      bool
	err          error
	width        int
	height       int
	mockMode     bool            // Add mock mode flag
	showDetails  bool            // Add side panel flag
	showPopup    bool            // Add popup flag
	popupService *models.Service // Service for popup
}

// KeyMap defines the key bindings
type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Start   key.Binding
	Stop    key.Binding
	Restart key.Binding
	Refresh key.Binding
	Quit    key.Binding
	Help    key.Binding
	Select  key.Binding
	Details key.Binding
	Status  key.Binding
}

// ShortHelp returns the short help text
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Start, k.Stop, k.Restart, k.Refresh, k.Quit}
}

// FullHelp returns the full help text
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Select, k.Details},
		{k.Start, k.Stop, k.Restart, k.Refresh},
		{k.Help, k.Quit},
	}
}

// DefaultKeyMap returns the default key bindings
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Start: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start"),
		),
		Stop: key.NewBinding(
			key.WithKeys("x"),
			key.WithHelp("x", "stop"),
		),
		Restart: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "restart"),
		),
		Refresh: key.NewBinding(
			key.WithKeys("R"),
			key.WithHelp("R", "refresh"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Details: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "details"),
		),
		Status: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "status"),
		),
	}
}

// NewModel creates a new application model
func NewModel(c *client.Client, mockMode bool) Model {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Name", Width: 20},
		{Title: "Status", Width: 12},
		{Title: "PID", Width: 6},
		{Title: "Uptime", Width: 12},
		{Title: "CPU%", Width: 6},
		{Title: "Memory%", Width: 8},
		{Title: "Description", Width: 30},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("235"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("11")).
		Bold(true)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("250")).
		Background(lipgloss.Color("0"))
	t.SetStyles(s)

	return Model{
		client:       c,
		table:        t,
		help:         help.New(),
		timer:        timer.NewWithInterval(5*time.Second, time.Second),
		services:     []*models.Service{},
		loading:      true,
		mockMode:     mockMode,
		showDetails:  false,
		showPopup:    false,
		popupService: nil,
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.timer.Init(),
		m.loadServices(),
	)
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap().Quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap().Refresh):
			return m, m.loadServices()
		case key.Matches(msg, DefaultKeyMap().Start):
			if m.selectedID > 0 {
				return m, m.startService(m.selectedID)
			}
		case key.Matches(msg, DefaultKeyMap().Stop):
			if m.selectedID > 0 {
				return m, m.stopService(m.selectedID)
			}
		case key.Matches(msg, DefaultKeyMap().Restart):
			if m.selectedID > 0 {
				return m, m.restartService(m.selectedID)
			}
		case key.Matches(msg, DefaultKeyMap().Select):
			if len(m.table.Rows()) > 0 {
				row := m.table.SelectedRow()
				if len(row) > 0 {
					if id, err := strconv.ParseUint(row[0], 10, 32); err == nil {
						m.selectedID = uint32(id)
						m.showDetails = !m.showDetails // Toggle side panel
					}
				}
			}
		case key.Matches(msg, DefaultKeyMap().Status):
			if m.selectedID > 0 {
				m.showStatusPopup()
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("escape"))):
			m.showDetails = false // Close side panel
			m.showPopup = false   // Close popup
		}

	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			// Check if click is on status column (column 2)
			if msg.X >= 24 && msg.X <= 36 { // Status column position
				// Find which row was clicked
				rowIndex := (msg.Y - 4) // Adjust for header and spacing
				if rowIndex >= 0 && rowIndex < len(m.services) {
					m.selectedID = m.services[rowIndex].Id
					m.showStatusPopup()
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Use full height minus space for title, status, and help
		tableHeight := msg.Height - 6 // 6 lines for title, status, spacing, and help
		m.table.SetHeight(tableHeight)

		// Update column widths to use full width
		columns := []table.Column{
			{Title: "ID", Width: 4},
			{Title: "Name", Width: 20},
			{Title: "Status", Width: 12},
			{Title: "PID", Width: 6},
			{Title: "Uptime", Width: 12},
			{Title: "CPU%", Width: 6},
			{Title: "Memory%", Width: 8},
			{Title: "Description", Width: msg.Width - 78}, // Use remaining width
		}
		m.table.SetColumns(columns)

	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, tea.Batch(cmd, m.loadServices())

	case servicesLoadedMsg:
		m.loading = false
		m.services = msg.services
		m.err = msg.err
		m.updateTable()

	case serviceActionMsg:
		m.err = msg.err
		return m, m.loadServices()
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View renders the UI
func (m Model) View() string {
	if m.loading {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
			lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Loading services..."))
	}

	if m.err != nil {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
			lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(fmt.Sprintf("Error: %v", m.err)))
	}

	// Mutt-style header
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("235")).
		Padding(0, 1).
		Render("HXE Service Manager")

	// Mutt-style status bar
	status := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Background(lipgloss.Color("0")).
		Padding(0, 1).
		Render(fmt.Sprintf("Services: %d | Selected: %d | %s",
			len(m.services), m.selectedID, m.timer.View()))

	var mainView string
	if m.showDetails {
		mainView = m.renderWithSidePanel(title, status)
	} else {
		table := m.table.View()
		help := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Background(lipgloss.Color("0")).
			Padding(0, 1).
			Render(m.help.View(DefaultKeyMap()))

		mainView = lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			status,
			"",
			table,
			"",
			help,
		)
	}

	// Show popup if active
	if m.showPopup && m.popupService != nil {
		return m.renderWithPopup(mainView)
	}

	return mainView
}

// Messages for async operations
type servicesLoadedMsg struct {
	services []*models.Service
	err      error
}

type serviceActionMsg struct {
	err error
}

// Commands
func (m Model) loadServices() tea.Cmd {
	return func() tea.Msg {
		if m.mockMode {
			// Return mock data
			return servicesLoadedMsg{services: createMockServices(), err: nil}
		}

		// Use real NATS client
		servicesResp, err := m.client.Service.List()
		if err != nil {
			return servicesLoadedMsg{services: nil, err: err}
		}
		return servicesLoadedMsg{services: servicesResp.Services, err: nil}
	}
}

func (m Model) startService(id uint32) tea.Cmd {
	return func() tea.Msg {
		if m.mockMode {
			// Simulate start action in mock mode
			return serviceActionMsg{err: nil}
		}
		_, err := m.client.Service.Start(uint(id))
		return serviceActionMsg{err: err}
	}
}

func (m Model) stopService(id uint32) tea.Cmd {
	return func() tea.Msg {
		if m.mockMode {
			// Simulate stop action in mock mode
			return serviceActionMsg{err: nil}
		}
		_, err := m.client.Service.Stop(uint(id))
		return serviceActionMsg{err: err}
	}
}

func (m Model) restartService(id uint32) tea.Cmd {
	return func() tea.Msg {
		if m.mockMode {
			// Simulate restart action in mock mode
			return serviceActionMsg{err: nil}
		}
		_, err := m.client.Service.Restart(uint(id))
		return serviceActionMsg{err: err}
	}
}

// updateTable updates the table with current service data
func (m *Model) updateTable() {
	var rows []table.Row
	for _, service := range m.services {
		// Get status with color coding
		status := getStatusString(service.Status)

		pid := strconv.Itoa(int(service.Pid))
		if service.Pid == 0 {
			pid = "-"
		}

		uptime := "-"
		if service.Uptime != nil {
			uptime = formatDuration(service.Uptime.AsDuration())
		}

		// Get CPU and Memory metrics from service data
		cpuPercent := "-"
		memoryPercent := "-"
		if service.Metrics != nil {
			if cpu, ok := service.Metrics["cpu_percent"]; ok {
				cpuPercent = fmt.Sprintf("%.1f", cpu)
			}
			if mem, ok := service.Metrics["memory_percent"]; ok {
				memoryPercent = fmt.Sprintf("%.1f", mem)
			}
		}

		rows = append(rows, table.Row{
			strconv.FormatUint(uint64(service.Id), 10),
			truncateString(service.Name, 18),
			status,
			pid,
			uptime,
			cpuPercent,
			memoryPercent,
			truncateString(service.Description, 28),
		})
	}
	m.table.SetRows(rows)
}

// Helper functions
func getStatusString(status models.ServiceStatus) string {
	switch status {
	case models.ServiceStatus_STATE_READY:
		return "Ready"
	case models.ServiceStatus_STATE_LOADING:
		return "Loading"
	case models.ServiceStatus_STATE_STARTING:
		return "Starting"
	case models.ServiceStatus_STATE_RESTARTING:
		return "Restarting"
	case models.ServiceStatus_STATE_RUNNING:
		return "Running"
	case models.ServiceStatus_STATE_STOPPING:
		return "Stopping"
	case models.ServiceStatus_STATE_STOPPED:
		return "Stopped"
	case models.ServiceStatus_STATE_FAILED:
		return "Failed"
	case models.ServiceStatus_STATE_SUCCESS:
		return "Success"
	default:
		return "Unknown"
	}
}

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}
	return fmt.Sprintf("%dh", int(d.Hours()))
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// createMockServices creates dummy service data for testing
func createMockServices() []*models.Service {
	now := time.Now()

	return []*models.Service{
		{
			Id:          1,
			Name:        "web-server",
			Description: "Web server for API endpoints",
			User:        1000,
			Group:       1000,
			Directory:   "/var/www",
			PreExec:     "mkdir -p /var/log/web",
			CmdExec:     "/usr/bin/python3 app.py",
			PostExec:    "echo 'Web server started'",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1234,
			ExitCode:    0,
			StartTime:   now.Add(-2 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 7200}, // 2 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    2.1,
				"memory_percent": 45.2,
				"disk_usage":     12.5,
			},
			CategoryId: 1,
			AppId:      1,
		},
		{
			Id:          2,
			Name:        "database",
			Description: "PostgreSQL database server",
			User:        999,
			Group:       999,
			Directory:   "/var/lib/postgresql",
			PreExec:     "initdb -D /var/lib/postgresql/data",
			CmdExec:     "/usr/bin/postgres -D /var/lib/postgresql/data",
			PostExec:    "pg_ctl start",
			Autostart:   true,
			Retries:     5,
			Enabled:     true,
			Pid:         1235,
			ExitCode:    0,
			StartTime:   now.Add(-90 * time.Minute).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 5400}, // 1.5 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    1.8,
				"memory_percent": 32.1,
				"disk_usage":     28.7,
			},
			CategoryId: 2,
			AppId:      1,
		},
		{
			Id:          3,
			Name:        "cache",
			Description: "Redis cache server",
			User:        6379,
			Group:       6379,
			Directory:   "/var/lib/redis",
			PreExec:     "mkdir -p /var/lib/redis",
			CmdExec:     "/usr/bin/redis-server /etc/redis/redis.conf",
			PostExec:    "redis-cli ping",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         0,
			ExitCode:    1,
			StartTime:   0,
			EndTime:     now.Add(-10 * time.Minute).Unix(),
			Uptime:      nil,
			Status:      models.ServiceStatus_STATE_STOPPED,
			Metrics: map[string]float64{
				"cpu_percent":    0.0,
				"memory_percent": 0.0,
				"disk_usage":     0.0,
			},
			CategoryId: 2,
			AppId:      1,
		},
		{
			Id:          4,
			Name:        "monitoring",
			Description: "System monitoring agent",
			User:        1001,
			Group:       1001,
			Directory:   "/opt/monitoring",
			PreExec:     "check_dependencies",
			CmdExec:     "/opt/monitoring/agent.py",
			PostExec:    "send_heartbeat",
			Autostart:   true,
			Retries:     10,
			Enabled:     true,
			Pid:         1237,
			ExitCode:    0,
			StartTime:   now.Add(-45 * time.Minute).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 2700}, // 45 minutes
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.5,
				"memory_percent": 12.3,
				"disk_usage":     5.2,
			},
			CategoryId: 3,
			AppId:      2,
		},
		{
			Id:          5,
			Name:        "backup",
			Description: "Automated backup service",
			User:        1002,
			Group:       1002,
			Directory:   "/var/backups",
			PreExec:     "check_disk_space",
			CmdExec:     "/usr/local/bin/backup.sh",
			PostExec:    "notify_completion",
			Autostart:   false,
			Retries:     2,
			Enabled:     true,
			Pid:         0,
			ExitCode:    127,
			StartTime:   now.Add(-30 * time.Minute).Unix(),
			EndTime:     now.Add(-25 * time.Minute).Unix(),
			Uptime:      nil,
			Status:      models.ServiceStatus_STATE_FAILED,
			Metrics: map[string]float64{
				"cpu_percent":    0.0,
				"memory_percent": 0.0,
				"disk_usage":     0.0,
			},
			CategoryId: 4,
			AppId:      3,
		},
		{
			Id:          6,
			Name:        "email-server",
			Description: "SMTP email server",
			User:        1003,
			Group:       1003,
			Directory:   "/var/mail",
			PreExec:     "check_certificates",
			CmdExec:     "/usr/sbin/postfix start",
			PostExec:    "test_smtp",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1238,
			ExitCode:    0,
			StartTime:   now.Add(-6 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 21600}, // 6 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.8,
				"memory_percent": 18.7,
				"disk_usage":     8.9,
			},
			CategoryId: 5,
			AppId:      4,
		},
		{
			Id:          7,
			Name:        "load-balancer",
			Description: "Nginx load balancer",
			User:        1004,
			Group:       1004,
			Directory:   "/etc/nginx",
			PreExec:     "nginx -t",
			CmdExec:     "nginx -g 'daemon off;'",
			PostExec:    "curl -f http://localhost/health",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1239,
			ExitCode:    0,
			StartTime:   now.Add(-1 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 3600}, // 1 hour
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    1.2,
				"memory_percent": 22.4,
				"disk_usage":     3.1,
			},
			CategoryId: 1,
			AppId:      1,
		},
		{
			Id:          8,
			Name:        "queue-worker",
			Description: "Background job processor",
			User:        1005,
			Group:       1005,
			Directory:   "/var/queue",
			PreExec:     "check_redis_connection",
			CmdExec:     "celery -A tasks worker --loglevel=info",
			PostExec:    "monitor_queue_size",
			Autostart:   true,
			Retries:     5,
			Enabled:     true,
			Pid:         0,
			ExitCode:    0,
			StartTime:   0,
			EndTime:     0,
			Uptime:      nil,
			Status:      models.ServiceStatus_STATE_READY,
			Metrics: map[string]float64{
				"cpu_percent":    0.0,
				"memory_percent": 0.0,
				"disk_usage":     0.0,
			},
			CategoryId: 6,
			AppId:      5,
		},
		{
			Id:          9,
			Name:        "file-server",
			Description: "FTP/SFTP file server",
			User:        1006,
			Group:       1006,
			Directory:   "/var/ftp",
			PreExec:     "create_ftp_users",
			CmdExec:     "/usr/sbin/vsftpd /etc/vsftpd.conf",
			PostExec:    "test_ftp_connection",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1240,
			ExitCode:    0,
			StartTime:   now.Add(-3 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 10800}, // 3 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.3,
				"memory_percent": 8.9,
				"disk_usage":     15.3,
			},
			CategoryId: 7,
			AppId:      6,
		},
		{
			Id:          10,
			Name:        "dns-server",
			Description: "DNS resolver and cache",
			User:        1007,
			Group:       1007,
			Directory:   "/etc/bind",
			PreExec:     "check_zone_files",
			CmdExec:     "/usr/sbin/named -u bind -g",
			PostExec:    "dig @localhost test.local",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1241,
			ExitCode:    0,
			StartTime:   now.Add(-12 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 43200}, // 12 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.2,
				"memory_percent": 6.7,
				"disk_usage":     2.1,
			},
			CategoryId: 8,
			AppId:      7,
		},
		{
			Id:          11,
			Name:        "log-aggregator",
			Description: "Centralized log collection",
			User:        1008,
			Group:       1008,
			Directory:   "/var/log/aggregator",
			PreExec:     "create_log_dirs",
			CmdExec:     "/usr/bin/fluentd -c /etc/fluentd/fluent.conf",
			PostExec:    "check_log_flow",
			Autostart:   true,
			Retries:     4,
			Enabled:     true,
			Pid:         1242,
			ExitCode:    0,
			StartTime:   now.Add(-30 * time.Minute).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 1800}, // 30 minutes
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    1.5,
				"memory_percent": 14.2,
				"disk_usage":     25.8,
			},
			CategoryId: 3,
			AppId:      2,
		},
		{
			Id:          12,
			Name:        "search-engine",
			Description: "Elasticsearch search service",
			User:        1009,
			Group:       1009,
			Directory:   "/var/lib/elasticsearch",
			PreExec:     "check_java_version",
			CmdExec:     "/usr/share/elasticsearch/bin/elasticsearch",
			PostExec:    "curl -f http://localhost:9200/_cluster/health",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1243,
			ExitCode:    0,
			StartTime:   now.Add(-4 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 14400}, // 4 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    3.2,
				"memory_percent": 38.9,
				"disk_usage":     42.1,
			},
			CategoryId: 9,
			AppId:      8,
		},
		{
			Id:          13,
			Name:        "message-broker",
			Description: "RabbitMQ message broker",
			User:        1010,
			Group:       1010,
			Directory:   "/var/lib/rabbitmq",
			PreExec:     "check_erlang",
			CmdExec:     "/usr/sbin/rabbitmq-server",
			PostExec:    "rabbitmqctl status",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1244,
			ExitCode:    0,
			StartTime:   now.Add(-5 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 18000}, // 5 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    1.7,
				"memory_percent": 28.3,
				"disk_usage":     18.7,
			},
			CategoryId: 6,
			AppId:      5,
		},
		{
			Id:          14,
			Name:        "webhook-handler",
			Description: "Webhook processing service",
			User:        1011,
			Group:       1011,
			Directory:   "/opt/webhooks",
			PreExec:     "validate_config",
			CmdExec:     "/usr/bin/node webhook-server.js",
			PostExec:    "test_webhook_endpoint",
			Autostart:   true,
			Retries:     5,
			Enabled:     true,
			Pid:         0,
			ExitCode:    0,
			StartTime:   0,
			EndTime:     0,
			Uptime:      nil,
			Status:      models.ServiceStatus_STATE_STARTING,
			Metrics: map[string]float64{
				"cpu_percent":    0.0,
				"memory_percent": 0.0,
				"disk_usage":     0.0,
			},
			CategoryId: 10,
			AppId:      9,
		},
		{
			Id:          15,
			Name:        "ssl-terminator",
			Description: "SSL/TLS termination proxy",
			User:        1012,
			Group:       1012,
			Directory:   "/etc/haproxy",
			PreExec:     "check_certificates",
			CmdExec:     "/usr/sbin/haproxy -f haproxy.cfg",
			PostExec:    "test_ssl_connection",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1245,
			ExitCode:    0,
			StartTime:   now.Add(-8 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 28800}, // 8 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.9,
				"memory_percent": 16.8,
				"disk_usage":     4.2,
			},
			CategoryId: 1,
			AppId:      1,
		},
		{
			Id:          16,
			Name:        "metrics-collector",
			Description: "Prometheus metrics collector",
			User:        1013,
			Group:       1013,
			Directory:   "/opt/prometheus",
			PreExec:     "validate_config",
			CmdExec:     "/opt/prometheus/prometheus --config.file=prometheus.yml",
			PostExec:    "check_targets",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1246,
			ExitCode:    0,
			StartTime:   now.Add(-7 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 25200}, // 7 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    1.1,
				"memory_percent": 24.6,
				"disk_usage":     31.5,
			},
			CategoryId: 3,
			AppId:      2,
		},
		{
			Id:          17,
			Name:        "api-gateway",
			Description: "Kong API gateway",
			User:        1014,
			Group:       1014,
			Directory:   "/etc/kong",
			PreExec:     "kong check",
			CmdExec:     "kong start",
			PostExec:    "kong health",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1247,
			ExitCode:    0,
			StartTime:   now.Add(-2 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 7200}, // 2 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    2.8,
				"memory_percent": 19.3,
				"disk_usage":     7.8,
			},
			CategoryId: 1,
			AppId:      1,
		},
		{
			Id:          18,
			Name:        "job-scheduler",
			Description: "Cron job scheduler",
			User:        1015,
			Group:       1015,
			Directory:   "/var/spool/cron",
			PreExec:     "check_cron_permissions",
			CmdExec:     "/usr/sbin/cron -f",
			PostExec:    "list_cron_jobs",
			Autostart:   true,
			Retries:     2,
			Enabled:     true,
			Pid:         1248,
			ExitCode:    0,
			StartTime:   now.Add(-24 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 86400}, // 24 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.1,
				"memory_percent": 2.4,
				"disk_usage":     1.2,
			},
			CategoryId: 11,
			AppId:      10,
		},
		{
			Id:          19,
			Name:        "config-manager",
			Description: "Configuration management service",
			User:        1016,
			Group:       1016,
			Directory:   "/etc/config-manager",
			PreExec:     "validate_configs",
			CmdExec:     "/usr/bin/config-manager --daemon",
			PostExec:    "reload_configs",
			Autostart:   true,
			Retries:     4,
			Enabled:     true,
			Pid:         0,
			ExitCode:    0,
			StartTime:   0,
			EndTime:     0,
			Uptime:      nil,
			Status:      models.ServiceStatus_STATE_LOADING,
			Metrics: map[string]float64{
				"cpu_percent":    0.0,
				"memory_percent": 0.0,
				"disk_usage":     0.0,
			},
			CategoryId: 12,
			AppId:      11,
		},
		{
			Id:          20,
			Name:        "health-checker",
			Description: "Service health monitoring",
			User:        1017,
			Group:       1017,
			Directory:   "/opt/health-checker",
			PreExec:     "setup_checks",
			CmdExec:     "/usr/bin/health-checker --interval=30s",
			PostExec:    "run_initial_checks",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1249,
			ExitCode:    0,
			StartTime:   now.Add(-15 * time.Minute).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 900}, // 15 minutes
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.4,
				"memory_percent": 5.7,
				"disk_usage":     2.8,
			},
			CategoryId: 3,
			AppId:      2,
		},
		{
			Id:          21,
			Name:        "data-sync",
			Description: "Data synchronization service",
			User:        1018,
			Group:       1018,
			Directory:   "/var/data/sync",
			PreExec:     "check_connections",
			CmdExec:     "/usr/bin/sync-daemon --config=sync.conf",
			PostExec:    "verify_sync_status",
			Autostart:   false,
			Retries:     3,
			Enabled:     true,
			Pid:         0,
			ExitCode:    0,
			StartTime:   0,
			EndTime:     0,
			Uptime:      nil,
			Status:      models.ServiceStatus_STATE_STOPPED,
			Metrics: map[string]float64{
				"cpu_percent":    0.0,
				"memory_percent": 0.0,
				"disk_usage":     0.0,
			},
			CategoryId: 13,
			AppId:      12,
		},
		{
			Id:          22,
			Name:        "notification",
			Description: "Push notification service",
			User:        1019,
			Group:       1019,
			Directory:   "/opt/notifications",
			PreExec:     "check_api_keys",
			CmdExec:     "/usr/bin/notification-server",
			PostExec:    "test_notification",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1250,
			ExitCode:    0,
			StartTime:   now.Add(-1 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 3600}, // 1 hour
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    1.3,
				"memory_percent": 11.2,
				"disk_usage":     6.4,
			},
			CategoryId: 14,
			AppId:      13,
		},
		{
			Id:          23,
			Name:        "analytics",
			Description: "Analytics data processor",
			User:        1020,
			Group:       1020,
			Directory:   "/var/analytics",
			PreExec:     "setup_analytics_db",
			CmdExec:     "/usr/bin/analytics-processor --workers=4",
			PostExec:    "check_processing_queue",
			Autostart:   true,
			Retries:     5,
			Enabled:     true,
			Pid:         1251,
			ExitCode:    0,
			StartTime:   now.Add(-3 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 10800}, // 3 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    4.7,
				"memory_percent": 35.8,
				"disk_usage":     48.2,
			},
			CategoryId: 15,
			AppId:      14,
		},
		{
			Id:          24,
			Name:        "auth-service",
			Description: "Authentication and authorization",
			User:        1021,
			Group:       1021,
			Directory:   "/opt/auth",
			PreExec:     "check_database",
			CmdExec:     "/usr/bin/auth-server --port=8081",
			PostExec:    "test_auth_endpoints",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1252,
			ExitCode:    0,
			StartTime:   now.Add(-4 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 14400}, // 4 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    1.9,
				"memory_percent": 26.4,
				"disk_usage":     9.1,
			},
			CategoryId: 16,
			AppId:      15,
		},
		{
			Id:          25,
			Name:        "rate-limiter",
			Description: "API rate limiting service",
			User:        1022,
			Group:       1022,
			Directory:   "/opt/rate-limiter",
			PreExec:     "init_redis_connection",
			CmdExec:     "/usr/bin/rate-limiter --config=limiter.conf",
			PostExec:    "test_rate_limits",
			Autostart:   true,
			Retries:     3,
			Enabled:     true,
			Pid:         1253,
			ExitCode:    0,
			StartTime:   now.Add(-2 * time.Hour).Unix(),
			EndTime:     0,
			Uptime:      &duration.Duration{Seconds: 7200}, // 2 hours
			Status:      models.ServiceStatus_STATE_RUNNING,
			Metrics: map[string]float64{
				"cpu_percent":    0.8,
				"memory_percent": 13.7,
				"disk_usage":     3.5,
			},
			CategoryId: 1,
			AppId:      1,
		},
	}
}

// renderWithSidePanel renders the UI with a side panel showing service details
func (m Model) renderWithSidePanel(title, status string) string {
	// Find the selected service
	var selectedService *models.Service
	for _, service := range m.services {
		if service.Id == m.selectedID {
			selectedService = service
			break
		}
	}

	if selectedService == nil {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
			lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Service not found"))
	}

	// Calculate layout
	tableWidth := m.width * 2 / 3
	panelWidth := m.width - tableWidth - 1 // -1 for border

	// Main table area
	table := m.table.View()
	table = lipgloss.NewStyle().Width(tableWidth).Render(table)

	// Side panel with mutt-style border
	panel := m.renderServiceDetails(selectedService, panelWidth)

	// Combine table and panel with mutt-style border
	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		table,
		lipgloss.NewStyle().
			Border(lipgloss.ThickBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color("240")).
			Render(panel),
	)

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Background(lipgloss.Color("0")).
		Padding(0, 1).
		Render("Press ESC to close details | " + m.help.View(DefaultKeyMap()))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		status,
		"",
		content,
		"",
		help,
	)
}

// renderServiceDetails renders detailed information about a service
func (m Model) renderServiceDetails(service *models.Service, width int) string {
	style := lipgloss.NewStyle().Width(width).Padding(1, 2).
		Foreground(lipgloss.Color("250")).
		Background(lipgloss.Color("0"))

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("235")).
		Padding(0, 1).
		Render("Service Details")

	// Section headers in mutt style
	sectionHeader := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("11")).
		Render

	// Basic info
	basicInfo := fmt.Sprintf(`ID: %d
Name: %s
Status: %s
PID: %d
Exit Code: %d
Enabled: %t
Autostart: %t
Retries: %d`,
		service.Id,
		service.Name,
		getStatusString(service.Status),
		service.Pid,
		service.ExitCode,
		service.Enabled,
		service.Autostart,
		service.Retries)

	// Timing info
	uptime := "-"
	if service.Uptime != nil {
		uptime = formatDuration(service.Uptime.AsDuration())
	}

	startTime := "-"
	if service.StartTime > 0 {
		startTime = time.Unix(service.StartTime, 0).Format("2006-01-02 15:04:05")
	}

	endTime := "-"
	if service.EndTime > 0 {
		endTime = time.Unix(service.EndTime, 0).Format("2006-01-02 15:04:05")
	}

	timingInfo := fmt.Sprintf(`Uptime: %s
Start Time: %s
End Time: %s`,
		uptime, startTime, endTime)

	// User/Group info
	userInfo := fmt.Sprintf(`User: %d
Group: %d
Directory: %s`,
		service.User, service.Group, service.Directory)

	// Commands
	commands := fmt.Sprintf(`Pre-Exec: %s
Command: %s
Post-Exec: %s`,
		service.PreExec, service.CmdExec, service.PostExec)

	// Metrics
	metrics := "Metrics:\n"
	if service.Metrics != nil {
		for key, value := range service.Metrics {
			metrics += fmt.Sprintf("  %s: %.2f\n", key, value)
		}
	} else {
		metrics += "  No metrics available"
	}

	// Categories
	categories := fmt.Sprintf(`Category ID: %d
App ID: %d`,
		service.CategoryId, service.AppId)

	// Combine all sections with mutt-style formatting
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		sectionHeader("Basic Information"),
		basicInfo,
		"",
		sectionHeader("Timing"),
		timingInfo,
		"",
		sectionHeader("User/Group"),
		userInfo,
		"",
		sectionHeader("Commands"),
		commands,
		"",
		sectionHeader("Categories"),
		categories,
		"",
		sectionHeader("Metrics"),
		metrics,
	)

	return style.Render(content)
}

// showStatusPopup shows the status popup for the selected service
func (m *Model) showStatusPopup() {
	for _, service := range m.services {
		if service.Id == m.selectedID {
			m.popupService = service
			m.showPopup = true
			break
		}
	}
}

// renderWithPopup renders the UI with a status popup overlay
func (m Model) renderWithPopup(mainView string) string {
	// Create the popup content
	popup := m.renderStatusPopup()

	// Create overlay with popup
	overlay := lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		popup,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("240")),
	)

	// Combine main view with overlay
	return lipgloss.JoinVertical(
		lipgloss.Top,
		mainView,
		overlay,
	)
}

// renderStatusPopup renders the status popup content
func (m Model) renderStatusPopup() string {
	service := m.popupService

	// Status color based on service status (mutt-style colors)
	var statusColor lipgloss.Color
	switch service.Status {
	case models.ServiceStatus_STATE_RUNNING:
		statusColor = lipgloss.Color("10") // Green
	case models.ServiceStatus_STATE_STOPPED:
		statusColor = lipgloss.Color("9") // Red
	case models.ServiceStatus_STATE_FAILED:
		statusColor = lipgloss.Color("1") // Bright red
	case models.ServiceStatus_STATE_STARTING, models.ServiceStatus_STATE_LOADING:
		statusColor = lipgloss.Color("11") // Yellow
	case models.ServiceStatus_STATE_READY:
		statusColor = lipgloss.Color("14") // Cyan
	default:
		statusColor = lipgloss.Color("7") // White
	}

	// Mutt-style popup styling
	popupStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("240")).
		Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("250")).
		Padding(1, 2).
		Margin(1, 1)

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("235")).
		Padding(0, 1).
		Render("Service Status")

	statusText := lipgloss.NewStyle().
		Foreground(statusColor).
		Bold(true).
		Render(getStatusString(service.Status))

	// Status details
	uptime := "-"
	if service.Uptime != nil {
		uptime = formatDuration(service.Uptime.AsDuration())
	}

	pid := "-"
	if service.Pid > 0 {
		pid = fmt.Sprintf("%d", service.Pid)
	}

	// Metrics
	cpuPercent := "-"
	memoryPercent := "-"
	if service.Metrics != nil {
		if cpu, ok := service.Metrics["cpu_percent"]; ok {
			cpuPercent = fmt.Sprintf("%.1f%%", cpu)
		}
		if mem, ok := service.Metrics["memory_percent"]; ok {
			memoryPercent = fmt.Sprintf("%.1f%%", mem)
		}
	}

	content := fmt.Sprintf(`%s

Service: %s
Status:  %s
PID:     %s
Uptime:  %s
CPU:     %s
Memory:  %s

Press ESC to close`,
		title,
		service.Name,
		statusText,
		pid,
		uptime,
		cpuPercent,
		memoryPercent,
	)

	return popupStyle.Render(content)
}
