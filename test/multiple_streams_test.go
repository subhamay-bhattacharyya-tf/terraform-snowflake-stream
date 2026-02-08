// File: test/multiple_streams_test.go
package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

// TestMultipleStreams tests creating multiple streams via the module
func TestMultipleStreams(t *testing.T) {
	t.Parallel()

	retrySleep := 5 * time.Second
	unique := strings.ToUpper(random.UniqueId())

	stream1Name := fmt.Sprintf("TT_ORDERS_%s", unique)
	stream2Name := fmt.Sprintf("TT_CUSTOMERS_%s", unique)
	stream3Name := fmt.Sprintf("TT_PRODUCTS_%s", unique)

	table1Name := fmt.Sprintf("TT_ORDERS_TBL_%s", unique)
	table2Name := fmt.Sprintf("TT_CUSTOMERS_TBL_%s", unique)
	table3Name := fmt.Sprintf("TT_PRODUCTS_TBL_%s", unique)

	tfDir := "../examples/multiple-streams"

	// Create test database, schema, and tables first
	db := openSnowflake(t)
	defer func() { _ = db.Close() }()

	setupTestTable(t, db, "TT_TEST_DB", "PUBLIC", table1Name)
	setupTestTable(t, db, "TT_TEST_DB", "PUBLIC", table2Name)
	setupTestTable(t, db, "TT_TEST_DB", "PUBLIC", table3Name)
	defer cleanupTestTable(t, db, "TT_TEST_DB", "PUBLIC", table1Name)
	defer cleanupTestTable(t, db, "TT_TEST_DB", "PUBLIC", table2Name)
	defer cleanupTestTable(t, db, "TT_TEST_DB", "PUBLIC", table3Name)

	streamConfigs := map[string]interface{}{
		"orders_stream": map[string]interface{}{
			"name":              stream1Name,
			"database":         "TT_TEST_DB",
			"schema":           "PUBLIC",
			"table":            table1Name,
			"append_only":      "false",
			"show_initial_rows": "false",
			"copy_grants":      false,
			"comment":          "Terratest orders stream",
		},
		"customers_stream": map[string]interface{}{
			"name":              stream2Name,
			"database":         "TT_TEST_DB",
			"schema":           "PUBLIC",
			"table":            table2Name,
			"append_only":      "true",
			"show_initial_rows": "false",
			"copy_grants":      false,
			"comment":          "Terratest customers stream",
		},
		"products_stream": map[string]interface{}{
			"name":              stream3Name,
			"database":         "TT_TEST_DB",
			"schema":           "PUBLIC",
			"table":            table3Name,
			"append_only":      "false",
			"show_initial_rows": "true",
			"copy_grants":      false,
			"comment":          "Terratest products stream",
		},
	}

	tfOptions := &terraform.Options{
		TerraformDir: tfDir,
		NoColor:      true,
		Vars: map[string]interface{}{
			"stream_configs":              streamConfigs,
			"snowflake_organization_name": os.Getenv("SNOWFLAKE_ORGANIZATION_NAME"),
			"snowflake_account_name":      os.Getenv("SNOWFLAKE_ACCOUNT_NAME"),
			"snowflake_user":              os.Getenv("SNOWFLAKE_USER"),
			"snowflake_role":              os.Getenv("SNOWFLAKE_ROLE"),
			"snowflake_private_key":       os.Getenv("SNOWFLAKE_PRIVATE_KEY"),
		},
	}

	defer terraform.Destroy(t, tfOptions)
	terraform.InitAndApply(t, tfOptions)

	time.Sleep(retrySleep)

	// Verify all three streams exist
	for _, streamName := range []string{stream1Name, stream2Name, stream3Name} {
		exists := streamExists(t, db, "TT_TEST_DB", "PUBLIC", streamName)
		require.True(t, exists, "Expected stream %q to exist in Snowflake", streamName)
	}

	// Verify properties of products stream (has show_initial_rows = true)
	props := fetchStreamProps(t, db, "TT_TEST_DB", "PUBLIC", stream3Name)
	require.Equal(t, stream3Name, props.Name)
	require.Contains(t, props.Comment, "Terratest products stream")
}
