package commit

type CommitOptions struct {
	Push        bool
	Emojis      bool
	Coauthored  bool
	SkipPrompts bool
	Branch      string
	Message     string
	Files       []string
}
