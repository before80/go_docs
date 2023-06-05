+++
title = "Cobra 在pkg.go.dev上的文档"
type = "docs"
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++



# Cobra 在pkg.go.dev上的文档

### Overview 

Package cobra is a commander providing a simple interface to create powerful modern CLI interfaces. In addition to providing an interface, Cobra simultaneously provides a controller to organize your application code.



### Constants 

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/bash_completions.go#L29)

```
const (
	BashCompFilenameExt     = "cobra_annotation_bash_completion_filename_extensions"
	BashCompCustom          = "cobra_annotation_bash_completion_custom"
	BashCompOneRequiredFlag = "cobra_annotation_bash_completion_one_required_flag"
	BashCompSubdirsInDir    = "cobra_annotation_bash_completion_subdirs_in_dir"
)
```

Annotations for Bash completion.

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/completions.go#L26)

```
const (
	// ShellCompRequestCmd is the name of the hidden command that is used to request
	// completion results from the program.  It is used by the shell completion scripts.
	ShellCompRequestCmd = "__complete"
	// ShellCompNoDescRequestCmd is the name of the hidden command that is used to request
	// completion results without their description.  It is used by the shell completion scripts.
	ShellCompNoDescRequestCmd = "__completeNoDesc"
)
```

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/command.go#L33)

```
const FlagSetByCobraAnnotation = "cobra_annotation_flag_set_by_cobra"
```

### Variables 

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L61)

```
var EnableCaseInsensitive = defaultCaseInsensitive
```

EnableCaseInsensitive allows case-insensitive commands names. (case sensitive by default)

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L58)

```
var EnableCommandSorting = defaultCommandSorting
```

EnableCommandSorting controls sorting of the slice of commands, which is turned on by default. To disable sorting, set it to false.

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L54)

```
var EnablePrefixMatching = defaultPrefixMatching
```

EnablePrefixMatching allows to set automatic prefix matching. Automatic prefix matching can be a dangerous thing to automatically enable in CLI tools. Set this to true to enable it.

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L76)

```
var MousetrapDisplayDuration = 5 * time.Second
```

MousetrapDisplayDuration controls how long the MousetrapHelpText message is displayed on Windows if the CLI is started from explorer.exe. Set to 0 to wait for the return key to be pressed. To disable the mousetrap, just set MousetrapHelpText to blank string (""). Works only on Microsoft Windows.

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L67)

```
var MousetrapHelpText = `This is a command line tool.

You need to open cmd.exe and run it from there.
`
```

MousetrapHelpText enables an information splash screen on Windows if the CLI is started from explorer.exe. To disable the mousetrap, just set this variable to blank string (""). Works only on Microsoft Windows.

### Functions 

#### func AddTemplateFunc 

```
func AddTemplateFunc(name string, tmplFunc interface{})
```

AddTemplateFunc adds a template function that's available to Usage and Help template generation.

#### func AddTemplateFuncs 

```
func AddTemplateFuncs(tmplFuncs template.FuncMap)
```

AddTemplateFuncs adds multiple template functions that are available to Usage and Help template generation.

#### func AppendActiveHelp  <- v1.5.0

```
func AppendActiveHelp(compArray []string, activeHelpStr string) []string
```

AppendActiveHelp adds the specified string to the specified array to be used as ActiveHelp. Such strings will be processed by the completion script and will be shown as ActiveHelp to the user. The array parameter should be the array that will contain the completions. This function can be called multiple times before and/or after completions are added to the array. Each time this function is called with the same array, the new ActiveHelp line will be shown below the previous ones when completion is triggered.

#### func ArbitraryArgs 

```
func ArbitraryArgs(cmd *Command, args []string) error
```

ArbitraryArgs never returns an error.

#### func CheckErr  <- v1.1.2

```
func CheckErr(msg interface{})
```

CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.

#### func CompDebug  <- v1.0.0

```
func CompDebug(msg string, printToStdErr bool)
```

CompDebug prints the specified string to the same file as where the completion script prints its logs. Note that completion printouts should never be on stdout as they would be wrongly interpreted as actual completion choices by the completion script.

#### func CompDebugln  <- v1.0.0

```
func CompDebugln(msg string, printToStdErr bool)
```

CompDebugln prints the specified string with a newline at the end to the same file as where the completion script prints its logs. Such logs are only printed when the user has set the environment variable BASH_COMP_DEBUG_FILE to the path of some file to be used.

#### func CompError  <- v1.0.0

```
func CompError(msg string)
```

CompError prints the specified completion message to stderr.

#### func CompErrorln  <- v1.0.0

```
func CompErrorln(msg string)
```

CompErrorln prints the specified completion message to stderr with a newline at the end.

#### func Eq 

```
func Eq(a interface{}, b interface{}) bool
```

Eq takes two types and checks whether they are equal. Supported types are int and string. Unsupported types will panic.

#### func FixedCompletions  <- v1.5.0

```
func FixedCompletions(choices []string, directive ShellCompDirective) func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)
```

FixedCompletions can be used to create a completion function which always returns the same results.

#### func GetActiveHelpConfig  <- v1.5.0

```
func GetActiveHelpConfig(cmd *Command) string
```

