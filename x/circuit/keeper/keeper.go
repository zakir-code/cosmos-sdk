package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/circuit/types"
)

// keeper definines the circuit module's keeper.
type Keeper struct {
	key       storetypes.StoreKey
	cdc       codec.Codec //used to marshall and unarshall structs to and from []byte
	authority []byte      /* string */
}

// contructs a new Circuit Keeper instance
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, authority []byte /* todo: byte or string */) Keeper {
	return Keeper{
		key:       storeKey,
		cdc:       cdc,
		authority: authority,
	}
}

func (k *Keeper) GetAuthority() []byte /* string */ {
	return k.authority
}

func (k *Keeper) GetPermissions(ctx sdk.Context, address []byte) (types.Permissions, error) {
	store := ctx.KVStore(k.key)

	bz := store.Get(address)

	perms := types.Permissions{}
	if err := k.cdc.Unmarshal(bz, &perms); err != nil {
		return types.Permissions{}, err
	}

	return perms, nil
}

func (k *Keeper) SetPermissions(ctx sdk.Context, address []byte, perms types.Permissions) error {
	store := ctx.KVStore(k.key)

	bz, err := k.cdc.Marshal(&perms)
	if err != nil {
		return err
	}

	key := types.CreateAddressPrefix(address)

	store.Set(key, bz)

	return nil
}

func (k *Keeper) IsMsgDisabled(ctx sdk.Context, msgURL string) bool {
	store := ctx.KVStore(k.key)
	return store.Has(types.CreateDisableMsgPrefix(msgURL))
}

func (k *Keeper) DisableMsg(ctx sdk.Context, msgURL string) {
	ctx.KVStore(k.key).Set(types.CreateDisableMsgPrefix(msgURL), []byte{})
}

func (k *Keeper) EnableMsg(ctx sdk.Context, msgURL string) {
	ctx.KVStore(k.key).Delete(types.CreateDisableMsgPrefix(msgURL))
}