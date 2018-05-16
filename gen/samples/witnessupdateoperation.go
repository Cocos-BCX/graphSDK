//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeWitnessUpdate

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeWitnessUpdate] =
		sampleDataWitnessUpdateOperation
}

var sampleDataWitnessUpdateOperation = `{
  "fee": {
    "amount": 4000000,
    "asset_id": "1.3.0"
  },
  "new_signing_key": "BTS7V2jfYZEToTjGFndctgTPipMuMcobqzwsd8RQv77c5Vzcexgt6",
  "new_url": "-",
  "witness": "1.6.23",
  "witness_account": "1.2.3284"
}`

//end of file