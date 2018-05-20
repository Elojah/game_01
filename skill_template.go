package game

// SkillTemplate alias an entity.
// It represents semi static data. When creating skills, those templates are used.
type SkillTemplate Skill

// SkillTemplateService is an interface for SkillTemplate object.
type SkillTemplateService interface {
	SetSkillTemplate(SkillTemplate) error
	GetSkillTemplate(SkillTemplateSubset) (SkillTemplate, error)
	ListSkillTemplate() ([]SkillTemplate, error)
}

// SkillTemplateSubset is a subset to retrieve one template.
type SkillTemplateSubset struct {
	Type SkillType
}
