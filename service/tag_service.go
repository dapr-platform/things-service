package service

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"things-service/entity"
	"things-service/model"
)

/*批量获取一个对象的tag*/
func BatchGetTag(ctx context.Context, relId string) (info entity.BatchTagInfo, err error) {
	mTags, err := common.DbQuery[model.Tag](ctx, common.GetDaprClient(), model.TagTableInfo.Name, model.Tag_FIELD_NAME_rel_id+"="+relId)
	if err != nil {
		err = errors.Wrap(err, "dbquery error")
		return
	}
	if len(mTags) > 0 {
		info.RelId = relId
		info.RelType = int(mTags[0].RelType)
		info.Tags = make([]entity.Tag, 0)
		for _, tag := range mTags {
			info.Tags = append(info.Tags, entity.Tag{
				Key:      tag.Key,
				Value:    tag.Value,
				Editable: int(tag.Editable),
			})
		}
	} else {
		info.RelId = relId
	}
	return

}

func BatchSaveTag(ctx context.Context, req entity.BatchTagInfo, deleteFlag bool) (err error) {
	relId := req.RelId
	if deleteFlag {
		err = common.DbDeleteByOps(ctx, common.GetDaprClient(), model.TagTableInfo.Name, []string{model.Tag_FIELD_NAME_rel_id}, []string{"=="}, []interface{}{relId})
		if err != nil {
			err = errors.Wrap(err, "delete relid error")
			return
		}
	}

	if len(req.Tags) == 0 {
		return
	}
	mTags := make([]model.Tag, 0)
	for _, tag := range req.Tags {
		if tag.Key == "" || tag.Value == "" {
			continue
		}
		id := common.GetMD5Hash(relId + "_" + tag.Key)
		mTags = append(mTags, model.Tag{
			ID:       id,
			RelID:    relId,
			Key:      tag.Key,
			Value:    cast.ToString(tag.Value),
			RelType:  int32(req.RelType),
			Editable: int32(tag.Editable),
		})
	}
	if len(mTags) == 0 {
		return
	}
	err = common.DbBatchUpsert[model.Tag](ctx, common.GetDaprClient(), mTags, model.TagTableInfo.Name, model.Tag_FIELD_NAME_id)
	if err != nil {
		err = errors.Wrap(err, "batch insert error")
	}
	return
}
