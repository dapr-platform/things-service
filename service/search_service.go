package service

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"sort"
	"strconv"
	"strings"

	"things-service/entity"
	"things-service/model"
)

func QueryAllTagKey(ctx context.Context, productId, relType string) (datas []string, err error) {
	qstr := "_select=" + model.Tag_FIELD_NAME_key + "&_distinct=true&rel_type=" + relType
	if productId != "" {
		qstr += "&product_id=" + productId
	}
	keys, err := common.DbQuery[map[string]any](ctx, common.GetDaprClient(), model.Tag_with_product_idTableInfo.Name, qstr)
	if err != nil {
		err = errors.Wrap(err, "db query error")
		return
	}
	for _, v := range keys {
		datas = append(datas, v[model.Tag_FIELD_NAME_key].(string))
	}
	return
}
func QueryAllTagKeyAndValue(ctx context.Context, productId, relType string) (datas []entity.TagInfo, err error) {
	qstr := "rel_type=" + relType
	if productId != "" {
		qstr += "&product_id=" + productId
	}
	entities, err := common.DbQuery[model.Tag_with_product_id](ctx, common.GetDaprClient(), model.Tag_with_product_idTableInfo.Name, qstr)
	if err != nil {
		err = errors.Wrap(err, "db query error")
		return
	}
	tagMap := make(map[string]map[string]any)
	for _, v := range entities {
		var valMap map[string]any
		valMap, ok := tagMap[v.Key]
		if !ok {
			valMap = make(map[string]any, 0)
		}
		valMap[v.Value] = true
		tagMap[v.Key] = valMap
	}
	for k, v := range tagMap {
		vals := make([]string, 0)
		for vv, _ := range v {
			vals = append(vals, vv)
		}
		sort.Strings(vals)
		datas = append(datas, entity.TagInfo{Key: k, Values: vals})
	}

	return
}

func QueryAllTagValue(ctx context.Context, tag string) (datas []string, err error) {
	qstr := "_select=" + model.Tag_FIELD_NAME_value + "&_distinct=true&_order=value&" + model.Tag_FIELD_NAME_key + "=" + tag
	values, err := common.DbQuery[map[string]any](ctx, common.GetDaprClient(), model.TagTableInfo.Name, qstr)
	if err != nil {
		err = errors.Wrap(err, "db query error")
		return
	}
	for _, v := range values {
		datas = append(datas, v[model.Tag_FIELD_NAME_value].(string))
	}
	return
}
func QueryPointNamesByTags(ctx context.Context, tags string) (names []string, err error) {
	selectStr := "distinct name"
	fromStr := "v_point_with_tag"
	whereStr := ""
	if tags != "" {
		arr := strings.Split(tags, ",")
		for _, s := range arr {
			if s != "" {
				whereStr += " '" + s + "'=any(tags)" + " and"
			}

		}
	}
	if whereStr != "" {
		whereStr = whereStr[:strings.LastIndex(whereStr, " and")]
	} else {
		whereStr = "1=1"
	}
	data, err := common.CustomSql[map[string]string](ctx, common.GetDaprClient(), selectStr, fromStr, whereStr)
	if err != nil {
		err = errors.Wrap(err, "select data error")
		return
	}
	for _, d := range data {
		names = append(names, d["name"])
	}
	return
}
func QueryDeviceByTagsAndProductIdAndNameIdentifier(ctx context.Context, page, pageSize int, tags string, order, id, name, identifier string, productId, productName string) (pageData common.PageGeneric[entity.DeviceInfo], err error) {
	selectStr := "*"
	countSelectStr := "count(*)"
	fromStr := "v_device_with_tag"
	whereStr := ""
	orderStr := ""
	if tags != "" {
		arr := strings.Split(tags, ",")
		for _, s := range arr {
			if s != "" {
				whereStr += " '" + s + "'=any(tags)" + " and"
			}

		}
	}
	if order != "" {
		if strings.Index(order, "-") == 0 {
			orderStr = "order by " + order[1:] + " desc"
		} else {
			orderStr = "order by " + order
		}
	}
	if id != "" {
		whereStr += " id='" + id + "'" + " and"
	}
	if name != "" {
		if strings.Index(name, "$like.") >= 0 {
			name = name[strings.Index(name, "$like.")+6:]
		}
		if strings.Index(name, "%") >= 0 {
			if strings.Index(name, "%25") >= 0 {
				whereStr += " name like '" + strings.Replace(name, "%25", "%", -1) + "'" + " and"
			} else {
				whereStr += " name like '" + name + "'" + " and"
			}

		} else {
			whereStr += " name='" + name + "'" + " and"
		}
	}
	if productId != "" {
		whereStr += " product_id='" + productId + "' and"
	}
	if productName != "" {
		whereStr += " product_name='" + productName + "' and"
	}
	if identifier != "" {
		if strings.Index(identifier, "$like.") >= 0 {
			identifier = identifier[strings.Index(identifier, "$like.")+6:]
		}
		if strings.Index(identifier, "%") >= 0 {
			if strings.Index(identifier, "%25") < 0 {
				whereStr += " identifier like '" + strings.Replace(identifier, "%", "%25", -1) + "'" + " and"
			} else {
				whereStr += " identifier like '" + identifier + "'" + " and"
			}
		} else {
			whereStr += " identifier='" + identifier + "'" + " and"
		}

	}
	if whereStr != "" {
		whereStr = whereStr[:strings.LastIndex(whereStr, " and")]
	} else {
		whereStr = "1=1"
	}

	countWhereStr := whereStr
	if orderStr != "" {
		whereStr += " " + orderStr
	}
	whereStr += " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa((page-1)*pageSize)
	common.Logger.Debug("searchDevice", selectStr, fromStr, whereStr)
	counts, err := common.CustomSql[entity.CountVo](ctx, common.GetDaprClient(), countSelectStr, fromStr, countWhereStr)
	if err != nil {
		err = errors.Wrap(err, "select count error")
		return
	}
	if len(counts) != 1 {
		err = errors.New("select count error,no data")
		return
	}
	count := counts[0].Count

	data, err := common.CustomSql[entity.DeviceInfo](ctx, common.GetDaprClient(), selectStr, fromStr, whereStr)
	if err != nil {
		err = errors.Wrap(err, "select data error")
		return
	}
	pageData = common.PageGeneric[entity.DeviceInfo]{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		Items:    data,
	}
	return
}

