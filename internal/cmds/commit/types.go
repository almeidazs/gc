package commit

type CommitOptions struct {
	Coauthored  bool
	SkipPrompts bool
	Branch      string
	Files       []string
}
