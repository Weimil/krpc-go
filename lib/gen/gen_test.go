package gen

import (
	"bytes"
	"go/format"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/require"
	"github.com/weimil/krpc-go/types"
)

func TestGenerateProcedure(t *testing.T) {
	tests := []struct {
		name        string
		procedure   *types.Procedure
		expectedOut string
	}{
		{
			name: "basic procedure",
			procedure: &types.Procedure{
				Name:          "MyProcedure",
				Documentation: "<summary>Test procedure generation.</summary>",
				Parameters: []*types.Parameter{
					{
						Name: "param1",
						Type: &types.Type{
							Code: types.Type_UINT64,
						},
					},
					{
						Name: "param2",
						Type: &types.Type{
							Code: types.Type_STRING,
						},
					},
				},
				ReturnType: &types.Type{
					Code: types.Type_BOOL,
				},
				GameScenes: []types.Procedure_GameScene{types.Procedure_FLIGHT},
			},
			expectedOut: testProcedure,
		},
		{
			name: "class setter",
			procedure: &types.Procedure{
				Name:          "MyClass_set_MyProperty",
				Documentation: "<summary>Test class setter generation.</summary>",
				Parameters: []*types.Parameter{
					{
						Name: "this",
						Type: &types.Type{
							Code:    types.Type_CLASS,
							Service: "MyService",
							Name:    "MyClass",
						},
					},
					{
						Name: "param1",
						Type: &types.Type{
							Code: types.Type_TUPLE,
							Types: []*types.Type{
								{
									Code: types.Type_STRING,
								},
								{
									Code: types.Type_UINT64,
								},
							},
						},
					},
				},
				ReturnType: &types.Type{
					Code: types.Type_NONE,
				},
			},
			expectedOut: testClassSetter,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			expectedOut, err := format.Source([]byte(tc.expectedOut))
			require.NoError(t, err)

			f := jen.NewFile("gentest")
			require.NoError(t, GenerateProcedure(f, "MyService", tc.procedure))

			var out bytes.Buffer
			require.NoError(t, f.Render(&out))
			require.Equal(t, string(expectedOut), out.String())
		})
	}
}

const testClass = `
package gentest

import (
	krpcgo "github.com/weimil/krpc-go"
	service "github.com/weimil/krpc-go/lib/service"
)

// Test - a test class.
type Test struct {
	service.BaseClass
}

// NewTest creates a new Test.
func NewTest(id uint64, client *krpcgo.KRPCClient) *Test {
	c := &Test{BaseClass: service.BaseClass{Client: client}}
	c.SetID(id)
	return c
}
`

func TestGenerateClass(t *testing.T) {
	expectedOut, err := format.Source([]byte(testClass))
	require.NoError(t, err)

	class := &types.Class{
		Name:          "Test",
		Documentation: "<summary>A test class.</summary>",
	}

	f := jen.NewFile("gentest")
	require.NoError(t, GenerateClass(f, class))

	var out bytes.Buffer
	require.NoError(t, f.Render(&out))
	require.Equal(t, string(expectedOut), out.String())
}

const testEnum = `
package gentest


// Test - a test enum.
type Test int32

const (
	// The first enum value.
	Test_One Test = 1
	// The second enum value.
	Test_Two Test = 2
	// The third enum value.
	Test_Three Test = 3
)

func (v Test) Value() int32 {
	return int32(v)
}
func (v *Test) SetValue(val int32) {
	*v = Test(val)
}
`

func TestGenerateEnum(t *testing.T) {
	expectedOut, err := format.Source([]byte(testEnum))
	require.NoError(t, err)

	enum := &types.Enumeration{
		Name:          "Test",
		Documentation: "<summary>A test enum.</summary>",
		Values: []*types.EnumerationValue{
			{
				Name:          "One",
				Value:         1,
				Documentation: "<summary>The first enum value.</summary>",
			},
			{
				Name:          "Two",
				Value:         2,
				Documentation: "<summary>The second enum value.</summary>",
			},
			{
				Name:          "Three",
				Value:         3,
				Documentation: "<summary>The third enum value.</summary>",
			},
		},
	}
	f := jen.NewFile("gentest")
	require.NoError(t, GenerateEnum(f, enum))

	var out bytes.Buffer
	require.NoError(t, f.Render(&out))
	require.Equal(t, string(expectedOut), out.String())
}

const testException = `
package gentest


// ErrTest - the exception generating code is being tested.
type ErrTest struct {
	msg string
}

// NewErrTest creates a new ErrTest.
func NewErrTest(msg string) *ErrTest {
	return &ErrTest{msg: msg}
}

// Error returns a human-readable error.
func (err ErrTest) Error() string {
	return err.msg
}
`

func TestGenerateException(t *testing.T) {
	expectedOut, err := format.Source([]byte(testException))
	require.NoError(t, err)

	exception := &types.Exception{
		Name:          "TestException",
		Documentation: "<summary>The exception generating code is being tested.</summary>",
	}
	f := jen.NewFile("gentest")
	require.NoError(t, GenerateException(f, exception))

	var out bytes.Buffer
	require.NoError(t, f.Render(&out))
	require.Equal(t, string(expectedOut), out.String())
}
