# gogo/log

Package gogo/log is a golang package for formatted log messages that works across Linux and Windows supporting ANSI-colors.

It provides a consistent API for formatting Log Messages (stdout) in CLI Projects for Users.


## Package Dependencies

[githun.com/logrusorgru/aurora](https://github.com/logrusorgru/aurora)
[github.com/mattn/go-colorable](https://github.com/onsi/ginkgo/tree/master/reporters/stenographer/support/go-colorable)
[github.com/mattn/go-isatty](https://github.com/onsi/ginkgo/tree/master/reporters/stenographer/support/go-isatty)


## Basic Usage

Here is a quick summary of basic usage

    import "github.com/knowntraveler/gogo/log"

    func main() {
        
        // Standard Log Messages
        log.Print("This is a standard log message")
        log.Printf("This is a standard log message with %s\n", "formatting")

        // Verbose Log Messages -> EnableVerbose()
        log.VPrint("This is a verbose log message")
        log.VPrintf("This is a verbose log message with %s\n", "formatting")

        // Debug Log Messages -> EnableDebug()
        log.Debug("This is a debug log message")
        log.Debug("This is a debug log message with %s\n", "formatting")

        // Trace Log Messages -> EnableTrace()
        log.Trace("This is a trace log message")
        log.Tracef("This is a trace log message with %s\n", "formatting")

        // Success Log Messages
        log.Success("This is a success log message")
        log.Successf("This is a success log message with %s\n", "formatting")

        // Warning Log Messages
        log.Warning("This is a warning log message")
        log.Warningf("This is a warning log message with %s\n", "formatting")

        // Failure Log Messages
        log.Failure("This is a failure log message")
        log.Failuref("This is a failure log message with %s\n", "formatting")

        // Error Log Messages
        log.Error("This is an error log message")
        log.Errorf("This is an errorlog message with %s\n", "formatting")

        // Panic Log Messages
        log.Panic("This is a panic log message")
        log.Panicf("This is a panic log message with %s\n", "formatting")

        // Fatal Log Messages
        log.Fatal("This is a fatal log message")
        log.Fatalf("This is a fatal log message with %s\n", "formatting")     
    }

## Practical Example

When developing Command Line Utilities (CLI) using Cobra/Viper you can do the following in root.go with package gogo/log

    // rootCmd represents the base command when called without any subcommands
    var rootCmd = &cobra.Command{
        Use:   "myCommand",
        Short: "A brief description of your application",
        Long: `A longer description that spans multiple lines and likely contains
    examples and usage of using your application. For example:

    Cobra is a CLI library for Go that empowers applications.
    This application is a tool to generate the needed files
    to quickly create a Cobra application.`,

        PersistentPreRun: func(cmd *cobra.Command, args []string) {

            // Configure Verbose Logging Option from CLI
            if options.verbose {
                log.EnableVerbose()
            }

            // Configure Logging options from Environment
            logLevel := os.Getenv("MYCOMMAND_DEBUG")

            if logLevel == "verbose" {
                log.EnableVerbose()
            }

            if logLevel == "debug" {
                log.EnableVerbose()
                log.EnableDebug()
            }

            if logLevel == "trace" {
                log.EnableVerbose()
                log.EnableDebug()
                log.EnableTrace()
            }

        },
    }

    type rootFlags struct {
    	verbose              bool
	    nonInteractive       bool
	    overrideConfirmation bool
    }   

    var options rootFlags

    func init() {
	    // Commands
	
	    // PersistentFlags
	    rootCmd.PersistentFlags().BoolVarP(&options.verbose, "verbose", "v", false, "Enable verbose output")
	    rootCmd.PersistentFlags().BoolVar(&options.nonInteractive, "non-interactive", false, "Disable interactive prompts")
	    rootCmd.PersistentFlags().BoolVarP(&options.overrideConfirmation, "yes", "y", false, "Override confirmations")
    }