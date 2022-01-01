package dal

import (
	"context"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/util"
	"time"
)

func CreateOrder(ctx context.Context, task *model.Order) (*model.Order, error) {
	org := util.GetOrgFromCtx(ctx)
	task.Org = org
	err := db.Create(task).Error
	return task, err
}

func UpdateOrder(ctx context.Context, task *model.Order) (*model.Order, error) {
	org := util.GetOrgFromCtx(ctx)
	task.Org = org
	err := db.Save(task).Error
	return task, err
}

func QueryOrders(ctx context.Context, req *req.OrderReq) ([]*model.Order, error) {
	org := util.GetOrgFromCtx(ctx)
	var tasks []*model.Order
	query := db.Where(map[string]interface{}{"period_id": req.PeriodId, "org": org})
	err := query.Find(&tasks).Error
	return tasks, err
}

func GetOrderById(ctx context.Context, id int64) (*model.Order, error) {
	org := util.GetOrgFromCtx(ctx)
	task := &model.Order{}
	query := db.Where("id=? and org=?", id, org)
	err := query.Find(task).Error
	return task, err

}

func QueryOrderByPriority(ctx context.Context, priority string) ([]*model.Order, error) {
	org := util.GetOrgFromCtx(ctx)
	tasks := make([]*model.Order, 0)
	query := db.Where("priority = ?and org=?", priority, org)
	err := query.Find(&tasks).Error
	return tasks, err
}

func QueryOrderByTime(ctx context.Context, startAt time.Time, endAt time.Time) ([]*model.Order, error) {
	org := util.GetOrgFromCtx(ctx)
	tasks := make([]*model.Order, 0)
	query := db.Where("created_at >? and created_at <=? and org=?", startAt, endAt, org)
	err := query.Find(&tasks).Error
	return tasks, err
}
