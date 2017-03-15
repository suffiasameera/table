package main

import (
	"encoding/json"
	"errors"
	"fmt"


	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SampleChaincode example Sample Chaincode implementation
type SampleChaincode struct {
}

// Init create tables for tests
func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Create table one
	err := createTableOne(stub)
	if err != nil {
		return nil, errors.New("Error creating table one during init.")
	}

	return nil, nil
}

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

  if function == "insertTableOne" {
	return insertTableOne(stub, args)

}
	return nil, nil
}

func insertTableOne(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

      if len(args) < 5 {
			return nil, errors.New("insertTableOne failed. Must include 5 column values")
		}

		col1Val := args[0]

		col2Val := args[1]

		col3Val := args[2]

    col4Val := args[3]

		col5Val := args[4]

		var columns []*shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
    col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
    col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}

		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
    columns = append(columns, &col5)

		row := shim.Row{Columns: columns}
		ok, err := stub.InsertRow("tableOne", row)
		if err != nil {
			 fmt.Printf("insertTableOne operation failed. %s", err)
			 return nil, err
		}
		if !ok {
			fmt.Printf("insertTableOne operation failed. Row with given key already exists")
			return nil, err
		}
		return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {

    case "getRowTableOne":
  		if len(args) < 1 {
  		 fmt.Printf("getRowTableOne failed. Must include key values")
			 	return nil, errors.New("Expected atleast 1 arguments for query")
  		}


		col1Val := args[0]
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
		columns = append(columns, col1)

		row, err := stub.GetRow("tableOne", columns)
		if err != nil {
			fmt.Printf("getRowTableOne operation failed. %s", err)
			return nil, err
		}

		rowString := fmt.Sprintf("%s", row)
		return []byte(rowString), nil
		/*
			var columns []shim.Column
			var op = args[1]

			if( op == "1" ) {
			col1Val:=args[0]

			  		col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}

			      columns = append(columns, col1)

			} else if( op == "2" ) {
				col2Val:=args[0]

				  		col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}

				      columns = append(columns, col2)

			} else if ( op == "3" ) {
				col3Val:=args[0]

				  		col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}

				      columns = append(columns, col3)

			} else if ( op == "4" ) {
				col4Val:=args[0]

				  		col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}

				      columns = append(columns, col4)
			} else {
				col5Val:=args[0]

				  		col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}

				      columns = append(columns, col5)
			}



  		row, err := stub.GetRow("tableOne", columns)
  		if err != nil {
  			 fmt.Printf("getRowTableOne operation failed. %s", err)
				 return nil, err
  		}

  		rowString := fmt.Sprintf("%s", row)
  		return []byte(rowString), nil
*/

  case "getRowsTableOne":
  if len(args) < 1 {
    return nil, errors.New("getRowsTableOne failed. Must include at least 1 key values")
  }

  var columns []shim.Column

  col1Val := args[0]
  col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
  columns = append(columns, col1)

  if len(args) > 1 {
    col2Val := args[1]
    col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
    columns = append(columns, col2)
  }

  rowChannel, err := stub.GetRows("tableOne", columns)
  if err != nil {
    return nil, errors.New("getRowsTableOne operation failed. ")
  }

  var rows []shim.Row
  for {
    select {
    case row, ok := <-rowChannel:
      if !ok {
        rowChannel = nil
      } else {
        rows = append(rows, row)
      }
    }
    if rowChannel == nil {
      break
    }
  }

  jsonRows, err := json.Marshal(rows)
  if err != nil {
    return nil, errors.New("getRowsTableOne operation failed. Error marshaling JSON:")
  }

  return jsonRows, nil
}
return nil, nil
}
func main() {
	err := shim.Start(new(SampleChaincode))
	if err != nil {
		fmt.Printf("Error starting Sample chaincode: %s", err)
	}
}

func createTableOne(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "colOneTableOne",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "colTwoTableOne",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "colThreeTableOne",
		Type: shim.ColumnDefinition_STRING, Key: true}
    columnFourTableOneDef := shim.ColumnDefinition{Name: "colFourTableOne",
  		Type: shim.ColumnDefinition_STRING, Key: true}
    columnFiveTableOneDef := shim.ColumnDefinition{Name: "colFiveTableOne",
    		Type: shim.ColumnDefinition_STRING, Key: true}

	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
  columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
  columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)

	return stub.CreateTable("tableOne", columnDefsTableOne)
}
