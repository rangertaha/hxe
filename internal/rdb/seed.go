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

package rdb

import (
	"gorm.io/gorm"
)

// SeedServices creates placeholder service records
func SeedServices() error {
	// Base services that we'll use to generate variations
	baseServices := []struct {
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

	var services []Service

	// Generate 100 services based on the base services
	for i := 0; i < 100; i++ {
		baseService := baseServices[i%len(baseServices)]

		service := Service{
			Name:        baseService.NamePrefix + " " + string('A'+rune(i%26)),
			Description: baseService.Description,
			CmdExec:     baseService.Exec,
			Directory:   "/tmp",
			User:        1000,
			Group:       1000,
			Status:      StateStopped,
			Autostart:   i%2 == 0,
			Enabled:     true,
			Retries:     i % 5,
		}
		services = append(services, service)
	}

	// Add the original service
	services = append(services, Service{
		Name:        "Database Backup Original",
		Description: "Daily backup of production database",
		CmdExec:     "pg_dump",
		Directory:   "/tmp",
		User:        0,
		Group:       0,
		Status:      StateStopped,
		Autostart:   true,
		Enabled:     true,
		Retries:     3,
	})

	// Create services only if they don't exist
	for _, service := range services {
		var existingService Service
		result := DB.Where("name = ?", service.Name).First(&existingService)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// Service doesn't exist, create it
				if err := DB.Create(&service).Error; err != nil {
					return err
				}
			} else {
				return result.Error
			}
		}
	}

	return nil
}
