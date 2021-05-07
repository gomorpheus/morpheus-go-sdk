package morpheus_test

import (
	"testing"

	"github.com/gomorpheus/morpheus-go-sdk"
)

<<<<<<< HEAD
=======
var (
	testTaskSetName = "golangtestworkflow"
)

>>>>>>> 2a75585 (Add additional resources)
func TestListTaskSets(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListTaskSets(req)
	assertResponse(t, resp, err)
}

func TestGetTaskSet(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListTaskSets(req)
	assertResponse(t, resp, err)

	// parse JSON and fetch the first one by ID

	result := resp.Result.(*morpheus.ListTaskSetsResult)
	recordCount := result.Meta.Total
	t.Logf("Found %d Task Sets.", recordCount)
	if recordCount != 0 {
		// Get by ID
		record := (*result.TaskSets)[0]
		resp, err = client.GetTaskSet(record.ID, &morpheus.Request{})
		assertResponse(t, resp, err)

		// List by name

<<<<<<< HEAD
	}
=======
	} else {

	}
}

func TestTaskSetsCRUD(t *testing.T) {
	client := getTestClient(t)
	//create
	// this has no uniqueness check on name, it probably should..
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"taskSet": map[string]interface{}{
				"name":             testNetworkDomainName,
				"description":      "a test domain",
				"publicZone":       false,
				"domainController": false,
				"visibility":       "private",
			},
		},
	}
	resp, err := client.CreateTaskSet(req)
	result := resp.Result.(*morpheus.CreateTaskSetResult)
	assertResponse(t, resp, err)
	assertNotNil(t, result)
	assertEqual(t, result.Success, true)
	assertNotNil(t, result.TaskSet)
	assertNotEqual(t, result.TaskSet.ID, 0)
	assertEqual(t, result.TaskSet.Name, testNetworkDomainName)

	// update
	updateReq := &morpheus.Request{
		Body: map[string]interface{}{
			"taskSet": map[string]interface{}{
				"description": "my new description",
			},
		},
	}
	updateResp, updateErr := client.UpdateTaskSet(result.TaskSet.ID, updateReq)
	updateResult := updateResp.Result.(*morpheus.UpdateTaskSetResult)
	assertResponse(t, updateResp, updateErr)
	assertNotNil(t, updateResult)
	assertEqual(t, updateResult.Success, true)
	assertNotNil(t, updateResult.TaskSet)
	assertNotEqual(t, updateResult.TaskSet.ID, 0)
	assertEqual(t, updateResult.TaskSet.Description, "my new description")

	// delete
	deleteReq := &morpheus.Request{}
	deleteResp, deleteErr := client.DeleteTaskSet(result.TaskSet.ID, deleteReq)
	deleteResult := deleteResp.Result.(*morpheus.DeleteTaskSetResult)
	assertResponse(t, deleteResp, deleteErr)
	assertNotNil(t, deleteResult)
	assertEqual(t, deleteResult.Success, true)

>>>>>>> 2a75585 (Add additional resources)
}
