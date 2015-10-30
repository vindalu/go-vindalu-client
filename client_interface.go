package vindalu

import (
	"encoding/json"
	"fmt"
	"net/http"
    "strings"

	"github.com/vindalu/vindalu/core"
)

func (c *Client) Create(atype, id string, data interface{}) (created map[string]string, err error) {
	var (
		resp  *http.Response
		b     []byte
		udata []byte
	)

	if data == nil {
		err = fmt.Errorf("Data cannot be `nil`")
		return
	}

	if udata, err = json.Marshal(data); err != nil {
		return
	}

	if resp, b, err = c.doRequest("POST", c.getOpaque(atype, id), udata); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s", b)
		return
	}

	created = map[string]string{}
	err = json.Unmarshal(b, &created)
	return
}

func (c *Client) Get(atype, id string, version int64) (ba core.BaseAsset, err error) {
	var (
		resp   *http.Response
		b      []byte
		opaque = c.getOpaque(atype, id)
	)

	if version != 0 {
		opaque = fmt.Sprintf("%s?version=%d", opaque, version)
	}

	if resp, b, err = c.doRequest("GET", opaque, nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	err = json.Unmarshal(b, &ba)
	return
}

func (c *Client) List(atype string, queryParams map[string]string, count int64) (ba []core.BaseAsset, err error) {
	var (
		resp   *http.Response
		b      []byte
		opaque = c.getOpaque(atype)
	)

	//if version != 0 {
	//	opaque = fmt.Sprintf("%s?version=%d", opaque, version)
	//}

    if len(queryParams) != 0 {
        first := true
        for k,v := range queryParams {
            if first {
                opaque = fmt.Sprintf("%s?%s=%s", opaque, k, v)
                first = false
            } else {
                opaque = fmt.Sprintf("%s&%s=%s", opaque, k, v)
            }
        }
    }

	if resp, b, err = c.doRequest("GET", opaque, nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	err = json.Unmarshal(b, &ba)
	return
}

func (c *Client) GetVersions(atype, id string) (versions []core.BaseAsset, err error) {
	var (
		resp   *http.Response
		b      []byte
		opaque = c.getOpaque(atype, id, "versions")
	)

	if resp, b, err = c.doRequest("GET", opaque, nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	err = json.Unmarshal(b, &versions)
	return
}

func (c *Client) GetVersionDiffs(atype, id string) (diffs []interface{}, err error) {
	var (
		resp   *http.Response
		b      []byte
		opaque = c.getOpaque(atype, id, "versions") + "?diff"
	)

	if resp, b, err = c.doRequest("GET", opaque, nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	diffs = []interface{}{}
	err = json.Unmarshal(b, &diffs)
	return
}

func (c *Client) Update(atype, id string, data interface{}, deletedFields ...string) (updated map[string]string, err error) {
	var (
		resp  *http.Response
		b     []byte
		udata []byte
	)

	if data == nil {
		err = fmt.Errorf("Data cannot be `nil`")
		return
	}

	if udata, err = json.Marshal(data); err != nil {
		return
	}

    urlPath := c.getOpaque(atype, id)
    if len(deletedFields) > 0 {
        urlPath = urlPath + "?delete_fields=" + strings.Join(deletedFields, ",")
    }

	if resp, b, err = c.doRequest("PUT", urlPath, udata); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	updated = map[string]string{}
	err = json.Unmarshal(b, &updated)
	return
}

func (c *Client) Delete(atype, id string) (deleted map[string]string, err error) {
	var (
		b    []byte
		resp *http.Response
	)

	if resp, b, err = c.doRequest("DELETE", c.getOpaque(atype, id), nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	deleted = map[string]string{}
	err = json.Unmarshal(b, &deleted)
	return
}

func (c *Client) GetTypes() (aggrs []core.AggregatedItem, err error) {
	var (
		b    []byte
		resp *http.Response
	)

	if resp, b, err = c.doRequest("GET", c.getOpaque(), nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	aggrs = []core.AggregatedItem{}
	err = json.Unmarshal(b, &aggrs)
	return
}

func (c *Client) ListTypeProperties(atype string) (props []string, err error) {
	var (
		b    []byte
		resp *http.Response
	)

	if resp, b, err = c.doRequest("GET", c.getOpaque(atype), nil); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%s - %s", resp.Status, b)
		return
	}

	props = []string{}
	err = json.Unmarshal(b, &props)
	return
}
