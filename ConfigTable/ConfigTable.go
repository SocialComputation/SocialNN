package main


import (
	
	"github.com/emer/emergent/patgen"
	"github.com/emer/etable/etable"
	"github.com/emer/etable/etensor"
	_ "github.com/emer/etable/etview" // include to get gui views

	
)

func main() {
	TheSim.New()
	TheSim.ConfigPats()
}

// Defining Pats as a Table type.  Pats is the variable pointing to the data file you are creating.  
// Could define multiple data files here.
type Sim struct {
	Pats         *etable.Table     `view:"no-inline" desc:"the training patterns to use"`
	}



// TheSim is the overall state for this simulation
var TheSim Sim

// New creates new blank elements and initializes defaults for the Sim. If you define multiple data files 
// in the struct, then you want to list all of them here.  
func (ss *Sim) New() {
	
	ss.Pats = &etable.Table{}
}

// This defines the structure of the datatable and the column headers and then writes out a .tsv file
// with the correct labels for the datatable

func (ss *Sim) ConfigPats() {
	dt := ss.Pats
	dt.SetMetaData("name", "TrainPats")
	dt.SetMetaData("desc", "Training patterns")
	sch := etable.Schema{
		{"Name", etensor.STRING, nil, nil},
		{"Features", etensor.FLOAT32, []int{1, 13}, []string{"Y", "X"}},
		{"Mammals", etensor.FLOAT32, []int{1, 2}, []string{"Y", "X"}},
		{"Species", etensor.FLOAT32, []int{1, 5}, []string{"Y", "X"}},

	}
	// creates 25 rows for the datatable
	dt.SetFromSchema(sch, 25)

	// patgen command is used to create random data, which can be easily replaced in Excel or a text editor
	// dt.Cols[x} indexes the column, skipping the Name column [0].
	// After dt.Cols[x] the first number is the number of randomly generated data points, the second is the specific 
	// value that is generated and the third is the number for the remaining data points. 
	patgen.PermutedBinaryRows(dt.Cols[1], 6, 1, 0)
	patgen.PermutedBinaryRows(dt.Cols[2], 1, 1, 0)
	patgen.PermutedBinaryRows(dt.Cols[3], 1, 1, 0)
	
	
	// saves the data file to disk in Tab delimited format, with emergent type headers.
	// obviously, you can pick your own data file name
	// could alternatively used COMMA delimited format.
	dt.SaveCSV("AnimalClassification.tsv", etable.Tab, etable.Headers)
}


