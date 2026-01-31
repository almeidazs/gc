package commit

type PushOptions struct {
	Force       bool
	SetUpstream bool
	Branch      string
}

type CommitOptions struct {
	Push        bool
	Force       bool
	Emojis      bool
	Coauthored  bool
	SkipPrompts bool
	SetUpstream bool
	Branch      string
	Message     string
	Files       []string
}