GetActiveHelpConfig returns the value of the ActiveHelp environment variable <PROGRAM>_ACTIVE_HELP where <PROGRAM> is the name of the root command in upper case, with all - replaced by _. It will always return "0" if the global environment variable COBRA_ACTIVE_HELP is set to "0".

#### func Gt 

```
func Gt(a interface{}, b interface{}) bool
```

Gt takes two types and checks whether the first type is greater than the second. In case of types Arrays, Chans, Maps and Slices, Gt will compare their lengths. Ints are compared directly while strings are first parsed as ints and then compared.

#### func MarkFlagCustom 

```
func MarkFlagCustom(flags *pflag.FlagSet, name string, f string) error
```

MarkFlagCustom adds the BashCompCustom annotation to the named flag, if it exists. The bash completion script will call the bash function f for the flag.

This will only work for bash completion. It is recommended to instead use c.RegisterFlagCompletionFunc(...) which allows to register a Go function which will work across all shells.

#### func MarkFlagDirname  <- v0.0.5

```
func MarkFlagDirname(flags *pflag.FlagSet, name string) error
```

MarkFlagDirname instructs the various shell completion implementations to limit completions for the named flag to directory names.

#### func MarkFlagFilename 

```
func MarkFlagFilename(flags *pflag.FlagSet, name string, extensions ...string) error
```

MarkFlagFilename instructs the various shell completion implementations to limit completions for the named flag to the specified file extensions.

#### func MarkFlagRequired 

```
func MarkFlagRequired(flags *pflag.FlagSet, name string) error
```

MarkFlagRequired instructs the various shell completion implementations to prioritize the named flag when performing completion, and causes your command to report an error if invoked without the flag.

#### func NoArgs 

```
func NoArgs(cmd *Command, args []string) error
```

NoArgs returns an error if any args are included.

#### func OnFinalize  <- v1.6.0

```
func OnFinalize(y ...func())
```

OnFinalize sets the passed functions to be run when each command's Execute method is terminated.

#### func OnInitialize 

```
func OnInitialize(y ...func())
```

OnInitialize sets the passed functions to be run when each command's Execute method is called.

#### func OnlyValidArgs 

```
func OnlyValidArgs(cmd *Command, args []string) error
```

OnlyValidArgs returns an error if there are any positional args that are not in the `ValidArgs` field of `Command`

#### func WriteStringAndCheck  <- v1.1.2

```
func WriteStringAndCheck(b io.StringWriter, s string)
```

WriteStringAndCheck writes a string into a buffer, and checks if the error is not nil.

### Types 

#### type Command 

```
type Command struct {
	// Use is the one-line usage message.
	// Recommended syntax is as follows:
	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//   ... indicates that you can specify multiple values for the previous argument.
	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//       optional, they are enclosed in brackets ([ ]).
	// Example: add [-F file | -D dir]... [-f format] profile
	Use string

	// Aliases is an array of aliases that can be used instead of the first word in Use.
	Aliases []string

	// SuggestFor is an array of command names for which this command will be suggested -
	// similar to aliases but only suggests.
	SuggestFor []string

	// Short is the short description shown in the 'help' output.
	Short string

	// The group id under which this subcommand is grouped in the 'help' output of its parent.
	GroupID string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// Example is examples of how to use the command.
	Example string

	// ValidArgs is list of all valid non-flag arguments that are accepted in shell completions
	ValidArgs []string
	// ValidArgsFunction is an optional function that provides valid non-flag arguments for shell completion.
	// It is a dynamic version of using ValidArgs.
	// Only one of ValidArgs and ValidArgsFunction can be used for a command.
	ValidArgsFunction func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)

	// Expected arguments
	Args PositionalArgs

	// ArgAliases is List of aliases for ValidArgs.
	// These are not suggested to the user in the shell completion,
	// but accepted if entered manually.
	ArgAliases []string

	// BashCompletionFunction is custom bash functions used by the legacy bash autocompletion generator.
	// For portability with other shells, it is recommended to instead use ValidArgsFunction
	BashCompletionFunction string

	// Deprecated defines, if this command is deprecated and should print this string when used.
	Deprecated string

	// Annotations are key/value pairs that can be used by applications to identify or
	// group commands.
	Annotations map[string]string

	// Version defines the version for this command. If this value is non-empty and the command does not
	// define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
	// will print content of the "Version" variable. A shorthand "v" flag will also be added if the
	// command does not define one.
	Version string

	// The *Run functions are executed in the following order:
	//   * PersistentPreRun()
	//   * PreRun()
	//   * Run()
	//   * PostRun()
	//   * PersistentPostRun()
	// All functions get the same args, the arguments after the command name.
	//
	// PersistentPreRun: children of this command will inherit and execute.
	PersistentPreRun func(cmd *Command, args []string)
	// PersistentPreRunE: PersistentPreRun but returns an error.
	PersistentPreRunE func(cmd *Command, args []string) error
	// PreRun: children of this command will not inherit.
	PreRun func(cmd *Command, args []string)
	// PreRunE: PreRun but returns an error.
	PreRunE func(cmd *Command, args []string) error
	// Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *Command, args []string)
	// RunE: Run but returns an error.
	RunE func(cmd *Command, args []string) error
	// PostRun: run after the Run command.
	PostRun func(cmd *Command, args []string)
	// PostRunE: PostRun but returns an error.
	PostRunE func(cmd *Command, args []string) error
	// PersistentPostRun: children of this command will inherit and execute after PostRun.
	PersistentPostRun func(cmd *Command, args []string)
	// PersistentPostRunE: PersistentPostRun but returns an error.
	PersistentPostRunE func(cmd *Command, args []string) error

	// FParseErrWhitelist flag parse errors to be ignored
	FParseErrWhitelist FParseErrWhitelist

	// CompletionOptions is a set of options to control the handling of shell completion
	CompletionOptions CompletionOptions

	// TraverseChildren parses flags on all parents before executing child command.
	TraverseChildren bool

	// Hidden defines, if this command is hidden and should NOT show up in the list of available commands.
	Hidden bool

	// SilenceErrors is an option to quiet errors down stream.
	SilenceErrors bool

	// SilenceUsage is an option to silence usage when an error occurs.
	SilenceUsage bool

	// DisableFlagParsing disables the flag parsing.
	// If this is true all flags will be passed to the command as arguments.
	DisableFlagParsing bool

	// DisableAutoGenTag defines, if gen tag ("Auto generated by spf13/cobra...")
	// will be printed by generating docs for this command.
	DisableAutoGenTag bool

	// DisableFlagsInUseLine will disable the addition of [flags] to the usage
	// line of a command when printing help or generating docs
	DisableFlagsInUseLine bool

	// DisableSuggestions disables the suggestions based on Levenshtein distance
	// that go along with 'unknown command' messages.
	DisableSuggestions bool

	// SuggestionsMinimumDistance defines minimum levenshtein distance to display suggestions.
	// Must be > 0.
	SuggestionsMinimumDistance int
	// contains filtered or unexported fields
}
```

