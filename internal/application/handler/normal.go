package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Normal(c *gin.Context) {
	var bodyMap map[string]interface{}

	if err := c.ShouldBindJSON(&bodyMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	// Generate random change number: CHG + 7 random digits
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	changeNumber := fmt.Sprintf("CHG%07d", r.Intn(10000000))

	// Generate random 32-character lowercase alphanumeric string for sys_id
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	sysID := make([]byte, 32)
	for i := range sysID {
		sysID[i] = charset[r.Intn(len(charset))]
	}

	response := fmt.Sprintf(`{
	  "result": {
		"reason": {
		  "display_value": "",
		  "value": ""
		},
		"parent": {
		  "display_value": "",
		  "value": ""
		},
		"watch_list": {
		  "display_value": "",
		  "value": ""
		},
		"upon_reject": {
		  "display_value": "Cancel all future Tasks",
		  "value": "cancel"
		},
		"sys_updated_on": {
		  "display_value": "2026-02-16 16:34:32",
		  "value": "2026-02-17 00:34:32",
		  "display_value_internal": "2026-02-16 16:34:32"
		},
		"type": {
		  "display_value": "Normal",
		  "value": "normal"
		},
		"approval_history": {
		  "display_value": "",
		  "value": ""
		},
		"test_plan": {
		  "display_value": "",
		  "value": ""
		},
		"number": {
		  "display_value": "%s",
		  "value": "%s"
		},
		"cab_delegate": {
		  "display_value": "",
		  "value": ""
		},
		"requested_by_date": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"state": {
		  "display_value": "New",
		  "value": -5.0
		},
		"sys_created_by": {
		  "display_value": "octopus",
		  "value": "octopus"
		},
		"knowledge": {
		  "display_value": "false",
		  "value": false
		},
		"order": {
		  "display_value": "",
		  "value": ""
		},
		"phase": {
		  "display_value": "Requested",
		  "value": "requested"
		},
		"cmdb_ci": {
		  "display_value": "",
		  "value": ""
		},
		"impact": {
		  "display_value": "3 - Low",
		  "value": 3.0
		},
		"contract": {
		  "display_value": "",
		  "value": ""
		},
		"active": {
		  "display_value": "true",
		  "value": true
		},
		"work_notes_list": {
		  "display_value": "",
		  "value": ""
		},
		"priority": {
		  "display_value": "4 - Low",
		  "value": 4.0
		},
		"sys_domain_path": {
		  "display_value": "/",
		  "value": "/"
		},
		"cab_recommendation": {
		  "display_value": "",
		  "value": ""
		},
		"production_system": {
		  "display_value": "false",
		  "value": false
		},
		"rejection_goto": {
		  "display_value": "",
		  "value": ""
		},
		"review_date": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"requested_by": {
		  "display_value": "",
		  "value": "c275cf9183c3b210d60a1a30ceaad3d9"
		},
		"business_duration": {
		  "display_value": "",
		  "value": ""
		},
		"group_list": {
		  "display_value": "",
		  "value": ""
		},
		"change_plan": {
		  "display_value": "",
		  "value": ""
		},
		"approval_set": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"copied_from": {
		  "display_value": "",
		  "value": ""
		},
		"implementation_plan": {
		  "display_value": "",
		  "value": ""
		},
		"universal_request": {
		  "display_value": "",
		  "value": ""
		},
		"end_date": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"short_description": {
		  "display_value": "Octopus: Deploy \"SNOW\" version 0.0.1 to \"Development\"",
		  "value": "Octopus: Deploy \"SNOW\" version 0.0.1 to \"Development\""
		},
		"correlation_display": {
		  "display_value": "",
		  "value": ""
		},
		"work_start": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"outside_maintenance_schedule": {
		  "display_value": "false",
		  "value": false
		},
		"additional_assignee_list": {
		  "display_value": "",
		  "value": ""
		},
		"std_change_producer_version": {
		  "display_value": "",
		  "value": ""
		},
		"sys_class_name": {
		  "display_value": "Change Request",
		  "value": "change_request"
		},
		"service_offering": {
		  "display_value": "",
		  "value": ""
		},
		"closed_by": {
		  "display_value": "",
		  "value": ""
		},
		"follow_up": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"review_status": {
		  "display_value": "",
		  "value": ""
		},
		"reassignment_count": {
		  "display_value": "0",
		  "value": 0.0
		},
		"start_date": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"assigned_to": {
		  "display_value": "",
		  "value": ""
		},
		"variables": {
		  "display_value": "variable_pool",
		  "value": "variable_pool"
		},
		"sla_due": {
		  "display_value": "UNKNOWN",
		  "value": "",
		  "display_value_internal": ""
		},
		"escalation": {
		  "display_value": "Normal",
		  "value": 0.0
		},
		"upon_approval": {
		  "display_value": "Proceed to Next Task",
		  "value": "proceed"
		},
		"correlation_id": {
		  "display_value": "",
		  "value": ""
		},
		"made_sla": {
		  "display_value": "true",
		  "value": true
		},
		"backout_plan": {
		  "display_value": "",
		  "value": ""
		},
		"conflict_status": {
		  "display_value": "Not Run",
		  "value": "Not Run"
		},
		"task_effective_number": {
		  "display_value": "%s",
		  "value": "%s"
		},
		"sys_updated_by": {
		  "display_value": "octopus",
		  "value": "octopus"
		},
		"opened_by": {
		  "display_value": "",
		  "value": "c275cf9183c3b210d60a1a30ceaad3d9"
		},
		"user_input": {
		  "display_value": "",
		  "value": ""
		},
		"sys_created_on": {
		  "display_value": "2026-02-16 16:34:32",
		  "value": "2026-02-17 00:34:32",
		  "display_value_internal": "2026-02-16 16:34:32"
		},
		"on_hold_task": {
		  "display_value": "",
		  "value": ""
		},
		"sys_domain": {
		  "display_value": "global",
		  "value": "global"
		},
		"route_reason": {
		  "display_value": "",
		  "value": ""
		},
		"closed_at": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"review_comments": {
		  "display_value": "",
		  "value": ""
		},
		"business_service": {
		  "display_value": "",
		  "value": ""
		},
		"time_worked": {
		  "display_value": "",
		  "value": ""
		},
		"chg_model": {
		  "display_value": "Normal",
		  "value": "007c4001c343101035ae3f52c1d3aeb2"
		},
		"expected_start": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"opened_at": {
		  "display_value": "2026-02-16 16:34:32",
		  "value": "2026-02-17 00:34:32",
		  "display_value_internal": "2026-02-16 16:34:32"
		},
		"work_end": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"phase_state": {
		  "display_value": "Open",
		  "value": "open"
		},
		"cab_date": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"work_notes": {
		  "display_value": "",
		  "value": ""
		},
		"close_code": {
		  "display_value": "",
		  "value": ""
		},
		"assignment_group": {
		  "display_value": "",
		  "value": ""
		},
		"description": {
		  "display_value": "",
		  "value": ""
		},
		"on_hold_reason": {
		  "display_value": "",
		  "value": ""
		},
		"calendar_duration": {
		  "display_value": "",
		  "value": ""
		},
		"close_notes": {
		  "display_value": "",
		  "value": ""
		},
		"sys_id": {
		  "display_value": "%s",
		  "value": "%s"
		},
		"contact_type": {
		  "display_value": "",
		  "value": ""
		},
		"cab_required": {
		  "display_value": "false",
		  "value": false
		},
		"urgency": {
		  "display_value": "3 - Low",
		  "value": 3.0
		},
		"scope": {
		  "display_value": "Medium",
		  "value": 3.0
		},
		"justification": {
		  "display_value": "",
		  "value": ""
		},
		"activity_due": {
		  "display_value": "UNKNOWN",
		  "value": "",
		  "display_value_internal": ""
		},
		"comments": {
		  "display_value": "",
		  "value": ""
		},
		"approval": {
		  "display_value": "Not Yet Requested",
		  "value": "not requested"
		},
		"due_date": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"sys_mod_count": {
		  "display_value": "0",
		  "value": 0.0
		},
		"on_hold": {
		  "display_value": "false",
		  "value": false
		},
		"sys_tags": {
		  "display_value": "",
		  "value": ""
		},
		"cab_date_time": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"conflict_last_run": {
		  "display_value": "",
		  "value": "",
		  "display_value_internal": ""
		},
		"unauthorized": {
		  "display_value": "false",
		  "value": false
		},
		"risk": {
		  "display_value": "",
		  "value": ""
		},
		"location": {
		  "display_value": "",
		  "value": ""
		},
		"category": {
		  "display_value": "Other",
		  "value": "Other"
		},
		"risk_impact_analysis": {
		  "display_value": "",
		  "value": ""
		},
		"__meta": {
		  "ignoredFields": []
		}
	  }
	}`, changeNumber, changeNumber, changeNumber, changeNumber, string(sysID), string(sysID))

	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(response))
}
