// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/types
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/querier
package node

import (
	"github.com/sentinel-official/hub/x/dvpn/node/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node/querier"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

const (
	Codespace               = types.Codespace
	EventTypeSetNode        = types.EventTypeSetNode
	EventTypeUpdateNode     = types.EventTypeUpdateNode
	EventTypeSetNodeStatus  = types.EventTypeSetNodeStatus
	AttributeKeyProvider    = types.AttributeKeyProvider
	AttributeKeyAddress     = types.AttributeKeyAddress
	AttributeKeyStatus      = types.AttributeKeyStatus
	ModuleName              = types.ModuleName
	ParamsSubspace          = types.ParamsSubspace
	QuerierRoute            = types.QuerierRoute
	CategoryUnknown         = types.CategoryUnknown
	DefaultInactiveDuration = types.DefaultInactiveDuration
	QueryNode               = types.QueryNode
	QueryNodes              = types.QueryNodes
	QueryNodesForProvider   = types.QueryNodesForProvider
)

var (
	// functions aliases
	RegisterCodec                  = types.RegisterCodec
	ErrorMarshal                   = types.ErrorMarshal
	ErrorUnmarshal                 = types.ErrorUnmarshal
	ErrorUnknownMsgType            = types.ErrorUnknownMsgType
	ErrorUnknownQueryType          = types.ErrorUnknownQueryType
	ErrorInvalidField              = types.ErrorInvalidField
	ErrorProviderDoesNotExist      = types.ErrorProviderDoesNotExist
	ErrorDuplicateNode             = types.ErrorDuplicateNode
	ErrorNodeDoesNotExist          = types.ErrorNodeDoesNotExist
	ErrorCanNotUpdate              = types.ErrorCanNotUpdate
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	NodeKey                        = types.NodeKey
	NodeForProviderKeyPrefix       = types.NodeForProviderKeyPrefix
	NodeForProviderKey             = types.NodeForProviderKey
	NewMsgRegisterNode             = types.NewMsgRegisterNode
	NewMsgUpdateNode               = types.NewMsgUpdateNode
	NewMsgSetNodeStatus            = types.NewMsgSetNodeStatus
	NodeCategoryFromString         = types.NodeCategoryFromString
	NewParams                      = types.NewParams
	DefaultParams                  = types.DefaultParams
	ParamsKeyTable                 = types.ParamsKeyTable
	NewQueryNodeParams             = types.NewQueryNodeParams
	NewQueryNodesForProviderParams = types.NewQueryNodesForProviderParams
	NewKeeper                      = keeper.NewKeeper
	Querier                        = querier.Querier

	// variable aliases
	ModuleCdc           = types.ModuleCdc
	RouterKey           = types.RouterKey
	StoreKey            = types.StoreKey
	NodeKeyPrefix       = types.NodeKeyPrefix
	KeyInactiveDuration = types.KeyInactiveDuration
)

type (
	GenesisState                = types.GenesisState
	MsgRegisterNode             = types.MsgRegisterNode
	MsgUpdateNode               = types.MsgUpdateNode
	MsgSetNodeStatus            = types.MsgSetNodeStatus
	NodeCategory                = types.NodeCategory
	Node                        = types.Node
	Nodes                       = types.Nodes
	Params                      = types.Params
	QueryNodeParams             = types.QueryNodeParams
	QueryNodesForProviderParams = types.QueryNodesForProviderParams
	Keeper                      = keeper.Keeper
)
