package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Change(c *gin.Context) {
	query := c.Query("sysparm_query")

	if strings.Contains(query, "numberIN") {
		crNumber := strings.TrimPrefix(query, "numberIN")
		searchChange(c, crNumber)
		return
	}

	descriptionSearch(c, query)
}

func descriptionSearch(c *gin.Context, query string) {
	// JSON escape the query string
	escapedQuery, _ := json.Marshal(query)
	// Remove the surrounding quotes added by json.Marshal
	escapedQueryStr := string(escapedQuery[1 : len(escapedQuery)-1])

	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(`{
		"result": [
			{
				"__meta": {
					"encodedQuery": "short_description=`+escapedQueryStr+`",
					"fields": {
						"applied": [
							"short_description"
						],
						"ignored": []
					}
				}
			}
		]
	}`))
}

func searchChange(c *gin.Context, crNumber string) {
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(`{
	  "result": [
		{
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
			"display_value": "2026-02-18 12:03:03",
			"value": "2026-02-18 20:03:03",
			"display_value_internal": "2026-02-18 12:03:03"
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
			"display_value": "`+crNumber+`",
			"value": "`+crNumber+`"
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
			"display_value": "Octopus: Deploy \"SNOW\" version 0.0.9 to \"Development\"",
			"value": "Octopus: Deploy \"SNOW\" version 0.0.9 to \"Development\""
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
			"display_value": "`+crNumber+`",
			"value": "`+crNumber+`"
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
			"display_value": "2026-02-18 12:03:03",
			"value": "2026-02-18 20:03:03",
			"display_value_internal": "2026-02-18 12:03:03"
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
			"display_value": "2026-02-18 12:03:03",
			"value": "2026-02-18 20:03:03",
			"display_value_internal": "2026-02-18 12:03:03"
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
			"display_value": "9f6e5bcb83033210d60a1a30ceaad3d2",
			"value": "9f6e5bcb83033210d60a1a30ceaad3d2"
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
		  }
		},
		{
		  "__meta": {
			"encodedQuery": "numberIN`+crNumber+`",
			"fields": {
			  "applied": [
				"number"
			  ],
			  "ignored": []
			}
		  }
		}
	  ]
	}`))
}
