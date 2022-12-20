package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func ListAlgoGroup(c context.Context) (groupVO []*vo.Group, err error) {
	groups, err := dal.ListAlgoGroup(c)
	if err != nil {
		return groupVO, nil
	}
	for _, g := range groups {
		groupVO = append(groupVO, GroupToVO(*g))
	}
	return groupVO, err
}

func CreateAlgoGroup(c context.Context, req req.Group) error {
	group := GroupToEntity(req)
	old, err := dal.QueryAlgoGroupByName(c, group.Name)
	if len(old) > 0 {
		return fmt.Errorf("algo <%v> already exist", old[0].Name)
	} else if err != nil {
		return fmt.Errorf("query old record failed %v", err)
	}
	userId, _ := utils.GetUserId(c)
	group.RecordMeta = model.RecordMeta{
		Id:        utils.GenerateId(),
		AccountId: userId,
		Org:       "",
		CreatedAt: time.Now(),
		CreatedBy: "Echo Bio",
	}
	_, err = dal.CreateAlgoGroup(c, group)
	return err
}

func DeleteAlgoGroupById(c context.Context, Ids []uint64) error {
	return dal.DeleteAlgoGroupById(c, Ids)
}

func UpdateAlgoGroup(c context.Context, req req.Group) error {
	group := GroupToEntity(req)
	group.UpdatedAt = time.Now()
	group.CreatedAt = time.Now()
	_, err := dal.UpdateAlgoGroup(c, group)
	return err
}

func GroupToEntity(req req.Group) *model.AlgoGroup {
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	req.RecordMeta.Id = uint64(id)
	return &model.AlgoGroup{
		RecordMeta: req.RecordMeta,
		Name:       req.Name,
		Label:      req.Label,
	}
}

func GroupToVO(group model.AlgoGroup) *vo.Group {
	return &vo.Group{
		RecordMeta: RecordMetaToVO(group.RecordMeta),
		Name:       group.Name,
		Label:      group.Label,
	}
}

func RecordMetaToVO(r model.RecordMeta) *vo.RecordMeta {
	return &vo.RecordMeta{
		Id:        fmt.Sprint(r.Id),
		AccountId: r.AccountId,
		Org:       r.Org,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		DeletedAt: r.DeletedAt,
		UpdatedBy: r.UpdatedBy,
		CreatedBy: r.CreatedBy,
		DeletedBy: r.DeletedBy,
	}
}
