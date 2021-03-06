@endpoint(incidents)
Feature: Incidents
  Manage incident response.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Incidents" API

  Scenario: Get a list of incidents returns "OK" response
    Given operation "ListIncidents" enabled
    And there is a valid "incident" in the system
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 200 OK

  Scenario: Create an incident returns "CREATED" response
    Given there is a valid "user" in the system
    And operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body {"data": {"type": "incidents", "attributes": {"title": "{{unique}}", "customer_impacted": false, "fields": {"state": {"type": "dropdown", "value": "resolved"}}}, "relationships": {"commander": {"data": {"type": "{{ user.data.type }}", "id": "{{ user.data.id }}"}}}}}
    When the request is sent
    Then the response status is 201 CREATED

  Scenario: Delete an existing incident returns "OK" response
    Given operation "DeleteIncident" enabled
    And there is a valid "incident" in the system
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    When the request is sent
    Then the response status is 204 OK

  Scenario: Get the details of an incident returns "OK" response
    Given operation "GetIncident" enabled
    And there is a valid "incident" in the system
    And new "GetIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.title" has the same value as "incident.data.attributes.title"

  Scenario: Update an existing incident returns "OK" response
    Given operation "UpdateIncident" enabled
    And there is a valid "incident" in the system
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    And body {"data": {"id": "{{incident.data.id}}", "type": "incidents", "attributes": {"fields": {"state": {"type": "dropdown", "value":"resolved"}}, "title": "{{ incident.data.attributes.title }}-updated"}}}
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.title" is equal to "{{ incident.data.attributes.title }}-updated"

  @generated @skip
  Scenario: Get a list of incidents returns "Bad Request" response
    Given operation "ListIncidents" enabled
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Get a list of incidents returns "Not Found" response
    Given operation "ListIncidents" enabled
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip
  Scenario: Create an incident returns "Bad Request" response
    Given operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body {}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Create an incident returns "Not Found" response
    Given operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body {}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip
  Scenario: Delete an existing incident returns "Bad Request" response
    Given operation "DeleteIncident" enabled
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Delete an existing incident returns "Not Found" response
    Given operation "DeleteIncident" enabled
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip
  Scenario: Get the details of an incident returns "Bad Request" response
    Given operation "GetIncident" enabled
    And new "GetIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Get the details of an incident returns "Not Found" response
    Given operation "GetIncident" enabled
    And new "GetIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip
  Scenario: Update an existing incident returns "Bad Request" response
    Given operation "UpdateIncident" enabled
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Update an existing incident returns "Not Found" response
    Given operation "UpdateIncident" enabled
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 404 Not Found
