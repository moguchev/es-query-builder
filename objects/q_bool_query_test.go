package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	objs "github.com/sdqri/effdsl/objects"
)

func TestBoolQuery(t *testing.T) {
	expectedBody := `{"bool":{"must":[{"fake_query1":"fake_value1"},{"fake_query2":"fake_value2"}],"filter":[{"fake_filter1":"fake_value1"},{"fake_filter2":"fake_value2"}],"must_not":[{"fake_query1":"fake_value1"},{"fake_query2":"fake_value2"}],"should":[{"fake_query1":"fake_value1"},{"fake_query2":"fake_value2"}]}}`
	boolQueryResult := objs.D.BoolQuery(
		objs.D.Must(objs.MockQuery(objs.M{"fake_query1": "fake_value1"})),
		objs.D.Must(objs.MockQuery(objs.M{"fake_query2": "fake_value2"})),
		objs.D.Filter(objs.MockQuery(objs.M{"fake_filter1": "fake_value1"})),
		objs.D.Filter(objs.MockQuery(objs.M{"fake_filter2": "fake_value2"})),
		objs.D.Should(objs.MockQuery(objs.M{"fake_query1": "fake_value1"})),
		objs.D.Should(objs.MockQuery(objs.M{"fake_query2": "fake_value2"})),
		objs.D.MustNot(objs.MockQuery(objs.M{"fake_query1": "fake_value1"})),
		objs.D.MustNot(objs.MockQuery(objs.M{"fake_query2": "fake_value2"})),
	)
	err := boolQueryResult.Err
	boolQuery := boolQueryResult.Ok
	assert.Nil(t, err)
	jsonBody, err := json.Marshal(boolQuery)
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(jsonBody))
}
