package models

import (
	"gorm.io/gorm"
)

// SeedPrograms creates placeholder program records
func SeedPrograms(db *gorm.DB) error {
	// Base programs that we'll use to generate variations
	basePrograms := []struct {
		NamePrefix  string
		Exec        string
		Description string
	}{
		{"Database Backup", "pg_dump -U postgres db_%d > /backups/db_%d.sql", "Automated database backup task"},
		{"Log Rotation", "logrotate /etc/logrotate.d/app_%d", "Rotate and compress application logs"},
		{"System Monitor", "python /scripts/monitor.py --service=%d", "Monitor system health and resources"},
		{"Data Sync", "rsync -avz /data/source_%d/ /data/dest_%d/", "Synchronize data between directories"},
		{"Security Scan", "nmap -sV localhost:808%d", "Security vulnerability scanning"},
		{"API Health Check", "curl -f http://api-%d.example.com/health", "Check API endpoint health"},
		{"Cache Cleanup", "redis-cli -n %d flushdb", "Clean up Redis cache instances"},
		{"Metrics Collection", "node /scripts/collect_metrics.js --instance=%d", "Collect system metrics"},
		{"Config Validation", "python /scripts/validate_config.py --env=%d", "Validate service configurations"},
		{"Backup Verification", "sha256sum /backups/backup_%d.tar", "Verify backup integrity"},
	}

	var programs []Program

	// Generate 100 programs based on the base programs
	for i := 0; i < 100; i++ {
		baseProgram := basePrograms[i%len(basePrograms)]

		program := Program{
			Name:        baseProgram.NamePrefix + " " + string('A'+rune(i%26)),
			Description: baseProgram.Description,
			Command:     baseProgram.Exec,
			Args:        "",
			WorkingDir:  "/tmp",
			User:        "nobody",
			Group:       "nobody",
			Status:      "stopped",
			Autostart:   i%2 == 0,
			Enabled:     true,
			Retries:     i % 5,
			MaxRetries:  5,
		}
		programs = append(programs, program)
	}

	// Add the original program
	programs = append(programs, Program{
		Name:        "Database Backup Original",
		Description: "Daily backup of production database",
		Command:     "pg_dump",
		Args:        "-U postgres mydb > /backups/db_$(date +%Y%m%d).sql",
		WorkingDir:  "/tmp",
		User:        "nobody",
		Group:       "nobody",
		Status:      "stopped",
		Autostart:   true,
		Enabled:     true,
		Retries:     3,
		MaxRetries:  5,
	})

	// Create programs only if they don't exist
	for _, program := range programs {
		var existingProgram Program
		result := db.Where("name = ?", program.Name).First(&existingProgram)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// Program doesn't exist, create it
				if err := db.Create(&program).Error; err != nil {
					return err
				}
			} else {
				return result.Error
			}
		}
	}

	return nil
}

// SeedAll runs all seed functions in the correct order
func SeedAll(db *gorm.DB) error {
	if err := SeedPrograms(db); err != nil {
		return err
	}

	return nil
}