func QueryDeviceByFuzzyText(ctx context.Context, page, pageSize int, tags, q string, order, productId, productName string) (pageData common.PageGeneric[entity.DeviceInfo], err error) {
	selectStr := "*"
	countSelectStr := "count(*)"
	fromStr := "v_device_with_tag_filter"
	whereStr := ""
	orderStr := ""

	if order != "" {
		if strings.Index(order, "-") == 0 {
			orderStr = "order by " + order[1:] + " desc"
		} else {
			orderStr = "order by " + order
		}
	}
	if tags != "" {
		arr := strings.Split(tags, ",")
		for _, s := range arr {
			if s != "" {
				whereStr += " '" + s + "'=any(tags)" + " and"
			}

		}
	}

	if productId != "" {
		whereStr += " product_id='" + productId + "' and"
	}
	if productName != "" {
		whereStr += " product_name='" + productName + "' and"
	}
	if q != "" {
		qarr := strings.Split(q, " ")
		for _, s := range qarr {
			if s != "" {
				whereStr += " filter_text like '%25" + s + "%25'" + " and" //通过url 调用db-service，需要将 %进行过滤
			}
		}
	}
	if whereStr != "" {
		whereStr = whereStr[:strings.LastIndex(whereStr, " and")]
	} else {
		whereStr = "1=1"
	}

	countWhereStr := whereStr
	if orderStr != "" {
		whereStr += " " + orderStr
	}
	whereStr += " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa((page-1)*pageSize)
	counts, err := common.CustomSql[entity.CountVo](ctx, common.GetDaprClient(), countSelectStr, fromStr, countWhereStr)
	if err != nil {
		err = errors.Wrap(err, "select count error")
		return
	}
	if len(counts) != 1 {
		err = errors.New("select count error,no data")
		return
	}
	count := counts[0].Count

	data, err := common.CustomSql[entity.DeviceInfo](ctx, common.GetDaprClient(), selectStr, fromStr, whereStr)
	if err != nil {
		err = errors.Wrap(err, "select data error")
		return
	}
	pageData = common.PageGeneric[entity.DeviceInfo]{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		Items:    data,
	}
	return
}

func QueryDeviceByTagsAndProductId(ctx context.Context, tags string, productId string) (data []entity.DeviceInfo, err error) {
	selectStr := "*"
	fromStr := "v_device_with_tag"
	whereStr := ""
	if tags != "" {
		arr := strings.Split(tags, ",")
		for _, s := range arr {
			if s != "" {
				whereStr += " '" + s + "'=any(tags)" + " and"
			}

		}
	}

	if productId != "" {
		whereStr += " product_id='" + productId + "' and"
	}

	if whereStr != "" {
		whereStr = whereStr[:strings.LastIndex(whereStr, " and")]
	} else {
		whereStr = "1=1"
	}

	data, err = common.CustomSql[entity.DeviceInfo](ctx, common.GetDaprClient(), selectStr, fromStr, whereStr)
	if err != nil {
		err = errors.Wrap(err, "select data error")
		return
	}

	return
}

func QueryDeviceByTagValueLike(ctx context.Context, page, pageSize int, queryStr string) (pageData common.PageGeneric[entity.DeviceCurrentData], err error) {
	selectStr := "v.*"
	countSelectStr := "count(v.*)"
	fromStr := ` v_device_current_data v,(
SELECT distinct o_device.id
FROM o_device
INNER JOIN o_point ON o_device.id = o_point.device_id
INNER JOIN o_tag ON (o_device.id = o_tag.rel_id) `
	if queryStr != "" {
		fromStr += " WHERE "
		arr := strings.Split(queryStr, " ")
		orArr := make([]string, 0)
		for _, s := range arr {
			if s != "" {
				orArr = append(orArr, " o_tag.value LIKE '%"+s+"%' ")
			}
		}
		fromStr += strings.Join(orArr, " OR ")
		fromStr += ") b"
	} else {
		fromStr += `) b `
	}
	whereStr := "v.id=b.id "

	countWhereStr := whereStr
	whereStr += " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa((page-1)*pageSize)
	counts, err := common.CustomSql[entity.CountVo](ctx, common.GetDaprClient(), countSelectStr, fromStr, countWhereStr)
	if err != nil {
		err = errors.Wrap(err, "select count error")
		return
	}
	if len(counts) != 1 {
		err = errors.New("select count error,no data")
		return
	}
	count := counts[0].Count

	data, err := common.CustomSql[entity.DeviceCurrentData](ctx, common.GetDaprClient(), selectStr, fromStr, whereStr)
	if err != nil {
		err = errors.Wrap(err, "select data error")
		return
	}
	pageData = common.PageGeneric[entity.DeviceCurrentData]{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		Items:    data,
	}
	return
}