Command is just that, a command for your application. E.g. 'go run ...' - 'run' is the command. Cobra requires you to define the usage and description as part of your command definition to ensure usability.

#### (*Command) AddCommand 

```
func (c *Command) AddCommand(cmds ...*Command)
```

AddCommand adds one or more commands to this parent command.

#### (*Command) AddGroup  <- v1.6.0

```
func (c *Command) AddGroup(groups ...*Group)
```

AddGroup adds one or more command groups to this parent command.

#### (*Command) AllChildCommandsHaveGroup  <- v1.6.0

```
func (c *Command) AllChildCommandsHaveGroup() bool
```

AllChildCommandsHaveGroup returns if all subcommands are assigned to a group

#### (*Command) ArgsLenAtDash 

```
func (c *Command) ArgsLenAtDash() int
```

ArgsLenAtDash will return the length of c.Flags().Args at the moment when a -- was found during args parsing.

#### (*Command) CalledAs  <- v0.0.2

```
func (c *Command) CalledAs() string
```

CalledAs returns the command name or alias that was used to invoke this command or an empty string if the command has not been called.

#### (*Command) CommandPath 

```
func (c *Command) CommandPath() string
```

CommandPath returns the full path to this command.

#### (*Command) CommandPathPadding 

```
func (c *Command) CommandPathPadding() int
```

CommandPathPadding return padding for the command path.

#### (*Command) Commands 

```
func (c *Command) Commands() []*Command
```

Commands returns a sorted slice of child commands.

#### (*Command) ContainsGroup  <- v1.6.0

```
func (c *Command) ContainsGroup(groupID string) bool
```

ContainsGroup return if groupID exists in the list of command groups.

#### (*Command) Context  <- v0.0.6

```
func (c *Command) Context() context.Context
```

Context returns underlying command context. If command was executed with ExecuteContext or the context was set with SetContext, the previously set context will be returned. Otherwise, nil is returned.

Notice that a call to Execute and ExecuteC will replace a nil context of a command with a context.Background, so a background context will be returned by Context after one of these functions has been called.

#### (*Command) DebugFlags 

```
func (c *Command) DebugFlags()
```

DebugFlags used to determine which flags have been assigned to which commands and which persist.

#### (*Command) ErrOrStderr  <- v0.0.5

```
func (c *Command) ErrOrStderr() io.Writer
```

ErrOrStderr returns output to stderr

#### (*Command) Execute 

```
func (c *Command) Execute() error
```

Execute uses the args (os.Args[1:] by default) and run through the command tree finding appropriate matches for commands and then corresponding flags.

#### (*Command) ExecuteC 

```
func (c *Command) ExecuteC() (cmd *Command, err error)
```

ExecuteC executes the command.

#### (*Command) ExecuteContext  <- v0.0.6

```
func (c *Command) ExecuteContext(ctx context.Context) error
```

ExecuteContext is the same as Execute(), but sets the ctx on the command. Retrieve ctx by calling cmd.Context() inside your *Run lifecycle or ValidArgs functions.

#### (*Command) ExecuteContextC  <- v1.2.0

```
func (c *Command) ExecuteContextC(ctx context.Context) (*Command, error)
```

ExecuteContextC is the same as ExecuteC(), but sets the ctx on the command. Retrieve ctx by calling cmd.Context() inside your *Run lifecycle or ValidArgs functions.

#### (*Command) Find 

```
func (c *Command) Find(args []string) (*Command, []string, error)
```

Find the target command given the args and command tree Meant to be run on the highest node. Only searches down.

#### (*Command) Flag 

