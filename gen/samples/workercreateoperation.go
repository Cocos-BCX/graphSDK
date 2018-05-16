//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeWorkerCreate

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeWorkerCreate] =
		sampleDataWorkerCreateOperation
}

var sampleDataWorkerCreateOperation = `{
  "daily_pay": 3966,
  "fee": {
    "amount": 200000,
    "asset_id": "1.3.0"
  },
  "initializer": [
    1,
    {
      "pay_vesting_period_days": 0
    }
  ],
  "name": "Debian/Ubuntu-based PPA",
  "owner": "1.2.6004",
  "url": "https://bitsharestalk.org/index.php/topic,19485.msg250031.html",
  "work_begin_date": "2015-10-28T00:00:00",
  "work_end_date": "2016-01-31T00:00:00"
}`

//end of file