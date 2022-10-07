// Code generated by ent, DO NOT EDIT.

package job

const (
	// Label holds the string label denoting the job type in the database.
	Label = "job"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldJobName holds the string denoting the job_name field in the database.
	FieldJobName = "job_name"
	// FieldCompanyName holds the string denoting the company_name field in the database.
	FieldCompanyName = "company_name"
	// FieldIsExist holds the string denoting the is_exist field in the database.
	FieldIsExist = "is_exist"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsRemote holds the string denoting the is_remote field in the database.
	FieldIsRemote = "is_remote"
	// FieldExp holds the string denoting the exp field in the database.
	FieldExp = "exp"
	// FieldArea holds the string denoting the area field in the database.
	FieldArea = "area"
	// Table holds the table name of the job in the database.
	Table = "jobs"
)

// Columns holds all SQL columns for job fields.
var Columns = []string{
	FieldID,
	FieldJobName,
	FieldCompanyName,
	FieldIsExist,
	FieldDescription,
	FieldIsRemote,
	FieldExp,
	FieldArea,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// JobNameValidator is a validator for the "job_name" field. It is called by the builders before save.
	JobNameValidator func(string) error
	// CompanyNameValidator is a validator for the "company_name" field. It is called by the builders before save.
	CompanyNameValidator func(string) error
	// DefaultIsExist holds the default value on creation for the "is_exist" field.
	DefaultIsExist bool
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// ExpValidator is a validator for the "exp" field. It is called by the builders before save.
	ExpValidator func(int8) error
	// AreaValidator is a validator for the "area" field. It is called by the builders before save.
	AreaValidator func(string) error
)
