package tail

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

// TestBuffer implements OutputWriter for testing
type TestBuffer struct {
	buffer *bytes.Buffer
}

func NewTestBuffer() *TestBuffer {
	return &TestBuffer{
		buffer: &bytes.Buffer{},
	}
}

func (tb *TestBuffer) Write(data []byte) (int, error) {
	return tb.buffer.Write(data)
}

func (tb *TestBuffer) String() string {
	return tb.buffer.String()
}

func (tb *TestBuffer) Reset() {
	tb.buffer.Reset()
}

// TestOpenFile tests the OpenFile function with various scenarios
func TestOpenFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		setup    func() string
		cleanup  func(string)
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "empty file path",
			filePath: "",
			setup:    func() string { return "" },
			cleanup:  func(s string) {},
			wantErr:  true,
			errMsg:   "file path cannot be empty",
		},
		{
			name:     "non-existent file",
			filePath: "nonexistent.txt",
			setup:    func() string { return "nonexistent.txt" },
			cleanup:  func(s string) {},
			wantErr:  true,
			errMsg:   "file does not exist",
		},
		{
			name:     "valid file",
			filePath: "testfile.txt",
			setup: func() string {
				content := "test content\n"
				err := os.WriteFile("testfile.txt", []byte(content), 0644)
				if err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return "testfile.txt"
			},
			cleanup: func(filename string) {
				os.Remove(filename)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := tt.setup()
			defer tt.cleanup(filename)

			file, err := OpenFile(tt.filePath)

			if tt.wantErr {
				if err == nil {
					t.Errorf("OpenFile() expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("OpenFile() error = %v, want error containing %v", err, tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("OpenFile() unexpected error = %v", err)
				return
			}

			if file == nil {
				t.Errorf("OpenFile() returned nil file")
				return
			}

			if file.Path != tt.filePath {
				t.Errorf("OpenFile() file.Path = %v, want %v", file.Path, tt.filePath)
			}

			if file.Handle == nil {
				t.Errorf("OpenFile() file.Handle is nil")
			}

			if file.Output == nil {
				t.Errorf("OpenFile() file.Output is nil")
			}

			// Clean up the file handle
			file.Close()
		})
	}
}

// TestOpenFileWithOutput tests the OpenFileWithOutput function
func TestOpenFileWithOutput(t *testing.T) {
	// Create a test file
	content := "test content\n"
	filename := "testfile_with_output.txt"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open file with custom output
	file, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("OpenFileWithOutput() error = %v", err)
	}
	defer file.Close()

	if file.Output != testBuffer {
		t.Errorf("OpenFileWithOutput() file.Output = %v, want %v", file.Output, testBuffer)
	}
}

// TestFile_SeekToEnd tests the SeekToEnd method
func TestFile_SeekToEnd(t *testing.T) {
	// Create a test file with some content
	content := "line 1\nline 2\nline 3\n"
	filename := "seektest.txt"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	// Open the file
	file, err := OpenFile(filename)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer file.Close()

	// Test SeekToEnd
	err = file.SeekToEnd()
	if err != nil {
		t.Errorf("SeekToEnd() error = %v", err)
	}

	// Verify we're at the end by trying to read
	buffer := make([]byte, 1)
	n, err := file.Handle.Read(buffer)
	if err == nil && n > 0 {
		t.Errorf("SeekToEnd() did not position at end of file, read %d bytes", n)
	}
}

// TestFile_Close tests the Close method
func TestFile_Close(t *testing.T) {
	// Create a test file
	content := "test content\n"
	filename := "closetest.txt"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	// Open the file
	file, err := OpenFile(filename)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	// Test Close
	err = file.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}

	// Test closing an already closed file (should not error)
	err = file.Close()
	if err != nil {
		t.Errorf("Close() on already closed file error = %v", err)
	}
}

// TestFile_Close_NilHandle tests Close with nil handle
func TestFile_Close_NilHandle(t *testing.T) {
	file := &File{
		Path:   "test.txt",
		Handle: nil,
	}

	err := file.Close()
	if err != nil {
		t.Errorf("Close() with nil handle error = %v", err)
	}
}

