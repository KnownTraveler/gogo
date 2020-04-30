// Copyright © 2020 Brian Hooper <knowntraveler.io>
// Author: Brian Hooper (@KnownTraveler)
// Project: gogo/fs

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package zip provides a uniform api for archive/zip related functions
package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Download Function for Downloading an Archive File (.zip) from a HTTP Source
func Download(source string, target string) error {

	// Parse source url and validate 'source' is a valid HTTP URL
	_, err := url.ParseRequestURI(source)
	if err != nil {
		return err
	}

	// Get the source data
	resp, err := http.Get(source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the .zip file on the local filesystem
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to .zip file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil

}

// Archive Function for Zipping an Archive File (.zip) from local filesystem
func Archive(source string, target string) error {

	// Validate Target Parameter
	if target == "" {
		return fmt.Errorf("The 'target' parameter was empty. A target is required to create a Zip Archive")
	}

	// Validate Source Parameter
	if source == "" {
		return fmt.Errorf("The 'source' parameter was empty. A source is required to create a Zip Archive")
	}

	// Create Archive
	err := createArchive(source, target)
	if err != nil {
		return err
	}

	return nil
}

// createArchive Function for Creating an Archive File (.zip) from a source on local filesystem
func createArchive(source string, target string) error {

	// Create Zip Archive File
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	// Create New Writer for Zipfile
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// Verify New Zip Archive File Exists
	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	// Verify Source is a Directory
	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	// Walk Source Filepath
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get File Header Info
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Verify Base Directory
		if baseDir != "" {
			if baseDir == "." {
				// Set Archive File Header
				header.Name = filepath.ToSlash(filepath.Join(baseDir, path))
			} else {
				// Set Archive File Header
				prefix := baseDir + "/"
				header.Name = filepath.ToSlash(path)
				header.Name = filepath.ToSlash(filepath.Join(strings.TrimPrefix(header.Name, prefix)))

				// ROOT DIRECTORY CHECK
				// Check if baseDir matches header.Name
				if baseDir == header.Name {
					return nil
				}
			}
		}

		// Check if Archive File Header is a Directory
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		// Create Header for Source File
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Open Source File
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy Source File to Archive (.zip)
		_, err = io.Copy(writer, file)
		return err
	})

	return nil
}

// Unarchive Function for Unzipping an Archive File (.zip)
func Unarchive(source string, target string) error {

	// Create a zipReader out of the Source (.zip)
	zipReader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// Iterate through each File/Directory found in Source Archive (.zip)
	for _, file := range zipReader.Reader.File {

		// Open the file inside the zip archive like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		// Specify what the extracted file name should be.
		// You can specify a full path or a prefix to move it to a different directory.
		var targetDir string
		if target == "" {
			targetDir = "./"
		} else {
			targetDir = target
		}

		// Set Extracted Filepath
		extractedFilePath := filepath.Join(targetDir, file.Name)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Check if Directory Path Exists
			if _, err := os.Stat(filepath.Dir(extractedFilePath)); os.IsNotExist(err) {
				// Directory Path Does Not Exist
				// Create Directory Path
				os.MkdirAll(filepath.Dir(extractedFilePath), 0755)
			}
			// Create directories to recreate directory structure inside the zip archive.
			// Also preserves permissions
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			// Check if File Path Exists
			if _, err := os.Stat(filepath.Dir(extractedFilePath)); os.IsNotExist(err) {
				// File Directory Path Does Not Exist
				// Create Directory Path
				os.MkdirAll(filepath.Dir(extractedFilePath), 0755)
			}

			// Create an output file for writing
			f, err := os.Create(extractedFilePath)
			if err != nil {
				return err
			}
			defer f.Close()

			// "Extract" the file by copying zipped file contents to the output file
			_, err = io.Copy(f, zippedFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
