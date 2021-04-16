package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/v0.5/types"
	"github.com/sentinel-official/hub/v0.5/x/subscription/types"
)

func (k Keeper) SetCount(ctx sdk.Context, count uint64) {
	key := types.CountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetCount(ctx sdk.Context) (count uint64) {
	store := k.Store(ctx)

	key := types.CountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	key := types.SubscriptionKey(subscription.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(subscription)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription types.Subscription, found bool) {
	store := k.Store(ctx)

	key := types.SubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &subscription)
	return subscription, true
}

func (k Keeper) GetSubscriptions(ctx sdk.Context, skip, limit int) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Subscription
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) {
	key := types.SubscriptionForNodeKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) HasSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool {
	key := types.SubscriptionForNodeKey(address, id)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k Keeper) GetSubscriptionsForNode(ctx sdk.Context, address hub.NodeAddress, skip, limit int) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetSubscriptionForNodeKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSubscription(ctx, types.IDFromSubscriptionForNodeKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetSubscriptionForPlan(ctx sdk.Context, plan, id uint64) {
	key := types.SubscriptionForPlanKey(plan, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) HasSubscriptionForPlan(ctx sdk.Context, plan, id uint64) bool {
	key := types.SubscriptionForPlanKey(plan, id)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k Keeper) GetSubscriptionsForPlan(ctx sdk.Context, plan uint64, skip, limit int) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, sdk.Uint64ToBigEndian(plan)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSubscription(ctx, types.IDFromSubscriptionForPlanKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetActiveSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.ActiveSubscriptionForAddressKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteActiveSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	store := k.Store(ctx)

	key := types.ActiveSubscriptionForAddressKey(address, id)
	store.Delete(key)
}

func (k Keeper) GetActiveSubscriptionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActiveSubscriptionForAddressKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSubscription(ctx, types.IDFromStatusSubscriptionForAddressKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetInactiveSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.InactiveSubscriptionForAddressKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInactiveSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	store := k.Store(ctx)

	key := types.InactiveSubscriptionForAddressKey(address, id)
	store.Delete(key)
}

func (k Keeper) GetInactiveSubscriptionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetInactiveSubscriptionForAddressKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSubscription(ctx, types.IDFromStatusSubscriptionForAddressKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) GetSubscriptionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActiveSubscriptionForAddressKeyPrefix(address)),
			sdk.KVStorePrefixIterator(store, types.GetInactiveSubscriptionForAddressKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSubscription(ctx, types.IDFromStatusSubscriptionForAddressKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetInactiveSubscriptionAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.InactiveSubscriptionAtKey(at, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInactiveSubscriptionAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.InactiveSubscriptionAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) IterateInactiveSubscriptions(ctx sdk.Context, end time.Time, fn func(index int, item types.Subscription) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.InactiveSubscriptionAtKeyPrefix, sdk.PrefixEndBytes(types.GetInactiveSubscriptionAtKeyPrefix(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		subscription, _ := k.GetSubscription(ctx, types.IDFromInactiveSubscriptionAtKey(iter.Key()))
		if stop := fn(i, subscription); stop {
			break
		}
		i++
	}
}
