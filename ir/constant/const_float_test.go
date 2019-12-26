package constant_test

import (
	"testing"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func TestNewFloatFromStringForPPCFP128(t *testing.T) {
	golden := []string{
		"0xM00000000000000000000000000000000",
		"0xM3DF00000000000000000000000000000",
		"0xM3FF00000000000000000000000000000",
		"0xM40000000000000000000000000000000",
		"0xM400C0000000000300000000010000000",
		"0xM400F000000000000BCB0000000000000",
		"0xM403B0000000000000000000000000000",
		"0xM405EDA5E353F7CEE0000000000000000",
		"0xM4093B400000000000000000000000000",
		"0xM41F00000000000000000000000000000",
		"0xM4D436562A0416DE00000000000000000",
		"0xM80000000000000000000000000000000",
		"0xM818F2887B9295809800000000032D000",
		"0xMC00547AE147AE1483CA47AE147AE147A",
	}
	for _, g := range golden {
		f, err := constant.NewFloatFromString(types.PPC_FP128, g)
		if err != nil {
			t.Errorf("%q: unable to create new float from string: %v", g, err)
			continue
		}
		got := f.Ident()
		want := g
		if want != got {
			t.Errorf("%q: floating-point value string mismatch; expected %q, got %q", g, want, got)
		}
	}
}

func TestNewFloatFromStringForFP128(t *testing.T) {
	golden := []struct{
		in   string
		want string
	}{
		{in: "0xL0", want: "0xL00000000000000000000000000000000"},
		{in: "0xL00000000000000000000000000000000", want: "0xL00000000000000000000000000000000"},
		{in: "0xL00000000000000000001000000000000", want: "0xL00000000000000000001000000000000"},
		{in: "0xL00000000000000003FFF000000000000", want: "0xL00000000000000003FFF000000000000"},
		{in: "0xL00000000000000003FFF000001000000", want: "0xL00000000000000003FFF000001000000"},
		{in: "0xL00000000000000003FFF000002000000", want: "0xL00000000000000003FFF000002000000"},
		{in: "0xL00000000000000004000000000000000", want: "0xL00000000000000004000000000000000"},
		{in: "0xL00000000000000004001400000000000", want: "0xL00000000000000004001400000000000"},
		{in: "0xL00000000000000004004C00000000000", want: "0xL00000000000000004004C00000000000"},
		{in: "0xL00000000000000004201000000000000", want: "0xL00000000000000004201000000000000"},
		{in: "0xL00000000000000005001000000000000", want: "0xL00000000000000005001000000000000"},
		{in: "0xL00000000000000007FFF000000000000", want: "0xL00000000000000007FFF000000000000"},
		{in: "0xL00000000000000007FFF800000000000", want: "0xL00000000000000007FFF800000000000"},
		{in: "0xL00000000000000008000000000000000", want: "0xL00000000000000008000000000000000"},
		{in: "0xL00000000000000013FFF000000000000", want: "0xL00000000000000013FFF000000000000"},
		{in: "0xL00000000000000018000000000000000", want: "0xL00000000000000018000000000000000"},
		{in: "0xL000FFFFF00000000000FFFFF00000000", want: "0xL000FFFFF00000000000FFFFF00000000"},
		{in: "0xL00FF00FF00FF00FF00FF00FF00FF00FF", want: "0xL00FF00FF00FF00FF00FF00FF00FF00FF"},
		{in: "0xL01", want: "0xL00000000000000000000000000000001"},
		{in: "0xL08000000000000003FFF000000000000", want: "0xL08000000000000003FFF000000000000"},
		{in: "0xL300000000000000040089CA8F5C28F5C", want: "0xL300000000000000040089CA8F5C28F5C"},
		{in: "0xL5000000000000000400E0C26324C8366", want: "0xL5000000000000000400E0C26324C8366"},
		{in: "0xL8000000000000000400A24E2E147AE14", want: "0xL8000000000000000400A24E2E147AE14"},
		{in: "0xL999999999999999A3FFB999999999999", want: "0xL999999999999999A3FFB999999999999"},
		{in: "0xLEB851EB851EB851F400091EB851EB851", want: "0xLEB851EB851EB851F400091EB851EB851"},
		{in: "0xLF000000000000000400808AB851EB851", want: "0xLF000000000000000400808AB851EB851"},
		{in: "0xLF8F8F8F8F8F8F8F8F8F8F8F8F8F8F8F8", want: "0xLF8F8F8F8F8F8F8F8F8F8F8F8F8F8F8F8"},
	}
	for _, g := range golden {
		f, err := constant.NewFloatFromString(types.FP128, g.in)
		if err != nil {
			t.Errorf("%q: unable to create new float from string: %v", g.in, err)
			continue
		}
		got := f.Ident()
		if g.want != got {
			t.Errorf("%q: floating-point value string mismatch; expected %q, got %q", g.in, g.want, got)
		}
	}
}
