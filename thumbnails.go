package proton

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// ThumbnailURL locates one encrypted thumbnail block for download.
type ThumbnailURL struct {
	ThumbnailID string
	BareURL     string
	Token       string
}

// GetThumbnails resolves download locations (bare URL + storage token) for a
// set of thumbnail IDs. Added locally for photon-migrate: upstream has no
// thumbnail download path, only upload. POST /drive/volumes/{volumeID}/thumbnails.
func (c *Client) GetThumbnails(ctx context.Context, volumeID string, thumbnailIDs []string) ([]ThumbnailURL, error) {
	var res struct {
		Code       int
		Thumbnails []ThumbnailURL
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).SetBody(map[string][]string{"ThumbnailIDs": thumbnailIDs}).
			Post("/drive/volumes/" + volumeID + "/thumbnails")
	}); err != nil {
		return nil, err
	}

	if res.Code != int(SuccessCode) {
		return nil, fmt.Errorf("GetThumbnails: unexpected response code %d", res.Code)
	}

	return res.Thumbnails, nil
}