// TestFile_Monitor tests the Monitor method
func TestFile_Monitor(t *testing.T) {
	// Create a test file
	filename := "monitortest.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	file.Close()
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open the file for monitoring with test buffer
	monitorFile, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Seek to end
	err = monitorFile.SeekToEnd()
	if err != nil {
		t.Fatalf("failed to seek to end: %v", err)
	}

	// Create stop channel
	stopChan := make(chan struct{})

	// Start monitoring in a goroutine
	monitorDone := make(chan error, 1)
	go func() {
		err := monitorFile.Monitor(stopChan)
		monitorDone <- err
	}()

	// Give some time for the monitor to start
	time.Sleep(50 * time.Millisecond)

	// Write some content to the file
	testContent := "test line 1\ntest line 2\n"
	err = os.WriteFile(filename, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("failed to write test content: %v", err)
	}

	// Give some time for the monitor to process
	time.Sleep(100 * time.Millisecond)

	// Stop monitoring
	close(stopChan)

	// Wait for monitor to finish
	select {
	case err := <-monitorDone:
		if err != nil {
			t.Errorf("Monitor() error = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("Monitor() did not stop within timeout")
	}

	// Validate the output content
	output := testBuffer.String()
	expectedOutput := "test line 1\ntest line 2\n"
	if output != expectedOutput {
		t.Errorf("Monitor() output = %q, want %q", output, expectedOutput)
	}
}

// TestFile_Monitor_StopImmediately tests Monitor with immediate stop
func TestFile_Monitor_StopImmediately(t *testing.T) {
	// Create a test file
	filename := "monitortest2.txt"
	content := "test content\n"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open the file for monitoring with test buffer
	monitorFile, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Seek to end
	err = monitorFile.SeekToEnd()
	if err != nil {
		t.Fatalf("failed to seek to end: %v", err)
	}

	// Create stop channel and close it immediately
	stopChan := make(chan struct{})
	close(stopChan)

	// Start monitoring
	err = monitorFile.Monitor(stopChan)
	if err != nil {
		t.Errorf("Monitor() with immediate stop error = %v", err)
	}

	// Validate no output was produced
	output := testBuffer.String()
	if output != "" {
		t.Errorf("Monitor() with immediate stop produced output: %q, want empty", output)
	}
}

// TestFile_Monitor_FileDeleted tests Monitor when file is deleted
func TestFile_Monitor_FileDeleted(t *testing.T) {
	// Create a test file
	filename := "deletetest.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	file.Close()
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open the file for monitoring with test buffer
	monitorFile, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Seek to end
	err = monitorFile.SeekToEnd()
	if err != nil {
		t.Fatalf("failed to seek to end: %v", err)
	}

	// Create stop channel
	stopChan := make(chan struct{})

	// Start monitoring in a goroutine
	monitorDone := make(chan error, 1)
	go func() {
		err := monitorFile.Monitor(stopChan)
		monitorDone <- err
	}()

	// Give some time for the monitor to start
	time.Sleep(50 * time.Millisecond)

	// Delete the file
	err = os.Remove(filename)
	if err != nil {
		t.Fatalf("failed to delete test file: %v", err)
	}

	// Wait for monitor to detect the deletion
	select {
	case err := <-monitorDone:
		if err == nil {
			t.Errorf("Monitor() should have returned error when file was deleted")
		}
		if !strings.Contains(err.Error(), "no longer exists") {
			t.Errorf("Monitor() error = %v, want error containing 'no longer exists'", err)
		}
	case <-time.After(3 * time.Second):
		t.Errorf("Monitor() did not detect file deletion within timeout")
	}
}

// TestFile_Monitor_Streaming tests Monitor with streaming content
func TestFile_Monitor_Streaming(t *testing.T) {
	// Create a test file
	filename := "streamtest.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	file.Close()
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open the file for monitoring with test buffer
	monitorFile, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Seek to end
	err = monitorFile.SeekToEnd()
	if err != nil {
		t.Fatalf("failed to seek to end: %v", err)
	}

	// Create stop channel
	stopChan := make(chan struct{})

	// Start monitoring in a goroutine
	monitorDone := make(chan error, 1)
	go func() {
		err := monitorFile.Monitor(stopChan)
		monitorDone <- err
	}()

	// Give some time for the monitor to start
	time.Sleep(50 * time.Millisecond)

	// Write content in chunks to simulate streaming
	testLines := []string{
		"line 1\n",
		"line 2\n",
		"line 3\n",
		"line 4\n",
	}

	for i, line := range testLines {
		// Append to file
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			t.Fatalf("failed to open file for writing: %v", err)
		}
		_, err = f.WriteString(line)
		f.Close()
		if err != nil {
			t.Fatalf("failed to write line %d: %v", i, err)
		}

		// Give some time for the monitor to process
		time.Sleep(100 * time.Millisecond)
	}

	// Stop monitoring
	close(stopChan)

	// Wait for monitor to finish
	select {
	case err := <-monitorDone:
		if err != nil {
			t.Errorf("Monitor() error = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("Monitor() did not stop within timeout")
	}

	// Validate the output content
	output := testBuffer.String()
	expectedOutput := "line 1\nline 2\nline 3\nline 4\n"
	if output != expectedOutput {
		t.Errorf("Monitor() streaming output = %q, want %q", output, expectedOutput)
	}
}

// TestFile_Monitor_PartialLines tests Monitor with partial lines
func TestFile_Monitor_PartialLines(t *testing.T) {
	// Create a test file
	filename := "partialtest.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	file.Close()
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open the file for monitoring with test buffer
	monitorFile, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Seek to end
	err = monitorFile.SeekToEnd()
	if err != nil {
		t.Fatalf("failed to seek to end: %v", err)
	}

	// Create stop channel
	stopChan := make(chan struct{})

	// Start monitoring in a goroutine
	monitorDone := make(chan error, 1)
	go func() {
		err := monitorFile.Monitor(stopChan)
		monitorDone <- err
	}()

	// Give some time for the monitor to start
	time.Sleep(50 * time.Millisecond)

	// Write partial line
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatalf("failed to open file for writing: %v", err)
	}
	_, err = f.WriteString("partial line")
	f.Close()
	if err != nil {
		t.Fatalf("failed to write partial line: %v", err)
	}

	// Give some time for the monitor to process
	time.Sleep(100 * time.Millisecond)

	// Complete the line
	f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatalf("failed to open file for writing: %v", err)
	}
	_, err = f.WriteString(" completed\n")
	f.Close()
	if err != nil {
		t.Fatalf("failed to complete line: %v", err)
	}

	// Give some time for the monitor to process
	time.Sleep(100 * time.Millisecond)

	// Stop monitoring
	close(stopChan)

	// Wait for monitor to finish
	select {
	case err := <-monitorDone:
		if err != nil {
			t.Errorf("Monitor() error = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("Monitor() did not stop within timeout")
	}

	// Validate the output content
	output := testBuffer.String()
	expectedOutput := "partial line completed\n"
	if output != expectedOutput {
		t.Errorf("Monitor() partial lines output = %q, want %q", output, expectedOutput)
	}
}

// TestFile_Monitor_FileTruncated tests Monitor when file is truncated
func TestFile_Monitor_FileTruncated(t *testing.T) {
	// Create a test file with initial content
	filename := "truncatetest.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	_, err = file.WriteString("initial line 1\ninitial line 2\n")
	file.Close()
	if err != nil {
		t.Fatalf("failed to write initial content: %v", err)
	}
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open the file for monitoring with test buffer
	monitorFile, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Seek to end
	err = monitorFile.SeekToEnd()
	if err != nil {
		t.Fatalf("failed to seek to end: %v", err)
	}

	// Create stop channel
	stopChan := make(chan struct{})

	// Start monitoring in a goroutine
	monitorDone := make(chan error, 1)
	go func() {
		err := monitorFile.Monitor(stopChan)
		monitorDone <- err
	}()

	// Give some time for the monitor to start
	time.Sleep(50 * time.Millisecond)

	// Truncate the file (simulate log rotation where content is deleted)
	err = os.Truncate(filename, 0)
	if err != nil {
		t.Fatalf("failed to truncate test file: %v", err)
	}

	// Wait for monitor to detect the truncation
	select {
	case err := <-monitorDone:
		if err == nil {
			t.Errorf("Monitor() should have returned error when file was truncated")
		}
		if !strings.Contains(err.Error(), "was truncated") {
			t.Errorf("Monitor() error = %v, want error containing 'was truncated'", err)
		}
	case <-time.After(3 * time.Second):
		t.Errorf("Monitor() did not detect file truncation within timeout")
	}
}

// TestFile_Monitor_FileReplaced tests Monitor when file is replaced (inode changes)
func TestFile_Monitor_FileReplaced(t *testing.T) {
	// Create a test file with initial content
	filename := "replacetest.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	_, err = file.WriteString("initial content\n")
	file.Close()
	if err != nil {
		t.Fatalf("failed to write initial content: %v", err)
	}
	defer os.Remove(filename)

	// Create test buffer
	testBuffer := NewTestBuffer()

	// Open the file for monitoring with test buffer
	monitorFile, err := OpenFileWithOutput(filename, testBuffer)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Seek to end
	err = monitorFile.SeekToEnd()
	if err != nil {
		t.Fatalf("failed to seek to end: %v", err)
	}

	// Create stop channel
	stopChan := make(chan struct{})

	// Start monitoring in a goroutine
	monitorDone := make(chan error, 1)
	go func() {
		err := monitorFile.Monitor(stopChan)
		monitorDone <- err
	}()

	// Give some time for the monitor to start
	time.Sleep(50 * time.Millisecond)

	// Replace the file (simulate log rotation where file is renamed and new one created)
	err = os.Remove(filename)
	if err != nil {
		t.Fatalf("failed to remove original file: %v", err)
	}

	// Create new file with same name (will have different inode)
	newFile, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create replacement file: %v", err)
	}
	_, err = newFile.WriteString("new content\n")
	newFile.Close()
	if err != nil {
		t.Fatalf("failed to write new content: %v", err)
	}

	// Wait for monitor to detect the replacement
	select {
	case err := <-monitorDone:
		if err == nil {
			t.Errorf("Monitor() should have returned error when file was replaced")
		}
		if !strings.Contains(err.Error(), "was renamed or replaced") {
			t.Errorf("Monitor() error = %v, want error containing 'was renamed or replaced'", err)
		}
	case <-time.After(3 * time.Second):
		t.Errorf("Monitor() did not detect file replacement within timeout")
	}
}

// TestFile_CheckFileState tests the checkFileState method
func TestFile_CheckFileState(t *testing.T) {
	// Create a test file
	filename := "statetest.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	_, err = file.WriteString("test content\n")
	file.Close()
	if err != nil {
		t.Fatalf("failed to write content: %v", err)
	}
	defer os.Remove(filename)

	// Open the file for monitoring
	monitorFile, err := OpenFile(filename)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}
	defer monitorFile.Close()

	// Test that checkFileState works normally
	err = monitorFile.checkFileState()
	if err != nil {
		t.Errorf("checkFileState() error = %v, want nil", err)
	}

	// Test that checkFileState detects truncation
	err = os.Truncate(filename, 0)
	if err != nil {
		t.Fatalf("failed to truncate file: %v", err)
	}

	err = monitorFile.checkFileState()
	if err == nil {
		t.Errorf("checkFileState() should have returned error when file was truncated")
	}
	if !strings.Contains(err.Error(), "was truncated") {
		t.Errorf("checkFileState() error = %v, want error containing 'was truncated'", err)
	}
}
