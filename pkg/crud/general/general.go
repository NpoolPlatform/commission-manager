package general

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/archivement-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/archivement-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/archivement-manager/pkg/tracer/general"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/archivement-manager/pkg/db"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.GeneralReq) (*ent.General, error) {
	var info *ent.General
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.General.Create()

		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.GoodID != nil {
			c.SetGoodID(uuid.MustParse(in.GetGoodID()))
		}
		if in.CoinTypeID != nil {
			c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
		}

		c.SetTotalAmount(decimal.NewFromInt(0))
		c.SetSelfAmount(decimal.NewFromInt(0))
		c.SetTotalUnits(0)
		c.SetSelfUnits(0)
		c.SetTotalCommission(decimal.NewFromInt(0))
		c.SetSelfCommission(decimal.NewFromInt(0))

		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.GeneralReq) ([]*ent.General, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.General{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.GeneralCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.General.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.GoodID != nil {
				bulk[i].SetGoodID(uuid.MustParse(info.GetGoodID()))
			}
			if info.CoinTypeID != nil {
				bulk[i].SetCoinTypeID(uuid.MustParse(info.GetCoinTypeID()))
			}
			bulk[i].SetTotalAmount(decimal.NewFromInt(0))
			bulk[i].SetSelfAmount(decimal.NewFromInt(0))
			bulk[i].SetTotalUnits(0)
			bulk[i].SetSelfUnits(0)
			bulk[i].SetTotalCommission(decimal.NewFromInt(0))
			bulk[i].SetSelfCommission(decimal.NewFromInt(0))
		}
		rows, err = tx.General.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func AddFields(ctx context.Context, in *npool.GeneralReq) (*ent.General, error) { //nolint
	var info *ent.General
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.General.Query().Where(general.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query general: %v", err)
		}

		totalAmount := decimal.NewFromInt(0)
		if in.TotalAmount != nil {
			totalAmount, err = decimal.NewFromString(in.GetTotalAmount())
			if err != nil {
				return err
			}
		}

		if totalAmount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("TotalAmount < 0")
		}

		selfAmount := decimal.NewFromInt(0)
		if in.SelfAmount != nil {
			selfAmount, err = decimal.NewFromString(in.GetSelfAmount())
			if err != nil {
				return err
			}
		}

		if selfAmount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("SelfAmount < 0")
		}

		totalCommission := decimal.NewFromInt(0)
		if in.TotalCommission != nil {
			totalCommission, err = decimal.NewFromString(in.GetTotalCommission())
			if err != nil {
				return err
			}
		}

		if totalCommission.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("TotalCommission < 0")
		}

		selfCommission := decimal.NewFromInt(0)
		if in.SelfCommission != nil {
			selfCommission, err = decimal.NewFromString(in.GetSelfCommission())
			if err != nil {
				return err
			}
		}

		if selfCommission.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("SelfCommission < 0")
		}

		stm := info.Update()

		if in.TotalAmount != nil {
			stm = stm.AddTotalAmount(totalAmount)
		}
		if in.SelfAmount != nil {
			stm = stm.AddSelfAmount(selfAmount)
		}
		if in.TotalUnits != nil {
			stm = stm.AddTotalUnits(int32(in.GetTotalUnits()))
		}
		if in.SelfUnits != nil {
			stm = stm.AddSelfUnits(int32(in.GetSelfUnits()))
		}
		if in.TotalCommission != nil {
			stm = stm.AddTotalCommission(totalCommission)
		}
		if in.SelfCommission != nil {
			stm = stm.AddSelfCommission(selfCommission)
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update general: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update general: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.General, error) {
	var info *ent.General
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.General.Query().Where(general.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.GeneralQuery, error) { //nolint
	stm := cli.General.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(general.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(general.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(general.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.UserIDs != nil {
		switch conds.GetUserIDs().GetOp() {
		case cruder.IN:
			users := []uuid.UUID{}
			for _, user := range conds.GetUserIDs().GetValue() {
				users = append(users, uuid.MustParse(user))
			}
			stm.Where(general.UserIDIn(users...))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(general.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(general.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.TotalAmount != nil {
		incoming, err := decimal.NewFromString(conds.GetTotalAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetTotalAmount().GetOp() {
		case cruder.LT:
			stm.Where(general.TotalAmountLT(incoming))
		case cruder.GT:
			stm.Where(general.TotalAmountGT(incoming))
		case cruder.EQ:
			stm.Where(general.TotalAmountEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.SelfAmount != nil {
		incoming, err := decimal.NewFromString(conds.GetSelfAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetSelfAmount().GetOp() {
		case cruder.LT:
			stm.Where(general.SelfAmountLT(incoming))
		case cruder.GT:
			stm.Where(general.SelfAmountGT(incoming))
		case cruder.EQ:
			stm.Where(general.SelfAmountEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.TotalUnits != nil {
		switch conds.GetTotalUnits().GetOp() {
		case cruder.LT:
			stm.Where(general.TotalUnitsLT(conds.GetTotalUnits().GetValue()))
		case cruder.GT:
			stm.Where(general.TotalUnitsGT(conds.GetTotalUnits().GetValue()))
		case cruder.EQ:
			stm.Where(general.TotalUnitsEQ(conds.GetTotalUnits().GetValue()))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.SelfUnits != nil {
		switch conds.GetSelfUnits().GetOp() {
		case cruder.LT:
			stm.Where(general.SelfUnitsLT(conds.GetSelfUnits().GetValue()))
		case cruder.GT:
			stm.Where(general.SelfUnitsGT(conds.GetSelfUnits().GetValue()))
		case cruder.EQ:
			stm.Where(general.SelfUnitsEQ(conds.GetSelfUnits().GetValue()))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.TotalCommission != nil {
		incoming, err := decimal.NewFromString(conds.GetTotalCommission().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetTotalCommission().GetOp() {
		case cruder.LT:
			stm.Where(general.TotalCommissionLT(incoming))
		case cruder.GT:
			stm.Where(general.TotalCommissionGT(incoming))
		case cruder.EQ:
			stm.Where(general.TotalCommissionEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	if conds.SelfCommission != nil {
		incoming, err := decimal.NewFromString(conds.GetSelfCommission().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetSelfCommission().GetOp() {
		case cruder.LT:
			stm.Where(general.SelfCommissionLT(incoming))
		case cruder.GT:
			stm.Where(general.SelfCommissionGT(incoming))
		case cruder.EQ:
			stm.Where(general.SelfCommissionEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid general field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.General, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.General{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(general.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.General, error) {
	var info *ent.General
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.General.Query().Where(general.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.General, error) {
	var info *ent.General
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.General.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
