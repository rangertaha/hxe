package models

import (
	"github.com/rangertaha/hxe/internal/db"
	"github.com/rangertaha/hxe/internal/log"
)


func AutoMigrate() (err error) {
	// Auto migrate models
	if err = db.AutoMigrate(
		Program{},
	); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate models")
	}

	Seed()
	return
}

func Seed() (err error) {
	programs := []Program{
		{
			Name:     "postgresql",
			Desc:     "PostgreSQL database server",
			Path:     "/usr/lib/postgresql/14/bin/postgres",
			User:     "postgres",
			Group:    "postgres", 
			Dir:      "/var/lib/postgresql/14/main",
			Args:     []string{"-D", "/var/lib/postgresql/14/main"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/postgresql",
			Exec:     "/usr/lib/postgresql/14/bin/postgres",
			PostExec: "rm -f /var/run/postgresql/14-main.pid",
			Enabled:  true,
		},
		{
			Name:     "mongodb",
			Desc:     "MongoDB database server",
			Path:     "/usr/bin/mongod",
			User:     "mongodb",
			Group:    "mongodb",
			Dir:      "/var/lib/mongodb",
			Args:     []string{"--config", "/etc/mongodb.conf"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/mongodb",
			Exec:     "/usr/bin/mongod",
			PostExec: "rm -f /var/run/mongodb.pid",
			Enabled:  true,
		},
		{
			Name:     "mysql",
			Desc:     "MySQL database server",
			Path:     "/usr/sbin/mysqld",
			User:     "mysql",
			Group:    "mysql",
			Dir:      "/var/lib/mysql",
			Args:     []string{"--defaults-file=/etc/mysql/my.cnf"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/mysqld",
			Exec:     "/usr/sbin/mysqld",
			PostExec: "rm -f /var/run/mysqld/mysqld.pid",
			Enabled:  true,
		},
		{
			Name:     "apache2",
			Desc:     "Apache HTTP Server",
			Path:     "/usr/sbin/apache2",
			User:     "www-data",
			Group:    "www-data",
			Dir:      "/etc/apache2",
			Args:     []string{"-DFOREGROUND"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/apache2",
			Exec:     "/usr/sbin/apache2",
			PostExec: "rm -f /var/run/apache2/apache2.pid",
			Enabled:  true,
		},
		{
			Name:     "elasticsearch",
			Desc:     "Elasticsearch search engine",
			Path:     "/usr/share/elasticsearch/bin/elasticsearch",
			User:     "elasticsearch",
			Group:    "elasticsearch",
			Dir:      "/usr/share/elasticsearch",
			Args:     []string{"-p", "/var/run/elasticsearch/elasticsearch.pid"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/elasticsearch",
			Exec:     "/usr/share/elasticsearch/bin/elasticsearch",
			PostExec: "rm -f /var/run/elasticsearch/elasticsearch.pid",
			Enabled:  true,
		},
		{
			Name:     "rabbitmq",
			Desc:     "RabbitMQ message broker",
			Path:     "/usr/sbin/rabbitmq-server",
			User:     "rabbitmq",
			Group:    "rabbitmq",
			Dir:      "/var/lib/rabbitmq",
			Args:     []string{},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/rabbitmq",
			Exec:     "/usr/sbin/rabbitmq-server",
			PostExec: "rm -f /var/run/rabbitmq/pid",
			Enabled:  true,
		},
		{
			Name:     "memcached",
			Desc:     "Memcached memory caching system",
			Path:     "/usr/bin/memcached",
			User:     "memcache",
			Group:    "memcache",
			Dir:      "/var/run/memcached",
			Args:     []string{"-m", "64", "-p", "11211", "-u", "memcache", "-l", "127.0.0.1"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/memcached",
			Exec:     "/usr/bin/memcached",
			PostExec: "rm -f /var/run/memcached/memcached.pid",
			Enabled:  true,
		},
		{
			Name:     "prometheus",
			Desc:     "Prometheus monitoring system",
			Path:     "/usr/local/bin/prometheus",
			User:     "prometheus",
			Group:    "prometheus",
			Dir:      "/etc/prometheus",
			Args:     []string{"--config.file=/etc/prometheus/prometheus.yml"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/prometheus",
			Exec:     "/usr/local/bin/prometheus",
			PostExec: "rm -f /var/run/prometheus/prometheus.pid",
			Enabled:  true,
		},
		{
			Name:     "grafana",
			Desc:     "Grafana analytics platform",
			Path:     "/usr/sbin/grafana-server",
			User:     "grafana",
			Group:    "grafana",
			Dir:      "/usr/share/grafana",
			Args:     []string{"--config", "/etc/grafana/grafana.ini"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/grafana",
			Exec:     "/usr/sbin/grafana-server",
			PostExec: "rm -f /var/run/grafana/grafana-server.pid",
			Enabled:  true,
		},
		{
			Name:     "jenkins",
			Desc:     "Jenkins automation server",
			Path:     "/usr/bin/java",
			User:     "jenkins",
			Group:    "jenkins",
			Dir:      "/var/lib/jenkins",
			Args:     []string{"-jar", "/usr/share/jenkins/jenkins.war"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/run/jenkins",
			Exec:     "/usr/bin/java",
			PostExec: "rm -f /var/run/jenkins/jenkins.pid",
			Enabled:  true,
		},
		{
			Name:     "nginx",
			Desc:     "Web server",
			Path:     "/usr/sbin/nginx",
			User:     "www-data",
			Group:    "www-data",
			Dir:      "/etc/nginx",
			Args:     []string{"-g", "daemon off;"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "mkdir -p /var/log/nginx",
			Exec:     "/usr/sbin/nginx",
			PostExec: "rm -f /var/run/nginx.pid",
			Enabled:  true,
		},
		{
			Name:     "redis",
			Desc:     "In-memory data store",
			Path:     "/usr/local/bin/redis-server",
			User:     "redis",
			Group:    "redis",
			Dir:      "/etc/redis",
			Args:     []string{"/etc/redis/redis.conf"},
			Env:      []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			PreExec:  "sysctl vm.overcommit_memory=1",
			Exec:     "/usr/local/bin/redis-server",
			PostExec: "rm -f /var/run/redis_*.pid",
			Enabled:  true,
		},
	}

	result := db.DB.Create(&programs)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("failed to seed programs")
		return result.Error
	}
	return nil
}