```
func (c *Command) Flag(name string) (flag *flag.Flag)
```

Flag climbs up the command tree looking for matching flag.

#### (*Command) FlagErrorFunc 

```
func (c *Command) FlagErrorFunc() (f func(*Command, error) error)
```

FlagErrorFunc returns either the function set by SetFlagErrorFunc for this command or a parent, or it returns a function which returns the original error.

#### (*Command) Flags 

```
func (c *Command) Flags() *flag.FlagSet
```

Flags returns the complete FlagSet that applies to this command (local and persistent declared here and by all parents).

#### (*Command) GenBashCompletion 

```
func (c *Command) GenBashCompletion(w io.Writer) error
```

GenBashCompletion generates bash completion file and writes to the passed writer.

#### (*Command) GenBashCompletionFile 

```
func (c *Command) GenBashCompletionFile(filename string) error
```

GenBashCompletionFile generates bash completion file.

#### (*Command) GenBashCompletionFileV2  <- v1.2.0

```
func (c *Command) GenBashCompletionFileV2(filename string, includeDesc bool) error
```

GenBashCompletionFileV2 generates Bash completion version 2.

#### (*Command) GenBashCompletionV2  <- v1.2.0

```
func (c *Command) GenBashCompletionV2(w io.Writer, includeDesc bool) error
```

GenBashCompletionV2 generates Bash completion file version 2 and writes it to the passed writer.

#### (*Command) GenFishCompletion  <- v1.0.0

```
func (c *Command) GenFishCompletion(w io.Writer, includeDesc bool) error
```

GenFishCompletion generates fish completion file and writes to the passed writer.

#### (*Command) GenFishCompletionFile  <- v1.0.0

```
func (c *Command) GenFishCompletionFile(filename string, includeDesc bool) error
```

GenFishCompletionFile generates fish completion file.

#### (*Command) GenPowerShellCompletion  <- v0.0.5

```
func (c *Command) GenPowerShellCompletion(w io.Writer) error
```

GenPowerShellCompletion generates powershell completion file without descriptions and writes it to the passed writer.

#### (*Command) GenPowerShellCompletionFile  <- v0.0.5

```
func (c *Command) GenPowerShellCompletionFile(filename string) error
```

GenPowerShellCompletionFile generates powershell completion file without descriptions.

#### (*Command) GenPowerShellCompletionFileWithDesc  <- v1.1.2

```
func (c *Command) GenPowerShellCompletionFileWithDesc(filename string) error
```

GenPowerShellCompletionFileWithDesc generates powershell completion file with descriptions.

#### (*Command) GenPowerShellCompletionWithDesc  <- v1.1.2

```
func (c *Command) GenPowerShellCompletionWithDesc(w io.Writer) error
```

GenPowerShellCompletionWithDesc generates powershell completion file with descriptions and writes it to the passed writer.

#### (*Command) GenZshCompletion 

```
func (c *Command) GenZshCompletion(w io.Writer) error
```

GenZshCompletion generates zsh completion file including descriptions and writes it to the passed writer.

#### (*Command) GenZshCompletionFile 

```
func (c *Command) GenZshCompletionFile(filename string) error
```

GenZshCompletionFile generates zsh completion file including descriptions.

#### (*Command) GenZshCompletionFileNoDesc  <- v1.1.0

```
func (c *Command) GenZshCompletionFileNoDesc(filename string) error
```

GenZshCompletionFileNoDesc generates zsh completion file without descriptions.

#### (*Command) GenZshCompletionNoDesc  <- v1.1.0

```
func (c *Command) GenZshCompletionNoDesc(w io.Writer) error
```

GenZshCompletionNoDesc generates zsh completion file without descriptions and writes it to the passed writer.

#### (*Command) GlobalNormalizationFunc 

```
func (c *Command) GlobalNormalizationFunc() func(f *flag.FlagSet, name string) flag.NormalizedName
```

GlobalNormalizationFunc returns the global normalization function or nil if it doesn't exist.

#### (*Command) Groups  <- v1.6.0

```
func (c *Command) Groups() []*Group
```

Groups returns a slice of child command groups.

#### (*Command) HasAlias 

```
func (c *Command) HasAlias(s string) bool
```

HasAlias determines if a given string is an alias of the command.

#### (*Command) HasAvailableFlags 

```
func (c *Command) HasAvailableFlags() bool
```

HasAvailableFlags checks if the command contains any flags (local plus persistent from the entire structure) which are not hidden or deprecated.

#### (*Command) HasAvailableInheritedFlags 

```
func (c *Command) HasAvailableInheritedFlags() bool
```

HasAvailableInheritedFlags checks if the command has flags inherited from its parent command which are not hidden or deprecated.

#### (*Command) HasAvailableLocalFlags 

```
func (c *Command) HasAvailableLocalFlags() bool
```

HasAvailableLocalFlags checks if the command has flags specifically declared locally which are not hidden or deprecated.

#### (*Command) HasAvailablePersistentFlags 

```
func (c *Command) HasAvailablePersistentFlags() bool
```

HasAvailablePersistentFlags checks if the command contains persistent flags which are not hidden or deprecated.

#### (*Command) HasAvailableSubCommands 

```
func (c *Command) HasAvailableSubCommands() bool
```

