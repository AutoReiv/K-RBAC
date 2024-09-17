package handlers

import (
	"net/http"
	"rbac/pkg/db"

	"github.com/labstack/echo/v4"
)

// AuditLog represents a single audit log entry.
type AuditLog struct {
	ID           int    `json:"id"`
	Action       string `json:"action"`
	ResourceName string `json:"resource_name"`
	Namespace    string `json:"namespace"`
	Timestamp    string `json:"timestamp"`
	Hash         string `json:"hash"`
}

// GetAuditLogsHandler handles the retrieval of audit logs.
func GetAuditLogsHandler(c echo.Context) error {
	rows, err := db.DB.Query("SELECT id, action, resource_name, namespace, timestamp, hash FROM audit_logs ORDER BY timestamp DESC")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve audit logs: " + err.Error()})
	}
	defer rows.Close()

	var logs []AuditLog
	for rows.Next() {
		var log AuditLog
		if err := rows.Scan(&log.ID, &log.Action, &log.ResourceName, &log.Namespace, &log.Timestamp, &log.Hash); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan audit log: " + err.Error()})
		}
		logs = append(logs, log)
	}

	return c.JSON(http.StatusOK, logs)
}
