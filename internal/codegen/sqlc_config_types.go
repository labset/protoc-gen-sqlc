package codegen

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
	Package                  string     `yaml:"package"`
	Out                      string     `yaml:"out"`
	EmitResultStructPointers bool       `yaml:"emit_result_struct_pointers"`
	EmitParamsStructPointers bool       `yaml:"emit_params_struct_pointers"`
	EmitPreparedQueries      bool       `yaml:"emit_prepared_queries"`
	EmitExportedQueries      bool       `yaml:"emit_exported_queries"`
	EmitInterface            bool       `yaml:"emit_interface"`
	EmitJsonTags             bool       `yaml:"emit_json_tags"`
	Overrides                []Override `yaml:"overrides"`
}

type Override struct {
	DbType string `yaml:"db_type"`
	GoType GoType `yaml:"go_type"`
}

type GoType struct {
	Import  string `yaml:"import"`
	Package string `yaml:"package"`
	Type    string `yaml:"type"`
}