HasAvailableSubCommands determines if a command has available sub commands that need to be shown in the usage/help default template under 'available commands'.

#### (*Command) HasExample 

```
func (c *Command) HasExample() bool
```

HasExample determines if the command has example.

#### (*Command) HasFlags 

```
func (c *Command) HasFlags() bool
```

HasFlags checks if the command contains any flags (local plus persistent from the entire structure).

#### (*Command) HasHelpSubCommands 

```
func (c *Command) HasHelpSubCommands() bool
```

HasHelpSubCommands determines if a command has any available 'help' sub commands that need to be shown in the usage/help default template under 'additional help topics'.

#### (*Command) HasInheritedFlags 

```
func (c *Command) HasInheritedFlags() bool
```

HasInheritedFlags checks if the command has flags inherited from its parent command.

#### (*Command) HasLocalFlags 

```
func (c *Command) HasLocalFlags() bool
```

HasLocalFlags checks if the command has flags specifically declared locally.

#### (*Command) HasParent 

```
func (c *Command) HasParent() bool
```

HasParent determines if the command is a child command.

#### (*Command) HasPersistentFlags 

```
func (c *Command) HasPersistentFlags() bool
```

HasPersistentFlags checks if the command contains persistent flags.

#### (*Command) HasSubCommands 

```
func (c *Command) HasSubCommands() bool
```

HasSubCommands determines if the command has children commands.

#### (*Command) Help 

```
func (c *Command) Help() error
```

Help puts out the help for the command. Used when a user calls help [command]. Can be defined by user by overriding HelpFunc.

#### (*Command) HelpFunc 

```
func (c *Command) HelpFunc() func(*Command, []string)
```

HelpFunc returns either the function set by SetHelpFunc for this command or a parent, or it returns a function with default help behavior.

#### (*Command) HelpTemplate 

```
func (c *Command) HelpTemplate() string
```

HelpTemplate return help template for the command.

#### (*Command) InOrStdin  <- v0.0.5

```
func (c *Command) InOrStdin() io.Reader
```

InOrStdin returns input to stdin

#### (*Command) InheritedFlags 

```
func (c *Command) InheritedFlags() *flag.FlagSet
```

InheritedFlags returns all flags which were inherited from parent commands.

#### (*Command) InitDefaultCompletionCmd  <- v1.6.0

```
func (c *Command) InitDefaultCompletionCmd()
```

InitDefaultCompletionCmd adds a default 'completion' command to c. This function will do nothing if any of the following is true: 1- the feature has been explicitly disabled by the program, 2- c has no subcommands (to avoid creating one), 3- c already has a 'completion' command provided by the program.

#### (*Command) InitDefaultHelpCmd 

```
func (c *Command) InitDefaultHelpCmd()
```

InitDefaultHelpCmd adds default help command to c. It is called automatically by executing the c or by calling help and usage. If c already has help command or c has no subcommands, it will do nothing.

#### (*Command) InitDefaultHelpFlag 

```
func (c *Command) InitDefaultHelpFlag()
```

InitDefaultHelpFlag adds default help flag to c. It is called automatically by executing the c or by calling help and usage. If c already has help flag, it will do nothing.

#### (*Command) InitDefaultVersionFlag  <- v0.0.2

```
func (c *Command) InitDefaultVersionFlag()
```

InitDefaultVersionFlag adds default version flag to c. It is called automatically by executing the c. If c already has a version flag, it will do nothing. If c.Version is empty, it will do nothing.

#### (*Command) IsAdditionalHelpTopicCommand 

```
func (c *Command) IsAdditionalHelpTopicCommand() bool
```

IsAdditionalHelpTopicCommand determines if a command is an additional help topic command; additional help topic command is determined by the fact that it is NOT runnable/hidden/deprecated, and has no sub commands that are runnable/hidden/deprecated. Concrete example: https://github.com/spf13/cobra/issues/393#issuecomment-282741924.

#### (*Command) IsAvailableCommand 

```
func (c *Command) IsAvailableCommand() bool
```

IsAvailableCommand determines if a command is available as a non-help command (this includes all non deprecated/hidden commands).

#### (*Command) LocalFlags 

```
func (c *Command) LocalFlags() *flag.FlagSet
```

LocalFlags returns the local FlagSet specifically set in the current command.

#### (*Command) LocalNonPersistentFlags 

```
func (c *Command) LocalNonPersistentFlags() *flag.FlagSet
```

LocalNonPersistentFlags are flags specific to this command which will NOT persist to subcommands.

#### (*Command) MarkFlagCustom 

```
func (c *Command) MarkFlagCustom(name string, f string) error
```

MarkFlagCustom adds the BashCompCustom annotation to the named flag, if it exists. The bash completion script will call the bash function f for the flag.

This will only work for bash completion. It is recommended to instead use c.RegisterFlagCompletionFunc(...) which allows to register a Go function which will work across all shells.

#### (*Command) MarkFlagDirname  <- v0.0.5

```
func (c *Command) MarkFlagDirname(name string) error
```

MarkFlagDirname instructs the various shell completion implementations to limit completions for the named flag to directory names.

#### (*Command) MarkFlagFilename 

```
func (c *Command) MarkFlagFilename(name string, extensions ...string) error
```

