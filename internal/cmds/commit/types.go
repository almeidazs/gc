package commit

type CommitOptions struct {
	Emojis      bool
	Coauthored  bool
	SkipPrompts bool
	Branch      string
	Files       []string
}
