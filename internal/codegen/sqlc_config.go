package codegen

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"gopkg.in/yaml.v3"
)

type SqlcConfig struct {
	Version string      `yaml:"version"`
	Sql     []SqlConfig `yaml:"sql"`
}

type SqlConfig struct {
	Engine  string    `yaml:"engine"`
	Name    string    `yaml:"name"`
	Queries string    `yaml:"queries"`
	Schema  string    `yaml:"schema"`
	Gen     GenConfig `yaml:"gen"`
}

type GenConfig struct {
	Go GoConfig `yaml:"go"`
}

type GoConfig struct {
	Package                     string      `yaml:"package"`
	Out                         string      `yaml:"out"`
	EmitResultStructPointers    bool        `yaml:"emit_result_struct_pointers"`
	EmitParamsStructPointers    bool        `yaml:"emit_params_struct_pointers"`
	EmitPreparedQueries         bool        `yaml:"emit_prepared_queries"`
	EmitExportedQueries         bool        `yaml:"emit_exported_queries"`
	EmitInterface               bool        `yaml:"emit_interface"`
	EmitJsonTags                bool        `yaml:"emit_json_tags"`
	Overrides                   []Override  `yaml:"overrides"`
}

type Override struct {
	DbType string   `yaml:"db_type"`
	GoType GoType   `yaml:"go_type"`
}

type GoType struct {
	Import  string `yaml:"import"`
	Package string `yaml:"package"`
	Type    string `yaml:"type"`
}

func GenerateSqlcConfigFile(file *descriptorpb.FileDescriptorProto) *pluginpb.CodeGeneratorResponse_File {
	packageName := file.GetPackage()
	if packageName == "" {
		return nil
	}

	domain := ExtractDomain(packageName)
	if domain == "" {
		return nil
	}

	return GenerateSqlcConfigFileForDomains([]string{domain})
}

func GenerateSqlcConfigFileForDomains(domains []string) *pluginpb.CodeGeneratorResponse_File {
	if len(domains) == 0 {
		return nil
	}

	// Remove duplicates
	uniqueDomains := make(map[string]bool)
	var sqlConfigs []SqlConfig
	
	for _, domain := range domains {
		if !uniqueDomains[domain] {
			uniqueDomains[domain] = true
			sqlConfigs = append(sqlConfigs, generateSqlConfig(domain))
		}
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

func ExtractDomain(packageName string) string {
	parts := strings.Split(packageName, ".")
	if len(parts) < 2 {
		return ""
	}
	
	// Support patterns:
	// domain.version (e.g., todo.v1) 
	// domain.subdomain.version (e.g., user.auth.v1)
	if len(parts) == 2 {
		return parts[0] // todo
	} else if len(parts) >= 3 {
		return parts[0] + "_" + parts[1] // user_auth
	}
	
	return ""
}

func extractDomainAndVersion(packageName string) (string, string) {
	parts := strings.Split(packageName, ".")
	if len(parts) < 2 {
		return "", ""
	}
	
	domain := parts[0]
	version := parts[len(parts)-1] // last part is always version
	
	return domain, version
}

func generateSqlConfig(domain string) SqlConfig {
	return SqlConfig{
		Engine:  "postgresql",
		Name:    domain,
		Queries: fmt.Sprintf("./data/queries/%s", domain),
		Schema:  fmt.Sprintf("./data/migrations/*_%s_*.sql", domain),
		Gen: GenConfig{
			Go: GoConfig{
				Package:                     fmt.Sprintf("%s_gendb", domain),
				Out:                         fmt.Sprintf("./internal/domain/%s/gendb", domain),
				EmitResultStructPointers:    true,
				EmitParamsStructPointers:    true,
				EmitPreparedQueries:         true,
				EmitExportedQueries:         true,
				EmitInterface:               true,
				EmitJsonTags:                true,
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