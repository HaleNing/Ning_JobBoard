// Code generated by ent, DO NOT EDIT.

package job

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// JobName applies equality check predicate on the "job_name" field. It's identical to JobNameEQ.
func JobName(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJobName), v))
	})
}

// CompanyName applies equality check predicate on the "company_name" field. It's identical to CompanyNameEQ.
func CompanyName(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompanyName), v))
	})
}

// IsExist applies equality check predicate on the "is_exist" field. It's identical to IsExistEQ.
func IsExist(v bool) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsExist), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// IsRemote applies equality check predicate on the "is_remote" field. It's identical to IsRemoteEQ.
func IsRemote(v bool) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsRemote), v))
	})
}

// Exp applies equality check predicate on the "exp" field. It's identical to ExpEQ.
func Exp(v int8) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExp), v))
	})
}

// Area applies equality check predicate on the "area" field. It's identical to AreaEQ.
func Area(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArea), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// JobNameEQ applies the EQ predicate on the "job_name" field.
func JobNameEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJobName), v))
	})
}

// JobNameNEQ applies the NEQ predicate on the "job_name" field.
func JobNameNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJobName), v))
	})
}

// JobNameIn applies the In predicate on the "job_name" field.
func JobNameIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldJobName), v...))
	})
}

// JobNameNotIn applies the NotIn predicate on the "job_name" field.
func JobNameNotIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldJobName), v...))
	})
}

// JobNameGT applies the GT predicate on the "job_name" field.
func JobNameGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJobName), v))
	})
}

// JobNameGTE applies the GTE predicate on the "job_name" field.
func JobNameGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJobName), v))
	})
}

// JobNameLT applies the LT predicate on the "job_name" field.
func JobNameLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJobName), v))
	})
}

// JobNameLTE applies the LTE predicate on the "job_name" field.
func JobNameLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJobName), v))
	})
}

// JobNameContains applies the Contains predicate on the "job_name" field.
func JobNameContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldJobName), v))
	})
}

// JobNameHasPrefix applies the HasPrefix predicate on the "job_name" field.
func JobNameHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldJobName), v))
	})
}

// JobNameHasSuffix applies the HasSuffix predicate on the "job_name" field.
func JobNameHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldJobName), v))
	})
}

// JobNameEqualFold applies the EqualFold predicate on the "job_name" field.
func JobNameEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldJobName), v))
	})
}

// JobNameContainsFold applies the ContainsFold predicate on the "job_name" field.
func JobNameContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldJobName), v))
	})
}

// CompanyNameEQ applies the EQ predicate on the "company_name" field.
func CompanyNameEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompanyName), v))
	})
}

// CompanyNameNEQ applies the NEQ predicate on the "company_name" field.
func CompanyNameNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompanyName), v))
	})
}

// CompanyNameIn applies the In predicate on the "company_name" field.
func CompanyNameIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCompanyName), v...))
	})
}

// CompanyNameNotIn applies the NotIn predicate on the "company_name" field.
func CompanyNameNotIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCompanyName), v...))
	})
}

// CompanyNameGT applies the GT predicate on the "company_name" field.
func CompanyNameGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompanyName), v))
	})
}

// CompanyNameGTE applies the GTE predicate on the "company_name" field.
func CompanyNameGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompanyName), v))
	})
}

// CompanyNameLT applies the LT predicate on the "company_name" field.
func CompanyNameLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompanyName), v))
	})
}

// CompanyNameLTE applies the LTE predicate on the "company_name" field.
func CompanyNameLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompanyName), v))
	})
}

// CompanyNameContains applies the Contains predicate on the "company_name" field.
func CompanyNameContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCompanyName), v))
	})
}

// CompanyNameHasPrefix applies the HasPrefix predicate on the "company_name" field.
func CompanyNameHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCompanyName), v))
	})
}

// CompanyNameHasSuffix applies the HasSuffix predicate on the "company_name" field.
func CompanyNameHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCompanyName), v))
	})
}

// CompanyNameEqualFold applies the EqualFold predicate on the "company_name" field.
func CompanyNameEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCompanyName), v))
	})
}

// CompanyNameContainsFold applies the ContainsFold predicate on the "company_name" field.
func CompanyNameContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCompanyName), v))
	})
}

// IsExistEQ applies the EQ predicate on the "is_exist" field.
func IsExistEQ(v bool) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsExist), v))
	})
}

// IsExistNEQ applies the NEQ predicate on the "is_exist" field.
func IsExistNEQ(v bool) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsExist), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// IsRemoteEQ applies the EQ predicate on the "is_remote" field.
func IsRemoteEQ(v bool) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsRemote), v))
	})
}

// IsRemoteNEQ applies the NEQ predicate on the "is_remote" field.
func IsRemoteNEQ(v bool) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsRemote), v))
	})
}

// ExpEQ applies the EQ predicate on the "exp" field.
func ExpEQ(v int8) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExp), v))
	})
}

// ExpNEQ applies the NEQ predicate on the "exp" field.
func ExpNEQ(v int8) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExp), v))
	})
}

// ExpIn applies the In predicate on the "exp" field.
func ExpIn(vs ...int8) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExp), v...))
	})
}

// ExpNotIn applies the NotIn predicate on the "exp" field.
func ExpNotIn(vs ...int8) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExp), v...))
	})
}

// ExpGT applies the GT predicate on the "exp" field.
func ExpGT(v int8) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExp), v))
	})
}

// ExpGTE applies the GTE predicate on the "exp" field.
func ExpGTE(v int8) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExp), v))
	})
}

// ExpLT applies the LT predicate on the "exp" field.
func ExpLT(v int8) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExp), v))
	})
}

// ExpLTE applies the LTE predicate on the "exp" field.
func ExpLTE(v int8) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExp), v))
	})
}

// AreaEQ applies the EQ predicate on the "area" field.
func AreaEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArea), v))
	})
}

// AreaNEQ applies the NEQ predicate on the "area" field.
func AreaNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArea), v))
	})
}

// AreaIn applies the In predicate on the "area" field.
func AreaIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldArea), v...))
	})
}

// AreaNotIn applies the NotIn predicate on the "area" field.
func AreaNotIn(vs ...string) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldArea), v...))
	})
}

// AreaGT applies the GT predicate on the "area" field.
func AreaGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArea), v))
	})
}

// AreaGTE applies the GTE predicate on the "area" field.
func AreaGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArea), v))
	})
}

// AreaLT applies the LT predicate on the "area" field.
func AreaLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArea), v))
	})
}

// AreaLTE applies the LTE predicate on the "area" field.
func AreaLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArea), v))
	})
}

// AreaContains applies the Contains predicate on the "area" field.
func AreaContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldArea), v))
	})
}

// AreaHasPrefix applies the HasPrefix predicate on the "area" field.
func AreaHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldArea), v))
	})
}

// AreaHasSuffix applies the HasSuffix predicate on the "area" field.
func AreaHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldArea), v))
	})
}

// AreaEqualFold applies the EqualFold predicate on the "area" field.
func AreaEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldArea), v))
	})
}

// AreaContainsFold applies the ContainsFold predicate on the "area" field.
func AreaContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldArea), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		p(s.Not())
	})
}
