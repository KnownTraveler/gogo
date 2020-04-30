// Copyright © 2020 Brian Hooper <knowntraveler.io>
// Author: Brian Hooper (@KnownTraveler)
// Project: gogo/fs

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package fs provides a uniform api for filesystem-related functions
package fs

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

// CreateDirectory simply checks if the directory path already
// exists before attempting to create the directory
func CreateDirectory(path string, mode os.FileMode) error {

	// Check IF Directory Exists
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("Directory '%v' already exists", path)
	}

	// Create Directory
	err = os.Mkdir(path, mode)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDirectory simply checks if the directory path already
// exists before attempting to delete the directory
func DeleteDirectory(path string) error {

	// Check IF Directory Exists
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Directory '%v' doesn't exist", path)
	}

	// Delete Directory
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDirectoryAll simply checks if the directory path already
// exists before attempting to delete the directory and any child paths
func DeleteDirectoryAll(path string) error {

	// Check IF Directory Exists
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Directory '%v' doesn't exist", path)
	}

	// Delete Directory All
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}

// CreateFile simply checks if the file path already
// exists before attempting to create the file
func CreateFile(path string, mode os.FileMode) error {

	// Check IF File Exists
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("File '%v' already exists", path)
	}

	// Create File
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Set File Permissions
	err = os.Chmod(path, mode)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFile simply deletes a file if it exists
func DeleteFile(path string) error {

	// Check IF File Exists
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("File '%v' doesn't exist", path)
	}

	// Delete File
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// ReadFile simply checks if the file path already
// exists before attempting to read the file
// if successful, returns a []byte array of data
func ReadFile(path string) ([]byte, error) {

	// Check IF File Exists
	_, err := os.Stat(path)
	if err != nil {
		return []byte{}, fmt.Errorf("File '%v' doesn't exist", path)
	}

	// Read File
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	return data, nil

}

// WriteFile simply checks if the file path already
// exists before attempting to create and write the file
func WriteFile(path string, mode os.FileMode, data []byte) error {

	// Check IF File Exists
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("File '%v' already exists", path)
	}

	// Create File
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write File
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	// Save File Changes
	err = file.Sync()
	if err != nil {
		return err
	}

	// Set File Permissions
	err = os.Chmod(path, mode)
	if err != nil {
		return err
	}

	return nil
}

// HomeDirectory returns the home directory for the executing user.
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
// SEE: https://github.com/mitchellh/go-homedir/blob/master/homedir.go
func HomeDirectory() (string, error) {
	// FIND USER HOME DIRECTORY
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// CreateSymlink simply creates a symbolic link after verifing
// the source exists
func CreateSymlink(source string, target string) error {

	// Check IF Source Exists
	_, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("Source '%v' doesn't exist", source)
	}

	err = os.Symlink(source, target)
	if err != nil {
		return err
	}

	return err

}
