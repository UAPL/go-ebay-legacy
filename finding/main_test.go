package finding

import (
	"net/http"
	"testing"
)

var (
	applicationId = "test"
)

type HttpRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

func TestFindingApi_FindItemsAdvanced(t *testing.T) {
	var aff = Affiliate{CustomId: "test-custom-id", TrackingId:"test-tracking-id", NetworkId: "test-network-id"}
	//var finding = New(applicationId)

	var request = FindItemsAdvancedRequest{}
	request.Affiliate = aff
	request.Pagination = Pagination{Page: 5, PerPage:10}
	request.Keywords = "test keywords"
	request.Categories = append(request.Categories, 123)
	request.Categories = append(request.Categories, 456)
	request.OutputSelectors = append(request.OutputSelectors, "test-output-selector")


}
