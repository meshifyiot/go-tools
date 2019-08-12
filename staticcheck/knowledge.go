package staticcheck

import (
	"reflect"

	"golang.org/x/tools/go/analysis"
	"github.com/meshifyiot/go-tools/internal/passes/buildssa"
	"github.com/meshifyiot/go-tools/ssa"
	"github.com/meshifyiot/go-tools/staticcheck/vrp"
)

var valueRangesAnalyzer = &analysis.Analyzer{
	Name: "vrp",
	Doc:  "calculate value ranges of functions",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		m := map[*ssa.Function]vrp.Ranges{}
		for _, ssafn := range pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs {
			vr := vrp.BuildGraph(ssafn).Solve()
			m[ssafn] = vr
		}
		return m, nil
	},
	Requires:   []*analysis.Analyzer{buildssa.Analyzer},
	ResultType: reflect.TypeOf(map[*ssa.Function]vrp.Ranges{}),
}
