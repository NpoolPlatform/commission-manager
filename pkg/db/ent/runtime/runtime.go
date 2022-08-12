// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/schema"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	detailMixin := schema.Detail{}.Mixin()
	detail.Policy = privacy.NewPolicies(detailMixin[0], schema.Detail{})
	detail.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := detail.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	detailMixinFields0 := detailMixin[0].Fields()
	_ = detailMixinFields0
	detailFields := schema.Detail{}.Fields()
	_ = detailFields
	// detailDescCreatedAt is the schema descriptor for created_at field.
	detailDescCreatedAt := detailMixinFields0[0].Descriptor()
	// detail.DefaultCreatedAt holds the default value on creation for the created_at field.
	detail.DefaultCreatedAt = detailDescCreatedAt.Default.(func() uint32)
	// detailDescUpdatedAt is the schema descriptor for updated_at field.
	detailDescUpdatedAt := detailMixinFields0[1].Descriptor()
	// detail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	detail.DefaultUpdatedAt = detailDescUpdatedAt.Default.(func() uint32)
	// detail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	detail.UpdateDefaultUpdatedAt = detailDescUpdatedAt.UpdateDefault.(func() uint32)
	// detailDescDeletedAt is the schema descriptor for deleted_at field.
	detailDescDeletedAt := detailMixinFields0[2].Descriptor()
	// detail.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	detail.DefaultDeletedAt = detailDescDeletedAt.Default.(func() uint32)
	// detailDescAppID is the schema descriptor for app_id field.
	detailDescAppID := detailFields[1].Descriptor()
	// detail.DefaultAppID holds the default value on creation for the app_id field.
	detail.DefaultAppID = detailDescAppID.Default.(func() uuid.UUID)
	// detailDescUserID is the schema descriptor for user_id field.
	detailDescUserID := detailFields[2].Descriptor()
	// detail.DefaultUserID holds the default value on creation for the user_id field.
	detail.DefaultUserID = detailDescUserID.Default.(func() uuid.UUID)
	// detailDescDirectContributorID is the schema descriptor for direct_contributor_id field.
	detailDescDirectContributorID := detailFields[3].Descriptor()
	// detail.DefaultDirectContributorID holds the default value on creation for the direct_contributor_id field.
	detail.DefaultDirectContributorID = detailDescDirectContributorID.Default.(func() uuid.UUID)
	// detailDescGoodID is the schema descriptor for good_id field.
	detailDescGoodID := detailFields[4].Descriptor()
	// detail.DefaultGoodID holds the default value on creation for the good_id field.
	detail.DefaultGoodID = detailDescGoodID.Default.(func() uuid.UUID)
	// detailDescOrderID is the schema descriptor for order_id field.
	detailDescOrderID := detailFields[5].Descriptor()
	// detail.DefaultOrderID holds the default value on creation for the order_id field.
	detail.DefaultOrderID = detailDescOrderID.Default.(func() uuid.UUID)
	// detailDescSelfOrder is the schema descriptor for self_order field.
	detailDescSelfOrder := detailFields[6].Descriptor()
	// detail.DefaultSelfOrder holds the default value on creation for the self_order field.
	detail.DefaultSelfOrder = detailDescSelfOrder.Default.(bool)
	// detailDescPaymentID is the schema descriptor for payment_id field.
	detailDescPaymentID := detailFields[7].Descriptor()
	// detail.DefaultPaymentID holds the default value on creation for the payment_id field.
	detail.DefaultPaymentID = detailDescPaymentID.Default.(func() uuid.UUID)
	// detailDescCoinTypeID is the schema descriptor for coin_type_id field.
	detailDescCoinTypeID := detailFields[8].Descriptor()
	// detail.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	detail.DefaultCoinTypeID = detailDescCoinTypeID.Default.(func() uuid.UUID)
	// detailDescPaymentCoinTypeID is the schema descriptor for payment_coin_type_id field.
	detailDescPaymentCoinTypeID := detailFields[9].Descriptor()
	// detail.DefaultPaymentCoinTypeID holds the default value on creation for the payment_coin_type_id field.
	detail.DefaultPaymentCoinTypeID = detailDescPaymentCoinTypeID.Default.(func() uuid.UUID)
	// detailDescPaymentCoinUsdCurrency is the schema descriptor for payment_coin_usd_currency field.
	detailDescPaymentCoinUsdCurrency := detailFields[10].Descriptor()
	// detail.DefaultPaymentCoinUsdCurrency holds the default value on creation for the payment_coin_usd_currency field.
	detail.DefaultPaymentCoinUsdCurrency = detailDescPaymentCoinUsdCurrency.Default.(decimal.Decimal)
	// detailDescUnits is the schema descriptor for units field.
	detailDescUnits := detailFields[11].Descriptor()
	// detail.DefaultUnits holds the default value on creation for the units field.
	detail.DefaultUnits = detailDescUnits.Default.(uint32)
	// detailDescAmount is the schema descriptor for amount field.
	detailDescAmount := detailFields[12].Descriptor()
	// detail.DefaultAmount holds the default value on creation for the amount field.
	detail.DefaultAmount = detailDescAmount.Default.(decimal.Decimal)
	// detailDescUsdAmount is the schema descriptor for usd_amount field.
	detailDescUsdAmount := detailFields[13].Descriptor()
	// detail.DefaultUsdAmount holds the default value on creation for the usd_amount field.
	detail.DefaultUsdAmount = detailDescUsdAmount.Default.(decimal.Decimal)
	// detailDescCommission is the schema descriptor for commission field.
	detailDescCommission := detailFields[14].Descriptor()
	// detail.DefaultCommission holds the default value on creation for the commission field.
	detail.DefaultCommission = detailDescCommission.Default.(decimal.Decimal)
	// detailDescID is the schema descriptor for id field.
	detailDescID := detailFields[0].Descriptor()
	// detail.DefaultID holds the default value on creation for the id field.
	detail.DefaultID = detailDescID.Default.(func() uuid.UUID)
	generalMixin := schema.General{}.Mixin()
	general.Policy = privacy.NewPolicies(generalMixin[0], schema.General{})
	general.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := general.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	generalMixinFields0 := generalMixin[0].Fields()
	_ = generalMixinFields0
	generalFields := schema.General{}.Fields()
	_ = generalFields
	// generalDescCreatedAt is the schema descriptor for created_at field.
	generalDescCreatedAt := generalMixinFields0[0].Descriptor()
	// general.DefaultCreatedAt holds the default value on creation for the created_at field.
	general.DefaultCreatedAt = generalDescCreatedAt.Default.(func() uint32)
	// generalDescUpdatedAt is the schema descriptor for updated_at field.
	generalDescUpdatedAt := generalMixinFields0[1].Descriptor()
	// general.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	general.DefaultUpdatedAt = generalDescUpdatedAt.Default.(func() uint32)
	// general.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	general.UpdateDefaultUpdatedAt = generalDescUpdatedAt.UpdateDefault.(func() uint32)
	// generalDescDeletedAt is the schema descriptor for deleted_at field.
	generalDescDeletedAt := generalMixinFields0[2].Descriptor()
	// general.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	general.DefaultDeletedAt = generalDescDeletedAt.Default.(func() uint32)
	// generalDescAppID is the schema descriptor for app_id field.
	generalDescAppID := generalFields[1].Descriptor()
	// general.DefaultAppID holds the default value on creation for the app_id field.
	general.DefaultAppID = generalDescAppID.Default.(func() uuid.UUID)
	// generalDescUserID is the schema descriptor for user_id field.
	generalDescUserID := generalFields[2].Descriptor()
	// general.DefaultUserID holds the default value on creation for the user_id field.
	general.DefaultUserID = generalDescUserID.Default.(func() uuid.UUID)
	// generalDescGoodID is the schema descriptor for good_id field.
	generalDescGoodID := generalFields[3].Descriptor()
	// general.DefaultGoodID holds the default value on creation for the good_id field.
	general.DefaultGoodID = generalDescGoodID.Default.(func() uuid.UUID)
	// generalDescCoinTypeID is the schema descriptor for coin_type_id field.
	generalDescCoinTypeID := generalFields[4].Descriptor()
	// general.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	general.DefaultCoinTypeID = generalDescCoinTypeID.Default.(func() uuid.UUID)
	// generalDescTotalUnits is the schema descriptor for total_units field.
	generalDescTotalUnits := generalFields[5].Descriptor()
	// general.DefaultTotalUnits holds the default value on creation for the total_units field.
	general.DefaultTotalUnits = generalDescTotalUnits.Default.(uint32)
	// generalDescSelfUnits is the schema descriptor for self_units field.
	generalDescSelfUnits := generalFields[6].Descriptor()
	// general.DefaultSelfUnits holds the default value on creation for the self_units field.
	general.DefaultSelfUnits = generalDescSelfUnits.Default.(uint32)
	// generalDescTotalAmount is the schema descriptor for total_amount field.
	generalDescTotalAmount := generalFields[7].Descriptor()
	// general.DefaultTotalAmount holds the default value on creation for the total_amount field.
	general.DefaultTotalAmount = generalDescTotalAmount.Default.(decimal.Decimal)
	// generalDescSelfAmount is the schema descriptor for self_amount field.
	generalDescSelfAmount := generalFields[8].Descriptor()
	// general.DefaultSelfAmount holds the default value on creation for the self_amount field.
	general.DefaultSelfAmount = generalDescSelfAmount.Default.(decimal.Decimal)
	// generalDescTotalCommission is the schema descriptor for total_commission field.
	generalDescTotalCommission := generalFields[9].Descriptor()
	// general.DefaultTotalCommission holds the default value on creation for the total_commission field.
	general.DefaultTotalCommission = generalDescTotalCommission.Default.(decimal.Decimal)
	// generalDescSelfCommission is the schema descriptor for self_commission field.
	generalDescSelfCommission := generalFields[10].Descriptor()
	// general.DefaultSelfCommission holds the default value on creation for the self_commission field.
	general.DefaultSelfCommission = generalDescSelfCommission.Default.(decimal.Decimal)
	// generalDescID is the schema descriptor for id field.
	generalDescID := generalFields[0].Descriptor()
	// general.DefaultID holds the default value on creation for the id field.
	general.DefaultID = generalDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)
