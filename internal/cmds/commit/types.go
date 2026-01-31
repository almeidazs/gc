package commit

type PushOptions struct {
	Branch      string
	SetUpstream bool
}

type CommitOptions struct {
	Push        bool
	Emojis      bool
	Coauthored  bool
	SkipPrompts bool
	SetUpstream bool
	Branch      string
	Message     string
	Files       []string
}