MarkFlagFilename instructs the various shell completion implementations to limit completions for the named flag to the specified file extensions.

#### (*Command) MarkFlagRequired 

```
func (c *Command) MarkFlagRequired(name string) error
```

MarkFlagRequired instructs the various shell completion implementations to prioritize the named flag when performing completion, and causes your command to report an error if invoked without the flag.

#### (*Command) MarkFlagsMutuallyExclusive  <- v1.5.0

```
func (c *Command) MarkFlagsMutuallyExclusive(flagNames ...string)
```

MarkFlagsMutuallyExclusive marks the given flags with annotations so that Cobra errors if the command is invoked with more than one flag from the given set of flags.

#### (*Command) MarkFlagsRequiredTogether  <- v1.5.0

```
func (c *Command) MarkFlagsRequiredTogether(flagNames ...string)
```

MarkFlagsRequiredTogether marks the given flags with annotations so that Cobra errors if the command is invoked with a subset (but not all) of the given flags.

#### (*Command) MarkPersistentFlagDirname  <- v0.0.5

```
func (c *Command) MarkPersistentFlagDirname(name string) error
```

MarkPersistentFlagDirname instructs the various shell completion implementations to limit completions for the named persistent flag to directory names.

#### (*Command) MarkPersistentFlagFilename 

```
func (c *Command) MarkPersistentFlagFilename(name string, extensions ...string) error
```

MarkPersistentFlagFilename instructs the various shell completion implementations to limit completions for the named persistent flag to the specified file extensions.

#### (*Command) MarkPersistentFlagRequired 

```
func (c *Command) MarkPersistentFlagRequired(name string) error
```

MarkPersistentFlagRequired instructs the various shell completion implementations to prioritize the named persistent flag when performing completion, and causes your command to report an error if invoked without the flag.

#### (*Command) MarkZshCompPositionalArgumentFile  <- v0.0.5

```
func (c *Command) MarkZshCompPositionalArgumentFile(argPosition int, patterns ...string) error
```

MarkZshCompPositionalArgumentFile only worked for zsh and its behavior was not consistent with Bash completion. It has therefore been disabled. Instead, when no other completion is specified, file completion is done by default for every argument. One can disable file completion on a per-argument basis by using ValidArgsFunction and ShellCompDirectiveNoFileComp. To achieve file extension filtering, one can use ValidArgsFunction and ShellCompDirectiveFilterFileExt.

Deprecated

#### (*Command) MarkZshCompPositionalArgumentWords  <- v0.0.5

```
func (c *Command) MarkZshCompPositionalArgumentWords(argPosition int, words ...string) error
```

MarkZshCompPositionalArgumentWords only worked for zsh. It has therefore been disabled. To achieve the same behavior across all shells, one can use ValidArgs (for the first argument only) or ValidArgsFunction for any argument (can include the first one also).

Deprecated

#### (*Command) Name 

```
func (c *Command) Name() string
```

Name returns the command's name: the first word in the use line.

#### (*Command) NameAndAliases 

```
func (c *Command) NameAndAliases() string
```

NameAndAliases returns a list of the command name and all aliases

#### (*Command) NamePadding 

```
func (c *Command) NamePadding() int
```

NamePadding returns padding for the name.

#### (*Command) NonInheritedFlags 

```
func (c *Command) NonInheritedFlags() *flag.FlagSet
```

NonInheritedFlags returns all flags which were not inherited from parent commands.

#### (*Command) OutOrStderr 

```
func (c *Command) OutOrStderr() io.Writer
```

OutOrStderr returns output to stderr

#### (*Command) OutOrStdout 

```
func (c *Command) OutOrStdout() io.Writer
```

OutOrStdout returns output to stdout.

#### (*Command) Parent 

```
func (c *Command) Parent() *Command
```

Parent returns a commands parent command.

#### (*Command) ParseFlags 

```
func (c *Command) ParseFlags(args []string) error
```

ParseFlags parses persistent flag tree and local flags.

#### (*Command) PersistentFlags 

```
func (c *Command) PersistentFlags() *flag.FlagSet
```

PersistentFlags returns the persistent FlagSet specifically set in the current command.

#### (*Command) Print 

```
func (c *Command) Print(i ...interface{})
```

Print is a convenience method to Print to the defined output, fallback to Stderr if not set.

#### (*Command) PrintErr  <- v0.0.5

```
func (c *Command) PrintErr(i ...interface{})
```

PrintErr is a convenience method to Print to the defined Err output, fallback to Stderr if not set.

#### (*Command) PrintErrf  <- v0.0.5

```
func (c *Command) PrintErrf(format string, i ...interface{})
```

PrintErrf is a convenience method to Printf to the defined Err output, fallback to Stderr if not set.

#### (*Command) PrintErrln  <- v0.0.5

```
func (c *Command) PrintErrln(i ...interface{})
```

PrintErrln is a convenience method to Println to the defined Err output, fallback to Stderr if not set.

#### (*Command) Printf 

```
func (c *Command) Printf(format string, i ...interface{})
```

Printf is a convenience method to Printf to the defined output, fallback to Stderr if not set.

#### (*Command) Println 

```
func (c *Command) Println(i ...interface{})
```

Println is a convenience method to Println to the defined output, fallback to Stderr if not set.

