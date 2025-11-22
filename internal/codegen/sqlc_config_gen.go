package codegen

import (
	"google.golang.org/protobuf/types/pluginpb"
	"gopkg.in/yaml.v3"
)

func SqlcConfigGen(domains map[string]bool) *pluginpb.CodeGeneratorResponse_File {
	if len(domains) == 0 {
		return nil
	}

	// Remove duplicates
	var sqlConfigs []SqlConfig

	for domain := range domains {
		sqlConfigs = append(sqlConfigs, generateSqlConfig(domain))
	}

	config := SqlcConfig{
		Version: "2",
		Sql:     sqlConfigs,
	}

	configContent, err := generateYamlContent(config)
	if err != nil {
		return nil
	}

	return &pluginpb.CodeGeneratorResponse_File{
		Name:    &[]string{"sqlc.yaml"}[0],
		Content: &configContent,
	}
}

func generateSqlConfig(domain string) SqlConfig {
	return SqlConfig{
		Engine:  "postgresql",
		Name:    domain,
		Queries: "./data/queries/" + domain,
		Schema:  "./data/migrations/*.sql",
		Gen: GenConfig{
			Go: GoConfig{
				Package:                  "gendb_" + domain,
				Out:                      "./internal/gendb/" + domain,
				EmitResultStructPointers: true,
				EmitParamsStructPointers: true,
				EmitPreparedQueries:      true,
				EmitExportedQueries:      true,
				EmitInterface:            true,
				EmitJsonTags:             true,
				Overrides: []Override{
					{
						DbType: "uuid.UUID",
						GoType: GoType{
							Import:  "github.com/gofrs/uuid/v5",
							Package: "uuid",
							Type:    "UUID",
						},
					},
					{
						DbType: "uuid.NullUUID",
						GoType: GoType{
							Import:  "github.com/gofrs/uuid/v5",
							Package: "uuid",
							Type:    "NullUUID",
						},
					},
				},
			},
		},
	}
}

func generateYamlContent(config SqlcConfig) (string, error) {
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}
