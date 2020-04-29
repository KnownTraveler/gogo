// Copyright © 2020 Brian Hooper <knowntraveler.io>
// Author: Brian Hooper (@KnownTraveler)
// Project: gogo/log

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package log provides a uniform logging api for formatting log output on linux/windows
package log

import (
	"fmt"
	"log"
	"os"

	auroraPackage "github.com/logrusorgru/aurora"
	colorable "github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	isatty "github.com/onsi/ginkgo/reporters/stenographer/support/go-isatty"
)

var aurora auroraPackage.Aurora

// Flag to Enable Verbose Logging
var verboseEnabled bool

// Flag to Enable Debug Logging
var debugEnabled bool

// Flog to Enable Trace Logging
var traceEnabled bool

func init() {
	aurora = auroraPackage.NewAurora(isatty.IsTerminal(os.Stdout.Fd()))
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFlags(0)
}

// Print logs a message at level Info
func Print(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightCyan("%v").String(), message))
}

// Printf logs a formatted message at level Info
func Printf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Printf(fmt.Sprintf(aurora.BrightCyan("%v").String(), message))
}

// VPrint logs a message at level Info when verboseEnabled is true
func VPrint(message string) {
	if verboseEnabled {
		log.Printf(fmt.Sprintf(aurora.BrightCyan("INFO: %v").String(), message))
	}
}

// VPrintf logs a message at level Info when verboseEnabled is true
func VPrintf(format string, args ...interface{}) {
	if verboseEnabled {
		message := fmt.Sprintf(format, args...)
		log.Printf(fmt.Sprintf(aurora.BrightCyan("INFO: %v").String(), message))
	}
}

// Success logs a message at level Info
func Success(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightGreen("SUCCESS: %v").String(), message))
}

// Successf logs a formatted message at level Info
func Successf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Printf(fmt.Sprintf(aurora.BrightGreen("SUCCESS: %v").String(), message))
}

// Warning logs a message at level Warn
func Warning(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightYellow("WARNING: %v").String(), message))
}

// Warningf logs a formatted message at level Warn
func Warningf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Printf(fmt.Sprintf(aurora.BrightYellow("WARNING: %v").String(), message))
}

// Failure logs a message at level Error
func Failure(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightRed("FAILURE: %v").String(), message))
}

// Failuref logs a formatted message at level Error
func Failuref(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Printf(fmt.Sprintf(aurora.BrightRed("FAILURE: %v").String(), message))
}

// Error logs a message at level Error
func Error(message string) {
	log.Printf(fmt.Sprintf(aurora.BrightRed("ERROR: %v").String(), message))
}

// Errorf logs a message at level Error
func Errorf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Printf(fmt.Sprintf(aurora.BrightRed("ERROR: %v").String(), message))
}

// Panic logs a message at level Panic
func Panic(message string) {
	log.Panicf(fmt.Sprintf(aurora.BrightRed("PANIC: %v").String(), message))
}

// Panicf logs a formatted message at level Panic
func Panicf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Panicf(fmt.Sprintf(aurora.BrightRed("PANIC: %v").String(), message))
}

// Fatal logs a message at level Fatal
func Fatal(message string) {
	log.Fatalf(fmt.Sprintf(aurora.BrightRed("FATAL: %v").String(), message))
}

// Fatalf logs a formatted message at level Fatal
func Fatalf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Fatalf(fmt.Sprintf(aurora.BrightRed("FATAL: %v").String(), message))
}

// Debug logs a message at level Debug
func Debug(message string) {
	if debugEnabled {
		log.Printf("DEBUG: %v", message)
	}
}

// Debugf logs a formatted message at level Debug
func Debugf(format string, args ...interface{}) {
	if debugEnabled {
		message := fmt.Sprintf(format, args...)
		log.Printf("DEBUG: %v", message)
	}
}

// Trace logs a message at level Trace
func Trace(message string) {
	if traceEnabled {
		log.Printf("TRACE: %v", message)
	}
}

// Tracef logs a formatted message at level Trace
func Tracef(format string, args ...interface{}) {
	if traceEnabled {
		message := fmt.Sprintf(format, args...)
		log.Printf("TRACE: %v", message)
	}
}

// EnableVerbose turns on verbose logging
func EnableVerbose() {
	verboseEnabled = true
}

// EnableDebug turns on enable logging
func EnableDebug() {
	debugEnabled = true
}

// EnableTrace turns on trace logging
func EnableTrace() {
	traceEnabled = true
}