#### (*Command) RegisterFlagCompletionFunc  <- v1.0.0

```
func (c *Command) RegisterFlagCompletionFunc(flagName string, f func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)) error
```

RegisterFlagCompletionFunc should be called to register a function to provide completion for a flag.

#### (*Command) RemoveCommand 

```
func (c *Command) RemoveCommand(cmds ...*Command)
```

RemoveCommand removes one or more commands from a parent command.

#### (*Command) ResetCommands 

```
func (c *Command) ResetCommands()
```

ResetCommands delete parent, subcommand and help command from c.

#### (*Command) ResetFlags 

```
func (c *Command) ResetFlags()
```

ResetFlags deletes all flags from command.

#### (*Command) Root 

```
func (c *Command) Root() *Command
```

Root finds root command.

#### (*Command) Runnable 

```
func (c *Command) Runnable() bool
```

Runnable determines if the command is itself runnable.

#### (*Command) SetArgs 

```
func (c *Command) SetArgs(a []string)
```

SetArgs sets arguments for the command. It is set to os.Args[1:] by default, if desired, can be overridden particularly useful when testing.

#### (*Command) SetCompletionCommandGroupID  <- v1.6.0

```
func (c *Command) SetCompletionCommandGroupID(groupID string)
```

SetCompletionCommandGroupID sets the group id of the completion command.

#### (*Command) SetContext  <- v1.5.0

```
func (c *Command) SetContext(ctx context.Context)
```

SetContext sets context for the command. This context will be overwritten by Command.ExecuteContext or Command.ExecuteContextC.

#### (*Command) SetErr  <- v0.0.5

```
func (c *Command) SetErr(newErr io.Writer)
```

SetErr sets the destination for error messages. If newErr is nil, os.Stderr is used.

#### (*Command) SetFlagErrorFunc 

```
func (c *Command) SetFlagErrorFunc(f func(*Command, error) error)
```

SetFlagErrorFunc sets a function to generate an error when flag parsing fails.

#### (*Command) SetGlobalNormalizationFunc 

```
func (c *Command) SetGlobalNormalizationFunc(n func(f *flag.FlagSet, name string) flag.NormalizedName)
```

SetGlobalNormalizationFunc sets a normalization function to all flag sets and also to child commands. The user should not have a cyclic dependency on commands.

#### (*Command) SetHelpCommand 

```
func (c *Command) SetHelpCommand(cmd *Command)
```

SetHelpCommand sets help command.

#### (*Command) SetHelpCommandGroupID  <- v1.6.0

```
func (c *Command) SetHelpCommandGroupID(groupID string)
```

SetHelpCommandGroupID sets the group id of the help command.

#### (*Command) SetHelpFunc 

```
func (c *Command) SetHelpFunc(f func(*Command, []string))
```

SetHelpFunc sets help function. Can be defined by Application.

#### (*Command) SetHelpTemplate 

```
func (c *Command) SetHelpTemplate(s string)
```

SetHelpTemplate sets help template to be used. Application can use it to set custom template.

#### (*Command) SetIn  <- v0.0.5

```
func (c *Command) SetIn(newIn io.Reader)
```

SetIn sets the source for input data If newIn is nil, os.Stdin is used.

#### (*Command) SetOut  <- v0.0.5

```
func (c *Command) SetOut(newOut io.Writer)
```

SetOut sets the destination for usage messages. If newOut is nil, os.Stdout is used.

#### (*Command) SetOutput 

```
func (c *Command) SetOutput(output io.Writer)
```

SetOutput sets the destination for usage and error messages. If output is nil, os.Stderr is used. Deprecated: Use SetOut and/or SetErr instead

#### (*Command) SetUsageFunc 

```
func (c *Command) SetUsageFunc(f func(*Command) error)
```

SetUsageFunc sets usage function. Usage can be defined by application.

#### (*Command) SetUsageTemplate 

```
func (c *Command) SetUsageTemplate(s string)
```

SetUsageTemplate sets usage template. Can be defined by Application.

#### (*Command) SetVersionTemplate  <- v0.0.2

```
func (c *Command) SetVersionTemplate(s string)
```

SetVersionTemplate sets version template to be used. Application can use it to set custom template.

#### (*Command) SuggestionsFor 

```
func (c *Command) SuggestionsFor(typedName string) []string
```

SuggestionsFor provides suggestions for the typedName.

#### (*Command) Traverse 

```
func (c *Command) Traverse(args []string) (*Command, []string, error)
```

Traverse the command tree to find the command, and parse args for each parent.

#### (*Command) Usage 

```
func (c *Command) Usage() error
```

Usage puts out the usage for the command. Used when a user provides invalid input. Can be defined by user by overriding UsageFunc.

#### (*Command) UsageFunc 

```
func (c *Command) UsageFunc() (f func(*Command) error)
```

UsageFunc returns either the function set by SetUsageFunc for this command or a parent, or it returns a default usage function.

#### (*Command) UsagePadding 

```
func (c *Command) UsagePadding() int
```

UsagePadding return padding for the usage.

#### (*Command) UsageString 

```
func (c *Command) UsageString() string
```

UsageString returns usage string.

#### (*Command) UsageTemplate 

```
func (c *Command) UsageTemplate() string
```

UsageTemplate returns usage template for the command.

