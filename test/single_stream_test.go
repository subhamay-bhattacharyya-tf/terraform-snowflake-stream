// File: test/single_stream_test.go
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

// TestSingleStream tests creating a single stream via the module
func TestSingleStream(t *testing.T) {
	t.Parallel()

	retrySleep := 5 * time.Second
	unique := strings.ToUpper(random.UniqueId())
	streamName := fmt.Sprintf("TT_STREAM_%s", unique)
	tableName := fmt.Sprintf("TT_TABLE_%s", unique)

	tfDir := "../examples/simple-stream"

	// Create test database, schema, and table first
	db := openSnowflake(t)
	defer func() { _ = db.Close() }()

	setupTestTable(t, db, "TT_TEST_DB", "PUBLIC", tableName)
	defer cleanupTestTable(t, db, "TT_TEST_DB", "PUBLIC", tableName)

	streamConfigs := map[string]interface{}{
		"test_stream": map[string]interface{}{
			"name":              streamName,
			"database":         "TT_TEST_DB",
			"schema":           "PUBLIC",
			"table":            tableName,
			"append_only":      "false",
			"show_initial_rows": "false",
			"copy_grants":      false,
			"comment":          "Terratest single stream test",
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

	exists := streamExists(t, db, "TT_TEST_DB", "PUBLIC", streamName)
	require.True(t, exists, "Expected stream %q to exist in Snowflake", streamName)

	props := fetchStreamProps(t, db, "TT_TEST_DB", "PUBLIC", streamName)
	require.Equal(t, streamName, props.Name)
	require.Contains(t, props.Comment, "Terratest single stream test")
}
