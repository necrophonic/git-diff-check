	"fmt"
	"os/exec"
	OK              bool

		if ok != tc.OK {
			t.Errorf("Got result %v, but got %v", ok, tc.OK)
	t.Logf("    %s should equal %s", field, expected)
	t.Logf("    %s should equal %d", field, expected)
	testCase{
		Name:            "a totally fine file",
		OK:              true,
		ExpectedReports: nil,
		Patch: []byte(`
diff --git a/diffcheck/README.md b/diffcheck/README.md
new file mode 100644
index 0000000..e69de29
			`),
	},
		OK:   false,