#### (*Command) UseLine 

```
func (c *Command) UseLine() string
```

UseLine puts out the full usage for a given command (including parents).

#### (*Command) ValidateArgs 

```
func (c *Command) ValidateArgs(args []string) error
```

#### (*Command) ValidateFlagGroups  <- v1.6.0

```
func (c *Command) ValidateFlagGroups() error
```

ValidateFlagGroups validates the mutuallyExclusive/requiredAsGroup logic and returns the first error encountered.

#### (*Command) ValidateRequiredFlags  <- v1.6.0

```
func (c *Command) ValidateRequiredFlags() error
```

ValidateRequiredFlags validates all required flags are present and returns an error otherwise

#### (*Command) VersionTemplate  <- v0.0.2

```
func (c *Command) VersionTemplate() string
```

VersionTemplate return version template for the command.

#### (*Command) VisitParents 

```
func (c *Command) VisitParents(fn func(*Command))
```

VisitParents visits all parents of the command and invokes fn on each parent.

#### type CompletionOptions  <- v1.2.0

```
type CompletionOptions struct {
	// DisableDefaultCmd prevents Cobra from creating a default 'completion' command
	DisableDefaultCmd bool
	// DisableNoDescFlag prevents Cobra from creating the '--no-descriptions' flag
	// for shells that support completion descriptions
	DisableNoDescFlag bool
	// DisableDescriptions turns off all completion descriptions for shells
	// that support them
	DisableDescriptions bool
	// HiddenDefaultCmd makes the default 'completion' command hidden
	HiddenDefaultCmd bool
}
```

CompletionOptions are the options to control shell completion

#### type FParseErrWhitelist  <- v0.0.3

```
type FParseErrWhitelist flag.ParseErrorsWhitelist
```

FParseErrWhitelist configures Flag parse errors to be ignored

#### type Group  <- v1.6.0

```
type Group struct {
	ID    string
	Title string
}
```

Group Structure to manage groups for commands

#### type PositionalArgs 

```
type PositionalArgs func(cmd *Command, args []string) error
```

#### func ExactArgs 

```
func ExactArgs(n int) PositionalArgs
```

ExactArgs returns an error if there are not exactly n args.

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="ExactValidArgs" data-kind="function" class="Documentation-typeFuncHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">func<a class="Documentation-source" href="https://github.com/spf13/cobra/blob/v1.7.0/args.go#L129" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">ExactValidArgs</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">DEPRECATED</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"><span class="Documentation-sinceVersionLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">added in</span><span>&nbsp;</span><span class="Documentation-sinceVersionVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">v0.0.4</span></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1208.38px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/builtin#int" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/github.com/spf13/cobra#PositionalArgs" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p></div></details>

#### func MatchAll  <- v1.3.0

```
func MatchAll(pargs ...PositionalArgs) PositionalArgs
```

MatchAll allows combining several PositionalArgs to work in concert.

#### func MaximumNArgs 

```
func MaximumNArgs(n int) PositionalArgs
```

MaximumNArgs returns an error if there are more than N args.

#### func MinimumNArgs 

```
func MinimumNArgs(n int) PositionalArgs
```

MinimumNArgs returns an error if there is not at least N args.

#### func RangeArgs 

```
func RangeArgs(min int, max int) PositionalArgs
```

RangeArgs returns an error if the number of args is not within the expected range.

#### type ShellCompDirective  <- v1.0.0

```
type ShellCompDirective int
```

ShellCompDirective is a bit map representing the different behaviors the shell can be instructed to have once completions have been provided.

```
const (
	// ShellCompDirectiveError indicates an error occurred and completions should be ignored.
	ShellCompDirectiveError ShellCompDirective = 1 << iota

	// ShellCompDirectiveNoSpace indicates that the shell should not add a space
	// after the completion even if there is a single completion provided.
	ShellCompDirectiveNoSpace

	// ShellCompDirectiveNoFileComp indicates that the shell should not provide
	// file completion even when no completion is provided.
	ShellCompDirectiveNoFileComp

	// ShellCompDirectiveFilterFileExt indicates that the provided completions
	// should be used as file extension filters.
	// For flags, using Command.MarkFlagFilename() and Command.MarkPersistentFlagFilename()
	// is a shortcut to using this directive explicitly.  The BashCompFilenameExt
	// annotation can also be used to obtain the same behavior for flags.
	ShellCompDirectiveFilterFileExt

	// ShellCompDirectiveFilterDirs indicates that only directory names should
	// be provided in file completion.  To request directory names within another
	// directory, the returned completions should specify the directory within
	// which to search.  The BashCompSubdirsInDir annotation can be used to
	// obtain the same behavior but only for flags.
	ShellCompDirectiveFilterDirs

	// ShellCompDirectiveKeepOrder indicates that the shell should preserve the order
	// in which the completions are provided
	ShellCompDirectiveKeepOrder

	// ShellCompDirectiveDefault indicates to let the shell perform its default
	// behavior after completions have been provided.
	// This one must be last to avoid messing up the iota count.
	ShellCompDirectiveDefault ShellCompDirective = 0
)
```

#### func NoFileCompletions  <- v1.2.0

```
func NoFileCompletions(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)
```

NoFileCompletions can be used to disable file completion for commands that should not trigger file completions.
