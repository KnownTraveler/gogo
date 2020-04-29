// Copyright © 2020 Brian Hooper <knowntraveler.io>
// Author: Brian Hooper (@KnownTraveler)
// Project: gogo/log

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package log provides a uniform logging api for formatting log output
package log

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TEST HELPER FUNCTIONS

// captureStdout() is a Helper Function for capturing Stdout into Buffer
func captureStdout(f func()) string {
	// Str Buffer needs no initialization.
	var str bytes.Buffer

	// Set the output destination for the standard logger
	log.SetOutput(&str)

	// Run the Log Function
	f()

	// Return the captured Stdout Log Message
	return str.String()
}

// STANDARD LOG MESSAGES

// TestPrint is a unit test for log.Print()
func TestPrint(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Print("Standard Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[96mStandard Log Message\x1b[0m\n", output)
}

// TestPrintf is a unit test for log.Printf()
func TestPrintf(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Printf("Standard Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[96mStandard Log Message with formatting\x1b[0m\n", output)
}

// VERBOSE LOG MESSAGES

// TestVPrint is a unit test for log.VPrint()
func TestVPrint(t *testing.T) {
	// Enable Verbose Logging
	EnableVerbose()
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		VPrint("Verbose Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[96mINFO: Verbose Log Message\x1b[0m\n", output)
}

// TestVPrintf is a unit test for log.VPrintf()
func TestVPrintf(t *testing.T) {
	// Enable Verbose Logging
	EnableVerbose()
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		VPrintf("Verbose Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[96mINFO: Verbose Log Message with formatting\x1b[0m\n", output)
}

// DEBUG LOG MESSAGES

// TestDebug is a unit test for log.Debug()
func TestDebug(t *testing.T) {
	// Enable Debug Logging
	EnableDebug()
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Debug("Debug Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "DEBUG: Debug Log Message\n", output)
}

// TestDebugf is a unit test for log.Debugf()
func TestDebugf(t *testing.T) {
	// Enable Debug Logging
	EnableDebug()
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Debugf("Debug Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "DEBUG: Debug Log Message with formatting\n", output)
}

// TRACE LOG MESSAGES

// TestTrace is a unit test for log.Trace()
func TestTrace(t *testing.T) {
	// Enable Trace Logging
	EnableTrace()
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Trace("Trace Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "TRACE: Trace Log Message\n", output)
}

// TestTracef is a unit test for log.Tracef()
func TestTracef(t *testing.T) {
	// Enable Trace Logging
	EnableTrace()
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Tracef("Trace Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "TRACE: Trace Log Message with formatting\n", output)
}

// SUCCESS LOG MESSAGES

// TestSuccess is a unit test for log.Success()
func TestSuccess(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Success("Success Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[92mSUCCESS: Success Log Message\x1b[0m\n", output)
}

// TestSuccessf is a unit test for log.Successf()
func TestSuccessf(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Successf("Success Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[92mSUCCESS: Success Log Message with formatting\x1b[0m\n", output)
}

// WARN LOG MESSAGES

// TestWarning is a unit test for log.Warning()
func TestWarning(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Warning("Warning Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[93mWARNING: Warning Log Message\x1b[0m\n", output)
}

// TestWarningf is a unit test for log.Warningf()
func TestWarningf(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Warningf("Warning Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[93mWARNING: Warning Log Message with formatting\x1b[0m\n", output)
}

// FAILURE LOG MESSAGES

// TestFailure is a unit test for log.Failure()
func TestFailure(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Failure("Failure Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[91mFAILURE: Failure Log Message\x1b[0m\n", output)
}

// TestFailuref is a unit test for log.Failuref()
func TestFailuref(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Failuref("Failure Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[91mFAILURE: Failure Log Message with formatting\x1b[0m\n", output)
}

// ERROR LOG MESSAGES

// TestError is a unit test for log.Error()
func TestError(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Error("Error Log Message")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[91mERROR: Error Log Message\x1b[0m\n", output)
}

// TestErrorf is a unit test for log.Errorf()
func TestErrorf(t *testing.T) {
	// Caputure Stdout for Log Message
	output := captureStdout(func() {
		Errorf("Error Log Message with %v", "formatting")
	})
	// Assert Unit Test
	assert.Equal(t, "\x1b[91mERROR: Error Log Message with formatting\x1b[0m\n", output)
}

// PANIC LOG MESSAGES

// // TestPanic is a unit test for log.Panic()
// func TestPanic(t *testing.T) {
// 	// Caputure Stdout for Log Message
// 	output := captureStdout(func() {
// 		Panic("Panic Log Message")
// 	})
// 	// Assert Unit Test
// 	assert.Equal(t, "\x1b[91mPANIC: Panic Log Message\x1b[0m\n", output)
// }

// // TestPanicf is a unit test for log.Panicf()
// func TestPanicf(t *testing.T) {
// 	// Caputure Stdout for Log Message
// 	output := captureStdout(func() {
// 		Panicf("Panic Log Message with %v", "formatting")
// 	})
// 	// Assert Unit Test
// 	assert.Equal(t, "\x1b[91mPANIC: Panic Log Message with formatting\x1b[0m\n", output)
// }

// FATAL LOG MESSAGES

// // TestFatal is a unit test for log.Fatal()
// func TestFatal(t *testing.T) {
// 	// Caputure Stdout for Log Message
// 	output := captureStdout(func() {
// 		Fatal("Fatal Log Message")
// 	})
// 	// Assert Unit Test
// 	assert.Equal(t, "\x1b[91mFATAL: Fatal Log Message\x1b[0m\n", output)
// }

// // TestFatalf is a unit test for log.Fatalf()
// func TestFatalf(t *testing.T) {
// 	// Caputure Stdout for Log Message
// 	output := captureStdout(func() {
// 		Fatalf("Fatal Log Message with %v", "formatting")
// 	})
// 	// Assert Unit Test
// 	assert.Equal(t, "\x1b[91mFATAL: Fatal Log Message with formatting\x1b[0m\n", output)
// }
